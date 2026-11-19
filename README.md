# AbuseZero Go SDK

**Abuse prevention API for SaaS**

Prevent users from creating extra accounts to repeatedly claim your trials, free plans, credits, and other benefits.

The AbuseZero Go SDK is the Go client for the [AbuseZero REST API](https://docs.abusezero.com/docs/api-reference). It has a method for every public server operation, retries transient failures within a fixed request budget, and lets you choose in advance what happens if AbuseZero can't be reached.

## Install the Go SDK

In your Go module, run:

```bash
go get github.com/abusezero/abusezero-go
```

Import it as `abusezero`:

```go
import abusezero "github.com/abusezero/abusezero-go"
```

The SDK supports Go 1.26 and 1.27 and uses only the Go standard library. It's proprietary: its license is in [`LICENSE`](LICENSE), and [`THIRD-PARTY-NOTICES`](THIRD-PARTY-NOTICES) confirms it bundles no third-party code.

## Create a Go SDK client

Store your [secret key or restricted key](https://docs.abusezero.com/docs/keys-and-access#secret-and-restricted-keys) in `ABUSEZERO_SECRET_KEY` on your server. Create one client when your server starts, reuse it for every request, and call `Close` at shutdown.

```go
func NewClient() (*abusezero.Client, error) {
	return abusezero.New(abusezero.Options{Key: os.Getenv("ABUSEZERO_SECRET_KEY"), Outage: "throw"})
}
```

### Client options

| Field | Type | Required | Description |
|---|---|---|---|
| `Key` | `string` | One of `Key` or `AccessToken` | A secret key (`sk_…`) or restricted key (`rk_…`). [Publishable keys](https://docs.abusezero.com/docs/keys-and-access#publishable-key) are rejected. |
| `AccessToken` | `string` | One of `Key` or `AccessToken` | An [OAuth](https://docs.abusezero.com/docs/keys-and-access#oauth) access token. Requires `Workspace`. |
| `Workspace` | `string` | With `AccessToken` | The ID of the [workspace](https://docs.abusezero.com/docs/api-reference/workspaces) this OAuth connection works in. |
| `Outage` | `"throw"`, `"allow"` or `"deny"` | No | What `Evaluate` returns when AbuseZero can't be reached. The default is `"throw"`. See [Outages](#outages). |
| `BaseURL` | `string` | No | The API origin. The default is `https://api.abusezero.com`. A custom origin must use HTTPS; HTTP is accepted only on loopback, for local testing. |
| `Transport` | `http.RoundTripper` | No | Sends every request through your own transport, for example through a proxy. It must honor context cancellation and verify TLS. |
| `Agent` | `string` | No | The [`AbuseZero-Agent`](https://docs.abusezero.com/docs/api-reference/authentication#identify-your-agent) value your agent reports on each request. |

`New` returns `CREDENTIAL_REQUIRED` for neither or both of `Key` and `AccessToken`, `INVALID_KEY` for a key that isn't a secret or restricted key, and `WORKSPACE_REQUIRED` for an access token without a workspace. See [Go SDK errors](#go-sdk-errors).

## Evaluate a claim and report the outcome

Use a stable claim ID, and keep the same [idempotency key](https://docs.abusezero.com/docs/api-reference#idempotency) when you retry the same action.

```go
package examples

import (
	"context"
	"errors"
	"os"

	abusezero "github.com/abusezero/abusezero-go"
)

func NewClient() (*abusezero.Client, error) {
	return abusezero.New(abusezero.Options{Key: os.Getenv("ABUSEZERO_SECRET_KEY"), Outage: "throw"})
}

func EvaluateAndGrant(ctx context.Context, client *abusezero.Client, claimID, email string,
	grant func(context.Context, string) error) (any, error) {
	if grant == nil {
		return nil, errors.New("provide a grant function")
	}
	result, err := client.Evaluate(ctx, abusezero.EvaluationsCreateInput{
		Benefit: "pro_trial", ClaimID: claimID,
		Subject: abusezero.EvaluationsCreateInputSubject{Email: email},
	}, &abusezero.CallOptions{IdempotencyKey: "evaluate:" + claimID})
	if err != nil {
		return nil, err
	}
	evaluation := result.(map[string]any)
	if evaluation["status"] == "skipped" || evaluation["source"] == "local_outage" || evaluation["result"] != "allow" {
		return result, nil
	}
	// Your transaction must finish successfully before returning.
	if err := grant(ctx, claimID); err != nil {
		return result, err
	}
	_, err = client.Outcome(ctx, abusezero.OutcomesReportInput{Outcome: "granted"},
		&abusezero.CallOptions{Path: map[string]any{"id": evaluation["id"]}, IdempotencyKey: "grant:" + claimID})
	return result, err
}
```

Call this helper from your authenticated handler with your own stored claim ID, the authenticated user's email and a transactional grant function. The grant function must be idempotent for that claim ID. The helper [reports the outcome](https://docs.abusezero.com/docs/concepts/outcomes) `granted` only after the grant function succeeds. If reporting fails afterward, retry with the same claim ID and idempotency keys. On `deny`, don't grant the benefit.

### Results

[`Evaluate`](https://docs.abusezero.com/docs/api-reference/evaluations#create-an-evaluation) returns a `map[string]any` in one of three shapes. Check `status` first, then `source`.

| Shape | Fields | Meaning |
|---|---|---|
| Evaluated | `source: "abusezero"`, `id`, `result`, `risk`, `reason`, `outcome` | AbuseZero evaluated the claim. `result` is `allow` or `deny`. |
| Skipped | `source: "abusezero"`, `status: "skipped"`, `reason` | The workspace couldn't run the evaluation, so there's no `result`. |
| Local outage | `source: "local_outage"`, `claim_id`, `result`, `error` | AbuseZero couldn't be reached. `result` is your `Outage` setting, and `error` holds the code, message and status behind it. |

A skipped evaluation is neither allow nor deny. Handle it as its own branch, even when its [skip reason](https://docs.abusezero.com/docs/api-reference/evaluations#skipped-evaluations) is new to your application. Your `Outage` setting never applies to a skip. `risk`, `reason` and `outcome` are described in [Evaluations](https://docs.abusezero.com/docs/concepts/evaluations).

## Outages

Set `Outage` to `throw`, `allow` or `deny`. With `throw`, `Evaluate` returns the error. With `allow` or `deny`, it returns a local outage result with that `result` when the call fails with a retryable error other than a rate limit. Only `Evaluate` returns local outage results. `DryRunEvaluation` and every other method return the error, and rate limits always remain errors.

## Retries

Calls make at most three attempts. Network errors, retryable API errors and HTTP 429, 502, 503 and 504 may be retried, and `Retry-After` sets the minimum wait. Every attempt sends the same serialized input and the same idempotency key.

The whole call, waits included, has one request budget: 1.8 seconds for `Evaluate` and 5 seconds for every other method, `DryRunEvaluation` included. An `Events` long poll adds 5 seconds to its requested wait. AbuseZero's own [time limits](https://docs.abusezero.com/docs/api-reference#time-limits) can be shorter. When the budget runs out, the call returns `DEADLINE_EXCEEDED`.

Every method takes a `context.Context`. Cancel it to stop an in-flight call; the method then returns `context.Canceled`.

## Methods and pagination

Every public server operation has a method, and [`API.md`](API.md) lists them all. Pass the operation's body or query fields in the input value, and its path fields in `CallOptions.Path`. Mutations require `CallOptions.IdempotencyKey`.

List methods accept `Cursor` and `Limit` where the operation has them. Process each page's `Data`, then pass its `NextCursor` to the same method until it's null. Keep [IDs and cursors](https://docs.abusezero.com/docs/api-reference#pagination) as strings.

```go
func EachEvaluation(ctx context.Context, client *abusezero.Client, handle func(abusezero.EvaluationsListResponseDataItem)) error {
	var cursor abusezero.Optional[string]
	for {
		page, err := client.Evaluations(ctx, abusezero.EvaluationsListInput{Cursor: cursor}, nil)
		if err != nil {
			return err
		}
		for _, item := range page.Data {
			handle(item)
		}
		next, ok := page.NextCursor.(string)
		if !ok {
			return nil
		}
		cursor = abusezero.Some(next)
	}
}
```

## Types

`Optional[T]` tells omission apart from null. Use `Some(value)` to send a value and `Null[T]()` to send null where the schema allows it. Omitting a field leaves it out of the request.

Response structs type the fields whose shape is fixed. Every other field, including some documented ones such as an outcome report's `state`, stays in `Extra` as raw JSON. Decode one with `encoding/json`, for example `json.Unmarshal(report.Extra["state"], &state)`. `Evaluate` and `DryRunEvaluation` return `map[string]any`.

Within `v1`, responses can gain new fields and new values for codes and enums. The SDK returns them unchanged, so handle a value your code doesn't recognize as a generic case.

## Credentials and public reads

Use `AccessToken` and `Workspace` for OAuth. Management operations need an OAuth access token; an API key alone can't run them. The selected workspace and the approving member's permissions still apply.

`NewPublic` creates a client for the public operations, such as pricing. It sends no credentials.

```go
func ReadPricing(ctx context.Context) (abusezero.PublicPricingResponse, error) {
	public, err := abusezero.NewPublic(abusezero.Options{})
	if err != nil {
		return abusezero.PublicPricingResponse{}, err
	}
	defer public.Close()
	return public.Pricing(ctx, map[string]any{}, nil)
}
```

## Verify webhooks

Verify the [signature](https://docs.abusezero.com/docs/api-reference/webhook-events#signatures) over the original request bytes before you read the event.

```go
// Pass the raw request bytes and the original request headers.
func VerifyEvent(rawBody []byte, header http.Header, currentSecret, previousSecret string) (map[string]any, error) {
	return abusezero.VerifyWebhook(rawBody, header,
		[]string{currentSecret, previousSecret}, 300*time.Second, time.Now())
}
```

Pass the timestamp tolerance explicitly, as shown with `300*time.Second`. Verification accepts the current and previous signing secrets and rejects a modified body. Deduplicate accepted events by their event ID in your application.

## Incoming requests and identifiers

`IncomingRequest` takes the HTTP method, absolute URL and headers of the request your server received. It removes credential headers and returns the value for the evaluation's `Request`. It doesn't infer a trusted client IP. Send the IP your trusted proxy configuration reports in the evaluation's `Context`, as `abusezero.Some[any](map[string]any{"ip": ip})`.

`CreatePhoneHash` requires an E.164 phone number. `CreateCustomIdentifier` requires a [custom identifier](https://docs.abusezero.com/docs/concepts/custom-identifiers)'s registered key and its value. Keep the hash secret and `key_version` the same across applications that should recognize the same identifier. `CreatePhoneHash` returns the value for `Identifiers.PhoneHash`; send `Identifiers.PhoneVerified` with it. `CreateCustomIdentifier` returns an item for `Identifiers.Custom`.

The helpers return the types `Evaluate` takes. For `DryRunEvaluation`, convert the request and phone hash, such as `abusezero.EvaluationsDryRunInputRequest(request)`, and build each custom identifier from its `Key` and its converted `ValueHash`.

## Browser protection

Load the [standalone browser script](https://docs.abusezero.com/docs/tools/browser-sdk#standalone-browser-script) on your page and pass its device token to your server. Send the token unchanged as `Identifiers.DeviceToken` in the evaluation input. The Go SDK is the only package to install. [`examples/frameworks.md`](examples/frameworks.md) shows a Go HTTP handler that sends the incoming request, the client IP and the device token.

## Transport and TLS

The client never follows redirects. The default transport verifies TLS, and a custom one must too. Never log credentials or raw authenticated requests.

## Go SDK errors

Failed calls return an error. Use `errors.As` with `*abusezero.Error` to read its fields.

| Field | Description |
|---|---|
| `Code` | The error code, such as `CURSOR_INVALID`. |
| `Message` | What to do next, written for a developer. |
| `Status` | The HTTP status. Errors the SDK raises before sending use `422`, or `403` for `OAUTH_REQUIRED`. `INVALID_RESPONSE` uses `502`, and `0` means no response arrived. |
| `RequestID` | The [`X-Request-ID`](https://docs.abusezero.com/docs/api-reference#request-ids) of the failed request, when there is one. Include it when you contact support. |
| `Retryable` | `true` when the same call, with the same idempotency key, can succeed. |
| `DocsURL` | The [error code reference](https://docs.abusezero.com/docs/api-reference/errors)'s section for this error, when AbuseZero sent one. |

```go
func CreateBenefit(ctx context.Context, client *abusezero.Client, input abusezero.BenefitsCreateInput) error {
	_, err := client.CreateBenefit(ctx, input, &abusezero.CallOptions{IdempotencyKey: "benefit:" + input.Key})
	var apiError *abusezero.Error
	if errors.As(err, &apiError) && apiError.Retryable {
		// Retry later with the same idempotency key.
	}
	return err
}
```

A canceled context returns `context.Canceled`, and `VerifyWebhook` returns a plain error. REST API error codes are listed in [REST API errors](https://docs.abusezero.com/docs/api-reference/errors), and how to fix the common ones in [Troubleshooting](https://docs.abusezero.com/docs/troubleshooting). The SDK adds these, raised before a request is sent or when a response can't be used:

| Code | Meaning |
|---|---|
| `CREDENTIAL_REQUIRED` | Pass exactly one of `Key` and `AccessToken`. |
| `INVALID_KEY` | Use a secret key (`sk_…`) or restricted key (`rk_…`). |
| `INVALID_TOKEN` | Pass the OAuth access token exactly as issued. |
| `WORKSPACE_REQUIRED` | Set `Workspace` with an OAuth access token. |
| `INVALID_WORKSPACE` | Use the workspace's ID. |
| `OUTAGE_BEHAVIOR_REQUIRED` | Set `Outage` to `"throw"`, `"allow"` or `"deny"`. |
| `INVALID_URL` | Use an HTTPS origin for `BaseURL`, or HTTP on loopback for local testing. |
| `OAUTH_REQUIRED` | This operation needs an OAuth access token. |
| `INVALID_IDEMPOTENCY_KEY` | Pass an idempotency key of 1 to 128 printable characters with no spaces. Mutations require one. |
| `INVALID_PATH` | Pass the operation's path fields in `CallOptions.Path`. |
| `INVALID_REQUEST` | The input doesn't match the API contract. |
| `INVALID_PHONE` | Normalize the phone number to E.164 before hashing it. |
| `INVALID_IDENTIFIER` | Use a lowercase identifier key that starts with a letter, and a non-empty value. |
| `INVALID_KEY_VERSION` | Use a printable key version with no spaces. |
| `DEADLINE_EXCEEDED` | The request budget ran out. Retry with the same idempotency key. |
| `TRANSPORT_ERROR` | AbuseZero couldn't be reached, or the response didn't arrive complete. Retry with the same idempotency key. |
| `INVALID_RESPONSE` | AbuseZero's response couldn't be used. Retry with the same idempotency key before acting on the claim. |
| `HTTP_ERROR` | AbuseZero answered with an HTTP error or a redirect that has no error code. Check `Status`. |
