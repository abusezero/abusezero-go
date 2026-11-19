// Package abusezero is the AbuseZero Go SDK. AbuseZero is an abuse prevention API for SaaS.
//
// Create one Client with a secret key, a restricted key or an OAuth access token, or a
// PublicClient for public reads, and reuse it across requests. README.md describes the
// client, and API.md lists every method.
package abusezero

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed operations.json
var operationsJSON []byte
var operations map[string]operation

func init() {
	if err := json.Unmarshal(operationsJSON, &operations); err != nil {
		panic(err)
	}
}

type operation struct {
	ID, Method, Path, Scope string
	Public, Mutation        bool
	Timeout                 int `json:"timeout_ms"`
	Paths                   []struct {
		Name   string
		Schema map[string]any
	}
	Input    map[string]any
	Response map[string]any
}

// Error is the error calls and the identifier helpers return for AbuseZero and SDK failures. Use errors.As to
// read it. A canceled context returns context.Canceled, and VerifyWebhook returns a plain error.
type Error struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
	RequestID string `json:"request_id,omitempty"`
	Retryable bool   `json:"retryable"`
	DocsURL   string `json:"docs_url,omitempty"`
}

// Error returns the code and message.
func (e *Error) Error() string            { return e.Code + ": " + e.Message }
func invalid(code, message string) *Error { return &Error{Code: code, Message: message, Status: 422} }

// Optional tells an omitted field apart from an explicit null. The zero value omits the field.
type Optional[T any] struct {
	Value   T
	Present bool
	Null    bool
}

// Some returns an Optional that sends value.
func Some[T any](value T) Optional[T] { return Optional[T]{Value: value, Present: true} }

// Null returns an Optional that sends null, where the schema allows it.
func Null[T any]() Optional[T] { return Optional[T]{Present: true, Null: true} }

// IsZero reports whether the field is omitted.
func (v Optional[T]) IsZero() bool { return !v.Present }
func (v Optional[T]) MarshalJSON() ([]byte, error) {
	if v.Null || !v.Present {
		return []byte("null"), nil
	}
	return json.Marshal(v.Value)
}
func (v *Optional[T]) UnmarshalJSON(b []byte) error {
	v.Present = true
	v.Null = string(b) == "null"
	if v.Null {
		return nil
	}
	return json.Unmarshal(b, &v.Value)
}
func marshalExtras(value any, extra map[string]json.RawMessage) ([]byte, error) {
	b, e := json.Marshal(value)
	if e != nil {
		return nil, e
	}
	var m map[string]json.RawMessage
	if e = json.Unmarshal(b, &m); e != nil {
		return nil, e
	}
	for k, v := range extra {
		if _, ok := m[k]; !ok {
			m[k] = v
		}
	}
	return json.Marshal(m)
}
func unmarshalExtras(b []byte, value any, extra *map[string]json.RawMessage, known []string) error {
	if e := json.Unmarshal(b, value); e != nil {
		return e
	}
	var m map[string]json.RawMessage
	if e := json.Unmarshal(b, &m); e != nil {
		return e
	}
	for _, k := range known {
		delete(m, k)
	}
	*extra = m
	return nil
}

// Options configures a client. Pass exactly one of Key and AccessToken to New; NewPublic needs neither.
type Options struct {
	Key, AccessToken, Workspace, Outage, Agent, BaseURL string
	Transport                                           http.RoundTripper
}

// CallOptions carries an operation's path fields and its idempotency key, which mutations require.
type CallOptions struct {
	Path           map[string]any
	IdempotencyKey string
}
type core struct {
	credential, workspace, outage, agent, base string
	key, public                                bool
	http                                       *http.Client
}

// Client calls the authenticated operations. Create one with New and reuse it.
type Client struct{ *core }

// PublicClient calls the public operations without credentials. Create one with NewPublic.
type PublicClient struct{ *core }

var serverKey = regexp.MustCompile(`^(sk|rk)_[1-9][0-9]{19}\.[A-Za-z0-9_-]{43}$`)
var opaque = regexp.MustCompile(`^[\x21-\x7e]{1,128}$`)
var publicID = regexp.MustCompile(`^[1-9][0-9]{19}$`)

// New creates a client for the authenticated operations, with a secret key, a restricted key or an OAuth
// access token and its workspace.
func New(options Options) (*Client, error) {
	c, e := newCore(options, false)
	if e != nil {
		return nil, e
	}
	return &Client{c}, nil
}

// NewPublic creates a client for the public operations. It sends no credentials.
func NewPublic(options Options) (*PublicClient, error) {
	c, e := newCore(options, true)
	if e != nil {
		return nil, e
	}
	return &PublicClient{c}, nil
}
func newCore(o Options, public bool) (*core, error) {
	if !public && ((o.Key == "") == (o.AccessToken == "")) {
		return nil, invalid("CREDENTIAL_REQUIRED", "Pass exactly one of Key and AccessToken.")
	}
	if o.Key != "" && !serverKey.MatchString(o.Key) {
		return nil, invalid("INVALID_KEY", "Use a secret key (sk_…) or restricted key (rk_…).")
	}
	if o.AccessToken != "" && (len(o.AccessToken) > 8192 || !regexp.MustCompile(`^[\x21-\x7e]+$`).MatchString(o.AccessToken)) {
		return nil, invalid("INVALID_TOKEN", "Pass the OAuth access token exactly as issued.")
	}
	if o.AccessToken != "" && o.Workspace == "" {
		return nil, invalid("WORKSPACE_REQUIRED", "Set Workspace with an OAuth access token.")
	}
	if o.Workspace != "" && !publicID.MatchString(o.Workspace) {
		return nil, invalid("INVALID_WORKSPACE", "Use the workspace's ID.")
	}
	if o.Outage == "" {
		o.Outage = "throw"
	}
	if o.Outage != "throw" && o.Outage != "allow" && o.Outage != "deny" {
		return nil, invalid("OUTAGE_BEHAVIOR_REQUIRED", "Set Outage to throw, allow or deny.")
	}
	if o.BaseURL == "" {
		o.BaseURL = "https://api.abusezero.com"
	}
	u, e := url.Parse(o.BaseURL)
	if e != nil || u.Hostname() == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" || (u.Path != "" && u.Path != "/") || (u.Scheme != "https" && !(u.Scheme == "http" && (u.Hostname() == "localhost" || u.Hostname() == "127.0.0.1" || u.Hostname() == "::1"))) {
		return nil, invalid("INVALID_URL", "Use an HTTPS origin for BaseURL, or HTTP on loopback for local testing.")
	}
	credential := o.Key
	if credential == "" {
		credential = o.AccessToken
	}
	transport := o.Transport
	if transport == nil {
		transport = http.DefaultTransport.(*http.Transport).Clone()
	}
	agent := strings.Map(func(r rune) rune {
		if r < 32 || r > 126 {
			return -1
		}
		return r
	}, o.Agent)
	if len(agent) > 512 {
		agent = agent[:512]
	}
	return &core{credential: credential, workspace: o.Workspace, outage: o.Outage, agent: agent, base: strings.TrimRight(o.BaseURL, "/"), key: o.Key != "", public: public, http: &http.Client{Transport: transport, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}}, nil
}

// Close releases idle connections. Call it when your server shuts down.
func (c *core) Close() { c.http.CloseIdleConnections() }
func (c *core) call(ctx context.Context, id string, input any, options *CallOptions) (json.RawMessage, error) {
	op, exists := operations[id]
	if !exists || op.Public != c.public {
		return nil, invalid("INVALID_OPERATION", "Use the matching API client.")
	}
	if options == nil {
		options = &CallOptions{}
	}
	if c.key && op.Scope == "operations:management" {
		return nil, &Error{Code: "OAUTH_REQUIRED", Message: "This operation needs an OAuth access token.", Status: 403}
	}
	if (op.Mutation && options.IdempotencyKey == "") || (options.IdempotencyKey != "" && !opaque.MatchString(options.IdempotencyKey)) {
		return nil, invalid("INVALID_IDEMPOTENCY_KEY", "Pass an idempotency key of 1 to 128 printable characters with no spaces. Mutations require one.")
	}
	route := op.Path
	if len(options.Path) != len(op.Paths) {
		return nil, invalid("INVALID_PATH", "Pass the operation's path fields in CallOptions.Path.")
	}
	for _, p := range op.Paths {
		v, ok := options.Path[p.Name]
		if !ok || validate(v, p.Schema, p.Schema) != nil {
			return nil, invalid("INVALID_PATH", "Check the path fields in CallOptions.Path.")
		}
		encoded := url.PathEscape(fmt.Sprint(v))
		if p.Name == "path" {
			encoded = strings.ReplaceAll(encoded, "%2F", "/")
		}
		route = strings.ReplaceAll(route, "{"+p.Name+"}", encoded)
	}
	b, e := json.Marshal(input)
	if e != nil {
		return nil, invalid("INVALID_REQUEST", "The input can't be encoded.")
	}
	var data map[string]any
	if e = json.Unmarshal(b, &data); e != nil {
		return nil, invalid("INVALID_REQUEST", "Pass an input object.")
	}
	if data == nil {
		data = map[string]any{}
	}
	if r, ok := data["request"].(map[string]any); ok {
		headers := http.Header{}
		if hs, ok := r["headers"].(map[string]any); ok {
			for k, v := range hs {
				if s, ok := v.(string); ok {
					headers.Set(k, s)
				}
			}
		}
		req, err := incomingRequest(fmt.Sprint(r["method"]), fmt.Sprint(r["url"]), headers)
		if err != nil {
			return nil, err
		}
		data["request"] = req
	}
	if validate(data, op.Input, op.Input) != nil {
		return nil, invalid("INVALID_REQUEST", "The input doesn't match the API contract.")
	}
	b, e = json.Marshal(data)
	if e != nil {
		return nil, invalid("INVALID_REQUEST", "The input can't be encoded.")
	}
	if id == "evaluations.create" && len(b) > 49152 {
		return nil, invalid("INVALID_REQUEST", "The evaluation input is too large.")
	}
	if op.Method == "GET" {
		query := url.Values{}
		for k, v := range data {
			switch x := v.(type) {
			case string:
				query.Set(k, x)
			default:
				j, _ := json.Marshal(v)
				query.Set(k, string(j))
			}
		}
		if len(query) > 0 {
			route += "?" + query.Encode()
		}
		b = nil
	}
	budget := time.Duration(op.Timeout) * time.Millisecond
	if id == "events.list" {
		if wait, ok := data["wait"].(float64); ok && time.Duration(wait*1000+5000)*time.Millisecond > budget {
			budget = time.Duration(wait*1000+5000) * time.Millisecond
		}
	}
	ctx, cancel := context.WithTimeout(ctx, budget)
	defer cancel()
	var failure, rateError *Error
retryLoop:
	for attempt := 0; attempt < 3; attempt++ {
		req, e := http.NewRequestWithContext(ctx, op.Method, c.base+route, bytes.NewReader(b))
		if e != nil {
			return nil, invalid("INVALID_REQUEST", "The request can't be sent.")
		}
		req.Header.Set("Accept", "application/json")
		if !c.public {
			req.Header.Set("Authorization", "Bearer "+c.credential)
			if c.workspace != "" {
				req.Header.Set("AbuseZero-Workspace", c.workspace)
			}
			if c.agent != "" {
				req.Header.Set("AbuseZero-Agent", c.agent)
			}
		}
		if b != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		if options.IdempotencyKey != "" {
			req.Header.Set("Idempotency-Key", options.IdempotencyKey)
		}
		response, e := c.http.Do(req)
		delay := []time.Duration{100 * time.Millisecond, 250 * time.Millisecond, 0}[attempt]
		if e != nil {
			failure = &Error{Code: "TRANSPORT_ERROR", Message: "AbuseZero couldn't be reached, or the response didn't arrive complete.", Retryable: true}
		} else {
			raw, readErr := io.ReadAll(response.Body)
			response.Body.Close()
			if readErr != nil {
				failure = &Error{Code: "TRANSPORT_ERROR", Message: "AbuseZero couldn't be reached, or the response didn't arrive complete.", Retryable: true}
			} else {
				result, err := readResponse(op, response, raw)
				if err == nil {
					return result, nil
				}
				failure = err
			}
			if n, e := strconv.ParseFloat(response.Header.Get("Retry-After"), 64); e == nil && n >= 0 && n < 1e10 {
				if d := time.Duration(n * float64(time.Second)); d > delay {
					delay = d
				}
			} else if t, e := http.ParseTime(response.Header.Get("Retry-After")); e == nil && time.Until(t) > delay {
				delay = time.Until(t)
			}
		}
		if failure.Status == 429 {
			rateError = failure
		}
		if ctx.Err() != nil {
			if errors.Is(ctx.Err(), context.Canceled) {
				return nil, ctx.Err()
			}
			failure = &Error{Code: "DEADLINE_EXCEEDED", Message: "The request budget ran out. Retry with the same idempotency key.", Retryable: true}
			break
		}
		if !failure.Retryable || attempt == 2 {
			break
		}
		if deadline, ok := ctx.Deadline(); ok && time.Now().Add(delay).After(deadline) {
			break
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			if errors.Is(ctx.Err(), context.Canceled) {
				return nil, ctx.Err()
			}
			failure = &Error{Code: "DEADLINE_EXCEEDED", Message: "The request budget ran out. Retry with the same idempotency key.", Retryable: true}
			break retryLoop
		case <-timer.C:
		}
	}
	if rateError != nil {
		failure = rateError
	}
	if id == "evaluations.create" && c.outage != "throw" && failure.Retryable && failure.Status != 429 {
		return json.Marshal(map[string]any{"source": "local_outage", "claim_id": data["claim_id"], "result": c.outage, "error": failure})
	}
	return nil, failure
}
func readResponse(op operation, response *http.Response, raw []byte) (json.RawMessage, *Error) {
	var body map[string]any
	e := json.Unmarshal(raw, &body)
	status := response.StatusCode
	if status >= 200 && status < 300 {
		if e != nil || body == nil {
			return nil, &Error{Code: "INVALID_RESPONSE", Message: "AbuseZero's response couldn't be used.", Status: 502, Retryable: true}
		}
		if !(body["status"] == "skipped" && (op.ID == "evaluations.create" || op.ID == "evaluations.dry_run")) && validateResponse(body, op.Response) != nil {
			return nil, &Error{Code: "INVALID_RESPONSE", Message: "AbuseZero's response couldn't be used.", Status: 502, Retryable: true}
		}
		if op.ID == "evaluations.create" || op.ID == "evaluations.dry_run" {
			reason, _ := body["reason"].(string)
			if body["status"] == "skipped" && regexp.MustCompile(`^[a-z][a-z0-9_]{0,63}$`).MatchString(reason) {
				if op.ID == "evaluations.create" {
					body["source"] = "abusezero"
				}
				b, _ := json.Marshal(body)
				return b, nil
			}
			risk, ok := body["risk"].(map[string]any)
			valid := ok && (body["result"] == "allow" || body["result"] == "deny")
			if ok {
				_, band := risk["band"].(string)
				valid = valid && band
				if score := risk["score"]; score != nil {
					s, ok := score.(float64)
					valid = valid && ok && s >= 0 && s <= 100
				}
			}
			if op.ID == "evaluations.create" {
				id, ok := body["id"].(string)
				valid = valid && ok && publicID.MatchString(id) && body["object"] == "evaluation"
				reported, ok := body["outcome"].(map[string]any)
				if body["outcome"] != nil && !ok {
					valid = false
				}
				state, _ := reported["state"].(string)
				if state == "" {
					state = response.Header.Get("AbuseZero-Outcome-State")
				}
				if state != "" {
					if !regexp.MustCompile(`^[a-z][a-z0-9_]{0,63}$`).MatchString(state) {
						valid = false
					}
					due := reported["due_at"]
					if due == nil && response.Header.Get("AbuseZero-Outcome-Due-At") != "" {
						due = response.Header.Get("AbuseZero-Outcome-Due-At")
					}
					body["outcome"] = map[string]any{"state": state, "due_at": due}
				} else {
					body["outcome"] = nil
				}
				body["source"] = "abusezero"
			}
			if op.ID == "evaluations.dry_run" {
				valid = valid && body["object"] == "evaluation_dry_run" && body["dry_run"] == true
			}
			if !valid {
				return nil, &Error{Code: "INVALID_RESPONSE", Message: "AbuseZero's evaluation couldn't be used.", Status: 502, Retryable: true}
			}
			b, _ := json.Marshal(body)
			return b, nil
		}
		return raw, nil
	}
	detail, _ := body["error"].(map[string]any)
	code, _ := detail["code"].(string)
	if code == "" {
		code = "HTTP_ERROR"
	}
	message, _ := detail["message"].(string)
	if message == "" {
		message = "AbuseZero couldn't finish the request."
	}
	docs, _ := detail["docs_url"].(string)
	return nil, &Error{Code: code, Message: message, Status: status, RequestID: response.Header.Get("X-Request-ID"), Retryable: status != 401 && status != 403 && !(status >= 300 && status < 400) && (status == 429 || status == 502 || status == 503 || status == 504 || detail["retryable"] == true), DocsURL: docs}
}
func validate(v any, s, root map[string]any) error { return check(v, s, root, false, "") }

// Responses can gain new values for codes and enums within v1, so a response accepts an enum or fixed value it
// doesn't know, except an object's type, and an object of a newer kind than a union's variants. Requests stay strict.
func validateResponse(v any, s map[string]any) error { return check(v, s, s, true, "") }

// resolve follows a schema's references within root.
func resolve(s, root map[string]any) (map[string]any, bool) {
	for depth := 0; depth < 32; depth++ {
		ref, ok := s["$ref"].(string)
		if !ok {
			return s, true
		}
		var x any = root
		for _, part := range strings.Split(strings.TrimPrefix(ref, "#/"), "/") {
			m, ok := x.(map[string]any)
			if !ok {
				return nil, false
			}
			x = m[strings.ReplaceAll(strings.ReplaceAll(part, "~1", "/"), "~0", "~")]
		}
		if s, ok = x.(map[string]any); !ok {
			return nil, false
		}
	}
	return nil, false
}

// newKind reports whether v is an object of a kind no variant of the union describes: every variant fixes the same
// discriminator, such as a lookup match's kind, v's value for it is a string none of them fixes, and v's object
// type is the one they share.
func newKind(v any, variants []any, root map[string]any) bool {
	m, ok := v.(map[string]any)
	if !ok || len(variants) == 0 {
		return false
	}
	var shared map[string]map[string]bool
	for _, variant := range variants {
		vs, ok := variant.(map[string]any)
		if !ok {
			return false
		}
		if vs, ok = resolve(vs, root); !ok {
			return false
		}
		props, _ := vs["properties"].(map[string]any)
		fixed := map[string]string{}
		for name, p := range props {
			if ps, ok := p.(map[string]any); ok {
				if ps, ok = resolve(ps, root); ok {
					if c, ok := ps["const"].(string); ok {
						fixed[name] = c
					}
				}
			}
		}
		if shared == nil {
			shared = map[string]map[string]bool{}
			for name, c := range fixed {
				shared[name] = map[string]bool{c: true}
			}
			continue
		}
		for name := range shared {
			if c, ok := fixed[name]; ok {
				shared[name][c] = true
			} else {
				delete(shared, name)
			}
		}
	}
	if types, ok := shared["object"]; ok && len(types) == 1 {
		object, _ := m["object"].(string)
		if !types[object] {
			return false
		}
	}
	for name, values := range shared {
		if value, ok := m[name].(string); ok && name != "object" && !values[value] {
			return true
		}
	}
	return false
}

func check(v any, s, root map[string]any, response bool, field string) error {
	fail := func() error { return errors.New("invalid value") }
	if _, ok := s["$ref"].(string); ok {
		target, ok := resolve(s, root)
		if !ok {
			return fail()
		}
		return check(v, target, root, response, field)
	}
	for _, key := range []string{"anyOf", "oneOf"} {
		if variants, ok := s[key].([]any); ok {
			count := 0
			for _, variant := range variants {
				if check(v, variant.(map[string]any), root, response, field) == nil {
					count++
				}
			}
			if (count == 0 && !(response && newKind(v, variants, root))) || (key == "oneOf" && !response && count != 1) {
				return fail()
			}
		}
	}
	matches := func(t string) bool {
		switch t {
		case "null":
			return v == nil
		case "string":
			_, ok := v.(string)
			return ok
		case "object":
			_, ok := v.(map[string]any)
			return ok
		case "array":
			_, ok := v.([]any)
			return ok
		case "boolean":
			_, ok := v.(bool)
			return ok
		case "number", "integer":
			switch x := v.(type) {
			case float64:
				return t == "number" || math.Trunc(x) == x
			case int, int64:
				return true
			}
		}
		return true
	}
	if t, ok := s["type"].(string); ok && !matches(t) {
		return fail()
	}
	if a, ok := s["type"].([]any); ok {
		found := false
		for _, t := range a {
			if matches(t.(string)) {
				found = true
			}
		}
		if !found {
			return fail()
		}
	}
	if x, ok := s["const"]; ok && (!response || field == "object") && !reflect.DeepEqual(x, v) {
		return fail()
	}
	if xs, ok := s["enum"].([]any); ok && !response {
		found := false
		for _, x := range xs {
			if reflect.DeepEqual(x, v) {
				found = true
			}
		}
		if !found {
			return fail()
		}
	}
	if m, ok := v.(map[string]any); ok {
		props, _ := s["properties"].(map[string]any)
		required, _ := s["required"].([]any)
		for _, r := range required {
			if _, ok := m[r.(string)]; !ok {
				// The contract lists a request field that has a default as required; a request may leave it out.
				if p, ok := props[r.(string)].(map[string]any); ok && !response {
					if p, ok = resolve(p, root); ok {
						if _, has := p["default"]; has {
							continue
						}
					}
				}
				return fail()
			}
		}
		for k, x := range m {
			if pn, ok := s["propertyNames"].(map[string]any); ok {
				if check(k, pn, root, response, "") != nil {
					return fail()
				}
			}
			child, exists := props[k]
			if !exists {
				child = s["additionalProperties"]
				if child == false {
					return fail()
				}
			}
			if c, ok := child.(map[string]any); ok {
				if e := check(x, c, root, response, k); e != nil {
					return e
				}
			}
		}
	}
	if text, ok := v.(string); ok {
		n := float64(len([]rune(text)))
		if min, ok := s["minLength"].(float64); ok && n < min {
			return fail()
		}
		if max, ok := s["maxLength"].(float64); ok && n > max {
			return fail()
		}
		if p, ok := s["pattern"].(string); ok {
			re, e := regexp.Compile(p)
			if e != nil || !re.MatchString(text) {
				return fail()
			}
		}
	}
	if a, ok := v.([]any); ok {
		if min, ok := s["minItems"].(float64); ok && float64(len(a)) < min {
			return fail()
		}
		if max, ok := s["maxItems"].(float64); ok && float64(len(a)) > max {
			return fail()
		}
		if child, ok := s["items"].(map[string]any); ok {
			for _, x := range a {
				if e := check(x, child, root, response, ""); e != nil {
					return e
				}
			}
		}
	}
	if x, ok := v.(float64); ok {
		if multiple, ok := s["multipleOf"].(float64); ok && math.Abs(x/multiple-math.Round(x/multiple)) > 1e-9 {
			return fail()
		}
		for _, k := range []string{"minimum", "maximum", "exclusiveMinimum", "exclusiveMaximum"} {
			if b, ok := s[k].(float64); ok && ((k == "minimum" && x < b) || (k == "maximum" && x > b) || (k == "exclusiveMinimum" && x <= b) || (k == "exclusiveMaximum" && x >= b)) {
				return fail()
			}
		}
	}
	return nil
}

var credentialHeader = regexp.MustCompile(`(?i)^(authorization|proxy-authorization|cookie|set-cookie|x-api-key|x-auth-token|x-csrf-token|x-xsrf-token)$|(^|-)(token|secret|session|password|api-key|apikey)(-|$)`)

// IncomingRequest describes the request your server received, for an evaluation's Request field. It removes
// credential headers. It doesn't infer a trusted client IP; send that in the evaluation's Context.
func IncomingRequest(method, rawURL string, headers http.Header) (EvaluationsCreateInputRequest, error) {
	r, err := incomingRequest(method, rawURL, headers)
	if err != nil {
		return EvaluationsCreateInputRequest{}, err
	}
	return EvaluationsCreateInputRequest{Method: method, URL: rawURL, Headers: r["headers"].(map[string]any)}, nil
}

func incomingRequest(method, rawURL string, headers http.Header) (map[string]any, error) {
	u, e := url.Parse(rawURL)
	if e != nil || u.Hostname() == "" || (u.Scheme != "http" && u.Scheme != "https") || u.User != nil || len(rawURL) > 4096 || !regexp.MustCompile(`^[A-Za-z]{1,16}$`).MatchString(method) {
		return nil, invalid("INVALID_REQUEST", "Provide an incoming HTTP method and URL without credentials.")
	}
	clean := map[string]any{}
	size := 0
	for k, vs := range headers {
		k = strings.ToLower(k)
		if credentialHeader.MatchString(k) || len(vs) == 0 || !regexp.MustCompile("^[!#$%&'*+.^_`|~0-9a-z-]{1,128}$").MatchString(k) {
			continue
		}
		v := strings.Join(vs, ", ")
		if strings.ContainsAny(v, "\r\n") {
			continue
		}
		if len(v) > 8192 {
			v = v[:8192]
		}
		if len(clean) >= 64 || size+len(k)+len(v) > 32768 {
			continue
		}
		clean[k] = v
		size += len(k) + len(v)
	}
	return map[string]any{"method": method, "url": rawURL, "headers": clean}, nil
}
func keyedHash(secret []byte, keyVersion, value string) (string, error) {
	if !opaque.MatchString(keyVersion) {
		return "", invalid("INVALID_KEY_VERSION", "Use a printable key version with no spaces.")
	}
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil)), nil
}

// CreatePhoneHash returns the keyed hash of an E.164 phone number, for an evaluation's Identifiers.PhoneHash.
func CreatePhoneHash(secret []byte, keyVersion, e164 string) (EvaluationsCreateInputIdentifiersPhoneHash, error) {
	if !regexp.MustCompile(`^\+[1-9][0-9]{7,14}$`).MatchString(e164) {
		return EvaluationsCreateInputIdentifiersPhoneHash{}, invalid("INVALID_PHONE", "Normalize the phone number to E.164 before hashing it.")
	}
	value, err := keyedHash(secret, keyVersion, "abusezero:phone:e164:v1:"+e164)
	if err != nil {
		return EvaluationsCreateInputIdentifiersPhoneHash{}, err
	}
	return EvaluationsCreateInputIdentifiersPhoneHash{Algorithm: "hmac-sha256", KeyVersion: keyVersion, Value: value}, nil
}

// CreateCustomIdentifier returns the keyed hash of a registered custom identifier's value, as an item for an
// evaluation's Identifiers.Custom.
func CreateCustomIdentifier(secret []byte, keyVersion, key, value string) (EvaluationsCreateInputIdentifiersCustomItem, error) {
	if !regexp.MustCompile(`^[a-z][a-z0-9_]{0,63}$`).MatchString(key) || value == "" {
		return EvaluationsCreateInputIdentifiersCustomItem{}, invalid("INVALID_IDENTIFIER", "Use a lowercase identifier key that starts with a letter, and a non-empty value.")
	}
	hash, err := keyedHash(secret, keyVersion, "abusezero:custom:"+key+":v1:"+value)
	if err != nil {
		return EvaluationsCreateInputIdentifiersCustomItem{}, err
	}
	return EvaluationsCreateInputIdentifiersCustomItem{Key: key, ValueHash: EvaluationsCreateInputIdentifiersCustomItemValueHash{Algorithm: "hmac-sha256", KeyVersion: keyVersion, Value: hash}}, nil
}

// VerifyWebhook checks a delivery's signature over the original request bytes against the current and previous
// signing secrets and its timestamp against tolerance, then returns the event. It returns a plain error.
func VerifyWebhook(payload []byte, headers http.Header, secrets []string, tolerance time.Duration, now time.Time) (map[string]any, error) {
	if tolerance < 0 {
		return nil, errors.New("use a nonnegative timestamp tolerance")
	}
	keys := [][]byte{}
	for _, s := range secrets {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if !regexp.MustCompile(`^whsk_[A-Za-z0-9+/]{43}=$`).MatchString(s) {
			return nil, errors.New("provide the complete webhook signing secret")
		}
		k, e := base64.StdEncoding.DecodeString(s[5:])
		if e != nil || len(k) != 32 {
			return nil, errors.New("invalid webhook signing secret")
		}
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return nil, errors.New("provide a webhook signing secret")
	}
	id, stamp := headers.Get("AbuseZero-Webhook-Id"), headers.Get("AbuseZero-Webhook-Timestamp")
	seconds, e := strconv.ParseInt(stamp, 10, 64)
	if e != nil || id == "" || !regexp.MustCompile(`^[0-9]{1,12}$`).MatchString(stamp) || math.Abs(float64(now.Unix()-seconds)) > tolerance.Seconds() {
		return nil, errors.New("invalid webhook timestamp or headers")
	}
	signed := append([]byte(id+"."+stamp+"."), payload...)
	matched := false
	for _, k := range keys {
		h := hmac.New(sha256.New, k)
		h.Write(signed)
		for _, offered := range strings.Fields(headers.Get("AbuseZero-Webhook-Signature")) {
			if !strings.HasPrefix(offered, "v1,") {
				continue
			}
			sig, e := base64.StdEncoding.DecodeString(offered[3:])
			if e == nil && subtle.ConstantTimeCompare(sig, h.Sum(nil)) == 1 {
				matched = true
			}
		}
	}
	if !matched {
		return nil, errors.New("the webhook signature doesn't match")
	}
	var event map[string]any
	if json.Unmarshal(payload, &event) != nil || event["object"] != "event" || event["id"] != id {
		return nil, errors.New("the webhook body is invalid")
	}
	return event, nil
}
