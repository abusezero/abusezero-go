package abusezero

import (
	"context"
	"encoding/json"
)

type PublicPricingResponsePlansItemPricesMonthVariant0 struct {
	AmountCents int64                      `json:"amount_cents"`
	Credits     int64                      `json:"credits"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingResponsePlansItemPricesMonthVariant0) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponsePlansItemPricesMonthVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponsePlansItemPricesMonthVariant0) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponsePlansItemPricesMonthVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"amount_cents", "credits"})
}

type PublicPricingResponsePlansItemPricesYearVariant0 struct {
	AmountCents int64                      `json:"amount_cents"`
	Credits     int64                      `json:"credits"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingResponsePlansItemPricesYearVariant0) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponsePlansItemPricesYearVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponsePlansItemPricesYearVariant0) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponsePlansItemPricesYearVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"amount_cents", "credits"})
}

type PublicPricingResponsePlansItemPrices struct {
	Month any                        `json:"month"`
	Year  any                        `json:"year"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingResponsePlansItemPrices) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponsePlansItemPrices
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponsePlansItemPrices) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponsePlansItemPrices
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"month", "year"})
}

type PublicPricingResponsePlansItemIncludesTraffic struct {
	RatePerSecond int64                      `json:"rate_per_second"`
	Burst         int64                      `json:"burst"`
	Concurrency   int64                      `json:"concurrency"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingResponsePlansItemIncludesTraffic) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponsePlansItemIncludesTraffic
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponsePlansItemIncludesTraffic) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponsePlansItemIncludesTraffic
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"rate_per_second", "burst", "concurrency"})
}

type PublicPricingResponsePlansItemIncludes struct {
	ProtectedBenefits int64                                         `json:"protected_benefits"`
	TeamSeats         int64                                         `json:"team_seats"`
	DataRetentionDays int64                                         `json:"data_retention_days"`
	Traffic           PublicPricingResponsePlansItemIncludesTraffic `json:"traffic"`
	Extra             map[string]json.RawMessage                    `json:"-"`
}

func (v PublicPricingResponsePlansItemIncludes) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponsePlansItemIncludes
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponsePlansItemIncludes) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponsePlansItemIncludes
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"protected_benefits", "team_seats", "data_retention_days", "traffic"})
}

type PublicPricingResponsePlansItem struct {
	Object   string                                 `json:"object"`
	Plan     string                                 `json:"plan"`
	Name     string                                 `json:"name"`
	Prices   PublicPricingResponsePlansItemPrices   `json:"prices"`
	Includes PublicPricingResponsePlansItemIncludes `json:"includes"`
	Extra    map[string]json.RawMessage             `json:"-"`
}

func (v PublicPricingResponsePlansItem) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponsePlansItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponsePlansItem) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponsePlansItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "plan", "name", "prices", "includes"})
}

type PublicPricingResponseCreditPacks struct {
	CreditsPerDollar        int64                      `json:"credits_per_dollar"`
	Amounts                 []int64                    `json:"amounts"`
	CustomAmountStep        int64                      `json:"custom_amount_step"`
	CustomAmountMax         int64                      `json:"custom_amount_max"`
	FirstPurchaseMultiplier int64                      `json:"first_purchase_multiplier"`
	Extra                   map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingResponseCreditPacks) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponseCreditPacks
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponseCreditPacks) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponseCreditPacks
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"credits_per_dollar", "amounts", "custom_amount_step", "custom_amount_max", "first_purchase_multiplier"})
}

type PublicPricingResponseYearly struct {
	RolloverPercent             int64                      `json:"rollover_percent"`
	ConvertedCreditExpiryMonths int64                      `json:"converted_credit_expiry_months"`
	Extra                       map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingResponseYearly) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponseYearly
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponseYearly) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponseYearly
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"rollover_percent", "converted_credit_expiry_months"})
}

type PublicPricingResponseEveryPlan struct {
	CustomPolicies                   int64                      `json:"custom_policies"`
	WebhookEndpoints                 int64                      `json:"webhook_endpoints"`
	WebhookResendsPerEndpointPerHour int64                      `json:"webhook_resends_per_endpoint_per_hour"`
	WebhookDeliveriesPerDay          int64                      `json:"webhook_deliveries_per_day"`
	Extra                            map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingResponseEveryPlan) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponseEveryPlan
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponseEveryPlan) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponseEveryPlan
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"custom_policies", "webhook_endpoints", "webhook_resends_per_endpoint_per_hour", "webhook_deliveries_per_day"})
}

type PublicPricingResponseEnterprise struct {
	Contact string                     `json:"contact"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingResponseEnterprise) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponseEnterprise
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponseEnterprise) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponseEnterprise
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"contact"})
}

type PublicPricingResponse struct {
	Object               string                           `json:"object"`
	Currency             string                           `json:"currency"`
	CreditsPerEvaluation int64                            `json:"credits_per_evaluation"`
	Plans                []PublicPricingResponsePlansItem `json:"plans"`
	CreditPacks          PublicPricingResponseCreditPacks `json:"credit_packs"`
	Yearly               PublicPricingResponseYearly      `json:"yearly"`
	EveryPlan            PublicPricingResponseEveryPlan   `json:"every_plan"`
	Enterprise           PublicPricingResponseEnterprise  `json:"enterprise"`
	URL                  string                           `json:"url"`
	Extra                map[string]json.RawMessage       `json:"-"`
}

func (v PublicPricingResponse) MarshalJSON() ([]byte, error) {
	type plain PublicPricingResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingResponse) UnmarshalJSON(b []byte) error {
	type plain PublicPricingResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "currency", "credits_per_evaluation", "plans", "credit_packs", "yearly", "every_plan", "enterprise", "url"})
}

type PublicPricingQuoteInput struct {
	Plan     string                     `json:"plan"`
	Interval Optional[string]           `json:"interval,omitzero"`
	Amount   Optional[int64]            `json:"amount,omitzero"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingQuoteInput) MarshalJSON() ([]byte, error) {
	type plain PublicPricingQuoteInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingQuoteInput) UnmarshalJSON(b []byte) error {
	type plain PublicPricingQuoteInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"plan", "interval", "amount"})
}

type PublicPricingQuoteResponse struct {
	Object       string                     `json:"object"`
	Plan         string                     `json:"plan"`
	Interval     any                        `json:"interval"`
	Currency     string                     `json:"currency"`
	AmountCents  int64                      `json:"amount_cents"`
	Credits      int64                      `json:"credits"`
	BonusCredits int64                      `json:"bonus_credits"`
	TotalCredits int64                      `json:"total_credits"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v PublicPricingQuoteResponse) MarshalJSON() ([]byte, error) {
	type plain PublicPricingQuoteResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicPricingQuoteResponse) UnmarshalJSON(b []byte) error {
	type plain PublicPricingQuoteResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "plan", "interval", "currency", "amount_cents", "credits", "bonus_credits", "total_credits"})
}

type PublicDocsInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PublicDocsInput) MarshalJSON() ([]byte, error) {
	type plain PublicDocsInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicDocsInput) UnmarshalJSON(b []byte) error {
	type plain PublicDocsInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type PublicDocsResponseDataItem struct {
	Object      string                     `json:"object"`
	Path        string                     `json:"path"`
	Title       string                     `json:"title"`
	Description any                        `json:"description"`
	URL         string                     `json:"url"`
	UpdatedAt   string                     `json:"updated_at"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicDocsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicDocsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicDocsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicDocsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "path", "title", "description", "url", "updated_at"})
}

type PublicDocsResponse struct {
	Object     string                       `json:"object"`
	Data       []PublicDocsResponseDataItem `json:"data"`
	NextCursor any                          `json:"next_cursor"`
	Extra      map[string]json.RawMessage   `json:"-"`
}

func (v PublicDocsResponse) MarshalJSON() ([]byte, error) {
	type plain PublicDocsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicDocsResponse) UnmarshalJSON(b []byte) error {
	type plain PublicDocsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicDocsSearchInput struct {
	Query string                     `json:"query"`
	Limit Optional[int64]            `json:"limit,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v PublicDocsSearchInput) MarshalJSON() ([]byte, error) {
	type plain PublicDocsSearchInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicDocsSearchInput) UnmarshalJSON(b []byte) error {
	type plain PublicDocsSearchInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"query", "limit"})
}

type PublicDocsSearchResponseDataItem struct {
	Object    string                     `json:"object"`
	Path      string                     `json:"path"`
	Title     string                     `json:"title"`
	URL       string                     `json:"url"`
	Snippet   string                     `json:"snippet"`
	UpdatedAt string                     `json:"updated_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v PublicDocsSearchResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicDocsSearchResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicDocsSearchResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicDocsSearchResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "path", "title", "url", "snippet", "updated_at"})
}

type PublicDocsSearchResponse struct {
	Object     string                             `json:"object"`
	Data       []PublicDocsSearchResponseDataItem `json:"data"`
	NextCursor any                                `json:"next_cursor"`
	Extra      map[string]json.RawMessage         `json:"-"`
}

func (v PublicDocsSearchResponse) MarshalJSON() ([]byte, error) {
	type plain PublicDocsSearchResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicDocsSearchResponse) UnmarshalJSON(b []byte) error {
	type plain PublicDocsSearchResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicDocsPageResponse struct {
	Object      string                     `json:"object"`
	Path        string                     `json:"path"`
	Title       string                     `json:"title"`
	Description any                        `json:"description"`
	URL         string                     `json:"url"`
	UpdatedAt   string                     `json:"updated_at"`
	Markdown    string                     `json:"markdown"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicDocsPageResponse) MarshalJSON() ([]byte, error) {
	type plain PublicDocsPageResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicDocsPageResponse) UnmarshalJSON(b []byte) error {
	type plain PublicDocsPageResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "path", "title", "description", "url", "updated_at", "markdown"})
}

type PublicHelpInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PublicHelpInput) MarshalJSON() ([]byte, error) {
	type plain PublicHelpInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicHelpInput) UnmarshalJSON(b []byte) error {
	type plain PublicHelpInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type PublicHelpResponseDataItem struct {
	Object      string                     `json:"object"`
	Path        string                     `json:"path"`
	Title       string                     `json:"title"`
	Description any                        `json:"description"`
	URL         string                     `json:"url"`
	UpdatedAt   string                     `json:"updated_at"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicHelpResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicHelpResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicHelpResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicHelpResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "path", "title", "description", "url", "updated_at"})
}

type PublicHelpResponse struct {
	Object     string                       `json:"object"`
	Data       []PublicHelpResponseDataItem `json:"data"`
	NextCursor any                          `json:"next_cursor"`
	Extra      map[string]json.RawMessage   `json:"-"`
}

func (v PublicHelpResponse) MarshalJSON() ([]byte, error) {
	type plain PublicHelpResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicHelpResponse) UnmarshalJSON(b []byte) error {
	type plain PublicHelpResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicHelpSearchInput struct {
	Query string                     `json:"query"`
	Limit Optional[int64]            `json:"limit,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v PublicHelpSearchInput) MarshalJSON() ([]byte, error) {
	type plain PublicHelpSearchInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicHelpSearchInput) UnmarshalJSON(b []byte) error {
	type plain PublicHelpSearchInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"query", "limit"})
}

type PublicHelpSearchResponseDataItem struct {
	Object    string                     `json:"object"`
	Path      string                     `json:"path"`
	Title     string                     `json:"title"`
	URL       string                     `json:"url"`
	Snippet   string                     `json:"snippet"`
	UpdatedAt string                     `json:"updated_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v PublicHelpSearchResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicHelpSearchResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicHelpSearchResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicHelpSearchResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "path", "title", "url", "snippet", "updated_at"})
}

type PublicHelpSearchResponse struct {
	Object     string                             `json:"object"`
	Data       []PublicHelpSearchResponseDataItem `json:"data"`
	NextCursor any                                `json:"next_cursor"`
	Extra      map[string]json.RawMessage         `json:"-"`
}

func (v PublicHelpSearchResponse) MarshalJSON() ([]byte, error) {
	type plain PublicHelpSearchResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicHelpSearchResponse) UnmarshalJSON(b []byte) error {
	type plain PublicHelpSearchResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicHelpArticleResponse struct {
	Object      string                     `json:"object"`
	Path        string                     `json:"path"`
	Title       string                     `json:"title"`
	Description any                        `json:"description"`
	URL         string                     `json:"url"`
	UpdatedAt   string                     `json:"updated_at"`
	Markdown    string                     `json:"markdown"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicHelpArticleResponse) MarshalJSON() ([]byte, error) {
	type plain PublicHelpArticleResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicHelpArticleResponse) UnmarshalJSON(b []byte) error {
	type plain PublicHelpArticleResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "path", "title", "description", "url", "updated_at", "markdown"})
}

type PublicOpenAPIResponse struct {
	Object   string                     `json:"object"`
	Version  string                     `json:"version"`
	Document map[string]any             `json:"document"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v PublicOpenAPIResponse) MarshalJSON() ([]byte, error) {
	type plain PublicOpenAPIResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicOpenAPIResponse) UnmarshalJSON(b []byte) error {
	type plain PublicOpenAPIResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "version", "document"})
}

type PublicActionsResponseDataItem struct {
	Object      string                     `json:"object"`
	ID          string                     `json:"id"`
	Summary     string                     `json:"summary"`
	Method      string                     `json:"method"`
	Path        string                     `json:"path"`
	Scope       string                     `json:"scope"`
	SignIn      bool                       `json:"sign_in"`
	MCPTool     any                        `json:"mcp_tool"`
	CLICommand  any                        `json:"cli_command"`
	Mutation    bool                       `json:"mutation"`
	Destructive bool                       `json:"destructive"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicActionsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicActionsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicActionsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicActionsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "summary", "method", "path", "scope", "sign_in", "mcp_tool", "cli_command", "mutation", "destructive"})
}

type PublicActionsResponse struct {
	Object     string                          `json:"object"`
	Data       []PublicActionsResponseDataItem `json:"data"`
	NextCursor any                             `json:"next_cursor"`
	Extra      map[string]json.RawMessage      `json:"-"`
}

func (v PublicActionsResponse) MarshalJSON() ([]byte, error) {
	type plain PublicActionsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicActionsResponse) UnmarshalJSON(b []byte) error {
	type plain PublicActionsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicErrorsResponseDataItem struct {
	Object    string                     `json:"object"`
	Code      string                     `json:"code"`
	Status    int64                      `json:"status"`
	Type      string                     `json:"type"`
	Meaning   string                     `json:"meaning"`
	Retryable bool                       `json:"retryable"`
	NextStep  string                     `json:"next_step"`
	DocsURL   string                     `json:"docs_url"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v PublicErrorsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicErrorsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicErrorsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicErrorsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "code", "status", "type", "meaning", "retryable", "next_step", "docs_url"})
}

type PublicErrorsResponse struct {
	Object     string                         `json:"object"`
	Data       []PublicErrorsResponseDataItem `json:"data"`
	NextCursor any                            `json:"next_cursor"`
	Extra      map[string]json.RawMessage     `json:"-"`
}

func (v PublicErrorsResponse) MarshalJSON() ([]byte, error) {
	type plain PublicErrorsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicErrorsResponse) UnmarshalJSON(b []byte) error {
	type plain PublicErrorsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicReasonsResponseDataItem struct {
	Object  string                     `json:"object"`
	Code    string                     `json:"code"`
	Label   string                     `json:"label"`
	Message string                     `json:"message"`
	Meaning any                        `json:"meaning"`
	DocsURL string                     `json:"docs_url"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PublicReasonsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicReasonsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicReasonsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicReasonsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "code", "label", "message", "meaning", "docs_url"})
}

type PublicReasonsResponse struct {
	Object     string                          `json:"object"`
	Data       []PublicReasonsResponseDataItem `json:"data"`
	NextCursor any                             `json:"next_cursor"`
	Extra      map[string]json.RawMessage      `json:"-"`
}

func (v PublicReasonsResponse) MarshalJSON() ([]byte, error) {
	type plain PublicReasonsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicReasonsResponse) UnmarshalJSON(b []byte) error {
	type plain PublicReasonsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicWebhookEventsResponseDataItemExample struct {
	Object     string                     `json:"object"`
	ID         string                     `json:"id"`
	Type       string                     `json:"type"`
	APIVersion string                     `json:"api_version"`
	CreatedAt  string                     `json:"created_at"`
	Workspace  string                     `json:"workspace"`
	Data       map[string]any             `json:"data"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v PublicWebhookEventsResponseDataItemExample) MarshalJSON() ([]byte, error) {
	type plain PublicWebhookEventsResponseDataItemExample
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicWebhookEventsResponseDataItemExample) UnmarshalJSON(b []byte) error {
	type plain PublicWebhookEventsResponseDataItemExample
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "type", "api_version", "created_at", "workspace", "data"})
}

type PublicWebhookEventsResponseDataItem struct {
	Object      string                                     `json:"object"`
	Type        string                                     `json:"type"`
	Description string                                     `json:"description"`
	Example     PublicWebhookEventsResponseDataItemExample `json:"example"`
	Extra       map[string]json.RawMessage                 `json:"-"`
}

func (v PublicWebhookEventsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicWebhookEventsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicWebhookEventsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicWebhookEventsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "type", "description", "example"})
}

type PublicWebhookEventsResponse struct {
	Object     string                                `json:"object"`
	Data       []PublicWebhookEventsResponseDataItem `json:"data"`
	NextCursor any                                   `json:"next_cursor"`
	Extra      map[string]json.RawMessage            `json:"-"`
}

func (v PublicWebhookEventsResponse) MarshalJSON() ([]byte, error) {
	type plain PublicWebhookEventsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicWebhookEventsResponse) UnmarshalJSON(b []byte) error {
	type plain PublicWebhookEventsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicConditionsResponseDataItem struct {
	Object  string                     `json:"object"`
	Key     string                     `json:"key"`
	Label   string                     `json:"label"`
	Group   string                     `json:"group"`
	Type    string                     `json:"type"`
	Values  Optional[[]string]         `json:"values,omitzero"`
	Min     Optional[float64]          `json:"min,omitzero"`
	Max     Optional[float64]          `json:"max,omitzero"`
	Integer Optional[bool]             `json:"integer,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PublicConditionsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicConditionsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicConditionsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicConditionsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "key", "label", "group", "type", "values", "min", "max", "integer"})
}

type PublicConditionsResponse struct {
	Object     string                             `json:"object"`
	Data       []PublicConditionsResponseDataItem `json:"data"`
	NextCursor any                                `json:"next_cursor"`
	Extra      map[string]json.RawMessage         `json:"-"`
}

func (v PublicConditionsResponse) MarshalJSON() ([]byte, error) {
	type plain PublicConditionsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicConditionsResponse) UnmarshalJSON(b []byte) error {
	type plain PublicConditionsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicGlossaryResponseTermsItem struct {
	Term       string                     `json:"term"`
	Definition string                     `json:"definition"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v PublicGlossaryResponseTermsItem) MarshalJSON() ([]byte, error) {
	type plain PublicGlossaryResponseTermsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicGlossaryResponseTermsItem) UnmarshalJSON(b []byte) error {
	type plain PublicGlossaryResponseTermsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"term", "definition"})
}

type PublicGlossaryResponseFaqItem struct {
	Section  any                        `json:"section"`
	Question string                     `json:"question"`
	Answer   string                     `json:"answer"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v PublicGlossaryResponseFaqItem) MarshalJSON() ([]byte, error) {
	type plain PublicGlossaryResponseFaqItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicGlossaryResponseFaqItem) UnmarshalJSON(b []byte) error {
	type plain PublicGlossaryResponseFaqItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"section", "question", "answer"})
}

type PublicGlossaryResponse struct {
	Object    string                            `json:"object"`
	Terms     []PublicGlossaryResponseTermsItem `json:"terms"`
	Faq       []PublicGlossaryResponseFaqItem   `json:"faq"`
	URL       string                            `json:"url"`
	FaqURL    string                            `json:"faq_url"`
	UpdatedAt string                            `json:"updated_at"`
	Extra     map[string]json.RawMessage        `json:"-"`
}

func (v PublicGlossaryResponse) MarshalJSON() ([]byte, error) {
	type plain PublicGlossaryResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicGlossaryResponse) UnmarshalJSON(b []byte) error {
	type plain PublicGlossaryResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "terms", "faq", "url", "faq_url", "updated_at"})
}

type PublicLegalResponseDataItem struct {
	Object        string                     `json:"object"`
	Document      string                     `json:"document"`
	Title         string                     `json:"title"`
	Version       string                     `json:"version"`
	EffectiveDate string                     `json:"effective_date"`
	URL           string                     `json:"url"`
	UpdatedAt     string                     `json:"updated_at"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v PublicLegalResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicLegalResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicLegalResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicLegalResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "document", "title", "version", "effective_date", "url", "updated_at"})
}

type PublicLegalResponse struct {
	Object     string                        `json:"object"`
	Data       []PublicLegalResponseDataItem `json:"data"`
	NextCursor any                           `json:"next_cursor"`
	Extra      map[string]json.RawMessage    `json:"-"`
}

func (v PublicLegalResponse) MarshalJSON() ([]byte, error) {
	type plain PublicLegalResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicLegalResponse) UnmarshalJSON(b []byte) error {
	type plain PublicLegalResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicLegalDocumentResponse struct {
	Object        string                     `json:"object"`
	Document      string                     `json:"document"`
	Title         string                     `json:"title"`
	Version       string                     `json:"version"`
	EffectiveDate string                     `json:"effective_date"`
	URL           string                     `json:"url"`
	UpdatedAt     string                     `json:"updated_at"`
	Markdown      string                     `json:"markdown"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v PublicLegalDocumentResponse) MarshalJSON() ([]byte, error) {
	type plain PublicLegalDocumentResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicLegalDocumentResponse) UnmarshalJSON(b []byte) error {
	type plain PublicLegalDocumentResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "document", "title", "version", "effective_date", "url", "updated_at", "markdown"})
}

type PublicChangelogResponseEntriesItem struct {
	ID       string                     `json:"id"`
	Title    string                     `json:"title"`
	Date     string                     `json:"date"`
	Markdown string                     `json:"markdown"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v PublicChangelogResponseEntriesItem) MarshalJSON() ([]byte, error) {
	type plain PublicChangelogResponseEntriesItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicChangelogResponseEntriesItem) UnmarshalJSON(b []byte) error {
	type plain PublicChangelogResponseEntriesItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "title", "date", "markdown"})
}

type PublicChangelogResponse struct {
	Object    string                               `json:"object"`
	URL       string                               `json:"url"`
	Entries   []PublicChangelogResponseEntriesItem `json:"entries"`
	UpdatedAt string                               `json:"updated_at"`
	Extra     map[string]json.RawMessage           `json:"-"`
}

func (v PublicChangelogResponse) MarshalJSON() ([]byte, error) {
	type plain PublicChangelogResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicChangelogResponse) UnmarshalJSON(b []byte) error {
	type plain PublicChangelogResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "url", "entries", "updated_at"})
}

type PublicBlogInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PublicBlogInput) MarshalJSON() ([]byte, error) {
	type plain PublicBlogInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicBlogInput) UnmarshalJSON(b []byte) error {
	type plain PublicBlogInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type PublicBlogResponseDataItem struct {
	Object      string                     `json:"object"`
	Slug        string                     `json:"slug"`
	Title       string                     `json:"title"`
	Description any                        `json:"description"`
	URL         string                     `json:"url"`
	PublishedOn any                        `json:"published_on"`
	RevisedOn   any                        `json:"revised_on"`
	Authors     []string                   `json:"authors"`
	Tags        []string                   `json:"tags"`
	UpdatedAt   string                     `json:"updated_at"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicBlogResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PublicBlogResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicBlogResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PublicBlogResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "slug", "title", "description", "url", "published_on", "revised_on", "authors", "tags", "updated_at"})
}

type PublicBlogResponse struct {
	Object     string                       `json:"object"`
	Data       []PublicBlogResponseDataItem `json:"data"`
	NextCursor any                          `json:"next_cursor"`
	Extra      map[string]json.RawMessage   `json:"-"`
}

func (v PublicBlogResponse) MarshalJSON() ([]byte, error) {
	type plain PublicBlogResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicBlogResponse) UnmarshalJSON(b []byte) error {
	type plain PublicBlogResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PublicBlogPostResponse struct {
	Object      string                     `json:"object"`
	Slug        string                     `json:"slug"`
	Title       string                     `json:"title"`
	Description any                        `json:"description"`
	URL         string                     `json:"url"`
	PublishedOn any                        `json:"published_on"`
	RevisedOn   any                        `json:"revised_on"`
	Authors     []string                   `json:"authors"`
	Tags        []string                   `json:"tags"`
	UpdatedAt   string                     `json:"updated_at"`
	Markdown    string                     `json:"markdown"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicBlogPostResponse) MarshalJSON() ([]byte, error) {
	type plain PublicBlogPostResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicBlogPostResponse) UnmarshalJSON(b []byte) error {
	type plain PublicBlogPostResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "slug", "title", "description", "url", "published_on", "revised_on", "authors", "tags", "updated_at", "markdown"})
}

type PublicVersionsResponseNpm struct {
	Package string                     `json:"package"`
	Version string                     `json:"version"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PublicVersionsResponseNpm) MarshalJSON() ([]byte, error) {
	type plain PublicVersionsResponseNpm
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicVersionsResponseNpm) UnmarshalJSON(b []byte) error {
	type plain PublicVersionsResponseNpm
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"package", "version"})
}

type PublicVersionsResponseBrowserScript struct {
	Version string                     `json:"version"`
	URL     string                     `json:"url"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PublicVersionsResponseBrowserScript) MarshalJSON() ([]byte, error) {
	type plain PublicVersionsResponseBrowserScript
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicVersionsResponseBrowserScript) UnmarshalJSON(b []byte) error {
	type plain PublicVersionsResponseBrowserScript
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"version", "url"})
}

type PublicVersionsResponseSkillVariant0 struct {
	Version string                     `json:"version"`
	URL     string                     `json:"url"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PublicVersionsResponseSkillVariant0) MarshalJSON() ([]byte, error) {
	type plain PublicVersionsResponseSkillVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicVersionsResponseSkillVariant0) UnmarshalJSON(b []byte) error {
	type plain PublicVersionsResponseSkillVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"version", "url"})
}

type PublicVersionsResponse struct {
	Object        string                              `json:"object"`
	API           string                              `json:"api"`
	Npm           PublicVersionsResponseNpm           `json:"npm"`
	BrowserScript PublicVersionsResponseBrowserScript `json:"browser_script"`
	Skill         any                                 `json:"skill"`
	Extra         map[string]json.RawMessage          `json:"-"`
}

func (v PublicVersionsResponse) MarshalJSON() ([]byte, error) {
	type plain PublicVersionsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicVersionsResponse) UnmarshalJSON(b []byte) error {
	type plain PublicVersionsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "api", "npm", "browser_script", "skill"})
}

type PublicStatusResponse struct {
	Object    string                     `json:"object"`
	Status    string                     `json:"status"`
	CheckedAt string                     `json:"checked_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v PublicStatusResponse) MarshalJSON() ([]byte, error) {
	type plain PublicStatusResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicStatusResponse) UnmarshalJSON(b []byte) error {
	type plain PublicStatusResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "checked_at"})
}

type PublicContactsResponseSupport struct {
	Email string                     `json:"email"`
	Form  any                        `json:"form"`
	API   string                     `json:"api"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v PublicContactsResponseSupport) MarshalJSON() ([]byte, error) {
	type plain PublicContactsResponseSupport
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicContactsResponseSupport) UnmarshalJSON(b []byte) error {
	type plain PublicContactsResponseSupport
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email", "form", "api"})
}

type PublicContactsResponseSales struct {
	Form  string                     `json:"form"`
	API   string                     `json:"api"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v PublicContactsResponseSales) MarshalJSON() ([]byte, error) {
	type plain PublicContactsResponseSales
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicContactsResponseSales) UnmarshalJSON(b []byte) error {
	type plain PublicContactsResponseSales
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"form", "api"})
}

type PublicContactsResponseSecurity struct {
	Email string                     `json:"email"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v PublicContactsResponseSecurity) MarshalJSON() ([]byte, error) {
	type plain PublicContactsResponseSecurity
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicContactsResponseSecurity) UnmarshalJSON(b []byte) error {
	type plain PublicContactsResponseSecurity
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email"})
}

type PublicContactsResponse struct {
	Object   string                         `json:"object"`
	Support  PublicContactsResponseSupport  `json:"support"`
	Sales    PublicContactsResponseSales    `json:"sales"`
	Security PublicContactsResponseSecurity `json:"security"`
	Extra    map[string]json.RawMessage     `json:"-"`
}

func (v PublicContactsResponse) MarshalJSON() ([]byte, error) {
	type plain PublicContactsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicContactsResponse) UnmarshalJSON(b []byte) error {
	type plain PublicContactsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "support", "sales", "security"})
}

type PublicContactSalesInput struct {
	Name    string                     `json:"name"`
	Email   string                     `json:"email"`
	Message string                     `json:"message"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PublicContactSalesInput) MarshalJSON() ([]byte, error) {
	type plain PublicContactSalesInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicContactSalesInput) UnmarshalJSON(b []byte) error {
	type plain PublicContactSalesInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "email", "message"})
}

type PublicContactSalesResponse struct {
	Object   string                     `json:"object"`
	Received bool                       `json:"received"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v PublicContactSalesResponse) MarshalJSON() ([]byte, error) {
	type plain PublicContactSalesResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicContactSalesResponse) UnmarshalJSON(b []byte) error {
	type plain PublicContactSalesResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "received"})
}

type PublicContactSupportInput struct {
	Email       string                     `json:"email"`
	Subject     string                     `json:"subject"`
	Message     string                     `json:"message"`
	Topic       string                     `json:"topic"`
	Name        Optional[string]           `json:"name,omitzero"`
	WorkspaceID Optional[string]           `json:"workspace_id,omitzero"`
	RequestID   Optional[string]           `json:"request_id,omitzero"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v PublicContactSupportInput) MarshalJSON() ([]byte, error) {
	type plain PublicContactSupportInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicContactSupportInput) UnmarshalJSON(b []byte) error {
	type plain PublicContactSupportInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email", "subject", "message", "topic", "name", "workspace_id", "request_id"})
}

type PublicContactSupportResponse struct {
	Object    string                     `json:"object"`
	Reference string                     `json:"reference"`
	Email     string                     `json:"email"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v PublicContactSupportResponse) MarshalJSON() ([]byte, error) {
	type plain PublicContactSupportResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicContactSupportResponse) UnmarshalJSON(b []byte) error {
	type plain PublicContactSupportResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "reference", "email"})
}

type PublicOnboardingStartInput struct {
	WorkspaceName     string                     `json:"workspace_name"`
	Plan              string                     `json:"plan"`
	Interval          Optional[string]           `json:"interval,omitzero"`
	CreditPackDollars Optional[int64]            `json:"credit_pack_dollars,omitzero"`
	Extra             map[string]json.RawMessage `json:"-"`
}

func (v PublicOnboardingStartInput) MarshalJSON() ([]byte, error) {
	type plain PublicOnboardingStartInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicOnboardingStartInput) UnmarshalJSON(b []byte) error {
	type plain PublicOnboardingStartInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"workspace_name", "plan", "interval", "credit_pack_dollars"})
}

type PublicOnboardingStartResponse struct {
	Object            string                     `json:"object"`
	ID                string                     `json:"id"`
	ApprovalURL       string                     `json:"approval_url"`
	Secret            string                     `json:"secret"`
	ExpiresAt         string                     `json:"expires_at"`
	WorkspaceName     string                     `json:"workspace_name"`
	Plan              string                     `json:"plan"`
	Interval          any                        `json:"interval"`
	CreditPackDollars any                        `json:"credit_pack_dollars"`
	AmountCents       int64                      `json:"amount_cents"`
	Extra             map[string]json.RawMessage `json:"-"`
}

func (v PublicOnboardingStartResponse) MarshalJSON() ([]byte, error) {
	type plain PublicOnboardingStartResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicOnboardingStartResponse) UnmarshalJSON(b []byte) error {
	type plain PublicOnboardingStartResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "approval_url", "secret", "expires_at", "workspace_name", "plan", "interval", "credit_pack_dollars", "amount_cents"})
}

type PublicOnboardingStatusInput struct {
	Secret string                     `json:"secret"`
	Wait   Optional[int64]            `json:"wait,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PublicOnboardingStatusInput) MarshalJSON() ([]byte, error) {
	type plain PublicOnboardingStatusInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicOnboardingStatusInput) UnmarshalJSON(b []byte) error {
	type plain PublicOnboardingStatusInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"secret", "wait"})
}

type PublicOnboardingStatusResponseSignInVariant0 struct {
	TokenEndpoint string                     `json:"token_endpoint"`
	GrantType     string                     `json:"grant_type"`
	ClientID      string                     `json:"client_id"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v PublicOnboardingStatusResponseSignInVariant0) MarshalJSON() ([]byte, error) {
	type plain PublicOnboardingStatusResponseSignInVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicOnboardingStatusResponseSignInVariant0) UnmarshalJSON(b []byte) error {
	type plain PublicOnboardingStatusResponseSignInVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"token_endpoint", "grant_type", "client_id"})
}

type PublicOnboardingStatusResponse struct {
	Object    string                     `json:"object"`
	ID        string                     `json:"id"`
	Status    string                     `json:"status"`
	ExpiresAt string                     `json:"expires_at"`
	Workspace any                        `json:"workspace"`
	SignIn    any                        `json:"sign_in"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v PublicOnboardingStatusResponse) MarshalJSON() ([]byte, error) {
	type plain PublicOnboardingStatusResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PublicOnboardingStatusResponse) UnmarshalJSON(b []byte) error {
	type plain PublicOnboardingStatusResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "status", "expires_at", "workspace", "sign_in"})
}

type WorkspacesListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WorkspacesListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain WorkspacesListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspacesListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain WorkspacesListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type WorkspacesListResponse struct {
	Object     string                           `json:"object"`
	Data       []WorkspacesListResponseDataItem `json:"data"`
	NextCursor any                              `json:"next_cursor"`
	Extra      map[string]json.RawMessage       `json:"-"`
}

func (v WorkspacesListResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspacesListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspacesListResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspacesListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type MeGetResponsePersonVariant0 struct {
	Name  any                        `json:"name"`
	Email any                        `json:"email"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v MeGetResponsePersonVariant0) MarshalJSON() ([]byte, error) {
	type plain MeGetResponsePersonVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *MeGetResponsePersonVariant0) UnmarshalJSON(b []byte) error {
	type plain MeGetResponsePersonVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "email"})
}

type MeGetResponseAppVariant0 struct {
	ID    string                     `json:"id"`
	Name  any                        `json:"name"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v MeGetResponseAppVariant0) MarshalJSON() ([]byte, error) {
	type plain MeGetResponseAppVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *MeGetResponseAppVariant0) UnmarshalJSON(b []byte) error {
	type plain MeGetResponseAppVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "name"})
}

type MeGetResponseAPIKeyVariant0 struct {
	ID    string                     `json:"id"`
	Name  any                        `json:"name"`
	Type  string                     `json:"type"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v MeGetResponseAPIKeyVariant0) MarshalJSON() ([]byte, error) {
	type plain MeGetResponseAPIKeyVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *MeGetResponseAPIKeyVariant0) UnmarshalJSON(b []byte) error {
	type plain MeGetResponseAPIKeyVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "name", "type"})
}

type MeGetResponseMachineClientVariant0 struct {
	ID    string                     `json:"id"`
	Name  any                        `json:"name"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v MeGetResponseMachineClientVariant0) MarshalJSON() ([]byte, error) {
	type plain MeGetResponseMachineClientVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *MeGetResponseMachineClientVariant0) UnmarshalJSON(b []byte) error {
	type plain MeGetResponseMachineClientVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "name"})
}

type MeGetResponseWorkspaceVariant0 struct {
	ID    string                     `json:"id"`
	Name  any                        `json:"name"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v MeGetResponseWorkspaceVariant0) MarshalJSON() ([]byte, error) {
	type plain MeGetResponseWorkspaceVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *MeGetResponseWorkspaceVariant0) UnmarshalJSON(b []byte) error {
	type plain MeGetResponseWorkspaceVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "name"})
}

type MeGetResponse struct {
	Object        string                     `json:"object"`
	Credential    string                     `json:"credential"`
	Person        any                        `json:"person"`
	App           any                        `json:"app"`
	APIKey        any                        `json:"api_key"`
	MachineClient any                        `json:"machine_client"`
	Workspace     any                        `json:"workspace"`
	Role          any                        `json:"role"`
	Permissions   []string                   `json:"permissions"`
	Scopes        []string                   `json:"scopes"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v MeGetResponse) MarshalJSON() ([]byte, error) {
	type plain MeGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *MeGetResponse) UnmarshalJSON(b []byte) error {
	type plain MeGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "credential", "person", "app", "api_key", "machine_client", "workspace", "role", "permissions", "scopes"})
}

type APIKeysListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v APIKeysListInput) MarshalJSON() ([]byte, error) {
	type plain APIKeysListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysListInput) UnmarshalJSON(b []byte) error {
	type plain APIKeysListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type APIKeysListResponseDataItem struct {
	Object               string                     `json:"object"`
	ID                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Type                 string                     `json:"type"`
	Scopes               []string                   `json:"scopes"`
	Prefix               any                        `json:"prefix"`
	LastFour             any                        `json:"last_four"`
	Value                any                        `json:"value"`
	Status               string                     `json:"status"`
	ReplacementAvailable bool                       `json:"replacement_available"`
	ExpiresAt            any                        `json:"expires_at"`
	LastUsedAt           any                        `json:"last_used_at"`
	CreatedAt            any                        `json:"created_at"`
	Extra                map[string]json.RawMessage `json:"-"`
}

func (v APIKeysListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain APIKeysListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain APIKeysListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "type", "scopes", "prefix", "last_four", "value", "status", "replacement_available", "expires_at", "last_used_at", "created_at"})
}

type APIKeysListResponse struct {
	Object     string                        `json:"object"`
	Data       []APIKeysListResponseDataItem `json:"data"`
	NextCursor any                           `json:"next_cursor"`
	Extra      map[string]json.RawMessage    `json:"-"`
}

func (v APIKeysListResponse) MarshalJSON() ([]byte, error) {
	type plain APIKeysListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysListResponse) UnmarshalJSON(b []byte) error {
	type plain APIKeysListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type APIKeysCreateInput struct {
	Name          string                     `json:"name"`
	Type          string                     `json:"type"`
	Scopes        Optional[[]string]         `json:"scopes,omitzero"`
	ExpiresInDays Optional[any]              `json:"expires_in_days,omitzero"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v APIKeysCreateInput) MarshalJSON() ([]byte, error) {
	type plain APIKeysCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysCreateInput) UnmarshalJSON(b []byte) error {
	type plain APIKeysCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "type", "scopes", "expires_in_days"})
}

type APIKeysCreateResponse struct {
	Object               string                     `json:"object"`
	ID                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Type                 string                     `json:"type"`
	Scopes               []string                   `json:"scopes"`
	Prefix               any                        `json:"prefix"`
	LastFour             any                        `json:"last_four"`
	Value                any                        `json:"value"`
	Status               string                     `json:"status"`
	ReplacementAvailable bool                       `json:"replacement_available"`
	ExpiresAt            any                        `json:"expires_at"`
	LastUsedAt           any                        `json:"last_used_at"`
	CreatedAt            any                        `json:"created_at"`
	Secret               string                     `json:"secret"`
	Extra                map[string]json.RawMessage `json:"-"`
}

func (v APIKeysCreateResponse) MarshalJSON() ([]byte, error) {
	type plain APIKeysCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysCreateResponse) UnmarshalJSON(b []byte) error {
	type plain APIKeysCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "type", "scopes", "prefix", "last_four", "value", "status", "replacement_available", "expires_at", "last_used_at", "created_at", "secret"})
}

type APIKeysUpdateInput struct {
	Name          Optional[string]           `json:"name,omitzero"`
	Scopes        Optional[[]string]         `json:"scopes,omitzero"`
	ExpiresInDays Optional[any]              `json:"expires_in_days,omitzero"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v APIKeysUpdateInput) MarshalJSON() ([]byte, error) {
	type plain APIKeysUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysUpdateInput) UnmarshalJSON(b []byte) error {
	type plain APIKeysUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "scopes", "expires_in_days"})
}

type APIKeysUpdateResponse struct {
	Object               string                     `json:"object"`
	ID                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Type                 string                     `json:"type"`
	Scopes               []string                   `json:"scopes"`
	Prefix               any                        `json:"prefix"`
	LastFour             any                        `json:"last_four"`
	Value                any                        `json:"value"`
	Status               string                     `json:"status"`
	ReplacementAvailable bool                       `json:"replacement_available"`
	ExpiresAt            any                        `json:"expires_at"`
	LastUsedAt           any                        `json:"last_used_at"`
	CreatedAt            any                        `json:"created_at"`
	Extra                map[string]json.RawMessage `json:"-"`
}

func (v APIKeysUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain APIKeysUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain APIKeysUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "type", "scopes", "prefix", "last_four", "value", "status", "replacement_available", "expires_at", "last_used_at", "created_at"})
}

type APIKeysRollInput struct {
	OverlapHours any                        `json:"overlap_hours"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v APIKeysRollInput) MarshalJSON() ([]byte, error) {
	type plain APIKeysRollInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysRollInput) UnmarshalJSON(b []byte) error {
	type plain APIKeysRollInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"overlap_hours"})
}

type APIKeysRollResponsePrevious struct {
	ID        any                        `json:"id"`
	ExpiresAt any                        `json:"expires_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v APIKeysRollResponsePrevious) MarshalJSON() ([]byte, error) {
	type plain APIKeysRollResponsePrevious
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysRollResponsePrevious) UnmarshalJSON(b []byte) error {
	type plain APIKeysRollResponsePrevious
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "expires_at"})
}

type APIKeysRollResponse struct {
	Object               string                      `json:"object"`
	ID                   string                      `json:"id"`
	Name                 string                      `json:"name"`
	Type                 string                      `json:"type"`
	Scopes               []string                    `json:"scopes"`
	Prefix               any                         `json:"prefix"`
	LastFour             any                         `json:"last_four"`
	Value                any                         `json:"value"`
	Status               string                      `json:"status"`
	ReplacementAvailable bool                        `json:"replacement_available"`
	ExpiresAt            any                         `json:"expires_at"`
	LastUsedAt           any                         `json:"last_used_at"`
	CreatedAt            any                         `json:"created_at"`
	Secret               string                      `json:"secret"`
	Previous             APIKeysRollResponsePrevious `json:"previous"`
	Extra                map[string]json.RawMessage  `json:"-"`
}

func (v APIKeysRollResponse) MarshalJSON() ([]byte, error) {
	type plain APIKeysRollResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysRollResponse) UnmarshalJSON(b []byte) error {
	type plain APIKeysRollResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "type", "scopes", "prefix", "last_four", "value", "status", "replacement_available", "expires_at", "last_used_at", "created_at", "secret", "previous"})
}

type APIKeysReplacementResponse struct {
	Object               string                     `json:"object"`
	ID                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Type                 string                     `json:"type"`
	Scopes               []string                   `json:"scopes"`
	Prefix               any                        `json:"prefix"`
	LastFour             any                        `json:"last_four"`
	Value                any                        `json:"value"`
	Status               string                     `json:"status"`
	ReplacementAvailable bool                       `json:"replacement_available"`
	ExpiresAt            any                        `json:"expires_at"`
	LastUsedAt           any                        `json:"last_used_at"`
	CreatedAt            any                        `json:"created_at"`
	Secret               string                     `json:"secret"`
	Extra                map[string]json.RawMessage `json:"-"`
}

func (v APIKeysReplacementResponse) MarshalJSON() ([]byte, error) {
	type plain APIKeysReplacementResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysReplacementResponse) UnmarshalJSON(b []byte) error {
	type plain APIKeysReplacementResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "type", "scopes", "prefix", "last_four", "value", "status", "replacement_available", "expires_at", "last_used_at", "created_at", "secret"})
}

type APIKeysRevokeResponse struct {
	Object               string                     `json:"object"`
	ID                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Type                 string                     `json:"type"`
	Scopes               []string                   `json:"scopes"`
	Prefix               any                        `json:"prefix"`
	LastFour             any                        `json:"last_four"`
	Value                any                        `json:"value"`
	Status               string                     `json:"status"`
	ReplacementAvailable bool                       `json:"replacement_available"`
	ExpiresAt            any                        `json:"expires_at"`
	LastUsedAt           any                        `json:"last_used_at"`
	CreatedAt            any                        `json:"created_at"`
	Extra                map[string]json.RawMessage `json:"-"`
}

func (v APIKeysRevokeResponse) MarshalJSON() ([]byte, error) {
	type plain APIKeysRevokeResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *APIKeysRevokeResponse) UnmarshalJSON(b []byte) error {
	type plain APIKeysRevokeResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "type", "scopes", "prefix", "last_four", "value", "status", "replacement_available", "expires_at", "last_used_at", "created_at"})
}

type MachineClientsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	View   Optional[string]           `json:"view,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v MachineClientsListInput) MarshalJSON() ([]byte, error) {
	type plain MachineClientsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsListInput) UnmarshalJSON(b []byte) error {
	type plain MachineClientsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit", "view"})
}

type MachineClientsListResponseDataItemMemberVariant0 struct {
	Name   any                        `json:"name"`
	Email  any                        `json:"email"`
	Role   any                        `json:"role"`
	You    bool                       `json:"you"`
	Former bool                       `json:"former"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v MachineClientsListResponseDataItemMemberVariant0) MarshalJSON() ([]byte, error) {
	type plain MachineClientsListResponseDataItemMemberVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsListResponseDataItemMemberVariant0) UnmarshalJSON(b []byte) error {
	type plain MachineClientsListResponseDataItemMemberVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "email", "role", "you", "former"})
}

type MachineClientsListResponseDataItem struct {
	Object     string                     `json:"object"`
	ID         any                        `json:"id"`
	Name       string                     `json:"name"`
	Scopes     []string                   `json:"scopes"`
	Member     any                        `json:"member"`
	CreatedAt  any                        `json:"created_at"`
	LastUsedAt any                        `json:"last_used_at"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v MachineClientsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain MachineClientsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain MachineClientsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "scopes", "member", "created_at", "last_used_at"})
}

type MachineClientsListResponse struct {
	Object     string                               `json:"object"`
	Data       []MachineClientsListResponseDataItem `json:"data"`
	NextCursor any                                  `json:"next_cursor"`
	Extra      map[string]json.RawMessage           `json:"-"`
}

func (v MachineClientsListResponse) MarshalJSON() ([]byte, error) {
	type plain MachineClientsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsListResponse) UnmarshalJSON(b []byte) error {
	type plain MachineClientsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type MachineClientsCreateInput struct {
	Name   string                     `json:"name"`
	Scopes []string                   `json:"scopes"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v MachineClientsCreateInput) MarshalJSON() ([]byte, error) {
	type plain MachineClientsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsCreateInput) UnmarshalJSON(b []byte) error {
	type plain MachineClientsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "scopes"})
}

type MachineClientsCreateResponseMemberVariant0 struct {
	Name   any                        `json:"name"`
	Email  any                        `json:"email"`
	Role   any                        `json:"role"`
	You    bool                       `json:"you"`
	Former bool                       `json:"former"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v MachineClientsCreateResponseMemberVariant0) MarshalJSON() ([]byte, error) {
	type plain MachineClientsCreateResponseMemberVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsCreateResponseMemberVariant0) UnmarshalJSON(b []byte) error {
	type plain MachineClientsCreateResponseMemberVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "email", "role", "you", "former"})
}

type MachineClientsCreateResponse struct {
	Object        string                     `json:"object"`
	ID            any                        `json:"id"`
	Name          string                     `json:"name"`
	Scopes        []string                   `json:"scopes"`
	Member        any                        `json:"member"`
	CreatedAt     any                        `json:"created_at"`
	LastUsedAt    any                        `json:"last_used_at"`
	Secret        string                     `json:"secret"`
	TokenEndpoint string                     `json:"token_endpoint"`
	Audience      string                     `json:"audience"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v MachineClientsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain MachineClientsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain MachineClientsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "scopes", "member", "created_at", "last_used_at", "secret", "token_endpoint", "audience"})
}

type MachineClientsRotateSecretResponse struct {
	Object string                     `json:"object"`
	ID     any                        `json:"id"`
	Secret string                     `json:"secret"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v MachineClientsRotateSecretResponse) MarshalJSON() ([]byte, error) {
	type plain MachineClientsRotateSecretResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsRotateSecretResponse) UnmarshalJSON(b []byte) error {
	type plain MachineClientsRotateSecretResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "secret"})
}

type MachineClientsRevokeResponse struct {
	Object  string                     `json:"object"`
	ID      any                        `json:"id"`
	Deleted bool                       `json:"deleted"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v MachineClientsRevokeResponse) MarshalJSON() ([]byte, error) {
	type plain MachineClientsRevokeResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *MachineClientsRevokeResponse) UnmarshalJSON(b []byte) error {
	type plain MachineClientsRevokeResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "deleted"})
}

type AuthorizedAppsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	View   Optional[string]           `json:"view,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AuthorizedAppsListInput) MarshalJSON() ([]byte, error) {
	type plain AuthorizedAppsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuthorizedAppsListInput) UnmarshalJSON(b []byte) error {
	type plain AuthorizedAppsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit", "view"})
}

type AuthorizedAppsListResponseDataItemMemberVariant0 struct {
	Name   any                        `json:"name"`
	Email  any                        `json:"email"`
	Role   any                        `json:"role"`
	You    bool                       `json:"you"`
	Former bool                       `json:"former"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AuthorizedAppsListResponseDataItemMemberVariant0) MarshalJSON() ([]byte, error) {
	type plain AuthorizedAppsListResponseDataItemMemberVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuthorizedAppsListResponseDataItemMemberVariant0) UnmarshalJSON(b []byte) error {
	type plain AuthorizedAppsListResponseDataItemMemberVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "email", "role", "you", "former"})
}

type AuthorizedAppsListResponseDataItem struct {
	Object      string                     `json:"object"`
	ID          any                        `json:"id"`
	Name        string                     `json:"name"`
	Trust       string                     `json:"trust"`
	Destination any                        `json:"destination"`
	Scopes      []string                   `json:"scopes"`
	Member      any                        `json:"member"`
	ApprovedAt  any                        `json:"approved_at"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v AuthorizedAppsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain AuthorizedAppsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuthorizedAppsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain AuthorizedAppsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "trust", "destination", "scopes", "member", "approved_at"})
}

type AuthorizedAppsListResponse struct {
	Object     string                               `json:"object"`
	Data       []AuthorizedAppsListResponseDataItem `json:"data"`
	NextCursor any                                  `json:"next_cursor"`
	Extra      map[string]json.RawMessage           `json:"-"`
}

func (v AuthorizedAppsListResponse) MarshalJSON() ([]byte, error) {
	type plain AuthorizedAppsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuthorizedAppsListResponse) UnmarshalJSON(b []byte) error {
	type plain AuthorizedAppsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type AuthorizedAppsDisconnectResponse struct {
	Object  string                     `json:"object"`
	ID      any                        `json:"id"`
	Deleted bool                       `json:"deleted"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v AuthorizedAppsDisconnectResponse) MarshalJSON() ([]byte, error) {
	type plain AuthorizedAppsDisconnectResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuthorizedAppsDisconnectResponse) UnmarshalJSON(b []byte) error {
	type plain AuthorizedAppsDisconnectResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "deleted"})
}

type MembersListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v MembersListInput) MarshalJSON() ([]byte, error) {
	type plain MembersListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *MembersListInput) UnmarshalJSON(b []byte) error {
	type plain MembersListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type MembersListResponseDataItem struct {
	Object       string                     `json:"object"`
	ID           any                        `json:"id"`
	Name         any                        `json:"name"`
	Email        string                     `json:"email"`
	Role         string                     `json:"role"`
	Permissions  []string                   `json:"permissions"`
	InvitedAt    any                        `json:"invited_at"`
	JoinedAt     any                        `json:"joined_at"`
	LastSignInAt any                        `json:"last_sign_in_at"`
	You          bool                       `json:"you"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v MembersListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain MembersListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *MembersListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain MembersListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "email", "role", "permissions", "invited_at", "joined_at", "last_sign_in_at", "you"})
}

type MembersListResponse struct {
	Object     string                        `json:"object"`
	Data       []MembersListResponseDataItem `json:"data"`
	NextCursor any                           `json:"next_cursor"`
	Extra      map[string]json.RawMessage    `json:"-"`
}

func (v MembersListResponse) MarshalJSON() ([]byte, error) {
	type plain MembersListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *MembersListResponse) UnmarshalJSON(b []byte) error {
	type plain MembersListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type MembersUpdateInput struct {
	Role        Optional[string]           `json:"role,omitzero"`
	Permissions Optional[[]string]         `json:"permissions,omitzero"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v MembersUpdateInput) MarshalJSON() ([]byte, error) {
	type plain MembersUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *MembersUpdateInput) UnmarshalJSON(b []byte) error {
	type plain MembersUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"role", "permissions"})
}

type MembersUpdateResponse struct {
	Object       string                     `json:"object"`
	ID           any                        `json:"id"`
	Name         any                        `json:"name"`
	Email        string                     `json:"email"`
	Role         string                     `json:"role"`
	Permissions  []string                   `json:"permissions"`
	InvitedAt    any                        `json:"invited_at"`
	JoinedAt     any                        `json:"joined_at"`
	LastSignInAt any                        `json:"last_sign_in_at"`
	You          bool                       `json:"you"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v MembersUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain MembersUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *MembersUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain MembersUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "email", "role", "permissions", "invited_at", "joined_at", "last_sign_in_at", "you"})
}

type MembersRemoveResponse struct {
	Object  string                     `json:"object"`
	ID      any                        `json:"id"`
	Deleted bool                       `json:"deleted"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v MembersRemoveResponse) MarshalJSON() ([]byte, error) {
	type plain MembersRemoveResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *MembersRemoveResponse) UnmarshalJSON(b []byte) error {
	type plain MembersRemoveResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "deleted"})
}

type InvitationsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v InvitationsListInput) MarshalJSON() ([]byte, error) {
	type plain InvitationsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *InvitationsListInput) UnmarshalJSON(b []byte) error {
	type plain InvitationsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type InvitationsListResponseDataItem struct {
	Object    string                     `json:"object"`
	ID        any                        `json:"id"`
	Email     string                     `json:"email"`
	Role      string                     `json:"role"`
	InvitedBy any                        `json:"invited_by"`
	Status    string                     `json:"status"`
	CreatedAt any                        `json:"created_at"`
	ExpiresAt any                        `json:"expires_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v InvitationsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain InvitationsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *InvitationsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain InvitationsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "email", "role", "invited_by", "status", "created_at", "expires_at"})
}

type InvitationsListResponse struct {
	Object     string                            `json:"object"`
	Data       []InvitationsListResponseDataItem `json:"data"`
	NextCursor any                               `json:"next_cursor"`
	Extra      map[string]json.RawMessage        `json:"-"`
}

func (v InvitationsListResponse) MarshalJSON() ([]byte, error) {
	type plain InvitationsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *InvitationsListResponse) UnmarshalJSON(b []byte) error {
	type plain InvitationsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type InvitationsCreateInput struct {
	Email string                     `json:"email"`
	Role  string                     `json:"role"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v InvitationsCreateInput) MarshalJSON() ([]byte, error) {
	type plain InvitationsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *InvitationsCreateInput) UnmarshalJSON(b []byte) error {
	type plain InvitationsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email", "role"})
}

type InvitationsCreateResponse struct {
	Object    string                     `json:"object"`
	ID        any                        `json:"id"`
	Email     string                     `json:"email"`
	Role      string                     `json:"role"`
	InvitedBy any                        `json:"invited_by"`
	Status    string                     `json:"status"`
	CreatedAt any                        `json:"created_at"`
	ExpiresAt any                        `json:"expires_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v InvitationsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain InvitationsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *InvitationsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain InvitationsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "email", "role", "invited_by", "status", "created_at", "expires_at"})
}

type InvitationsResendResponse struct {
	Object    string                     `json:"object"`
	ID        any                        `json:"id"`
	Email     string                     `json:"email"`
	Role      string                     `json:"role"`
	InvitedBy any                        `json:"invited_by"`
	Status    string                     `json:"status"`
	CreatedAt any                        `json:"created_at"`
	ExpiresAt any                        `json:"expires_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v InvitationsResendResponse) MarshalJSON() ([]byte, error) {
	type plain InvitationsResendResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *InvitationsResendResponse) UnmarshalJSON(b []byte) error {
	type plain InvitationsResendResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "email", "role", "invited_by", "status", "created_at", "expires_at"})
}

type InvitationsCancelResponse struct {
	Object    string                     `json:"object"`
	ID        any                        `json:"id"`
	Email     string                     `json:"email"`
	Role      string                     `json:"role"`
	InvitedBy any                        `json:"invited_by"`
	Status    string                     `json:"status"`
	CreatedAt any                        `json:"created_at"`
	ExpiresAt any                        `json:"expires_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v InvitationsCancelResponse) MarshalJSON() ([]byte, error) {
	type plain InvitationsCancelResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *InvitationsCancelResponse) UnmarshalJSON(b []byte) error {
	type plain InvitationsCancelResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "email", "role", "invited_by", "status", "created_at", "expires_at"})
}

type WorkspaceGetResponseIconVariant0 struct {
	UpdatedAt any                        `json:"updated_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceGetResponseIconVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceGetResponseIconVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceGetResponseIconVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceGetResponseIconVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"updated_at"})
}

type WorkspaceGetResponseProductTypeVariant0 struct {
	Type  string                     `json:"type"`
	Other Optional[string]           `json:"other,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceGetResponseProductTypeVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceGetResponseProductTypeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceGetResponseProductTypeVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceGetResponseProductTypeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "other"})
}

type WorkspaceGetResponseDataRetention struct {
	EvidenceDays int64                      `json:"evidence_days"`
	MaxDays      int64                      `json:"max_days"`
	Options      []int64                    `json:"options"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceGetResponseDataRetention) MarshalJSON() ([]byte, error) {
	type plain WorkspaceGetResponseDataRetention
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceGetResponseDataRetention) UnmarshalJSON(b []byte) error {
	type plain WorkspaceGetResponseDataRetention
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"evidence_days", "max_days", "options"})
}

type WorkspaceGetResponse struct {
	Object        string                            `json:"object"`
	ID            string                            `json:"id"`
	Name          string                            `json:"name"`
	Icon          any                               `json:"icon"`
	TimeZone      any                               `json:"time_zone"`
	Website       any                               `json:"website"`
	ProductType   any                               `json:"product_type"`
	DataRetention WorkspaceGetResponseDataRetention `json:"data_retention"`
	Extra         map[string]json.RawMessage        `json:"-"`
}

func (v WorkspaceGetResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceGetResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "icon", "time_zone", "website", "product_type", "data_retention"})
}

type WorkspaceUpdateInputProductTypeVariant0 struct {
	Type  string                     `json:"type"`
	Other Optional[string]           `json:"other,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceUpdateInputProductTypeVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceUpdateInputProductTypeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceUpdateInputProductTypeVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceUpdateInputProductTypeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "other"})
}

type WorkspaceUpdateInput struct {
	Name        Optional[string]           `json:"name,omitzero"`
	Website     Optional[any]              `json:"website,omitzero"`
	ProductType Optional[any]              `json:"product_type,omitzero"`
	TimeZone    Optional[string]           `json:"time_zone,omitzero"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceUpdateInput) MarshalJSON() ([]byte, error) {
	type plain WorkspaceUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceUpdateInput) UnmarshalJSON(b []byte) error {
	type plain WorkspaceUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "website", "product_type", "time_zone"})
}

type WorkspaceUpdateResponseIconVariant0 struct {
	UpdatedAt any                        `json:"updated_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceUpdateResponseIconVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceUpdateResponseIconVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceUpdateResponseIconVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceUpdateResponseIconVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"updated_at"})
}

type WorkspaceUpdateResponseProductTypeVariant0 struct {
	Type  string                     `json:"type"`
	Other Optional[string]           `json:"other,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceUpdateResponseProductTypeVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceUpdateResponseProductTypeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceUpdateResponseProductTypeVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceUpdateResponseProductTypeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "other"})
}

type WorkspaceUpdateResponseDataRetention struct {
	EvidenceDays int64                      `json:"evidence_days"`
	MaxDays      int64                      `json:"max_days"`
	Options      []int64                    `json:"options"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceUpdateResponseDataRetention) MarshalJSON() ([]byte, error) {
	type plain WorkspaceUpdateResponseDataRetention
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceUpdateResponseDataRetention) UnmarshalJSON(b []byte) error {
	type plain WorkspaceUpdateResponseDataRetention
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"evidence_days", "max_days", "options"})
}

type WorkspaceUpdateResponse struct {
	Object        string                               `json:"object"`
	ID            string                               `json:"id"`
	Name          string                               `json:"name"`
	Icon          any                                  `json:"icon"`
	TimeZone      any                                  `json:"time_zone"`
	Website       any                                  `json:"website"`
	ProductType   any                                  `json:"product_type"`
	DataRetention WorkspaceUpdateResponseDataRetention `json:"data_retention"`
	Extra         map[string]json.RawMessage           `json:"-"`
}

func (v WorkspaceUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "icon", "time_zone", "website", "product_type", "data_retention"})
}

type WorkspaceIconUploadInput struct {
	Image string                     `json:"image"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceIconUploadInput) MarshalJSON() ([]byte, error) {
	type plain WorkspaceIconUploadInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceIconUploadInput) UnmarshalJSON(b []byte) error {
	type plain WorkspaceIconUploadInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"image"})
}

type WorkspaceIconUploadResponse struct {
	Object string                     `json:"object"`
	ID     string                     `json:"id"`
	State  string                     `json:"state"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceIconUploadResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceIconUploadResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceIconUploadResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceIconUploadResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "state"})
}

type WorkspaceIconRemoveResponseIconVariant0 struct {
	UpdatedAt any                        `json:"updated_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceIconRemoveResponseIconVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceIconRemoveResponseIconVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceIconRemoveResponseIconVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceIconRemoveResponseIconVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"updated_at"})
}

type WorkspaceIconRemoveResponseProductTypeVariant0 struct {
	Type  string                     `json:"type"`
	Other Optional[string]           `json:"other,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceIconRemoveResponseProductTypeVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceIconRemoveResponseProductTypeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceIconRemoveResponseProductTypeVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceIconRemoveResponseProductTypeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "other"})
}

type WorkspaceIconRemoveResponseDataRetention struct {
	EvidenceDays int64                      `json:"evidence_days"`
	MaxDays      int64                      `json:"max_days"`
	Options      []int64                    `json:"options"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceIconRemoveResponseDataRetention) MarshalJSON() ([]byte, error) {
	type plain WorkspaceIconRemoveResponseDataRetention
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceIconRemoveResponseDataRetention) UnmarshalJSON(b []byte) error {
	type plain WorkspaceIconRemoveResponseDataRetention
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"evidence_days", "max_days", "options"})
}

type WorkspaceIconRemoveResponse struct {
	Object        string                                   `json:"object"`
	ID            string                                   `json:"id"`
	Name          string                                   `json:"name"`
	Icon          any                                      `json:"icon"`
	TimeZone      any                                      `json:"time_zone"`
	Website       any                                      `json:"website"`
	ProductType   any                                      `json:"product_type"`
	DataRetention WorkspaceIconRemoveResponseDataRetention `json:"data_retention"`
	Extra         map[string]json.RawMessage               `json:"-"`
}

func (v WorkspaceIconRemoveResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceIconRemoveResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceIconRemoveResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceIconRemoveResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "icon", "time_zone", "website", "product_type", "data_retention"})
}

type WorkspaceIconUploadGetResponse struct {
	Object string                     `json:"object"`
	ID     string                     `json:"id"`
	State  string                     `json:"state"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceIconUploadGetResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceIconUploadGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceIconUploadGetResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceIconUploadGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "state"})
}

type WorkspaceRetentionInput struct {
	EvidenceDays int64                      `json:"evidence_days"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceRetentionInput) MarshalJSON() ([]byte, error) {
	type plain WorkspaceRetentionInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceRetentionInput) UnmarshalJSON(b []byte) error {
	type plain WorkspaceRetentionInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"evidence_days"})
}

type WorkspaceRetentionResponseIconVariant0 struct {
	UpdatedAt any                        `json:"updated_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceRetentionResponseIconVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceRetentionResponseIconVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceRetentionResponseIconVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceRetentionResponseIconVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"updated_at"})
}

type WorkspaceRetentionResponseProductTypeVariant0 struct {
	Type  string                     `json:"type"`
	Other Optional[string]           `json:"other,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceRetentionResponseProductTypeVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceRetentionResponseProductTypeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceRetentionResponseProductTypeVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceRetentionResponseProductTypeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "other"})
}

type WorkspaceRetentionResponseDataRetention struct {
	EvidenceDays int64                      `json:"evidence_days"`
	MaxDays      int64                      `json:"max_days"`
	Options      []int64                    `json:"options"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceRetentionResponseDataRetention) MarshalJSON() ([]byte, error) {
	type plain WorkspaceRetentionResponseDataRetention
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceRetentionResponseDataRetention) UnmarshalJSON(b []byte) error {
	type plain WorkspaceRetentionResponseDataRetention
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"evidence_days", "max_days", "options"})
}

type WorkspaceRetentionResponse struct {
	Object        string                                  `json:"object"`
	ID            string                                  `json:"id"`
	Name          string                                  `json:"name"`
	Icon          any                                     `json:"icon"`
	TimeZone      any                                     `json:"time_zone"`
	Website       any                                     `json:"website"`
	ProductType   any                                     `json:"product_type"`
	DataRetention WorkspaceRetentionResponseDataRetention `json:"data_retention"`
	Extra         map[string]json.RawMessage              `json:"-"`
}

func (v WorkspaceRetentionResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceRetentionResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceRetentionResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceRetentionResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "icon", "time_zone", "website", "product_type", "data_retention"})
}

type WorkspaceSetupResponseStepsItem struct {
	Step  string                     `json:"step"`
	Done  bool                       `json:"done"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceSetupResponseStepsItem) MarshalJSON() ([]byte, error) {
	type plain WorkspaceSetupResponseStepsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceSetupResponseStepsItem) UnmarshalJSON(b []byte) error {
	type plain WorkspaceSetupResponseStepsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"step", "done"})
}

type WorkspaceSetupResponse struct {
	Object    string                            `json:"object"`
	Steps     []WorkspaceSetupResponseStepsItem `json:"steps"`
	Complete  bool                              `json:"complete"`
	Finished  bool                              `json:"finished"`
	Dismissed any                               `json:"dismissed"`
	Extra     map[string]json.RawMessage        `json:"-"`
}

func (v WorkspaceSetupResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceSetupResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceSetupResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceSetupResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "steps", "complete", "finished", "dismissed"})
}

type WorkspaceSetupFinishInputProductTypeVariant0 struct {
	Type  string                     `json:"type"`
	Other Optional[string]           `json:"other,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceSetupFinishInputProductTypeVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceSetupFinishInputProductTypeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceSetupFinishInputProductTypeVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceSetupFinishInputProductTypeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "other"})
}

type WorkspaceSetupFinishInput struct {
	Website     Optional[any]              `json:"website,omitzero"`
	ProductType Optional[any]              `json:"product_type,omitzero"`
	TimeZone    Optional[any]              `json:"time_zone,omitzero"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceSetupFinishInput) MarshalJSON() ([]byte, error) {
	type plain WorkspaceSetupFinishInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceSetupFinishInput) UnmarshalJSON(b []byte) error {
	type plain WorkspaceSetupFinishInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"website", "product_type", "time_zone"})
}

type WorkspaceSetupFinishResponseStepsItem struct {
	Step  string                     `json:"step"`
	Done  bool                       `json:"done"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceSetupFinishResponseStepsItem) MarshalJSON() ([]byte, error) {
	type plain WorkspaceSetupFinishResponseStepsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceSetupFinishResponseStepsItem) UnmarshalJSON(b []byte) error {
	type plain WorkspaceSetupFinishResponseStepsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"step", "done"})
}

type WorkspaceSetupFinishResponse struct {
	Object    string                                  `json:"object"`
	Steps     []WorkspaceSetupFinishResponseStepsItem `json:"steps"`
	Complete  bool                                    `json:"complete"`
	Finished  bool                                    `json:"finished"`
	Dismissed any                                     `json:"dismissed"`
	Extra     map[string]json.RawMessage              `json:"-"`
}

func (v WorkspaceSetupFinishResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceSetupFinishResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceSetupFinishResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceSetupFinishResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "steps", "complete", "finished", "dismissed"})
}

type WorkspaceSetupDismissResponseStepsItem struct {
	Step  string                     `json:"step"`
	Done  bool                       `json:"done"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceSetupDismissResponseStepsItem) MarshalJSON() ([]byte, error) {
	type plain WorkspaceSetupDismissResponseStepsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceSetupDismissResponseStepsItem) UnmarshalJSON(b []byte) error {
	type plain WorkspaceSetupDismissResponseStepsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"step", "done"})
}

type WorkspaceSetupDismissResponse struct {
	Object    string                                   `json:"object"`
	Steps     []WorkspaceSetupDismissResponseStepsItem `json:"steps"`
	Complete  bool                                     `json:"complete"`
	Finished  bool                                     `json:"finished"`
	Dismissed any                                      `json:"dismissed"`
	Extra     map[string]json.RawMessage               `json:"-"`
}

func (v WorkspaceSetupDismissResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceSetupDismissResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceSetupDismissResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceSetupDismissResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "steps", "complete", "finished", "dismissed"})
}

type WorkspaceLeaveResponse struct {
	Object    string                     `json:"object"`
	Workspace string                     `json:"workspace"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceLeaveResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceLeaveResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceLeaveResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceLeaveResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "workspace"})
}

type EvaluationsListInput struct {
	Cursor  Optional[string]           `json:"cursor,omitzero"`
	Limit   Optional[int64]            `json:"limit,omitzero"`
	Result  Optional[string]           `json:"result,omitzero"`
	Benefit Optional[string]           `json:"benefit,omitzero"`
	Query   Optional[string]           `json:"query,omitzero"`
	Days    Optional[int64]            `json:"days,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsListInput) MarshalJSON() ([]byte, error) {
	type plain EvaluationsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsListInput) UnmarshalJSON(b []byte) error {
	type plain EvaluationsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit", "result", "benefit", "query", "days"})
}

type EvaluationsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain EvaluationsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain EvaluationsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type EvaluationsListResponse struct {
	Object     string                            `json:"object"`
	Data       []EvaluationsListResponseDataItem `json:"data"`
	NextCursor any                               `json:"next_cursor"`
	Extra      map[string]json.RawMessage        `json:"-"`
}

func (v EvaluationsListResponse) MarshalJSON() ([]byte, error) {
	type plain EvaluationsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsListResponse) UnmarshalJSON(b []byte) error {
	type plain EvaluationsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type EvaluationsCreateInputSubject struct {
	ExternalID    Optional[string]           `json:"external_id,omitzero"`
	Email         string                     `json:"email"`
	EmailVerified Optional[bool]             `json:"email_verified,omitzero"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateInputSubject) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInputSubject
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInputSubject) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInputSubject
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"external_id", "email", "email_verified"})
}

type EvaluationsCreateInputContextVariant0 struct {
	IP        Optional[any]              `json:"ip,omitzero"`
	UserAgent Optional[string]           `json:"user_agent,omitzero"`
	SessionID Optional[string]           `json:"session_id,omitzero"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateInputContextVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInputContextVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInputContextVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInputContextVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"ip", "user_agent", "session_id"})
}

type EvaluationsCreateInputIdentifiersPayment struct {
	Provider    string                     `json:"provider"`
	Scope       string                     `json:"scope"`
	Fingerprint string                     `json:"fingerprint"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateInputIdentifiersPayment) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInputIdentifiersPayment
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInputIdentifiersPayment) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInputIdentifiersPayment
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"provider", "scope", "fingerprint"})
}

type EvaluationsCreateInputIdentifiersPhoneHash struct {
	Algorithm  string                     `json:"algorithm"`
	KeyVersion string                     `json:"key_version"`
	Value      string                     `json:"value"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateInputIdentifiersPhoneHash) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInputIdentifiersPhoneHash
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInputIdentifiersPhoneHash) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInputIdentifiersPhoneHash
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"algorithm", "key_version", "value"})
}

type EvaluationsCreateInputIdentifiersCustomItemValueHash struct {
	Algorithm  string                     `json:"algorithm"`
	KeyVersion string                     `json:"key_version"`
	Value      string                     `json:"value"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateInputIdentifiersCustomItemValueHash) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInputIdentifiersCustomItemValueHash
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInputIdentifiersCustomItemValueHash) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInputIdentifiersCustomItemValueHash
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"algorithm", "key_version", "value"})
}

type EvaluationsCreateInputIdentifiersCustomItem struct {
	Key       string                                               `json:"key"`
	ValueHash EvaluationsCreateInputIdentifiersCustomItemValueHash `json:"value_hash"`
	Extra     map[string]json.RawMessage                           `json:"-"`
}

func (v EvaluationsCreateInputIdentifiersCustomItem) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInputIdentifiersCustomItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInputIdentifiersCustomItem) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInputIdentifiersCustomItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"key", "value_hash"})
}

type EvaluationsCreateInputIdentifiers struct {
	DeviceID        Optional[string]                                        `json:"device_id,omitzero"`
	DeviceToken     Optional[string]                                        `json:"device_token,omitzero"`
	BrowserCookieID Optional[string]                                        `json:"browser_cookie_id,omitzero"`
	Payment         Optional[EvaluationsCreateInputIdentifiersPayment]      `json:"payment,omitzero"`
	PhoneHash       Optional[EvaluationsCreateInputIdentifiersPhoneHash]    `json:"phone_hash,omitzero"`
	PhoneVerified   Optional[bool]                                          `json:"phone_verified,omitzero"`
	Custom          Optional[[]EvaluationsCreateInputIdentifiersCustomItem] `json:"custom,omitzero"`
	Extra           map[string]json.RawMessage                              `json:"-"`
}

func (v EvaluationsCreateInputIdentifiers) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInputIdentifiers
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInputIdentifiers) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInputIdentifiers
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"device_id", "device_token", "browser_cookie_id", "payment", "phone_hash", "phone_verified", "custom"})
}

type EvaluationsCreateInputRequest struct {
	Method  string                     `json:"method"`
	URL     string                     `json:"url"`
	Headers map[string]any             `json:"headers"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateInputRequest) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInputRequest
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInputRequest) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInputRequest
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"method", "url", "headers"})
}

type EvaluationsCreateInput struct {
	Benefit     string                                      `json:"benefit"`
	ClaimID     string                                      `json:"claim_id"`
	Subject     EvaluationsCreateInputSubject               `json:"subject"`
	Context     Optional[any]                               `json:"context,omitzero"`
	Identifiers Optional[EvaluationsCreateInputIdentifiers] `json:"identifiers,omitzero"`
	Request     Optional[EvaluationsCreateInputRequest]     `json:"request,omitzero"`
	Extra       map[string]json.RawMessage                  `json:"-"`
}

func (v EvaluationsCreateInput) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateInput) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"benefit", "claim_id", "subject", "context", "identifiers", "request"})
}

type EvaluationsCreateResponseVariant0Risk struct {
	Score any                        `json:"score"`
	Band  string                     `json:"band"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateResponseVariant0Risk) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateResponseVariant0Risk
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateResponseVariant0Risk) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateResponseVariant0Risk
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"score", "band"})
}

type EvaluationsCreateResponseVariant0ReasonVariant0 struct {
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Rule    Optional[string]           `json:"rule,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateResponseVariant0ReasonVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateResponseVariant0ReasonVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateResponseVariant0ReasonVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateResponseVariant0ReasonVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code", "message", "rule"})
}

type EvaluationsCreateResponseVariant0OutcomeVariant0 struct {
	State string                     `json:"state"`
	DueAt any                        `json:"due_at"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateResponseVariant0OutcomeVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateResponseVariant0OutcomeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateResponseVariant0OutcomeVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateResponseVariant0OutcomeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "due_at"})
}

type EvaluationsCreateResponseVariant0 struct {
	Object  string                                `json:"object"`
	ID      string                                `json:"id"`
	Result  string                                `json:"result"`
	Risk    EvaluationsCreateResponseVariant0Risk `json:"risk"`
	Reason  any                                   `json:"reason"`
	Outcome Optional[any]                         `json:"outcome,omitzero"`
	Extra   map[string]json.RawMessage            `json:"-"`
}

func (v EvaluationsCreateResponseVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateResponseVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateResponseVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateResponseVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "result", "risk", "reason", "outcome"})
}

type EvaluationsCreateResponseVariant1 struct {
	Object string                     `json:"object"`
	Status string                     `json:"status"`
	Reason string                     `json:"reason"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsCreateResponseVariant1) MarshalJSON() ([]byte, error) {
	type plain EvaluationsCreateResponseVariant1
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsCreateResponseVariant1) UnmarshalJSON(b []byte) error {
	type plain EvaluationsCreateResponseVariant1
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "reason"})
}

type EvaluationsFiltersInput struct {
	Days  Optional[int64]            `json:"days,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsFiltersInput) MarshalJSON() ([]byte, error) {
	type plain EvaluationsFiltersInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsFiltersInput) UnmarshalJSON(b []byte) error {
	type plain EvaluationsFiltersInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"days"})
}

type EvaluationsFiltersResponseReasonsItem struct {
	Code  string                     `json:"code"`
	Label string                     `json:"label"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsFiltersResponseReasonsItem) MarshalJSON() ([]byte, error) {
	type plain EvaluationsFiltersResponseReasonsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsFiltersResponseReasonsItem) UnmarshalJSON(b []byte) error {
	type plain EvaluationsFiltersResponseReasonsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code", "label"})
}

type EvaluationsFiltersResponse struct {
	Object   string                                  `json:"object"`
	Days     any                                     `json:"days"`
	Benefits []string                                `json:"benefits"`
	Results  []string                                `json:"results"`
	Reasons  []EvaluationsFiltersResponseReasonsItem `json:"reasons"`
	Extra    map[string]json.RawMessage              `json:"-"`
}

func (v EvaluationsFiltersResponse) MarshalJSON() ([]byte, error) {
	type plain EvaluationsFiltersResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsFiltersResponse) UnmarshalJSON(b []byte) error {
	type plain EvaluationsFiltersResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "days", "benefits", "results", "reasons"})
}

type EvaluationsGetResponseReasonVariant0 struct {
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Rule    Optional[string]           `json:"rule,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsGetResponseReasonVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsGetResponseReasonVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsGetResponseReasonVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsGetResponseReasonVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code", "message", "rule"})
}

type EvaluationsGetResponseOutcome struct {
	State any                        `json:"state"`
	DueAt Optional[any]              `json:"due_at,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsGetResponseOutcome) MarshalJSON() ([]byte, error) {
	type plain EvaluationsGetResponseOutcome
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsGetResponseOutcome) UnmarshalJSON(b []byte) error {
	type plain EvaluationsGetResponseOutcome
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "due_at"})
}

type EvaluationsGetResponseRisk struct {
	Score any                        `json:"score"`
	Band  string                     `json:"band"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsGetResponseRisk) MarshalJSON() ([]byte, error) {
	type plain EvaluationsGetResponseRisk
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsGetResponseRisk) UnmarshalJSON(b []byte) error {
	type plain EvaluationsGetResponseRisk
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"score", "band"})
}

type EvaluationsGetResponseInputs struct {
	Email   string                     `json:"email"`
	IP      string                     `json:"ip"`
	UserID  string                     `json:"user_id"`
	Phone   string                     `json:"phone"`
	Payment string                     `json:"payment"`
	Device  string                     `json:"device"`
	Browser string                     `json:"browser"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsGetResponseInputs) MarshalJSON() ([]byte, error) {
	type plain EvaluationsGetResponseInputs
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsGetResponseInputs) UnmarshalJSON(b []byte) error {
	type plain EvaluationsGetResponseInputs
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email", "ip", "user_id", "phone", "payment", "device", "browser"})
}

type EvaluationsGetResponsePolicy struct {
	ID      any                        `json:"id"`
	Version any                        `json:"version"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsGetResponsePolicy) MarshalJSON() ([]byte, error) {
	type plain EvaluationsGetResponsePolicy
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsGetResponsePolicy) UnmarshalJSON(b []byte) error {
	type plain EvaluationsGetResponsePolicy
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "version"})
}

type EvaluationsGetResponseAccessListEntryVariant0 struct {
	ID    any                        `json:"id"`
	Kind  string                     `json:"kind"`
	Type  any                        `json:"type"`
	Value any                        `json:"value"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsGetResponseAccessListEntryVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsGetResponseAccessListEntryVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsGetResponseAccessListEntryVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsGetResponseAccessListEntryVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "kind", "type", "value"})
}

type EvaluationsGetResponse struct {
	Object          string                        `json:"object"`
	ID              string                        `json:"id"`
	Benefit         any                           `json:"benefit"`
	ClaimID         any                           `json:"claim_id"`
	Result          string                        `json:"result"`
	Reason          any                           `json:"reason"`
	Outcome         EvaluationsGetResponseOutcome `json:"outcome"`
	PolicyVersion   any                           `json:"policy_version"`
	CreatedAt       any                           `json:"created_at"`
	Mode            string                        `json:"mode"`
	ObservedResult  Optional[any]                 `json:"observed_result,omitzero"`
	Risk            EvaluationsGetResponseRisk    `json:"risk"`
	Inputs          EvaluationsGetResponseInputs  `json:"inputs"`
	BrowserData     any                           `json:"browser_data"`
	Policy          EvaluationsGetResponsePolicy  `json:"policy"`
	AccessListEntry any                           `json:"access_list_entry"`
	Extra           map[string]json.RawMessage    `json:"-"`
}

func (v EvaluationsGetResponse) MarshalJSON() ([]byte, error) {
	type plain EvaluationsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsGetResponse) UnmarshalJSON(b []byte) error {
	type plain EvaluationsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "benefit", "claim_id", "result", "reason", "outcome", "policy_version", "created_at", "mode", "observed_result", "risk", "inputs", "browser_data", "policy", "access_list_entry"})
}

type EvaluationsGraphResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsGraphResponse) MarshalJSON() ([]byte, error) {
	type plain EvaluationsGraphResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsGraphResponse) UnmarshalJSON(b []byte) error {
	type plain EvaluationsGraphResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type EvaluationsEvidenceResponseReasonVariant0 struct {
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Rule    Optional[string]           `json:"rule,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsEvidenceResponseReasonVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponseReasonVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponseReasonVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponseReasonVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code", "message", "rule"})
}

type EvaluationsEvidenceResponseOutcome struct {
	State any                        `json:"state"`
	DueAt Optional[any]              `json:"due_at,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsEvidenceResponseOutcome) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponseOutcome
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponseOutcome) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponseOutcome
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "due_at"})
}

type EvaluationsEvidenceResponseRisk struct {
	Score any                        `json:"score"`
	Band  string                     `json:"band"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsEvidenceResponseRisk) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponseRisk
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponseRisk) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponseRisk
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"score", "band"})
}

type EvaluationsEvidenceResponseSignalsItem struct {
	Family     string                     `json:"family"`
	Assessment string                     `json:"assessment"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsEvidenceResponseSignalsItem) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponseSignalsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponseSignalsItem) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponseSignalsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"family", "assessment"})
}

type EvaluationsEvidenceResponseInputs struct {
	Email   string                     `json:"email"`
	IP      string                     `json:"ip"`
	UserID  string                     `json:"user_id"`
	Phone   string                     `json:"phone"`
	Payment string                     `json:"payment"`
	Device  string                     `json:"device"`
	Browser string                     `json:"browser"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsEvidenceResponseInputs) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponseInputs
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponseInputs) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponseInputs
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email", "ip", "user_id", "phone", "payment", "device", "browser"})
}

type EvaluationsEvidenceResponsePolicy struct {
	ID      any                        `json:"id"`
	Version any                        `json:"version"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsEvidenceResponsePolicy) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponsePolicy
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponsePolicy) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponsePolicy
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "version"})
}

type EvaluationsEvidenceResponseAccessListEntryVariant0 struct {
	ID    any                        `json:"id"`
	Kind  string                     `json:"kind"`
	Type  any                        `json:"type"`
	Value any                        `json:"value"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsEvidenceResponseAccessListEntryVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponseAccessListEntryVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponseAccessListEntryVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponseAccessListEntryVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "kind", "type", "value"})
}

type EvaluationsEvidenceResponseEvidenceDeviceProfile struct {
	Status    string                     `json:"status"`
	ExpiresAt any                        `json:"expires_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsEvidenceResponseEvidenceDeviceProfile) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponseEvidenceDeviceProfile
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponseEvidenceDeviceProfile) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponseEvidenceDeviceProfile
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"status", "expires_at"})
}

type EvaluationsEvidenceResponseEvidence struct {
	Request       any                                              `json:"request"`
	DeviceProfile EvaluationsEvidenceResponseEvidenceDeviceProfile `json:"device_profile"`
	Extra         map[string]json.RawMessage                       `json:"-"`
}

func (v EvaluationsEvidenceResponseEvidence) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponseEvidence
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponseEvidence) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponseEvidence
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"request", "device_profile"})
}

type EvaluationsEvidenceResponse struct {
	Object          string                                   `json:"object"`
	ID              string                                   `json:"id"`
	Benefit         any                                      `json:"benefit"`
	ClaimID         any                                      `json:"claim_id"`
	Result          string                                   `json:"result"`
	Reason          any                                      `json:"reason"`
	Outcome         EvaluationsEvidenceResponseOutcome       `json:"outcome"`
	PolicyVersion   any                                      `json:"policy_version"`
	CreatedAt       any                                      `json:"created_at"`
	Mode            string                                   `json:"mode"`
	ObservedResult  Optional[any]                            `json:"observed_result,omitzero"`
	Risk            EvaluationsEvidenceResponseRisk          `json:"risk"`
	Signals         []EvaluationsEvidenceResponseSignalsItem `json:"signals"`
	Inputs          EvaluationsEvidenceResponseInputs        `json:"inputs"`
	BrowserData     any                                      `json:"browser_data"`
	Policy          EvaluationsEvidenceResponsePolicy        `json:"policy"`
	AccessListEntry any                                      `json:"access_list_entry"`
	Evidence        EvaluationsEvidenceResponseEvidence      `json:"evidence"`
	Extra           map[string]json.RawMessage               `json:"-"`
}

func (v EvaluationsEvidenceResponse) MarshalJSON() ([]byte, error) {
	type plain EvaluationsEvidenceResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsEvidenceResponse) UnmarshalJSON(b []byte) error {
	type plain EvaluationsEvidenceResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "benefit", "claim_id", "result", "reason", "outcome", "policy_version", "created_at", "mode", "observed_result", "risk", "signals", "inputs", "browser_data", "policy", "access_list_entry", "evidence"})
}

type OutcomesListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v OutcomesListInput) MarshalJSON() ([]byte, error) {
	type plain OutcomesListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *OutcomesListInput) UnmarshalJSON(b []byte) error {
	type plain OutcomesListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type OutcomesListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v OutcomesListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain OutcomesListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *OutcomesListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain OutcomesListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type OutcomesListResponse struct {
	Object     string                         `json:"object"`
	Data       []OutcomesListResponseDataItem `json:"data"`
	NextCursor any                            `json:"next_cursor"`
	Extra      map[string]json.RawMessage     `json:"-"`
}

func (v OutcomesListResponse) MarshalJSON() ([]byte, error) {
	type plain OutcomesListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *OutcomesListResponse) UnmarshalJSON(b []byte) error {
	type plain OutcomesListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type OutcomesGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v OutcomesGetResponse) MarshalJSON() ([]byte, error) {
	type plain OutcomesGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *OutcomesGetResponse) UnmarshalJSON(b []byte) error {
	type plain OutcomesGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type OutcomesReportInput struct {
	Outcome           string                     `json:"outcome"`
	Reason            Optional[string]           `json:"reason,omitzero"`
	ExternalOutcomeID Optional[string]           `json:"external_outcome_id,omitzero"`
	Extra             map[string]json.RawMessage `json:"-"`
}

func (v OutcomesReportInput) MarshalJSON() ([]byte, error) {
	type plain OutcomesReportInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *OutcomesReportInput) UnmarshalJSON(b []byte) error {
	type plain OutcomesReportInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"outcome", "reason", "external_outcome_id"})
}

type OutcomesReportResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v OutcomesReportResponse) MarshalJSON() ([]byte, error) {
	type plain OutcomesReportResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *OutcomesReportResponse) UnmarshalJSON(b []byte) error {
	type plain OutcomesReportResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type EvaluationsDryRunInputSubject struct {
	ExternalID    Optional[string]           `json:"external_id,omitzero"`
	Email         string                     `json:"email"`
	EmailVerified Optional[bool]             `json:"email_verified,omitzero"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunInputSubject) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInputSubject
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInputSubject) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInputSubject
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"external_id", "email", "email_verified"})
}

type EvaluationsDryRunInputContextVariant0 struct {
	IP        Optional[any]              `json:"ip,omitzero"`
	UserAgent Optional[string]           `json:"user_agent,omitzero"`
	SessionID Optional[string]           `json:"session_id,omitzero"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunInputContextVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInputContextVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInputContextVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInputContextVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"ip", "user_agent", "session_id"})
}

type EvaluationsDryRunInputIdentifiersPayment struct {
	Provider    string                     `json:"provider"`
	Scope       string                     `json:"scope"`
	Fingerprint string                     `json:"fingerprint"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunInputIdentifiersPayment) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInputIdentifiersPayment
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInputIdentifiersPayment) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInputIdentifiersPayment
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"provider", "scope", "fingerprint"})
}

type EvaluationsDryRunInputIdentifiersPhoneHash struct {
	Algorithm  string                     `json:"algorithm"`
	KeyVersion string                     `json:"key_version"`
	Value      string                     `json:"value"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunInputIdentifiersPhoneHash) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInputIdentifiersPhoneHash
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInputIdentifiersPhoneHash) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInputIdentifiersPhoneHash
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"algorithm", "key_version", "value"})
}

type EvaluationsDryRunInputIdentifiersCustomItemValueHash struct {
	Algorithm  string                     `json:"algorithm"`
	KeyVersion string                     `json:"key_version"`
	Value      string                     `json:"value"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunInputIdentifiersCustomItemValueHash) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInputIdentifiersCustomItemValueHash
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInputIdentifiersCustomItemValueHash) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInputIdentifiersCustomItemValueHash
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"algorithm", "key_version", "value"})
}

type EvaluationsDryRunInputIdentifiersCustomItem struct {
	Key       string                                               `json:"key"`
	ValueHash EvaluationsDryRunInputIdentifiersCustomItemValueHash `json:"value_hash"`
	Extra     map[string]json.RawMessage                           `json:"-"`
}

func (v EvaluationsDryRunInputIdentifiersCustomItem) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInputIdentifiersCustomItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInputIdentifiersCustomItem) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInputIdentifiersCustomItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"key", "value_hash"})
}

type EvaluationsDryRunInputIdentifiers struct {
	DeviceID        Optional[string]                                        `json:"device_id,omitzero"`
	DeviceToken     Optional[string]                                        `json:"device_token,omitzero"`
	BrowserCookieID Optional[string]                                        `json:"browser_cookie_id,omitzero"`
	Payment         Optional[EvaluationsDryRunInputIdentifiersPayment]      `json:"payment,omitzero"`
	PhoneHash       Optional[EvaluationsDryRunInputIdentifiersPhoneHash]    `json:"phone_hash,omitzero"`
	PhoneVerified   Optional[bool]                                          `json:"phone_verified,omitzero"`
	Custom          Optional[[]EvaluationsDryRunInputIdentifiersCustomItem] `json:"custom,omitzero"`
	Extra           map[string]json.RawMessage                              `json:"-"`
}

func (v EvaluationsDryRunInputIdentifiers) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInputIdentifiers
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInputIdentifiers) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInputIdentifiers
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"device_id", "device_token", "browser_cookie_id", "payment", "phone_hash", "phone_verified", "custom"})
}

type EvaluationsDryRunInputRequest struct {
	Method  string                     `json:"method"`
	URL     string                     `json:"url"`
	Headers map[string]any             `json:"headers"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunInputRequest) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInputRequest
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInputRequest) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInputRequest
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"method", "url", "headers"})
}

type EvaluationsDryRunInput struct {
	Benefit     string                                      `json:"benefit"`
	ClaimID     string                                      `json:"claim_id"`
	Subject     EvaluationsDryRunInputSubject               `json:"subject"`
	Context     Optional[any]                               `json:"context,omitzero"`
	Identifiers Optional[EvaluationsDryRunInputIdentifiers] `json:"identifiers,omitzero"`
	Request     Optional[EvaluationsDryRunInputRequest]     `json:"request,omitzero"`
	Extra       map[string]json.RawMessage                  `json:"-"`
}

func (v EvaluationsDryRunInput) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunInput) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"benefit", "claim_id", "subject", "context", "identifiers", "request"})
}

type EvaluationsDryRunResponseVariant0Risk struct {
	Score any                        `json:"score"`
	Band  string                     `json:"band"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunResponseVariant0Risk) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunResponseVariant0Risk
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunResponseVariant0Risk) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunResponseVariant0Risk
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"score", "band"})
}

type EvaluationsDryRunResponseVariant0ReasonVariant0 struct {
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Rule    Optional[string]           `json:"rule,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunResponseVariant0ReasonVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunResponseVariant0ReasonVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunResponseVariant0ReasonVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunResponseVariant0ReasonVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code", "message", "rule"})
}

type EvaluationsDryRunResponseVariant0Inputs struct {
	Email   string                     `json:"email"`
	IP      string                     `json:"ip"`
	UserID  string                     `json:"user_id"`
	Phone   string                     `json:"phone"`
	Payment string                     `json:"payment"`
	Device  string                     `json:"device"`
	Browser string                     `json:"browser"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunResponseVariant0Inputs) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunResponseVariant0Inputs
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunResponseVariant0Inputs) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunResponseVariant0Inputs
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email", "ip", "user_id", "phone", "payment", "device", "browser"})
}

type EvaluationsDryRunResponseVariant0 struct {
	Object string                                  `json:"object"`
	DryRun bool                                    `json:"dry_run"`
	Result string                                  `json:"result"`
	Risk   EvaluationsDryRunResponseVariant0Risk   `json:"risk"`
	Reason any                                     `json:"reason"`
	Inputs EvaluationsDryRunResponseVariant0Inputs `json:"inputs"`
	Extra  map[string]json.RawMessage              `json:"-"`
}

func (v EvaluationsDryRunResponseVariant0) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunResponseVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunResponseVariant0) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunResponseVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "dry_run", "result", "risk", "reason", "inputs"})
}

type EvaluationsDryRunResponseVariant1 struct {
	Object string                     `json:"object"`
	Status string                     `json:"status"`
	Reason string                     `json:"reason"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v EvaluationsDryRunResponseVariant1) MarshalJSON() ([]byte, error) {
	type plain EvaluationsDryRunResponseVariant1
	return marshalExtras(plain(v), v.Extra)
}
func (v *EvaluationsDryRunResponseVariant1) UnmarshalJSON(b []byte) error {
	type plain EvaluationsDryRunResponseVariant1
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "reason"})
}

type SimulationsSequenceInput struct {
	BenefitID string                     `json:"benefit_id"`
	PolicyKey string                     `json:"policy_key"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v SimulationsSequenceInput) MarshalJSON() ([]byte, error) {
	type plain SimulationsSequenceInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *SimulationsSequenceInput) UnmarshalJSON(b []byte) error {
	type plain SimulationsSequenceInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"benefit_id", "policy_key"})
}

type SimulationsSequenceResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v SimulationsSequenceResponse) MarshalJSON() ([]byte, error) {
	type plain SimulationsSequenceResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *SimulationsSequenceResponse) UnmarshalJSON(b []byte) error {
	type plain SimulationsSequenceResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type SimulationsReplayInput struct {
	BenefitID string                     `json:"benefit_id"`
	PolicyKey string                     `json:"policy_key"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v SimulationsReplayInput) MarshalJSON() ([]byte, error) {
	type plain SimulationsReplayInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *SimulationsReplayInput) UnmarshalJSON(b []byte) error {
	type plain SimulationsReplayInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"benefit_id", "policy_key"})
}

type SimulationsReplayResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v SimulationsReplayResponse) MarshalJSON() ([]byte, error) {
	type plain SimulationsReplayResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *SimulationsReplayResponse) UnmarshalJSON(b []byte) error {
	type plain SimulationsReplayResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type BenefitsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BenefitsListInput) MarshalJSON() ([]byte, error) {
	type plain BenefitsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsListInput) UnmarshalJSON(b []byte) error {
	type plain BenefitsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type BenefitsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BenefitsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain BenefitsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain BenefitsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type BenefitsListResponse struct {
	Object     string                         `json:"object"`
	Data       []BenefitsListResponseDataItem `json:"data"`
	NextCursor any                            `json:"next_cursor"`
	Extra      map[string]json.RawMessage     `json:"-"`
}

func (v BenefitsListResponse) MarshalJSON() ([]byte, error) {
	type plain BenefitsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsListResponse) UnmarshalJSON(b []byte) error {
	type plain BenefitsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type BenefitsCreateInputConfigEligibilityVariant0 struct {
	Kind  string                     `json:"kind"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BenefitsCreateInputConfigEligibilityVariant0) MarshalJSON() ([]byte, error) {
	type plain BenefitsCreateInputConfigEligibilityVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsCreateInputConfigEligibilityVariant0) UnmarshalJSON(b []byte) error {
	type plain BenefitsCreateInputConfigEligibilityVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind"})
}

type BenefitsCreateInputConfigEligibilityVariant1 struct {
	Kind     string                     `json:"kind"`
	Period   string                     `json:"period"`
	Timezone string                     `json:"timezone"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v BenefitsCreateInputConfigEligibilityVariant1) MarshalJSON() ([]byte, error) {
	type plain BenefitsCreateInputConfigEligibilityVariant1
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsCreateInputConfigEligibilityVariant1) UnmarshalJSON(b []byte) error {
	type plain BenefitsCreateInputConfigEligibilityVariant1
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "period", "timezone"})
}

type BenefitsCreateInputConfigEligibilityVariant2 struct {
	Kind  string                     `json:"kind"`
	Days  int64                      `json:"days"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BenefitsCreateInputConfigEligibilityVariant2) MarshalJSON() ([]byte, error) {
	type plain BenefitsCreateInputConfigEligibilityVariant2
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsCreateInputConfigEligibilityVariant2) UnmarshalJSON(b []byte) error {
	type plain BenefitsCreateInputConfigEligibilityVariant2
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "days"})
}

type BenefitsCreateInputConfigValue struct {
	Amount   float64                    `json:"amount"`
	Currency string                     `json:"currency"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v BenefitsCreateInputConfigValue) MarshalJSON() ([]byte, error) {
	type plain BenefitsCreateInputConfigValue
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsCreateInputConfigValue) UnmarshalJSON(b []byte) error {
	type plain BenefitsCreateInputConfigValue
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"amount", "currency"})
}

type BenefitsCreateInputConfig struct {
	Automation             Optional[string]                         `json:"automation,omitzero"`
	OutcomeTracking        Optional[string]                         `json:"outcome_tracking,omitzero"`
	MissingOutcomeBehavior Optional[string]                         `json:"missing_outcome_behavior,omitzero"`
	Eligibility            Optional[any]                            `json:"eligibility,omitzero"`
	Value                  Optional[BenefitsCreateInputConfigValue] `json:"value,omitzero"`
	Extra                  map[string]json.RawMessage               `json:"-"`
}

func (v BenefitsCreateInputConfig) MarshalJSON() ([]byte, error) {
	type plain BenefitsCreateInputConfig
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsCreateInputConfig) UnmarshalJSON(b []byte) error {
	type plain BenefitsCreateInputConfig
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"automation", "outcome_tracking", "missing_outcome_behavior", "eligibility", "value"})
}

type BenefitsCreateInput struct {
	Name           string                     `json:"name"`
	Key            string                     `json:"key"`
	Config         BenefitsCreateInputConfig  `json:"config"`
	PolicyKey      string                     `json:"policy_key"`
	ProtectionMode string                     `json:"protection_mode"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BenefitsCreateInput) MarshalJSON() ([]byte, error) {
	type plain BenefitsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsCreateInput) UnmarshalJSON(b []byte) error {
	type plain BenefitsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "key", "config", "policy_key", "protection_mode"})
}

type BenefitsCreateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BenefitsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain BenefitsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain BenefitsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type BenefitsGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BenefitsGetResponse) MarshalJSON() ([]byte, error) {
	type plain BenefitsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsGetResponse) UnmarshalJSON(b []byte) error {
	type plain BenefitsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type BenefitsUpdateInputConfigEligibilityVariant0 struct {
	Kind  string                     `json:"kind"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BenefitsUpdateInputConfigEligibilityVariant0) MarshalJSON() ([]byte, error) {
	type plain BenefitsUpdateInputConfigEligibilityVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsUpdateInputConfigEligibilityVariant0) UnmarshalJSON(b []byte) error {
	type plain BenefitsUpdateInputConfigEligibilityVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind"})
}

type BenefitsUpdateInputConfigEligibilityVariant1 struct {
	Kind     string                     `json:"kind"`
	Period   string                     `json:"period"`
	Timezone string                     `json:"timezone"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v BenefitsUpdateInputConfigEligibilityVariant1) MarshalJSON() ([]byte, error) {
	type plain BenefitsUpdateInputConfigEligibilityVariant1
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsUpdateInputConfigEligibilityVariant1) UnmarshalJSON(b []byte) error {
	type plain BenefitsUpdateInputConfigEligibilityVariant1
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "period", "timezone"})
}

type BenefitsUpdateInputConfigEligibilityVariant2 struct {
	Kind  string                     `json:"kind"`
	Days  int64                      `json:"days"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BenefitsUpdateInputConfigEligibilityVariant2) MarshalJSON() ([]byte, error) {
	type plain BenefitsUpdateInputConfigEligibilityVariant2
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsUpdateInputConfigEligibilityVariant2) UnmarshalJSON(b []byte) error {
	type plain BenefitsUpdateInputConfigEligibilityVariant2
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "days"})
}

type BenefitsUpdateInputConfigValue struct {
	Amount   float64                    `json:"amount"`
	Currency string                     `json:"currency"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v BenefitsUpdateInputConfigValue) MarshalJSON() ([]byte, error) {
	type plain BenefitsUpdateInputConfigValue
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsUpdateInputConfigValue) UnmarshalJSON(b []byte) error {
	type plain BenefitsUpdateInputConfigValue
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"amount", "currency"})
}

type BenefitsUpdateInputConfig struct {
	Automation             Optional[string]                         `json:"automation,omitzero"`
	OutcomeTracking        Optional[string]                         `json:"outcome_tracking,omitzero"`
	MissingOutcomeBehavior Optional[string]                         `json:"missing_outcome_behavior,omitzero"`
	Eligibility            Optional[any]                            `json:"eligibility,omitzero"`
	Value                  Optional[BenefitsUpdateInputConfigValue] `json:"value,omitzero"`
	Extra                  map[string]json.RawMessage               `json:"-"`
}

func (v BenefitsUpdateInputConfig) MarshalJSON() ([]byte, error) {
	type plain BenefitsUpdateInputConfig
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsUpdateInputConfig) UnmarshalJSON(b []byte) error {
	type plain BenefitsUpdateInputConfig
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"automation", "outcome_tracking", "missing_outcome_behavior", "eligibility", "value"})
}

type BenefitsUpdateInput struct {
	Name           string                     `json:"name"`
	Config         BenefitsUpdateInputConfig  `json:"config"`
	PolicyKey      string                     `json:"policy_key"`
	ProtectionMode string                     `json:"protection_mode"`
	Revision       int64                      `json:"revision"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BenefitsUpdateInput) MarshalJSON() ([]byte, error) {
	type plain BenefitsUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsUpdateInput) UnmarshalJSON(b []byte) error {
	type plain BenefitsUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "config", "policy_key", "protection_mode", "revision"})
}

type BenefitsUpdateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BenefitsUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain BenefitsUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain BenefitsUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type BenefitsDeleteResponse struct {
	Object  string                     `json:"object"`
	ID      any                        `json:"id"`
	Deleted bool                       `json:"deleted"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v BenefitsDeleteResponse) MarshalJSON() ([]byte, error) {
	type plain BenefitsDeleteResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BenefitsDeleteResponse) UnmarshalJSON(b []byte) error {
	type plain BenefitsDeleteResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "deleted"})
}

type PoliciesListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesListInput) MarshalJSON() ([]byte, error) {
	type plain PoliciesListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesListInput) UnmarshalJSON(b []byte) error {
	type plain PoliciesListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type PoliciesListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PoliciesListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PoliciesListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type PoliciesListResponse struct {
	Object     string                         `json:"object"`
	Data       []PoliciesListResponseDataItem `json:"data"`
	NextCursor any                            `json:"next_cursor"`
	Extra      map[string]json.RawMessage     `json:"-"`
}

func (v PoliciesListResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesListResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PoliciesCreateInputDefinitionRulesItem struct {
	ID      string                     `json:"id"`
	Name    string                     `json:"name"`
	Result  string                     `json:"result"`
	When    any                        `json:"when"`
	Enabled Optional[bool]             `json:"enabled,omitzero"`
	Note    Optional[string]           `json:"note,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PoliciesCreateInputDefinitionRulesItem) MarshalJSON() ([]byte, error) {
	type plain PoliciesCreateInputDefinitionRulesItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesCreateInputDefinitionRulesItem) UnmarshalJSON(b []byte) error {
	type plain PoliciesCreateInputDefinitionRulesItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "name", "result", "when", "enabled", "note"})
}

type PoliciesCreateInputDefinition struct {
	Base          Optional[string]                         `json:"base,omitzero"`
	RiskThreshold Optional[any]                            `json:"risk_threshold,omitzero"`
	DefaultResult string                                   `json:"default_result"`
	Rules         []PoliciesCreateInputDefinitionRulesItem `json:"rules"`
	Extra         map[string]json.RawMessage               `json:"-"`
}

func (v PoliciesCreateInputDefinition) MarshalJSON() ([]byte, error) {
	type plain PoliciesCreateInputDefinition
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesCreateInputDefinition) UnmarshalJSON(b []byte) error {
	type plain PoliciesCreateInputDefinition
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"base", "risk_threshold", "default_result", "rules"})
}

type PoliciesCreateInput struct {
	Name       string                        `json:"name"`
	Key        string                        `json:"key"`
	Definition PoliciesCreateInputDefinition `json:"definition"`
	Extra      map[string]json.RawMessage    `json:"-"`
}

func (v PoliciesCreateInput) MarshalJSON() ([]byte, error) {
	type plain PoliciesCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesCreateInput) UnmarshalJSON(b []byte) error {
	type plain PoliciesCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "key", "definition"})
}

type PoliciesCreateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesCreateResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesCreateResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type PoliciesGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesGetResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesGetResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type PoliciesUpdateInputDefinitionRulesItem struct {
	ID      string                     `json:"id"`
	Name    string                     `json:"name"`
	Result  string                     `json:"result"`
	When    any                        `json:"when"`
	Enabled Optional[bool]             `json:"enabled,omitzero"`
	Note    Optional[string]           `json:"note,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PoliciesUpdateInputDefinitionRulesItem) MarshalJSON() ([]byte, error) {
	type plain PoliciesUpdateInputDefinitionRulesItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesUpdateInputDefinitionRulesItem) UnmarshalJSON(b []byte) error {
	type plain PoliciesUpdateInputDefinitionRulesItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "name", "result", "when", "enabled", "note"})
}

type PoliciesUpdateInputDefinition struct {
	Base          Optional[string]                         `json:"base,omitzero"`
	RiskThreshold Optional[any]                            `json:"risk_threshold,omitzero"`
	DefaultResult string                                   `json:"default_result"`
	Rules         []PoliciesUpdateInputDefinitionRulesItem `json:"rules"`
	Extra         map[string]json.RawMessage               `json:"-"`
}

func (v PoliciesUpdateInputDefinition) MarshalJSON() ([]byte, error) {
	type plain PoliciesUpdateInputDefinition
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesUpdateInputDefinition) UnmarshalJSON(b []byte) error {
	type plain PoliciesUpdateInputDefinition
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"base", "risk_threshold", "default_result", "rules"})
}

type PoliciesUpdateInput struct {
	Name         string                        `json:"name"`
	Definition   PoliciesUpdateInputDefinition `json:"definition"`
	Revision     int64                         `json:"revision"`
	RestoredFrom Optional[int64]               `json:"restored_from,omitzero"`
	Extra        map[string]json.RawMessage    `json:"-"`
}

func (v PoliciesUpdateInput) MarshalJSON() ([]byte, error) {
	type plain PoliciesUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesUpdateInput) UnmarshalJSON(b []byte) error {
	type plain PoliciesUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "definition", "revision", "restored_from"})
}

type PoliciesUpdateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type PoliciesDeleteResponse struct {
	Object  string                     `json:"object"`
	ID      any                        `json:"id"`
	Deleted bool                       `json:"deleted"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v PoliciesDeleteResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesDeleteResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesDeleteResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesDeleteResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "deleted"})
}

type PoliciesDuplicateInput struct {
	Name  string                     `json:"name"`
	Key   string                     `json:"key"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v PoliciesDuplicateInput) MarshalJSON() ([]byte, error) {
	type plain PoliciesDuplicateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesDuplicateInput) UnmarshalJSON(b []byte) error {
	type plain PoliciesDuplicateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "key"})
}

type PoliciesDuplicateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesDuplicateResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesDuplicateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesDuplicateResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesDuplicateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type PoliciesVersionsResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesVersionsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PoliciesVersionsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesVersionsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PoliciesVersionsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type PoliciesVersionsResponse struct {
	Object     string                             `json:"object"`
	Data       []PoliciesVersionsResponseDataItem `json:"data"`
	NextCursor any                                `json:"next_cursor"`
	Extra      map[string]json.RawMessage         `json:"-"`
}

func (v PoliciesVersionsResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesVersionsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesVersionsResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesVersionsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type PoliciesVersionResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesVersionResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesVersionResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesVersionResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesVersionResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type PoliciesSignalsResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v PoliciesSignalsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain PoliciesSignalsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesSignalsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain PoliciesSignalsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type PoliciesSignalsResponse struct {
	Object     string                            `json:"object"`
	Data       []PoliciesSignalsResponseDataItem `json:"data"`
	NextCursor any                               `json:"next_cursor"`
	Extra      map[string]json.RawMessage        `json:"-"`
}

func (v PoliciesSignalsResponse) MarshalJSON() ([]byte, error) {
	type plain PoliciesSignalsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *PoliciesSignalsResponse) UnmarshalJSON(b []byte) error {
	type plain PoliciesSignalsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type AccessListsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Value  Optional[string]           `json:"value,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AccessListsListInput) MarshalJSON() ([]byte, error) {
	type plain AccessListsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsListInput) UnmarshalJSON(b []byte) error {
	type plain AccessListsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit", "value"})
}

type AccessListsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AccessListsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain AccessListsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain AccessListsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type AccessListsListResponse struct {
	Object     string                            `json:"object"`
	Data       []AccessListsListResponseDataItem `json:"data"`
	NextCursor any                               `json:"next_cursor"`
	Extra      map[string]json.RawMessage        `json:"-"`
}

func (v AccessListsListResponse) MarshalJSON() ([]byte, error) {
	type plain AccessListsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsListResponse) UnmarshalJSON(b []byte) error {
	type plain AccessListsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type AccessListsCreateInput struct {
	Kind          string                     `json:"kind"`
	Type          string                     `json:"type"`
	Value         Optional[string]           `json:"value,omitzero"`
	Values        Optional[[]string]         `json:"values,omitzero"`
	BenefitIDs    Optional[[]string]         `json:"benefit_ids,omitzero"`
	Reason        string                     `json:"reason"`
	ExpiresAt     Optional[any]              `json:"expires_at,omitzero"`
	ExpiresInDays Optional[any]              `json:"expires_in_days,omitzero"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v AccessListsCreateInput) MarshalJSON() ([]byte, error) {
	type plain AccessListsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsCreateInput) UnmarshalJSON(b []byte) error {
	type plain AccessListsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "type", "value", "values", "benefit_ids", "reason", "expires_at", "expires_in_days"})
}

type AccessListsCreateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AccessListsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain AccessListsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain AccessListsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type AccessListsGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AccessListsGetResponse) MarshalJSON() ([]byte, error) {
	type plain AccessListsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsGetResponse) UnmarshalJSON(b []byte) error {
	type plain AccessListsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type AccessListsDeleteResponse struct {
	Object  string                     `json:"object"`
	ID      any                        `json:"id"`
	Deleted bool                       `json:"deleted"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v AccessListsDeleteResponse) MarshalJSON() ([]byte, error) {
	type plain AccessListsDeleteResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsDeleteResponse) UnmarshalJSON(b []byte) error {
	type plain AccessListsDeleteResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "deleted"})
}

type AccessListsUpdateInput struct {
	Kind          Optional[string]           `json:"kind,omitzero"`
	BenefitIDs    Optional[[]string]         `json:"benefit_ids,omitzero"`
	Reason        Optional[string]           `json:"reason,omitzero"`
	ExpiresAt     Optional[any]              `json:"expires_at,omitzero"`
	ExpiresInDays Optional[any]              `json:"expires_in_days,omitzero"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v AccessListsUpdateInput) MarshalJSON() ([]byte, error) {
	type plain AccessListsUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsUpdateInput) UnmarshalJSON(b []byte) error {
	type plain AccessListsUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "benefit_ids", "reason", "expires_at", "expires_in_days"})
}

type AccessListsUpdateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AccessListsUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain AccessListsUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccessListsUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain AccessListsUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type IdentifiersLookupInput struct {
	Value string                     `json:"value"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupInput) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupInput) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"value"})
}

type IdentifiersLookupResponseMatchesItemVariant0 struct {
	Object  string                     `json:"object"`
	Kind    string                     `json:"kind"`
	ID      string                     `json:"id"`
	Result  any                        `json:"result"`
	Benefit any                        `json:"benefit"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant0) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant0) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "result", "benefit"})
}

type IdentifiersLookupResponseMatchesItemVariant1 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	Count  int64                      `json:"count"`
	Capped bool                       `json:"capped"`
	Days   float64                    `json:"days"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant1) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant1
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant1) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant1
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "count", "capped", "days"})
}

type IdentifiersLookupResponseMatchesItemVariant2 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	Count  int64                      `json:"count"`
	Capped bool                       `json:"capped"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant2) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant2
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant2) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant2
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "count", "capped"})
}

type IdentifiersLookupResponseMatchesItemVariant3 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	ID     string                     `json:"id"`
	Key    string                     `json:"key"`
	Name   string                     `json:"name"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant3) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant3
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant3) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant3
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "key", "name"})
}

type IdentifiersLookupResponseMatchesItemVariant4 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	ID     string                     `json:"id"`
	Key    string                     `json:"key"`
	Name   string                     `json:"name"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant4) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant4
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant4) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant4
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "key", "name"})
}

type IdentifiersLookupResponseMatchesItemVariant5 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	ID     string                     `json:"id"`
	Key    string                     `json:"key"`
	Name   string                     `json:"name"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant5) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant5
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant5) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant5
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "key", "name"})
}

type IdentifiersLookupResponseMatchesItemVariant6 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	ID     string                     `json:"id"`
	List   string                     `json:"list"`
	Type   string                     `json:"type"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant6) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant6
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant6) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant6
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "list", "type"})
}

type IdentifiersLookupResponseMatchesItemVariant7 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	ID     string                     `json:"id"`
	Name   string                     `json:"name"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant7) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant7
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant7) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant7
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "name"})
}

type IdentifiersLookupResponseMatchesItemVariant8 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	ID     string                     `json:"id"`
	URL    string                     `json:"url"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant8) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant8
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant8) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant8
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "url"})
}

type IdentifiersLookupResponseMatchesItemVariant9 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	ID     string                     `json:"id"`
	Email  string                     `json:"email"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant9) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant9
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant9) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant9
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "email"})
}

type IdentifiersLookupResponseMatchesItemVariant10 struct {
	Object   string                     `json:"object"`
	Kind     string                     `json:"kind"`
	ID       string                     `json:"id"`
	FileName any                        `json:"file_name"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant10) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant10
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant10) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant10
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "file_name"})
}

type IdentifiersLookupResponseMatchesItemVariant11 struct {
	Object string                     `json:"object"`
	Kind   string                     `json:"kind"`
	ID     string                     `json:"id"`
	Key    string                     `json:"key"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponseMatchesItemVariant11) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponseMatchesItemVariant11
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponseMatchesItemVariant11) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponseMatchesItemVariant11
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "id", "key"})
}

type IdentifiersLookupResponse struct {
	Object  string                     `json:"object"`
	Type    any                        `json:"type"`
	Matches []any                      `json:"matches"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v IdentifiersLookupResponse) MarshalJSON() ([]byte, error) {
	type plain IdentifiersLookupResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *IdentifiersLookupResponse) UnmarshalJSON(b []byte) error {
	type plain IdentifiersLookupResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "type", "matches"})
}

type CustomIdentifiersListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v CustomIdentifiersListInput) MarshalJSON() ([]byte, error) {
	type plain CustomIdentifiersListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *CustomIdentifiersListInput) UnmarshalJSON(b []byte) error {
	type plain CustomIdentifiersListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type CustomIdentifiersListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v CustomIdentifiersListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain CustomIdentifiersListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *CustomIdentifiersListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain CustomIdentifiersListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type CustomIdentifiersListResponse struct {
	Object     string                                  `json:"object"`
	Data       []CustomIdentifiersListResponseDataItem `json:"data"`
	NextCursor any                                     `json:"next_cursor"`
	Extra      map[string]json.RawMessage              `json:"-"`
}

func (v CustomIdentifiersListResponse) MarshalJSON() ([]byte, error) {
	type plain CustomIdentifiersListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *CustomIdentifiersListResponse) UnmarshalJSON(b []byte) error {
	type plain CustomIdentifiersListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type CustomIdentifiersCreateInput struct {
	Key         string                     `json:"key"`
	KeyVersion  string                     `json:"key_version"`
	Description string                     `json:"description"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v CustomIdentifiersCreateInput) MarshalJSON() ([]byte, error) {
	type plain CustomIdentifiersCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *CustomIdentifiersCreateInput) UnmarshalJSON(b []byte) error {
	type plain CustomIdentifiersCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"key", "key_version", "description"})
}

type CustomIdentifiersCreateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v CustomIdentifiersCreateResponse) MarshalJSON() ([]byte, error) {
	type plain CustomIdentifiersCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *CustomIdentifiersCreateResponse) UnmarshalJSON(b []byte) error {
	type plain CustomIdentifiersCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type CustomIdentifiersGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v CustomIdentifiersGetResponse) MarshalJSON() ([]byte, error) {
	type plain CustomIdentifiersGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *CustomIdentifiersGetResponse) UnmarshalJSON(b []byte) error {
	type plain CustomIdentifiersGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type CustomIdentifiersDeleteResponse struct {
	Object  string                     `json:"object"`
	ID      any                        `json:"id"`
	Deleted bool                       `json:"deleted"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v CustomIdentifiersDeleteResponse) MarshalJSON() ([]byte, error) {
	type plain CustomIdentifiersDeleteResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *CustomIdentifiersDeleteResponse) UnmarshalJSON(b []byte) error {
	type plain CustomIdentifiersDeleteResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "deleted"})
}

type ImportsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsListInput) MarshalJSON() ([]byte, error) {
	type plain ImportsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsListInput) UnmarshalJSON(b []byte) error {
	type plain ImportsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type ImportsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain ImportsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain ImportsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ImportsListResponse struct {
	Object     string                        `json:"object"`
	Data       []ImportsListResponseDataItem `json:"data"`
	NextCursor any                           `json:"next_cursor"`
	Extra      map[string]json.RawMessage    `json:"-"`
}

func (v ImportsListResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsListResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type ImportsCreateInputRowsItemSubject struct {
	Email      string                     `json:"email"`
	ExternalID Optional[string]           `json:"external_id,omitzero"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v ImportsCreateInputRowsItemSubject) MarshalJSON() ([]byte, error) {
	type plain ImportsCreateInputRowsItemSubject
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCreateInputRowsItemSubject) UnmarshalJSON(b []byte) error {
	type plain ImportsCreateInputRowsItemSubject
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email", "external_id"})
}

type ImportsCreateInputRowsItemIdentifiersPayment struct {
	Provider    string                     `json:"provider"`
	Scope       string                     `json:"scope"`
	Fingerprint string                     `json:"fingerprint"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v ImportsCreateInputRowsItemIdentifiersPayment) MarshalJSON() ([]byte, error) {
	type plain ImportsCreateInputRowsItemIdentifiersPayment
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCreateInputRowsItemIdentifiersPayment) UnmarshalJSON(b []byte) error {
	type plain ImportsCreateInputRowsItemIdentifiersPayment
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"provider", "scope", "fingerprint"})
}

type ImportsCreateInputRowsItemIdentifiers struct {
	Payment Optional[ImportsCreateInputRowsItemIdentifiersPayment] `json:"payment,omitzero"`
	Extra   map[string]json.RawMessage                             `json:"-"`
}

func (v ImportsCreateInputRowsItemIdentifiers) MarshalJSON() ([]byte, error) {
	type plain ImportsCreateInputRowsItemIdentifiers
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCreateInputRowsItemIdentifiers) UnmarshalJSON(b []byte) error {
	type plain ImportsCreateInputRowsItemIdentifiers
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"payment"})
}

type ImportsCreateInputRowsItem struct {
	Benefit     Optional[string]                                `json:"benefit,omitzero"`
	Subject     ImportsCreateInputRowsItemSubject               `json:"subject"`
	GrantedOn   Optional[string]                                `json:"granted_on,omitzero"`
	GrantedTime Optional[string]                                `json:"granted_time,omitzero"`
	OccurredAt  Optional[string]                                `json:"occurred_at,omitzero"`
	ClaimID     Optional[string]                                `json:"claim_id,omitzero"`
	Identifiers Optional[ImportsCreateInputRowsItemIdentifiers] `json:"identifiers,omitzero"`
	Extra       map[string]json.RawMessage                      `json:"-"`
}

func (v ImportsCreateInputRowsItem) MarshalJSON() ([]byte, error) {
	type plain ImportsCreateInputRowsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCreateInputRowsItem) UnmarshalJSON(b []byte) error {
	type plain ImportsCreateInputRowsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"benefit", "subject", "granted_on", "granted_time", "occurred_at", "claim_id", "identifiers"})
}

type ImportsCreateInput struct {
	Rows     []ImportsCreateInputRowsItem `json:"rows"`
	Benefit  Optional[string]             `json:"benefit,omitzero"`
	TimeZone Optional[string]             `json:"time_zone,omitzero"`
	FileName Optional[string]             `json:"file_name,omitzero"`
	Confirm  Optional[bool]               `json:"confirm,omitzero"`
	Extra    map[string]json.RawMessage   `json:"-"`
}

func (v ImportsCreateInput) MarshalJSON() ([]byte, error) {
	type plain ImportsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCreateInput) UnmarshalJSON(b []byte) error {
	type plain ImportsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"rows", "benefit", "time_zone", "file_name", "confirm"})
}

type ImportsCreateResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ImportsCSVInputPayment struct {
	Provider string                     `json:"provider"`
	Scope    string                     `json:"scope"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v ImportsCSVInputPayment) MarshalJSON() ([]byte, error) {
	type plain ImportsCSVInputPayment
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCSVInputPayment) UnmarshalJSON(b []byte) error {
	type plain ImportsCSVInputPayment
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"provider", "scope"})
}

type ImportsCSVInput struct {
	Benefit  Optional[string]                 `json:"benefit,omitzero"`
	TimeZone Optional[string]                 `json:"time_zone,omitzero"`
	FileName Optional[string]                 `json:"file_name,omitzero"`
	Rows     Optional[int64]                  `json:"rows,omitzero"`
	Payment  Optional[ImportsCSVInputPayment] `json:"payment,omitzero"`
	Confirm  Optional[bool]                   `json:"confirm,omitzero"`
	Extra    map[string]json.RawMessage       `json:"-"`
}

func (v ImportsCSVInput) MarshalJSON() ([]byte, error) {
	type plain ImportsCSVInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCSVInput) UnmarshalJSON(b []byte) error {
	type plain ImportsCSVInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"benefit", "time_zone", "file_name", "rows", "payment", "confirm"})
}

type ImportsCSVResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsCSVResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsCSVResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsCSVResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsCSVResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ImportsGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsGetResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsGetResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ImportsProblemsInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsProblemsInput) MarshalJSON() ([]byte, error) {
	type plain ImportsProblemsInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsProblemsInput) UnmarshalJSON(b []byte) error {
	type plain ImportsProblemsInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type ImportsProblemsResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsProblemsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain ImportsProblemsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsProblemsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain ImportsProblemsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ImportsProblemsResponse struct {
	Object     string                            `json:"object"`
	Data       []ImportsProblemsResponseDataItem `json:"data"`
	NextCursor any                               `json:"next_cursor"`
	Extra      map[string]json.RawMessage        `json:"-"`
}

func (v ImportsProblemsResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsProblemsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsProblemsResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsProblemsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type ImportsProblemsDownloadResponse struct {
	Object string                     `json:"object"`
	URL    string                     `json:"url"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsProblemsDownloadResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsProblemsDownloadResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsProblemsDownloadResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsProblemsDownloadResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "url"})
}

type ImportsConfirmResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsConfirmResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsConfirmResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsConfirmResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsConfirmResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ImportsDiscardResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsDiscardResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsDiscardResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsDiscardResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsDiscardResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ImportsFixRowInputCells struct {
	Email              Optional[string]           `json:"email,omitzero"`
	Benefit            Optional[string]           `json:"benefit,omitzero"`
	GrantedOn          Optional[string]           `json:"granted_on,omitzero"`
	GrantedTime        Optional[string]           `json:"granted_time,omitzero"`
	OccurredAt         Optional[string]           `json:"occurred_at,omitzero"`
	UserID             Optional[string]           `json:"user_id,omitzero"`
	ClaimID            Optional[string]           `json:"claim_id,omitzero"`
	PaymentFingerprint Optional[string]           `json:"payment_fingerprint,omitzero"`
	Extra              map[string]json.RawMessage `json:"-"`
}

func (v ImportsFixRowInputCells) MarshalJSON() ([]byte, error) {
	type plain ImportsFixRowInputCells
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsFixRowInputCells) UnmarshalJSON(b []byte) error {
	type plain ImportsFixRowInputCells
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"email", "benefit", "granted_on", "granted_time", "occurred_at", "user_id", "claim_id", "payment_fingerprint"})
}

type ImportsFixRowInput struct {
	Row   int64                      `json:"row"`
	Cells ImportsFixRowInputCells    `json:"cells"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v ImportsFixRowInput) MarshalJSON() ([]byte, error) {
	type plain ImportsFixRowInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsFixRowInput) UnmarshalJSON(b []byte) error {
	type plain ImportsFixRowInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"row", "cells"})
}

type ImportsFixRowResponse struct {
	Object string                     `json:"object"`
	Import string                     `json:"import"`
	Row    int64                      `json:"row"`
	State  string                     `json:"state"`
	Kind   any                        `json:"kind"`
	Field  any                        `json:"field"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsFixRowResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsFixRowResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsFixRowResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsFixRowResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "import", "row", "state", "kind", "field"})
}

type ImportsSkipRowInput struct {
	Row   int64                      `json:"row"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v ImportsSkipRowInput) MarshalJSON() ([]byte, error) {
	type plain ImportsSkipRowInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsSkipRowInput) UnmarshalJSON(b []byte) error {
	type plain ImportsSkipRowInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"row"})
}

type ImportsSkipRowResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ImportsSkipRowResponse) MarshalJSON() ([]byte, error) {
	type plain ImportsSkipRowResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ImportsSkipRowResponse) UnmarshalJSON(b []byte) error {
	type plain ImportsSkipRowResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ExportsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ExportsListInput) MarshalJSON() ([]byte, error) {
	type plain ExportsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *ExportsListInput) UnmarshalJSON(b []byte) error {
	type plain ExportsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type ExportsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ExportsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain ExportsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *ExportsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain ExportsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ExportsListResponse struct {
	Object     string                        `json:"object"`
	Data       []ExportsListResponseDataItem `json:"data"`
	NextCursor any                           `json:"next_cursor"`
	Extra      map[string]json.RawMessage    `json:"-"`
}

func (v ExportsListResponse) MarshalJSON() ([]byte, error) {
	type plain ExportsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ExportsListResponse) UnmarshalJSON(b []byte) error {
	type plain ExportsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type ExportsCreateInput struct {
	Days            Optional[int64]            `json:"days,omitzero"`
	IncludeEvidence Optional[bool]             `json:"include_evidence,omitzero"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v ExportsCreateInput) MarshalJSON() ([]byte, error) {
	type plain ExportsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *ExportsCreateInput) UnmarshalJSON(b []byte) error {
	type plain ExportsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"days", "include_evidence"})
}

type ExportsCreateResponse struct {
	Object    string                     `json:"object"`
	ID        string                     `json:"id"`
	Status    string                     `json:"status"`
	StatusURL string                     `json:"status_url"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v ExportsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain ExportsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ExportsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain ExportsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "status", "status_url"})
}

type ExportsGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ExportsGetResponse) MarshalJSON() ([]byte, error) {
	type plain ExportsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ExportsGetResponse) UnmarshalJSON(b []byte) error {
	type plain ExportsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type ExportsDownloadResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v ExportsDownloadResponse) MarshalJSON() ([]byte, error) {
	type plain ExportsDownloadResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *ExportsDownloadResponse) UnmarshalJSON(b []byte) error {
	type plain ExportsDownloadResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type WorkspaceExportsCreateInput struct {
	Password   string                     `json:"password"`
	Encryption Optional[string]           `json:"encryption,omitzero"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceExportsCreateInput) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsCreateInput) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"password", "encryption"})
}

type WorkspaceExportsCreateResponseExportVariant0 struct {
	ID             string                     `json:"id"`
	State          string                     `json:"state"`
	RequestedByYou bool                       `json:"requested_by_you"`
	RequestedAt    any                        `json:"requested_at"`
	CompletedAt    any                        `json:"completed_at"`
	ExpiresAt      any                        `json:"expires_at"`
	Parts          any                        `json:"parts"`
	Bytes          any                        `json:"bytes"`
	CanDelete      bool                       `json:"can_delete"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceExportsCreateResponseExportVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsCreateResponseExportVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsCreateResponseExportVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsCreateResponseExportVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "state", "requested_by_you", "requested_at", "completed_at", "expires_at", "parts", "bytes", "can_delete"})
}

type WorkspaceExportsCreateResponse struct {
	Object        string                     `json:"object"`
	Export        any                        `json:"export"`
	NextAllowedAt any                        `json:"next_allowed_at"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceExportsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "export", "next_allowed_at"})
}

type WorkspaceExportsStatusResponseExportVariant0 struct {
	ID             string                     `json:"id"`
	State          string                     `json:"state"`
	RequestedByYou bool                       `json:"requested_by_you"`
	RequestedAt    any                        `json:"requested_at"`
	CompletedAt    any                        `json:"completed_at"`
	ExpiresAt      any                        `json:"expires_at"`
	Parts          any                        `json:"parts"`
	Bytes          any                        `json:"bytes"`
	CanDelete      bool                       `json:"can_delete"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceExportsStatusResponseExportVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsStatusResponseExportVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsStatusResponseExportVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsStatusResponseExportVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "state", "requested_by_you", "requested_at", "completed_at", "expires_at", "parts", "bytes", "can_delete"})
}

type WorkspaceExportsStatusResponse struct {
	Object        string                     `json:"object"`
	Export        any                        `json:"export"`
	NextAllowedAt any                        `json:"next_allowed_at"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceExportsStatusResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsStatusResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsStatusResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsStatusResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "export", "next_allowed_at"})
}

type WorkspaceExportsDownloadResponsePartsItem struct {
	Part     int64                      `json:"part"`
	Bytes    int64                      `json:"bytes"`
	FileName string                     `json:"file_name"`
	URL      string                     `json:"url"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceExportsDownloadResponsePartsItem) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsDownloadResponsePartsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsDownloadResponsePartsItem) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsDownloadResponsePartsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"part", "bytes", "file_name", "url"})
}

type WorkspaceExportsDownloadResponse struct {
	Object        string                                      `json:"object"`
	ID            string                                      `json:"id"`
	ExpiresAt     any                                         `json:"expires_at"`
	LinksExpireAt any                                         `json:"links_expire_at"`
	Parts         []WorkspaceExportsDownloadResponsePartsItem `json:"parts"`
	Extra         map[string]json.RawMessage                  `json:"-"`
}

func (v WorkspaceExportsDownloadResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsDownloadResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsDownloadResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsDownloadResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "expires_at", "links_expire_at", "parts"})
}

type WorkspaceExportsDeleteResponseExportVariant0 struct {
	ID             string                     `json:"id"`
	State          string                     `json:"state"`
	RequestedByYou bool                       `json:"requested_by_you"`
	RequestedAt    any                        `json:"requested_at"`
	CompletedAt    any                        `json:"completed_at"`
	ExpiresAt      any                        `json:"expires_at"`
	Parts          any                        `json:"parts"`
	Bytes          any                        `json:"bytes"`
	CanDelete      bool                       `json:"can_delete"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceExportsDeleteResponseExportVariant0) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsDeleteResponseExportVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsDeleteResponseExportVariant0) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsDeleteResponseExportVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "state", "requested_by_you", "requested_at", "completed_at", "expires_at", "parts", "bytes", "can_delete"})
}

type WorkspaceExportsDeleteResponse struct {
	Object        string                     `json:"object"`
	Export        any                        `json:"export"`
	NextAllowedAt any                        `json:"next_allowed_at"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v WorkspaceExportsDeleteResponse) MarshalJSON() ([]byte, error) {
	type plain WorkspaceExportsDeleteResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WorkspaceExportsDeleteResponse) UnmarshalJSON(b []byte) error {
	type plain WorkspaceExportsDeleteResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "export", "next_allowed_at"})
}

type DeletionsDryRunInputSelectorVariant0 struct {
	Kind  string                     `json:"kind"`
	User  string                     `json:"user"`
	Email bool                       `json:"email"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v DeletionsDryRunInputSelectorVariant0) MarshalJSON() ([]byte, error) {
	type plain DeletionsDryRunInputSelectorVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsDryRunInputSelectorVariant0) UnmarshalJSON(b []byte) error {
	type plain DeletionsDryRunInputSelectorVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "user", "email"})
}

type DeletionsDryRunInputSelectorVariant1 struct {
	Kind  string                     `json:"kind"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v DeletionsDryRunInputSelectorVariant1) MarshalJSON() ([]byte, error) {
	type plain DeletionsDryRunInputSelectorVariant1
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsDryRunInputSelectorVariant1) UnmarshalJSON(b []byte) error {
	type plain DeletionsDryRunInputSelectorVariant1
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind"})
}

type DeletionsDryRunInput struct {
	Selector any                        `json:"selector"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v DeletionsDryRunInput) MarshalJSON() ([]byte, error) {
	type plain DeletionsDryRunInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsDryRunInput) UnmarshalJSON(b []byte) error {
	type plain DeletionsDryRunInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"selector"})
}

type DeletionsDryRunResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v DeletionsDryRunResponse) MarshalJSON() ([]byte, error) {
	type plain DeletionsDryRunResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsDryRunResponse) UnmarshalJSON(b []byte) error {
	type plain DeletionsDryRunResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type DeletionsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v DeletionsListInput) MarshalJSON() ([]byte, error) {
	type plain DeletionsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsListInput) UnmarshalJSON(b []byte) error {
	type plain DeletionsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type DeletionsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v DeletionsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain DeletionsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain DeletionsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type DeletionsListResponse struct {
	Object     string                          `json:"object"`
	Data       []DeletionsListResponseDataItem `json:"data"`
	NextCursor any                             `json:"next_cursor"`
	Extra      map[string]json.RawMessage      `json:"-"`
}

func (v DeletionsListResponse) MarshalJSON() ([]byte, error) {
	type plain DeletionsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsListResponse) UnmarshalJSON(b []byte) error {
	type plain DeletionsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type DeletionsCreateInputSelectorVariant0 struct {
	Kind  string                     `json:"kind"`
	User  string                     `json:"user"`
	Email bool                       `json:"email"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v DeletionsCreateInputSelectorVariant0) MarshalJSON() ([]byte, error) {
	type plain DeletionsCreateInputSelectorVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsCreateInputSelectorVariant0) UnmarshalJSON(b []byte) error {
	type plain DeletionsCreateInputSelectorVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "user", "email"})
}

type DeletionsCreateInputSelectorVariant1 struct {
	Kind  string                     `json:"kind"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v DeletionsCreateInputSelectorVariant1) MarshalJSON() ([]byte, error) {
	type plain DeletionsCreateInputSelectorVariant1
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsCreateInputSelectorVariant1) UnmarshalJSON(b []byte) error {
	type plain DeletionsCreateInputSelectorVariant1
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind"})
}

type DeletionsCreateInput struct {
	Selector any                        `json:"selector"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v DeletionsCreateInput) MarshalJSON() ([]byte, error) {
	type plain DeletionsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsCreateInput) UnmarshalJSON(b []byte) error {
	type plain DeletionsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"selector"})
}

type DeletionsCreateResponse struct {
	Object    string                     `json:"object"`
	ID        string                     `json:"id"`
	Status    string                     `json:"status"`
	StatusURL string                     `json:"status_url"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v DeletionsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain DeletionsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain DeletionsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "status", "status_url"})
}

type DeletionsGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v DeletionsGetResponse) MarshalJSON() ([]byte, error) {
	type plain DeletionsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *DeletionsGetResponse) UnmarshalJSON(b []byte) error {
	type plain DeletionsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type AnalyticsListInput struct {
	Days  Optional[int64]            `json:"days,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v AnalyticsListInput) MarshalJSON() ([]byte, error) {
	type plain AnalyticsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsListInput) UnmarshalJSON(b []byte) error {
	type plain AnalyticsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"days"})
}

type AnalyticsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AnalyticsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain AnalyticsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain AnalyticsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type AnalyticsListResponse struct {
	Object     string                          `json:"object"`
	Data       []AnalyticsListResponseDataItem `json:"data"`
	NextCursor any                             `json:"next_cursor"`
	Extra      map[string]json.RawMessage      `json:"-"`
}

func (v AnalyticsListResponse) MarshalJSON() ([]byte, error) {
	type plain AnalyticsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsListResponse) UnmarshalJSON(b []byte) error {
	type plain AnalyticsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type OverviewGetResponseEvaluations struct {
	Total                  int64                      `json:"total"`
	Allowed                int64                      `json:"allowed"`
	Denied                 int64                      `json:"denied"`
	WaitingForConfirmation int64                      `json:"waiting_for_confirmation"`
	Extra                  map[string]json.RawMessage `json:"-"`
}

func (v OverviewGetResponseEvaluations) MarshalJSON() ([]byte, error) {
	type plain OverviewGetResponseEvaluations
	return marshalExtras(plain(v), v.Extra)
}
func (v *OverviewGetResponseEvaluations) UnmarshalJSON(b []byte) error {
	type plain OverviewGetResponseEvaluations
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"total", "allowed", "denied", "waiting_for_confirmation"})
}

type OverviewGetResponsePrevious struct {
	Total                  int64                      `json:"total"`
	Allowed                int64                      `json:"allowed"`
	Denied                 int64                      `json:"denied"`
	WaitingForConfirmation int64                      `json:"waiting_for_confirmation"`
	Extra                  map[string]json.RawMessage `json:"-"`
}

func (v OverviewGetResponsePrevious) MarshalJSON() ([]byte, error) {
	type plain OverviewGetResponsePrevious
	return marshalExtras(plain(v), v.Extra)
}
func (v *OverviewGetResponsePrevious) UnmarshalJSON(b []byte) error {
	type plain OverviewGetResponsePrevious
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"total", "allowed", "denied", "waiting_for_confirmation"})
}

type OverviewGetResponseAPIRequests struct {
	Total  int64                      `json:"total"`
	Errors int64                      `json:"errors"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v OverviewGetResponseAPIRequests) MarshalJSON() ([]byte, error) {
	type plain OverviewGetResponseAPIRequests
	return marshalExtras(plain(v), v.Extra)
}
func (v *OverviewGetResponseAPIRequests) UnmarshalJSON(b []byte) error {
	type plain OverviewGetResponseAPIRequests
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"total", "errors"})
}

type OverviewGetResponseLatestEvaluationsItemReasonVariant0 struct {
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v OverviewGetResponseLatestEvaluationsItemReasonVariant0) MarshalJSON() ([]byte, error) {
	type plain OverviewGetResponseLatestEvaluationsItemReasonVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *OverviewGetResponseLatestEvaluationsItemReasonVariant0) UnmarshalJSON(b []byte) error {
	type plain OverviewGetResponseLatestEvaluationsItemReasonVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code", "message"})
}

type OverviewGetResponseLatestEvaluationsItemOutcome struct {
	State any                        `json:"state"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v OverviewGetResponseLatestEvaluationsItemOutcome) MarshalJSON() ([]byte, error) {
	type plain OverviewGetResponseLatestEvaluationsItemOutcome
	return marshalExtras(plain(v), v.Extra)
}
func (v *OverviewGetResponseLatestEvaluationsItemOutcome) UnmarshalJSON(b []byte) error {
	type plain OverviewGetResponseLatestEvaluationsItemOutcome
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state"})
}

type OverviewGetResponseLatestEvaluationsItem struct {
	Object         string                                          `json:"object"`
	ID             string                                          `json:"id"`
	Benefit        string                                          `json:"benefit"`
	ClaimID        string                                          `json:"claim_id"`
	Result         string                                          `json:"result"`
	Mode           string                                          `json:"mode"`
	ObservedResult any                                             `json:"observed_result"`
	Reason         any                                             `json:"reason"`
	Outcome        OverviewGetResponseLatestEvaluationsItemOutcome `json:"outcome"`
	PolicyVersion  any                                             `json:"policy_version"`
	BrowserData    any                                             `json:"browser_data"`
	CreatedAt      any                                             `json:"created_at"`
	Extra          map[string]json.RawMessage                      `json:"-"`
}

func (v OverviewGetResponseLatestEvaluationsItem) MarshalJSON() ([]byte, error) {
	type plain OverviewGetResponseLatestEvaluationsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *OverviewGetResponseLatestEvaluationsItem) UnmarshalJSON(b []byte) error {
	type plain OverviewGetResponseLatestEvaluationsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "benefit", "claim_id", "result", "mode", "observed_result", "reason", "outcome", "policy_version", "browser_data", "created_at"})
}

type OverviewGetResponse struct {
	Object            string                                     `json:"object"`
	Hours             float64                                    `json:"hours"`
	Evaluations       OverviewGetResponseEvaluations             `json:"evaluations"`
	Previous          OverviewGetResponsePrevious                `json:"previous"`
	APIRequests       OverviewGetResponseAPIRequests             `json:"api_requests"`
	LastEvaluationAt  any                                        `json:"last_evaluation_at"`
	BrowserIssue      any                                        `json:"browser_issue"`
	LatestEvaluations []OverviewGetResponseLatestEvaluationsItem `json:"latest_evaluations"`
	Extra             map[string]json.RawMessage                 `json:"-"`
}

func (v OverviewGetResponse) MarshalJSON() ([]byte, error) {
	type plain OverviewGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *OverviewGetResponse) UnmarshalJSON(b []byte) error {
	type plain OverviewGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "hours", "evaluations", "previous", "api_requests", "last_evaluation_at", "browser_issue", "latest_evaluations"})
}

type AnalyticsValueInput struct {
	Days     Optional[int64]            `json:"days,omitzero"`
	Currency Optional[string]           `json:"currency,omitzero"`
	Benefit  Optional[string]           `json:"benefit,omitzero"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v AnalyticsValueInput) MarshalJSON() ([]byte, error) {
	type plain AnalyticsValueInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsValueInput) UnmarshalJSON(b []byte) error {
	type plain AnalyticsValueInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"days", "currency", "benefit"})
}

type AnalyticsValueResponseBenefitsItemValueVariant0 struct {
	Amount   float64                    `json:"amount"`
	Currency string                     `json:"currency"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v AnalyticsValueResponseBenefitsItemValueVariant0) MarshalJSON() ([]byte, error) {
	type plain AnalyticsValueResponseBenefitsItemValueVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsValueResponseBenefitsItemValueVariant0) UnmarshalJSON(b []byte) error {
	type plain AnalyticsValueResponseBenefitsItemValueVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"amount", "currency"})
}

type AnalyticsValueResponseBenefitsItem struct {
	Benefit   any                        `json:"benefit"`
	Value     any                        `json:"value"`
	Delivered float64                    `json:"delivered"`
	Saved     float64                    `json:"saved"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v AnalyticsValueResponseBenefitsItem) MarshalJSON() ([]byte, error) {
	type plain AnalyticsValueResponseBenefitsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsValueResponseBenefitsItem) UnmarshalJSON(b []byte) error {
	type plain AnalyticsValueResponseBenefitsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"benefit", "value", "delivered", "saved"})
}

type AnalyticsValueResponse struct {
	Object    string                               `json:"object"`
	Currency  string                               `json:"currency"`
	From      string                               `json:"from"`
	To        string                               `json:"to"`
	Delivered float64                              `json:"delivered"`
	Saved     float64                              `json:"saved"`
	Benefits  []AnalyticsValueResponseBenefitsItem `json:"benefits"`
	Extra     map[string]json.RawMessage           `json:"-"`
}

func (v AnalyticsValueResponse) MarshalJSON() ([]byte, error) {
	type plain AnalyticsValueResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsValueResponse) UnmarshalJSON(b []byte) error {
	type plain AnalyticsValueResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "currency", "from", "to", "delivered", "saved", "benefits"})
}

type AnalyticsDenialReasonsInput struct {
	Days    Optional[int64]            `json:"days,omitzero"`
	Benefit Optional[string]           `json:"benefit,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v AnalyticsDenialReasonsInput) MarshalJSON() ([]byte, error) {
	type plain AnalyticsDenialReasonsInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsDenialReasonsInput) UnmarshalJSON(b []byte) error {
	type plain AnalyticsDenialReasonsInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"days", "benefit"})
}

type AnalyticsDenialReasonsResponseReasonsItem struct {
	Code  string                     `json:"code"`
	Label string                     `json:"label"`
	Count int64                      `json:"count"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v AnalyticsDenialReasonsResponseReasonsItem) MarshalJSON() ([]byte, error) {
	type plain AnalyticsDenialReasonsResponseReasonsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsDenialReasonsResponseReasonsItem) UnmarshalJSON(b []byte) error {
	type plain AnalyticsDenialReasonsResponseReasonsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code", "label", "count"})
}

type AnalyticsDenialReasonsResponse struct {
	Object  string                                      `json:"object"`
	Benefit any                                         `json:"benefit"`
	From    string                                      `json:"from"`
	To      string                                      `json:"to"`
	Denied  int64                                       `json:"denied"`
	Reasons []AnalyticsDenialReasonsResponseReasonsItem `json:"reasons"`
	Extra   map[string]json.RawMessage                  `json:"-"`
}

func (v AnalyticsDenialReasonsResponse) MarshalJSON() ([]byte, error) {
	type plain AnalyticsDenialReasonsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AnalyticsDenialReasonsResponse) UnmarshalJSON(b []byte) error {
	type plain AnalyticsDenialReasonsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "benefit", "from", "to", "denied", "reasons"})
}

type RequestLogsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Status Optional[string]           `json:"status,omitzero"`
	KeyID  Optional[string]           `json:"key_id,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v RequestLogsListInput) MarshalJSON() ([]byte, error) {
	type plain RequestLogsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *RequestLogsListInput) UnmarshalJSON(b []byte) error {
	type plain RequestLogsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit", "status", "key_id"})
}

type RequestLogsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v RequestLogsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain RequestLogsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *RequestLogsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain RequestLogsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type RequestLogsListResponse struct {
	Object     string                            `json:"object"`
	Data       []RequestLogsListResponseDataItem `json:"data"`
	NextCursor any                               `json:"next_cursor"`
	Extra      map[string]json.RawMessage        `json:"-"`
}

func (v RequestLogsListResponse) MarshalJSON() ([]byte, error) {
	type plain RequestLogsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *RequestLogsListResponse) UnmarshalJSON(b []byte) error {
	type plain RequestLogsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type RequestLogsGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v RequestLogsGetResponse) MarshalJSON() ([]byte, error) {
	type plain RequestLogsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *RequestLogsGetResponse) UnmarshalJSON(b []byte) error {
	type plain RequestLogsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type AuditLogsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AuditLogsListInput) MarshalJSON() ([]byte, error) {
	type plain AuditLogsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuditLogsListInput) UnmarshalJSON(b []byte) error {
	type plain AuditLogsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type AuditLogsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AuditLogsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain AuditLogsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuditLogsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain AuditLogsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type AuditLogsListResponse struct {
	Object     string                          `json:"object"`
	Data       []AuditLogsListResponseDataItem `json:"data"`
	NextCursor any                             `json:"next_cursor"`
	Extra      map[string]json.RawMessage      `json:"-"`
}

func (v AuditLogsListResponse) MarshalJSON() ([]byte, error) {
	type plain AuditLogsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuditLogsListResponse) UnmarshalJSON(b []byte) error {
	type plain AuditLogsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type AuditLogsGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AuditLogsGetResponse) MarshalJSON() ([]byte, error) {
	type plain AuditLogsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AuditLogsGetResponse) UnmarshalJSON(b []byte) error {
	type plain AuditLogsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type AccountStatusGetResponsePlan struct {
	Plan       string                     `json:"plan"`
	Name       string                     `json:"name"`
	Interval   any                        `json:"interval"`
	Enterprise bool                       `json:"enterprise"`
	RenewsAt   any                        `json:"renews_at"`
	EndsAt     any                        `json:"ends_at"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v AccountStatusGetResponsePlan) MarshalJSON() ([]byte, error) {
	type plain AccountStatusGetResponsePlan
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountStatusGetResponsePlan) UnmarshalJSON(b []byte) error {
	type plain AccountStatusGetResponsePlan
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"plan", "name", "interval", "enterprise", "renews_at", "ends_at"})
}

type AccountStatusGetResponseCredits struct {
	Available       int64                      `json:"available"`
	Allowance       int64                      `json:"allowance"`
	PayAsYouGo      int64                      `json:"pay_as_you_go"`
	PerEvaluation   int64                      `json:"per_evaluation"`
	EvaluationsLeft int64                      `json:"evaluations_left"`
	State           string                     `json:"state"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v AccountStatusGetResponseCredits) MarshalJSON() ([]byte, error) {
	type plain AccountStatusGetResponseCredits
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountStatusGetResponseCredits) UnmarshalJSON(b []byte) error {
	type plain AccountStatusGetResponseCredits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"available", "allowance", "pay_as_you_go", "per_evaluation", "evaluations_left", "state"})
}

type AccountStatusGetResponseLimitsProtectedBenefits struct {
	Used  int64                      `json:"used"`
	Limit int64                      `json:"limit"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v AccountStatusGetResponseLimitsProtectedBenefits) MarshalJSON() ([]byte, error) {
	type plain AccountStatusGetResponseLimitsProtectedBenefits
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountStatusGetResponseLimitsProtectedBenefits) UnmarshalJSON(b []byte) error {
	type plain AccountStatusGetResponseLimitsProtectedBenefits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "limit"})
}

type AccountStatusGetResponseLimitsCustomPolicies struct {
	Used  int64                      `json:"used"`
	Limit int64                      `json:"limit"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v AccountStatusGetResponseLimitsCustomPolicies) MarshalJSON() ([]byte, error) {
	type plain AccountStatusGetResponseLimitsCustomPolicies
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountStatusGetResponseLimitsCustomPolicies) UnmarshalJSON(b []byte) error {
	type plain AccountStatusGetResponseLimitsCustomPolicies
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "limit"})
}

type AccountStatusGetResponseLimitsTeamSeats struct {
	Used  int64                      `json:"used"`
	Limit int64                      `json:"limit"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v AccountStatusGetResponseLimitsTeamSeats) MarshalJSON() ([]byte, error) {
	type plain AccountStatusGetResponseLimitsTeamSeats
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountStatusGetResponseLimitsTeamSeats) UnmarshalJSON(b []byte) error {
	type plain AccountStatusGetResponseLimitsTeamSeats
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "limit"})
}

type AccountStatusGetResponseLimits struct {
	ProtectedBenefits AccountStatusGetResponseLimitsProtectedBenefits `json:"protected_benefits"`
	CustomPolicies    AccountStatusGetResponseLimitsCustomPolicies    `json:"custom_policies"`
	TeamSeats         AccountStatusGetResponseLimitsTeamSeats         `json:"team_seats"`
	RetentionDays     int64                                           `json:"retention_days"`
	Rate              int64                                           `json:"rate"`
	Extra             map[string]json.RawMessage                      `json:"-"`
}

func (v AccountStatusGetResponseLimits) MarshalJSON() ([]byte, error) {
	type plain AccountStatusGetResponseLimits
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountStatusGetResponseLimits) UnmarshalJSON(b []byte) error {
	type plain AccountStatusGetResponseLimits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"protected_benefits", "custom_policies", "team_seats", "retention_days", "rate"})
}

type AccountStatusGetResponseAutoReload struct {
	Enabled bool                       `json:"enabled"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v AccountStatusGetResponseAutoReload) MarshalJSON() ([]byte, error) {
	type plain AccountStatusGetResponseAutoReload
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountStatusGetResponseAutoReload) UnmarshalJSON(b []byte) error {
	type plain AccountStatusGetResponseAutoReload
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"enabled"})
}

type AccountStatusGetResponse struct {
	Object     string                             `json:"object"`
	Workspace  string                             `json:"workspace"`
	Plan       AccountStatusGetResponsePlan       `json:"plan"`
	Credits    AccountStatusGetResponseCredits    `json:"credits"`
	Limits     AccountStatusGetResponseLimits     `json:"limits"`
	AutoReload AccountStatusGetResponseAutoReload `json:"auto_reload"`
	Extra      map[string]json.RawMessage         `json:"-"`
}

func (v AccountStatusGetResponse) MarshalJSON() ([]byte, error) {
	type plain AccountStatusGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountStatusGetResponse) UnmarshalJSON(b []byte) error {
	type plain AccountStatusGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "workspace", "plan", "credits", "limits", "auto_reload"})
}

type IntegrationStatusGetResponseAPIKeys struct {
	State    string                     `json:"state"`
	NextStep any                        `json:"next_step"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v IntegrationStatusGetResponseAPIKeys) MarshalJSON() ([]byte, error) {
	type plain IntegrationStatusGetResponseAPIKeys
	return marshalExtras(plain(v), v.Extra)
}
func (v *IntegrationStatusGetResponseAPIKeys) UnmarshalJSON(b []byte) error {
	type plain IntegrationStatusGetResponseAPIKeys
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "next_step"})
}

type IntegrationStatusGetResponseBrowserScript struct {
	State    string                     `json:"state"`
	NextStep any                        `json:"next_step"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v IntegrationStatusGetResponseBrowserScript) MarshalJSON() ([]byte, error) {
	type plain IntegrationStatusGetResponseBrowserScript
	return marshalExtras(plain(v), v.Extra)
}
func (v *IntegrationStatusGetResponseBrowserScript) UnmarshalJSON(b []byte) error {
	type plain IntegrationStatusGetResponseBrowserScript
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "next_step"})
}

type IntegrationStatusGetResponseEvaluations struct {
	State    string                     `json:"state"`
	NextStep any                        `json:"next_step"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v IntegrationStatusGetResponseEvaluations) MarshalJSON() ([]byte, error) {
	type plain IntegrationStatusGetResponseEvaluations
	return marshalExtras(plain(v), v.Extra)
}
func (v *IntegrationStatusGetResponseEvaluations) UnmarshalJSON(b []byte) error {
	type plain IntegrationStatusGetResponseEvaluations
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "next_step"})
}

type IntegrationStatusGetResponseOutcomes struct {
	State    string                     `json:"state"`
	NextStep any                        `json:"next_step"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v IntegrationStatusGetResponseOutcomes) MarshalJSON() ([]byte, error) {
	type plain IntegrationStatusGetResponseOutcomes
	return marshalExtras(plain(v), v.Extra)
}
func (v *IntegrationStatusGetResponseOutcomes) UnmarshalJSON(b []byte) error {
	type plain IntegrationStatusGetResponseOutcomes
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "next_step"})
}

type IntegrationStatusGetResponseWebhooks struct {
	State    string                     `json:"state"`
	NextStep any                        `json:"next_step"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v IntegrationStatusGetResponseWebhooks) MarshalJSON() ([]byte, error) {
	type plain IntegrationStatusGetResponseWebhooks
	return marshalExtras(plain(v), v.Extra)
}
func (v *IntegrationStatusGetResponseWebhooks) UnmarshalJSON(b []byte) error {
	type plain IntegrationStatusGetResponseWebhooks
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "next_step"})
}

type IntegrationStatusGetResponseCredits struct {
	State    string                     `json:"state"`
	NextStep any                        `json:"next_step"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v IntegrationStatusGetResponseCredits) MarshalJSON() ([]byte, error) {
	type plain IntegrationStatusGetResponseCredits
	return marshalExtras(plain(v), v.Extra)
}
func (v *IntegrationStatusGetResponseCredits) UnmarshalJSON(b []byte) error {
	type plain IntegrationStatusGetResponseCredits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "next_step"})
}

type IntegrationStatusGetResponse struct {
	Object        string                                    `json:"object"`
	Status        string                                    `json:"status"`
	APIKeys       IntegrationStatusGetResponseAPIKeys       `json:"api_keys"`
	BrowserScript IntegrationStatusGetResponseBrowserScript `json:"browser_script"`
	Evaluations   IntegrationStatusGetResponseEvaluations   `json:"evaluations"`
	Outcomes      IntegrationStatusGetResponseOutcomes      `json:"outcomes"`
	Webhooks      IntegrationStatusGetResponseWebhooks      `json:"webhooks"`
	Credits       IntegrationStatusGetResponseCredits       `json:"credits"`
	CheckedAt     any                                       `json:"checked_at"`
	Extra         map[string]json.RawMessage                `json:"-"`
}

func (v IntegrationStatusGetResponse) MarshalJSON() ([]byte, error) {
	type plain IntegrationStatusGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *IntegrationStatusGetResponse) UnmarshalJSON(b []byte) error {
	type plain IntegrationStatusGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "api_keys", "browser_script", "evaluations", "outcomes", "webhooks", "credits", "checked_at"})
}

type EventsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Wait   Optional[int64]            `json:"wait,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v EventsListInput) MarshalJSON() ([]byte, error) {
	type plain EventsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListInput) UnmarshalJSON(b []byte) error {
	type plain EventsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit", "wait"})
}

type EventsListResponseDataItemDataVariant0 struct {
	Object          string                     `json:"object"`
	State           string                     `json:"state"`
	Available       int64                      `json:"available"`
	EvaluationsLeft int64                      `json:"evaluations_left"`
	DaysLeft        any                        `json:"days_left"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant0) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant0) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "state", "available", "evaluations_left", "days_left"})
}

type EventsListResponseDataItemDataVariant1 struct {
	Object          string                     `json:"object"`
	State           string                     `json:"state"`
	Available       int64                      `json:"available"`
	EvaluationsLeft int64                      `json:"evaluations_left"`
	DaysLeft        any                        `json:"days_left"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant1) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant1
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant1) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant1
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "state", "available", "evaluations_left", "days_left"})
}

type EventsListResponseDataItemDataVariant2 struct {
	Object  string                     `json:"object"`
	Status  string                     `json:"status"`
	Purpose string                     `json:"purpose"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant2) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant2
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant2) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant2
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "purpose"})
}

type EventsListResponseDataItemDataVariant3 struct {
	Object               string                     `json:"object"`
	Status               string                     `json:"status"`
	Purpose              string                     `json:"purpose"`
	ConfirmationRequired bool                       `json:"confirmation_required"`
	Extra                map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant3) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant3
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant3) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant3
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "purpose", "confirmation_required"})
}

type EventsListResponseDataItemDataVariant4PreviousVariant0 struct {
	Plan     string                     `json:"plan"`
	Name     string                     `json:"name"`
	Interval any                        `json:"interval"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant4PreviousVariant0) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant4PreviousVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant4PreviousVariant0) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant4PreviousVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"plan", "name", "interval"})
}

type EventsListResponseDataItemDataVariant4 struct {
	Plan     string                     `json:"plan"`
	Name     string                     `json:"name"`
	Interval any                        `json:"interval"`
	Object   string                     `json:"object"`
	Previous any                        `json:"previous"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant4) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant4
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant4) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant4
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"plan", "name", "interval", "object", "previous"})
}

type EventsListResponseDataItemDataVariant5 struct {
	Object string                     `json:"object"`
	ID     string                     `json:"id"`
	Status string                     `json:"status"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant5) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant5
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant5) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant5
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "status"})
}

type EventsListResponseDataItemDataVariant6 struct {
	Object    string                     `json:"object"`
	ID        string                     `json:"id"`
	Name      string                     `json:"name"`
	Type      string                     `json:"type"`
	ExpiresAt any                        `json:"expires_at"`
	CreatedAt any                        `json:"created_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant6) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant6
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant6) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant6
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "type", "expires_at", "created_at"})
}

type EventsListResponseDataItemDataVariant7 struct {
	Object    string                     `json:"object"`
	ID        string                     `json:"id"`
	Name      string                     `json:"name"`
	Type      string                     `json:"type"`
	ExpiresAt any                        `json:"expires_at"`
	CreatedAt any                        `json:"created_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant7) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant7
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant7) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant7
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "type", "expires_at", "created_at"})
}

type EventsListResponseDataItemDataVariant8 struct {
	Object   string                     `json:"object"`
	ID       string                     `json:"id"`
	Name     any                        `json:"name"`
	Role     string                     `json:"role"`
	JoinedAt any                        `json:"joined_at"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant8) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant8
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant8) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant8
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "name", "role", "joined_at"})
}

type EventsListResponseDataItemDataVariant9OwnerVariant0 struct {
	ID    string                     `json:"id"`
	Name  any                        `json:"name"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant9OwnerVariant0) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant9OwnerVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant9OwnerVariant0) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant9OwnerVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "name"})
}

type EventsListResponseDataItemDataVariant9NewOwnerVariant0 struct {
	ID    string                     `json:"id"`
	Name  any                        `json:"name"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant9NewOwnerVariant0) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant9NewOwnerVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant9NewOwnerVariant0) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant9NewOwnerVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "name"})
}

type EventsListResponseDataItemDataVariant9 struct {
	Object   string                     `json:"object"`
	Status   string                     `json:"status"`
	Owner    any                        `json:"owner"`
	NewOwner any                        `json:"new_owner"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant9) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant9
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant9) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant9
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "owner", "new_owner"})
}

type EventsListResponseDataItemDataVariant10 struct {
	Object string                     `json:"object"`
	Day    string                     `json:"day"`
	Limit  int64                      `json:"limit"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItemDataVariant10) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItemDataVariant10
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItemDataVariant10) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItemDataVariant10
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "day", "limit"})
}

type EventsListResponseDataItem struct {
	Object     string                     `json:"object"`
	ID         string                     `json:"id"`
	Type       string                     `json:"type"`
	APIVersion string                     `json:"api_version"`
	CreatedAt  any                        `json:"created_at"`
	Workspace  string                     `json:"workspace"`
	Test       bool                       `json:"test"`
	Data       any                        `json:"data"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v EventsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain EventsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain EventsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "type", "api_version", "created_at", "workspace", "test", "data"})
}

type EventsListResponse struct {
	Object     string                       `json:"object"`
	Data       []EventsListResponseDataItem `json:"data"`
	NextCursor string                       `json:"next_cursor"`
	HasMore    bool                         `json:"has_more"`
	Extra      map[string]json.RawMessage   `json:"-"`
}

func (v EventsListResponse) MarshalJSON() ([]byte, error) {
	type plain EventsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *EventsListResponse) UnmarshalJSON(b []byte) error {
	type plain EventsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor", "has_more"})
}

type BrowserCollectionGetResponseSitesItem struct {
	Origin      string                     `json:"origin"`
	Listed      bool                       `json:"listed"`
	FirstSeenAt any                        `json:"first_seen_at"`
	LastSeenAt  any                        `json:"last_seen_at"`
	SDKVersion  any                        `json:"sdk_version"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BrowserCollectionGetResponseSitesItem) MarshalJSON() ([]byte, error) {
	type plain BrowserCollectionGetResponseSitesItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BrowserCollectionGetResponseSitesItem) UnmarshalJSON(b []byte) error {
	type plain BrowserCollectionGetResponseSitesItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"origin", "listed", "first_seen_at", "last_seen_at", "sdk_version"})
}

type BrowserCollectionGetResponseCoverage struct {
	Days            float64                    `json:"days"`
	Total           int64                      `json:"total"`
	Collected       int64                      `json:"collected"`
	ConsentDeclined int64                      `json:"consent_declined"`
	NotLoaded       int64                      `json:"not_loaded"`
	Blocked         int64                      `json:"blocked"`
	TimedOut        int64                      `json:"timed_out"`
	Failed          int64                      `json:"failed"`
	SiteNotListed   int64                      `json:"site_not_listed"`
	SDKOutdated     int64                      `json:"sdk_outdated"`
	Expired         int64                      `json:"expired"`
	NotPassedOn     int64                      `json:"not_passed_on"`
	NotSent         int64                      `json:"not_sent"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v BrowserCollectionGetResponseCoverage) MarshalJSON() ([]byte, error) {
	type plain BrowserCollectionGetResponseCoverage
	return marshalExtras(plain(v), v.Extra)
}
func (v *BrowserCollectionGetResponseCoverage) UnmarshalJSON(b []byte) error {
	type plain BrowserCollectionGetResponseCoverage
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"days", "total", "collected", "consent_declined", "not_loaded", "blocked", "timed_out", "failed", "site_not_listed", "sdk_outdated", "expired", "not_passed_on", "not_sent"})
}

type BrowserCollectionGetResponse struct {
	Object     string                                  `json:"object"`
	Origins    []string                                `json:"origins"`
	Revision   int64                                   `json:"revision"`
	Sites      []BrowserCollectionGetResponseSitesItem `json:"sites"`
	Issue      any                                     `json:"issue"`
	IssueSince any                                     `json:"issue_since"`
	Coverage   BrowserCollectionGetResponseCoverage    `json:"coverage"`
	Notices    []string                                `json:"notices"`
	Extra      map[string]json.RawMessage              `json:"-"`
}

func (v BrowserCollectionGetResponse) MarshalJSON() ([]byte, error) {
	type plain BrowserCollectionGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BrowserCollectionGetResponse) UnmarshalJSON(b []byte) error {
	type plain BrowserCollectionGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "origins", "revision", "sites", "issue", "issue_since", "coverage", "notices"})
}

type BrowserCollectionUpdateInput struct {
	Origins  []string                   `json:"origins"`
	Revision Optional[int64]            `json:"revision,omitzero"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v BrowserCollectionUpdateInput) MarshalJSON() ([]byte, error) {
	type plain BrowserCollectionUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BrowserCollectionUpdateInput) UnmarshalJSON(b []byte) error {
	type plain BrowserCollectionUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"origins", "revision"})
}

type BrowserCollectionUpdateResponseSitesItem struct {
	Origin      string                     `json:"origin"`
	Listed      bool                       `json:"listed"`
	FirstSeenAt any                        `json:"first_seen_at"`
	LastSeenAt  any                        `json:"last_seen_at"`
	SDKVersion  any                        `json:"sdk_version"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BrowserCollectionUpdateResponseSitesItem) MarshalJSON() ([]byte, error) {
	type plain BrowserCollectionUpdateResponseSitesItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BrowserCollectionUpdateResponseSitesItem) UnmarshalJSON(b []byte) error {
	type plain BrowserCollectionUpdateResponseSitesItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"origin", "listed", "first_seen_at", "last_seen_at", "sdk_version"})
}

type BrowserCollectionUpdateResponseCoverage struct {
	Days            float64                    `json:"days"`
	Total           int64                      `json:"total"`
	Collected       int64                      `json:"collected"`
	ConsentDeclined int64                      `json:"consent_declined"`
	NotLoaded       int64                      `json:"not_loaded"`
	Blocked         int64                      `json:"blocked"`
	TimedOut        int64                      `json:"timed_out"`
	Failed          int64                      `json:"failed"`
	SiteNotListed   int64                      `json:"site_not_listed"`
	SDKOutdated     int64                      `json:"sdk_outdated"`
	Expired         int64                      `json:"expired"`
	NotPassedOn     int64                      `json:"not_passed_on"`
	NotSent         int64                      `json:"not_sent"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v BrowserCollectionUpdateResponseCoverage) MarshalJSON() ([]byte, error) {
	type plain BrowserCollectionUpdateResponseCoverage
	return marshalExtras(plain(v), v.Extra)
}
func (v *BrowserCollectionUpdateResponseCoverage) UnmarshalJSON(b []byte) error {
	type plain BrowserCollectionUpdateResponseCoverage
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"days", "total", "collected", "consent_declined", "not_loaded", "blocked", "timed_out", "failed", "site_not_listed", "sdk_outdated", "expired", "not_passed_on", "not_sent"})
}

type BrowserCollectionUpdateResponse struct {
	Object     string                                     `json:"object"`
	Origins    []string                                   `json:"origins"`
	Revision   int64                                      `json:"revision"`
	Sites      []BrowserCollectionUpdateResponseSitesItem `json:"sites"`
	Issue      any                                        `json:"issue"`
	IssueSince any                                        `json:"issue_since"`
	Coverage   BrowserCollectionUpdateResponseCoverage    `json:"coverage"`
	Notices    []string                                   `json:"notices"`
	Extra      map[string]json.RawMessage                 `json:"-"`
}

func (v BrowserCollectionUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain BrowserCollectionUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BrowserCollectionUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain BrowserCollectionUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "origins", "revision", "sites", "issue", "issue_since", "coverage", "notices"})
}

type BrowserCollectionScriptResponse struct {
	Object         string                     `json:"object"`
	Script         any                        `json:"script"`
	SDKVersion     string                     `json:"sdk_version"`
	PublishableKey any                        `json:"publishable_key"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BrowserCollectionScriptResponse) MarshalJSON() ([]byte, error) {
	type plain BrowserCollectionScriptResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BrowserCollectionScriptResponse) UnmarshalJSON(b []byte) error {
	type plain BrowserCollectionScriptResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "script", "sdk_version", "publishable_key"})
}

type WebhookEndpointsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsListInput) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsListInput) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type WebhookEndpointsListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type WebhookEndpointsListResponse struct {
	Object     string                                 `json:"object"`
	Data       []WebhookEndpointsListResponseDataItem `json:"data"`
	NextCursor any                                    `json:"next_cursor"`
	Extra      map[string]json.RawMessage             `json:"-"`
}

func (v WebhookEndpointsListResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsListResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type WebhookEndpointsCreateInput struct {
	URL        string                     `json:"url"`
	Events     []string                   `json:"events"`
	BenefitIDs Optional[[]string]         `json:"benefit_ids,omitzero"`
	DeniedOnly Optional[bool]             `json:"denied_only,omitzero"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsCreateInput) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsCreateInput) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"url", "events", "benefit_ids", "denied_only"})
}

type WebhookEndpointsCreateResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	URL            string                     `json:"url"`
	Description    any                        `json:"description"`
	Events         []string                   `json:"events"`
	Benefits       []any                      `json:"benefits"`
	DeniedOnly     bool                       `json:"denied_only"`
	Status         string                     `json:"status"`
	DisabledReason any                        `json:"disabled_reason"`
	LastSuccessAt  any                        `json:"last_success_at"`
	Revision       int64                      `json:"revision"`
	CreatedAt      any                        `json:"created_at"`
	UpdatedAt      any                        `json:"updated_at"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "url", "description", "events", "benefits", "denied_only", "status", "disabled_reason", "last_success_at", "revision", "created_at", "updated_at"})
}

type WebhookEndpointsGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsGetResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsGetResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type WebhookEndpointsUpdateInput struct {
	Revision   int64                      `json:"revision"`
	URL        Optional[string]           `json:"url,omitzero"`
	Events     Optional[[]string]         `json:"events,omitzero"`
	BenefitIDs Optional[[]string]         `json:"benefit_ids,omitzero"`
	DeniedOnly Optional[bool]             `json:"denied_only,omitzero"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsUpdateInput) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsUpdateInput) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"revision", "url", "events", "benefit_ids", "denied_only"})
}

type WebhookEndpointsUpdateResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	URL            string                     `json:"url"`
	Description    any                        `json:"description"`
	Events         []string                   `json:"events"`
	Benefits       []any                      `json:"benefits"`
	DeniedOnly     bool                       `json:"denied_only"`
	Status         string                     `json:"status"`
	DisabledReason any                        `json:"disabled_reason"`
	LastSuccessAt  any                        `json:"last_success_at"`
	Revision       int64                      `json:"revision"`
	CreatedAt      any                        `json:"created_at"`
	UpdatedAt      any                        `json:"updated_at"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "url", "description", "events", "benefits", "denied_only", "status", "disabled_reason", "last_success_at", "revision", "created_at", "updated_at"})
}

type WebhookEndpointsDeleteResponse struct {
	Object  string                     `json:"object"`
	ID      any                        `json:"id"`
	Deleted bool                       `json:"deleted"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsDeleteResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsDeleteResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsDeleteResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsDeleteResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "deleted"})
}

type WebhookEndpointsTestInput struct {
	Type  Optional[string]           `json:"type,omitzero"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsTestInput) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsTestInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsTestInput) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsTestInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type"})
}

type WebhookEndpointsTestResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsTestResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsTestResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsTestResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsTestResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type WebhookEndpointsEnableResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	URL            string                     `json:"url"`
	Description    any                        `json:"description"`
	Events         []string                   `json:"events"`
	Benefits       []any                      `json:"benefits"`
	DeniedOnly     bool                       `json:"denied_only"`
	Status         string                     `json:"status"`
	DisabledReason any                        `json:"disabled_reason"`
	LastSuccessAt  any                        `json:"last_success_at"`
	Revision       int64                      `json:"revision"`
	CreatedAt      any                        `json:"created_at"`
	UpdatedAt      any                        `json:"updated_at"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsEnableResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsEnableResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsEnableResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsEnableResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "url", "description", "events", "benefits", "denied_only", "status", "disabled_reason", "last_success_at", "revision", "created_at", "updated_at"})
}

type WebhookEndpointsDisableResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	URL            string                     `json:"url"`
	Description    any                        `json:"description"`
	Events         []string                   `json:"events"`
	Benefits       []any                      `json:"benefits"`
	DeniedOnly     bool                       `json:"denied_only"`
	Status         string                     `json:"status"`
	DisabledReason any                        `json:"disabled_reason"`
	LastSuccessAt  any                        `json:"last_success_at"`
	Revision       int64                      `json:"revision"`
	CreatedAt      any                        `json:"created_at"`
	UpdatedAt      any                        `json:"updated_at"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsDisableResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsDisableResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsDisableResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsDisableResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "url", "description", "events", "benefits", "denied_only", "status", "disabled_reason", "last_success_at", "revision", "created_at", "updated_at"})
}

type WebhookEndpointsSecretResponse struct {
	Object   string                     `json:"object"`
	Endpoint any                        `json:"endpoint"`
	Secret   string                     `json:"secret"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsSecretResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsSecretResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsSecretResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsSecretResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "endpoint", "secret"})
}

type WebhookEndpointsRotateSecretResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	URL            string                     `json:"url"`
	Description    any                        `json:"description"`
	Events         []string                   `json:"events"`
	Benefits       []any                      `json:"benefits"`
	DeniedOnly     bool                       `json:"denied_only"`
	Status         string                     `json:"status"`
	DisabledReason any                        `json:"disabled_reason"`
	LastSuccessAt  any                        `json:"last_success_at"`
	Revision       int64                      `json:"revision"`
	CreatedAt      any                        `json:"created_at"`
	UpdatedAt      any                        `json:"updated_at"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v WebhookEndpointsRotateSecretResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookEndpointsRotateSecretResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookEndpointsRotateSecretResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookEndpointsRotateSecretResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "url", "description", "events", "benefits", "denied_only", "status", "disabled_reason", "last_success_at", "revision", "created_at", "updated_at"})
}

type WebhookDeliveriesUsageResponse struct {
	Object       string                     `json:"object"`
	Day          string                     `json:"day"`
	Deliveries   int64                      `json:"deliveries"`
	Limit        int64                      `json:"limit"`
	LimitReached bool                       `json:"limit_reached"`
	ResetsAt     any                        `json:"resets_at"`
	Extra        map[string]json.RawMessage `json:"-"`
}

func (v WebhookDeliveriesUsageResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookDeliveriesUsageResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookDeliveriesUsageResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookDeliveriesUsageResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "day", "deliveries", "limit", "limit_reached", "resets_at"})
}

type WebhookDeliveriesListInput struct {
	Cursor     Optional[string]           `json:"cursor,omitzero"`
	Limit      Optional[int64]            `json:"limit,omitzero"`
	EndpointID Optional[string]           `json:"endpoint_id,omitzero"`
	Status     Optional[string]           `json:"status,omitzero"`
	Type       Optional[string]           `json:"type,omitzero"`
	EventID    Optional[string]           `json:"event_id,omitzero"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v WebhookDeliveriesListInput) MarshalJSON() ([]byte, error) {
	type plain WebhookDeliveriesListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookDeliveriesListInput) UnmarshalJSON(b []byte) error {
	type plain WebhookDeliveriesListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit", "endpoint_id", "status", "type", "event_id"})
}

type WebhookDeliveriesListResponseDataItem struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WebhookDeliveriesListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain WebhookDeliveriesListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookDeliveriesListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain WebhookDeliveriesListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type WebhookDeliveriesListResponse struct {
	Object     string                                  `json:"object"`
	Data       []WebhookDeliveriesListResponseDataItem `json:"data"`
	NextCursor any                                     `json:"next_cursor"`
	Extra      map[string]json.RawMessage              `json:"-"`
}

func (v WebhookDeliveriesListResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookDeliveriesListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookDeliveriesListResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookDeliveriesListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type WebhookDeliveriesGetResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WebhookDeliveriesGetResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookDeliveriesGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookDeliveriesGetResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookDeliveriesGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type WebhookDeliveriesResendInput struct {
	Target Optional[string]           `json:"target,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WebhookDeliveriesResendInput) MarshalJSON() ([]byte, error) {
	type plain WebhookDeliveriesResendInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookDeliveriesResendInput) UnmarshalJSON(b []byte) error {
	type plain WebhookDeliveriesResendInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"target"})
}

type WebhookDeliveriesResendResponse struct {
	Object string                     `json:"object"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v WebhookDeliveriesResendResponse) MarshalJSON() ([]byte, error) {
	type plain WebhookDeliveriesResendResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *WebhookDeliveriesResendResponse) UnmarshalJSON(b []byte) error {
	type plain WebhookDeliveriesResendResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object"})
}

type BillingPlanResponseScheduledChangeVariant0 struct {
	Plan     string                     `json:"plan"`
	Name     string                     `json:"name"`
	Interval any                        `json:"interval"`
	At       any                        `json:"at"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v BillingPlanResponseScheduledChangeVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingPlanResponseScheduledChangeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPlanResponseScheduledChangeVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingPlanResponseScheduledChangeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"plan", "name", "interval", "at"})
}

type BillingPlanResponsePayment struct {
	State string                     `json:"state"`
	DueAt any                        `json:"due_at"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingPlanResponsePayment) MarshalJSON() ([]byte, error) {
	type plain BillingPlanResponsePayment
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPlanResponsePayment) UnmarshalJSON(b []byte) error {
	type plain BillingPlanResponsePayment
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "due_at"})
}

type BillingPlanResponseLimits struct {
	ProtectedBenefits int64                      `json:"protected_benefits"`
	CustomPolicies    int64                      `json:"custom_policies"`
	TeamSeats         int64                      `json:"team_seats"`
	RetentionDays     int64                      `json:"retention_days"`
	Rate              int64                      `json:"rate"`
	Extra             map[string]json.RawMessage `json:"-"`
}

func (v BillingPlanResponseLimits) MarshalJSON() ([]byte, error) {
	type plain BillingPlanResponseLimits
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPlanResponseLimits) UnmarshalJSON(b []byte) error {
	type plain BillingPlanResponseLimits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"protected_benefits", "custom_policies", "team_seats", "retention_days", "rate"})
}

type BillingPlanResponseOverLimitVariant0ItemsItem struct {
	Kind      string                     `json:"kind"`
	ID        any                        `json:"id"`
	Name      string                     `json:"name"`
	Paused    bool                       `json:"paused"`
	WillPause bool                       `json:"will_pause"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v BillingPlanResponseOverLimitVariant0ItemsItem) MarshalJSON() ([]byte, error) {
	type plain BillingPlanResponseOverLimitVariant0ItemsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPlanResponseOverLimitVariant0ItemsItem) UnmarshalJSON(b []byte) error {
	type plain BillingPlanResponseOverLimitVariant0ItemsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "id", "name", "paused", "will_pause"})
}

type BillingPlanResponseOverLimitVariant0 struct {
	State      string                                          `json:"state"`
	GraceUntil any                                             `json:"grace_until"`
	Items      []BillingPlanResponseOverLimitVariant0ItemsItem `json:"items"`
	Extra      map[string]json.RawMessage                      `json:"-"`
}

func (v BillingPlanResponseOverLimitVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingPlanResponseOverLimitVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPlanResponseOverLimitVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingPlanResponseOverLimitVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"state", "grace_until", "items"})
}

type BillingPlanResponse struct {
	Object          string                     `json:"object"`
	Plan            string                     `json:"plan"`
	Name            string                     `json:"name"`
	Interval        any                        `json:"interval"`
	Subscribed      bool                       `json:"subscribed"`
	RenewsAt        any                        `json:"renews_at"`
	EndsAt          any                        `json:"ends_at"`
	ScheduledChange any                        `json:"scheduled_change"`
	Enterprise      bool                       `json:"enterprise"`
	Payment         BillingPlanResponsePayment `json:"payment"`
	Limits          BillingPlanResponseLimits  `json:"limits"`
	OverLimit       any                        `json:"over_limit"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v BillingPlanResponse) MarshalJSON() ([]byte, error) {
	type plain BillingPlanResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPlanResponse) UnmarshalJSON(b []byte) error {
	type plain BillingPlanResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "plan", "name", "interval", "subscribed", "renews_at", "ends_at", "scheduled_change", "enterprise", "payment", "limits", "over_limit"})
}

type BillingUsageResponseCredits struct {
	Available     int64                      `json:"available"`
	Allowance     int64                      `json:"allowance"`
	PayAsYouGo    int64                      `json:"pay_as_you_go"`
	PerEvaluation int64                      `json:"per_evaluation"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseCredits) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseCredits
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseCredits) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseCredits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"available", "allowance", "pay_as_you_go", "per_evaluation"})
}

type BillingUsageResponseUsageAbsoluteAllowanceCredits struct {
	Used  int64                      `json:"used"`
	Total int64                      `json:"total"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseUsageAbsoluteAllowanceCredits) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAbsoluteAllowanceCredits
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAbsoluteAllowanceCredits) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAbsoluteAllowanceCredits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "total"})
}

type BillingUsageResponseUsageAbsolutePayAsYouGoCredits struct {
	Used  int64                      `json:"used"`
	Total int64                      `json:"total"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseUsageAbsolutePayAsYouGoCredits) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAbsolutePayAsYouGoCredits
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAbsolutePayAsYouGoCredits) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAbsolutePayAsYouGoCredits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "total"})
}

type BillingUsageResponseUsageAbsoluteProtectedBenefits struct {
	Used  int64                      `json:"used"`
	Total int64                      `json:"total"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseUsageAbsoluteProtectedBenefits) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAbsoluteProtectedBenefits
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAbsoluteProtectedBenefits) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAbsoluteProtectedBenefits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "total"})
}

type BillingUsageResponseUsageAbsoluteTeamSeats struct {
	Used  int64                      `json:"used"`
	Total int64                      `json:"total"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseUsageAbsoluteTeamSeats) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAbsoluteTeamSeats
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAbsoluteTeamSeats) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAbsoluteTeamSeats
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "total"})
}

type BillingUsageResponseUsageAbsolute struct {
	AllowanceCredits  BillingUsageResponseUsageAbsoluteAllowanceCredits  `json:"allowance_credits"`
	PayAsYouGoCredits BillingUsageResponseUsageAbsolutePayAsYouGoCredits `json:"pay_as_you_go_credits"`
	ProtectedBenefits BillingUsageResponseUsageAbsoluteProtectedBenefits `json:"protected_benefits"`
	TeamSeats         BillingUsageResponseUsageAbsoluteTeamSeats         `json:"team_seats"`
	Extra             map[string]json.RawMessage                         `json:"-"`
}

func (v BillingUsageResponseUsageAbsolute) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAbsolute
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAbsolute) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAbsolute
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"allowance_credits", "pay_as_you_go_credits", "protected_benefits", "team_seats"})
}

type BillingUsageResponseUsageAllowanceVariant0AllowanceCredits struct {
	Used  int64                      `json:"used"`
	Total int64                      `json:"total"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseUsageAllowanceVariant0AllowanceCredits) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAllowanceVariant0AllowanceCredits
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAllowanceVariant0AllowanceCredits) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAllowanceVariant0AllowanceCredits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "total"})
}

type BillingUsageResponseUsageAllowanceVariant0PayAsYouGoCredits struct {
	Used  int64                      `json:"used"`
	Total int64                      `json:"total"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseUsageAllowanceVariant0PayAsYouGoCredits) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAllowanceVariant0PayAsYouGoCredits
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAllowanceVariant0PayAsYouGoCredits) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAllowanceVariant0PayAsYouGoCredits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "total"})
}

type BillingUsageResponseUsageAllowanceVariant0ProtectedBenefits struct {
	Used  int64                      `json:"used"`
	Total int64                      `json:"total"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseUsageAllowanceVariant0ProtectedBenefits) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAllowanceVariant0ProtectedBenefits
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAllowanceVariant0ProtectedBenefits) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAllowanceVariant0ProtectedBenefits
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "total"})
}

type BillingUsageResponseUsageAllowanceVariant0TeamSeats struct {
	Used  int64                      `json:"used"`
	Total int64                      `json:"total"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseUsageAllowanceVariant0TeamSeats) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAllowanceVariant0TeamSeats
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAllowanceVariant0TeamSeats) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAllowanceVariant0TeamSeats
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"used", "total"})
}

type BillingUsageResponseUsageAllowanceVariant0 struct {
	AllowanceCredits  BillingUsageResponseUsageAllowanceVariant0AllowanceCredits  `json:"allowance_credits"`
	PayAsYouGoCredits BillingUsageResponseUsageAllowanceVariant0PayAsYouGoCredits `json:"pay_as_you_go_credits"`
	ProtectedBenefits BillingUsageResponseUsageAllowanceVariant0ProtectedBenefits `json:"protected_benefits"`
	TeamSeats         BillingUsageResponseUsageAllowanceVariant0TeamSeats         `json:"team_seats"`
	Extra             map[string]json.RawMessage                                  `json:"-"`
}

func (v BillingUsageResponseUsageAllowanceVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsageAllowanceVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsageAllowanceVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsageAllowanceVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"allowance_credits", "pay_as_you_go_credits", "protected_benefits", "team_seats"})
}

type BillingUsageResponseUsage struct {
	Absolute  BillingUsageResponseUsageAbsolute `json:"absolute"`
	Allowance any                               `json:"allowance"`
	Extra     map[string]json.RawMessage        `json:"-"`
}

func (v BillingUsageResponseUsage) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseUsage
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseUsage) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseUsage
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"absolute", "allowance"})
}

type BillingUsageResponseAutoReload struct {
	Enabled           bool                       `json:"enabled"`
	Capped            bool                       `json:"capped"`
	Threshold         any                        `json:"threshold"`
	AmountDollars     any                        `json:"amount_dollars"`
	MonthlyCapDollars any                        `json:"monthly_cap_dollars"`
	Extra             map[string]json.RawMessage `json:"-"`
}

func (v BillingUsageResponseAutoReload) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponseAutoReload
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponseAutoReload) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponseAutoReload
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"enabled", "capped", "threshold", "amount_dollars", "monthly_cap_dollars"})
}

type BillingUsageResponse struct {
	Object     string                         `json:"object"`
	Credits    BillingUsageResponseCredits    `json:"credits"`
	Usage      BillingUsageResponseUsage      `json:"usage"`
	AutoReload BillingUsageResponseAutoReload `json:"auto_reload"`
	Extra      map[string]json.RawMessage     `json:"-"`
}

func (v BillingUsageResponse) MarshalJSON() ([]byte, error) {
	type plain BillingUsageResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUsageResponse) UnmarshalJSON(b []byte) error {
	type plain BillingUsageResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "credits", "usage", "auto_reload"})
}

type BillingCreditHistoryInput struct {
	Cursor    Optional[string]           `json:"cursor,omitzero"`
	Activity  Optional[string]           `json:"activity,omitzero"`
	Direction Optional[string]           `json:"direction,omitzero"`
	Days      Optional[any]              `json:"days,omitzero"`
	From      Optional[string]           `json:"from,omitzero"`
	To        Optional[string]           `json:"to,omitzero"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v BillingCreditHistoryInput) MarshalJSON() ([]byte, error) {
	type plain BillingCreditHistoryInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCreditHistoryInput) UnmarshalJSON(b []byte) error {
	type plain BillingCreditHistoryInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "activity", "direction", "days", "from", "to"})
}

type BillingCreditHistoryResponseDataItem struct {
	Object          string                     `json:"object"`
	Date            string                     `json:"date"`
	DailyTotal      bool                       `json:"daily_total"`
	Activity        string                     `json:"activity"`
	Sources         []string                   `json:"sources"`
	Credits         int64                      `json:"credits"`
	Plans           []string                   `json:"plans"`
	EnterprisePlans []string                   `json:"enterprise_plans"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v BillingCreditHistoryResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain BillingCreditHistoryResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCreditHistoryResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain BillingCreditHistoryResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "date", "daily_total", "activity", "sources", "credits", "plans", "enterprise_plans"})
}

type BillingCreditHistoryResponse struct {
	Object     string                                 `json:"object"`
	Data       []BillingCreditHistoryResponseDataItem `json:"data"`
	NextCursor any                                    `json:"next_cursor"`
	Extra      map[string]json.RawMessage             `json:"-"`
}

func (v BillingCreditHistoryResponse) MarshalJSON() ([]byte, error) {
	type plain BillingCreditHistoryResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCreditHistoryResponse) UnmarshalJSON(b []byte) error {
	type plain BillingCreditHistoryResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type BillingDailyUsageResponseDaysItemTopBenefitVariant0 struct {
	Name        string                     `json:"name"`
	Evaluations int64                      `json:"evaluations"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BillingDailyUsageResponseDaysItemTopBenefitVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingDailyUsageResponseDaysItemTopBenefitVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDailyUsageResponseDaysItemTopBenefitVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingDailyUsageResponseDaysItemTopBenefitVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "evaluations"})
}

type BillingDailyUsageResponseDaysItem struct {
	Day        string                     `json:"day"`
	Credits    int64                      `json:"credits"`
	TopBenefit any                        `json:"top_benefit"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v BillingDailyUsageResponseDaysItem) MarshalJSON() ([]byte, error) {
	type plain BillingDailyUsageResponseDaysItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDailyUsageResponseDaysItem) UnmarshalJSON(b []byte) error {
	type plain BillingDailyUsageResponseDaysItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"day", "credits", "top_benefit"})
}

type BillingDailyUsageResponse struct {
	Object   string                              `json:"object"`
	TimeZone string                              `json:"time_zone"`
	Days     []BillingDailyUsageResponseDaysItem `json:"days"`
	Extra    map[string]json.RawMessage          `json:"-"`
}

func (v BillingDailyUsageResponse) MarshalJSON() ([]byte, error) {
	type plain BillingDailyUsageResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDailyUsageResponse) UnmarshalJSON(b []byte) error {
	type plain BillingDailyUsageResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "time_zone", "days"})
}

type BillingEnterprisePlansResponseDataItemGrants struct {
	Credits           int64                      `json:"credits"`
	CreditCadence     any                        `json:"credit_cadence"`
	CreditExpiry      any                        `json:"credit_expiry"`
	CreditExpiryDays  any                        `json:"credit_expiry_days"`
	TeamSeats         int64                      `json:"team_seats"`
	ProtectedBenefits int64                      `json:"protected_benefits"`
	CustomPolicies    any                        `json:"custom_policies"`
	Rate              any                        `json:"rate"`
	Burst             any                        `json:"burst"`
	Concurrency       any                        `json:"concurrency"`
	RetentionDays     any                        `json:"retention_days"`
	Extra             map[string]json.RawMessage `json:"-"`
}

func (v BillingEnterprisePlansResponseDataItemGrants) MarshalJSON() ([]byte, error) {
	type plain BillingEnterprisePlansResponseDataItemGrants
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingEnterprisePlansResponseDataItemGrants) UnmarshalJSON(b []byte) error {
	type plain BillingEnterprisePlansResponseDataItemGrants
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"credits", "credit_cadence", "credit_expiry", "credit_expiry_days", "team_seats", "protected_benefits", "custom_policies", "rate", "burst", "concurrency", "retention_days"})
}

type BillingEnterprisePlansResponseDataItem struct {
	Object  string                                       `json:"object"`
	Name    string                                       `json:"name"`
	Active  bool                                         `json:"active"`
	AddedAt any                                          `json:"added_at"`
	EndsAt  any                                          `json:"ends_at"`
	Grants  BillingEnterprisePlansResponseDataItemGrants `json:"grants"`
	Extra   map[string]json.RawMessage                   `json:"-"`
}

func (v BillingEnterprisePlansResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain BillingEnterprisePlansResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingEnterprisePlansResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain BillingEnterprisePlansResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "name", "active", "added_at", "ends_at", "grants"})
}

type BillingEnterprisePlansResponse struct {
	Object     string                                   `json:"object"`
	Data       []BillingEnterprisePlansResponseDataItem `json:"data"`
	NextCursor any                                      `json:"next_cursor"`
	Extra      map[string]json.RawMessage               `json:"-"`
}

func (v BillingEnterprisePlansResponse) MarshalJSON() ([]byte, error) {
	type plain BillingEnterprisePlansResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingEnterprisePlansResponse) UnmarshalJSON(b []byte) error {
	type plain BillingEnterprisePlansResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type BillingCodesPreviewInput struct {
	Code  string                     `json:"code"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingCodesPreviewInput) MarshalJSON() ([]byte, error) {
	type plain BillingCodesPreviewInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCodesPreviewInput) UnmarshalJSON(b []byte) error {
	type plain BillingCodesPreviewInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code"})
}

type BillingCodesPreviewResponsePlan struct {
	Name              string                     `json:"name"`
	Description       any                        `json:"description"`
	Credits           int64                      `json:"credits"`
	CreditCadence     any                        `json:"credit_cadence"`
	CreditExpiry      any                        `json:"credit_expiry"`
	CreditExpiryDays  any                        `json:"credit_expiry_days"`
	TeamSeats         int64                      `json:"team_seats"`
	ProtectedBenefits int64                      `json:"protected_benefits"`
	CustomPolicies    any                        `json:"custom_policies"`
	Rate              any                        `json:"rate"`
	Burst             any                        `json:"burst"`
	Concurrency       any                        `json:"concurrency"`
	RetentionDays     any                        `json:"retention_days"`
	DurationMonths    any                        `json:"duration_months"`
	Extra             map[string]json.RawMessage `json:"-"`
}

func (v BillingCodesPreviewResponsePlan) MarshalJSON() ([]byte, error) {
	type plain BillingCodesPreviewResponsePlan
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCodesPreviewResponsePlan) UnmarshalJSON(b []byte) error {
	type plain BillingCodesPreviewResponsePlan
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "description", "credits", "credit_cadence", "credit_expiry", "credit_expiry_days", "team_seats", "protected_benefits", "custom_policies", "rate", "burst", "concurrency", "retention_days", "duration_months"})
}

type BillingCodesPreviewResponse struct {
	Object string                          `json:"object"`
	Plan   BillingCodesPreviewResponsePlan `json:"plan"`
	Extra  map[string]json.RawMessage      `json:"-"`
}

func (v BillingCodesPreviewResponse) MarshalJSON() ([]byte, error) {
	type plain BillingCodesPreviewResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCodesPreviewResponse) UnmarshalJSON(b []byte) error {
	type plain BillingCodesPreviewResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "plan"})
}

type BillingUpcomingResponseDataItem struct {
	Object        string                     `json:"object"`
	Kind          string                     `json:"kind"`
	AmountCents   int64                      `json:"amount_cents"`
	At            any                        `json:"at"`
	PaymentMethod any                        `json:"payment_method"`
	Discounted    bool                       `json:"discounted"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v BillingUpcomingResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain BillingUpcomingResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUpcomingResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain BillingUpcomingResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "kind", "amount_cents", "at", "payment_method", "discounted"})
}

type BillingUpcomingResponse struct {
	Object     string                            `json:"object"`
	Data       []BillingUpcomingResponseDataItem `json:"data"`
	NextCursor any                               `json:"next_cursor"`
	Extra      map[string]json.RawMessage        `json:"-"`
}

func (v BillingUpcomingResponse) MarshalJSON() ([]byte, error) {
	type plain BillingUpcomingResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingUpcomingResponse) UnmarshalJSON(b []byte) error {
	type plain BillingUpcomingResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type BillingInvoicesInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BillingInvoicesInput) MarshalJSON() ([]byte, error) {
	type plain BillingInvoicesInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingInvoicesInput) UnmarshalJSON(b []byte) error {
	type plain BillingInvoicesInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor"})
}

type BillingInvoicesResponseDataItemPeriodVariant0 struct {
	Start any                        `json:"start"`
	End   any                        `json:"end"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingInvoicesResponseDataItemPeriodVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingInvoicesResponseDataItemPeriodVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingInvoicesResponseDataItemPeriodVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingInvoicesResponseDataItemPeriodVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"start", "end"})
}

type BillingInvoicesResponseDataItem struct {
	Object        string                     `json:"object"`
	Number        any                        `json:"number"`
	CreatedAt     any                        `json:"created_at"`
	Status        string                     `json:"status"`
	Description   string                     `json:"description"`
	Period        any                        `json:"period"`
	TotalCents    int64                      `json:"total_cents"`
	PaidCents     int64                      `json:"paid_cents"`
	DueCents      int64                      `json:"due_cents"`
	RefundedCents int64                      `json:"refunded_cents"`
	Payable       bool                       `json:"payable"`
	URL           any                        `json:"url"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v BillingInvoicesResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain BillingInvoicesResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingInvoicesResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain BillingInvoicesResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "number", "created_at", "status", "description", "period", "total_cents", "paid_cents", "due_cents", "refunded_cents", "payable", "url"})
}

type BillingInvoicesResponse struct {
	Object           string                            `json:"object"`
	Data             []BillingInvoicesResponseDataItem `json:"data"`
	NextCursor       any                               `json:"next_cursor"`
	OutstandingCents any                               `json:"outstanding_cents"`
	Extra            map[string]json.RawMessage        `json:"-"`
}

func (v BillingInvoicesResponse) MarshalJSON() ([]byte, error) {
	type plain BillingInvoicesResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingInvoicesResponse) UnmarshalJSON(b []byte) error {
	type plain BillingInvoicesResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor", "outstanding_cents"})
}

type BillingPaymentMethodsResponseDataItem struct {
	Object   string                     `json:"object"`
	ID       any                        `json:"id"`
	Type     string                     `json:"type"`
	Brand    any                        `json:"brand"`
	Last4    any                        `json:"last4"`
	ExpMonth any                        `json:"exp_month"`
	ExpYear  any                        `json:"exp_year"`
	Label    string                     `json:"label"`
	Default  bool                       `json:"default"`
	Backup   any                        `json:"backup"`
	Declined bool                       `json:"declined"`
	Expired  bool                       `json:"expired"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v BillingPaymentMethodsResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "type", "brand", "last4", "exp_month", "exp_year", "label", "default", "backup", "declined", "expired"})
}

type BillingPaymentMethodsResponse struct {
	Object     string                                  `json:"object"`
	Data       []BillingPaymentMethodsResponseDataItem `json:"data"`
	NextCursor any                                     `json:"next_cursor"`
	Extra      map[string]json.RawMessage              `json:"-"`
}

func (v BillingPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsResponse) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type BillingDetailsResponseAddressVariant0 struct {
	Line1      string                     `json:"line1"`
	Line2      string                     `json:"line2"`
	City       string                     `json:"city"`
	State      string                     `json:"state"`
	PostalCode string                     `json:"postal_code"`
	Country    string                     `json:"country"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v BillingDetailsResponseAddressVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsResponseAddressVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsResponseAddressVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsResponseAddressVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"line1", "line2", "city", "state", "postal_code", "country"})
}

type BillingDetailsResponseTaxIDVariant0 struct {
	Type  string                     `json:"type"`
	Value string                     `json:"value"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingDetailsResponseTaxIDVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsResponseTaxIDVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsResponseTaxIDVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsResponseTaxIDVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "value"})
}

type BillingDetailsResponse struct {
	Object           string                     `json:"object"`
	Name             string                     `json:"name"`
	Email            string                     `json:"email"`
	UsesAccountEmail bool                       `json:"uses_account_email"`
	Address          any                        `json:"address"`
	TaxID            any                        `json:"tax_id"`
	Extra            map[string]json.RawMessage `json:"-"`
}

func (v BillingDetailsResponse) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsResponse) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "name", "email", "uses_account_email", "address", "tax_id"})
}

type BillingDetailsUpdateInputAddress struct {
	Line1      string                     `json:"line1"`
	Line2      string                     `json:"line2"`
	City       string                     `json:"city"`
	State      string                     `json:"state"`
	PostalCode string                     `json:"postal_code"`
	Country    any                        `json:"country"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v BillingDetailsUpdateInputAddress) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsUpdateInputAddress
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsUpdateInputAddress) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsUpdateInputAddress
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"line1", "line2", "city", "state", "postal_code", "country"})
}

type BillingDetailsUpdateInputTaxIDVariant0 struct {
	Type  string                     `json:"type"`
	Value string                     `json:"value"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingDetailsUpdateInputTaxIDVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsUpdateInputTaxIDVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsUpdateInputTaxIDVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsUpdateInputTaxIDVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "value"})
}

type BillingDetailsUpdateInput struct {
	Name    string                           `json:"name"`
	Email   any                              `json:"email"`
	Address BillingDetailsUpdateInputAddress `json:"address"`
	TaxID   any                              `json:"tax_id"`
	Extra   map[string]json.RawMessage       `json:"-"`
}

func (v BillingDetailsUpdateInput) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsUpdateInput) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name", "email", "address", "tax_id"})
}

type BillingDetailsUpdateResponseAddressVariant0 struct {
	Line1      string                     `json:"line1"`
	Line2      string                     `json:"line2"`
	City       string                     `json:"city"`
	State      string                     `json:"state"`
	PostalCode string                     `json:"postal_code"`
	Country    string                     `json:"country"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v BillingDetailsUpdateResponseAddressVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsUpdateResponseAddressVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsUpdateResponseAddressVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsUpdateResponseAddressVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"line1", "line2", "city", "state", "postal_code", "country"})
}

type BillingDetailsUpdateResponseTaxIDVariant0 struct {
	Type  string                     `json:"type"`
	Value string                     `json:"value"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingDetailsUpdateResponseTaxIDVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsUpdateResponseTaxIDVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsUpdateResponseTaxIDVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsUpdateResponseTaxIDVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"type", "value"})
}

type BillingDetailsUpdateResponse struct {
	Object           string                     `json:"object"`
	Name             string                     `json:"name"`
	Email            string                     `json:"email"`
	UsesAccountEmail bool                       `json:"uses_account_email"`
	Address          any                        `json:"address"`
	TaxID            any                        `json:"tax_id"`
	Extra            map[string]json.RawMessage `json:"-"`
}

func (v BillingDetailsUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain BillingDetailsUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDetailsUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain BillingDetailsUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "name", "email", "uses_account_email", "address", "tax_id"})
}

type BillingPaymentsGetResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	Kind           string                     `json:"kind"`
	Status         string                     `json:"status"`
	PaymentPageURL any                        `json:"payment_page_url"`
	Credits        any                        `json:"credits"`
	BonusCredits   any                        `json:"bonus_credits"`
	Scheduled      bool                       `json:"scheduled"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BillingPaymentsGetResponse) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentsGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentsGetResponse) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentsGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "kind", "status", "payment_page_url", "credits", "bonus_credits", "scheduled"})
}

type BillingQuotesCreateInput struct {
	Plan      string                     `json:"plan"`
	Interval  Optional[any]              `json:"interval,omitzero"`
	PromoCode Optional[string]           `json:"promo_code,omitzero"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v BillingQuotesCreateInput) MarshalJSON() ([]byte, error) {
	type plain BillingQuotesCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingQuotesCreateInput) UnmarshalJSON(b []byte) error {
	type plain BillingQuotesCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"plan", "interval", "promo_code"})
}

type BillingQuotesCreateResponseNextChargeVariant0 struct {
	AmountCents int64                      `json:"amount_cents"`
	At          any                        `json:"at"`
	Discounted  bool                       `json:"discounted"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BillingQuotesCreateResponseNextChargeVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingQuotesCreateResponseNextChargeVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingQuotesCreateResponseNextChargeVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingQuotesCreateResponseNextChargeVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"amount_cents", "at", "discounted"})
}

type BillingQuotesCreateResponsePromotionVariant0 struct {
	Code        string                     `json:"code"`
	AmountCents int64                      `json:"amount_cents"`
	FullCents   int64                      `json:"full_cents"`
	Duration    string                     `json:"duration"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BillingQuotesCreateResponsePromotionVariant0) MarshalJSON() ([]byte, error) {
	type plain BillingQuotesCreateResponsePromotionVariant0
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingQuotesCreateResponsePromotionVariant0) UnmarshalJSON(b []byte) error {
	type plain BillingQuotesCreateResponsePromotionVariant0
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code", "amount_cents", "full_cents", "duration"})
}

type BillingQuotesCreateResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	Plan           string                     `json:"plan"`
	Interval       any                        `json:"interval"`
	Direction      string                     `json:"direction"`
	Timing         string                     `json:"timing"`
	Confirmable    bool                       `json:"confirmable"`
	ChargeNowCents int64                      `json:"charge_now_cents"`
	RenewalCents   int64                      `json:"renewal_cents"`
	CreditsNow     int64                      `json:"credits_now"`
	EffectiveAt    any                        `json:"effective_at"`
	ExpiresAt      any                        `json:"expires_at"`
	NextCharge     any                        `json:"next_charge"`
	Promotion      any                        `json:"promotion"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BillingQuotesCreateResponse) MarshalJSON() ([]byte, error) {
	type plain BillingQuotesCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingQuotesCreateResponse) UnmarshalJSON(b []byte) error {
	type plain BillingQuotesCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "plan", "interval", "direction", "timing", "confirmable", "charge_now_cents", "renewal_cents", "credits_now", "effective_at", "expires_at", "next_charge", "promotion"})
}

type BillingLowCreditAlertUpdateInput struct {
	Credits any                        `json:"credits"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v BillingLowCreditAlertUpdateInput) MarshalJSON() ([]byte, error) {
	type plain BillingLowCreditAlertUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingLowCreditAlertUpdateInput) UnmarshalJSON(b []byte) error {
	type plain BillingLowCreditAlertUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"credits"})
}

type BillingLowCreditAlertUpdateResponse struct {
	Object  string                     `json:"object"`
	Credits any                        `json:"credits"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v BillingLowCreditAlertUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain BillingLowCreditAlertUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingLowCreditAlertUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain BillingLowCreditAlertUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "credits"})
}

type BillingOverLimitChooseInput struct {
	Kind       string                     `json:"kind"`
	KeepActive string                     `json:"keep_active"`
	Pause      string                     `json:"pause"`
	Extra      map[string]json.RawMessage `json:"-"`
}

func (v BillingOverLimitChooseInput) MarshalJSON() ([]byte, error) {
	type plain BillingOverLimitChooseInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingOverLimitChooseInput) UnmarshalJSON(b []byte) error {
	type plain BillingOverLimitChooseInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "keep_active", "pause"})
}

type BillingOverLimitChooseResponse struct {
	Object string                     `json:"object"`
	Status string                     `json:"status"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BillingOverLimitChooseResponse) MarshalJSON() ([]byte, error) {
	type plain BillingOverLimitChooseResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingOverLimitChooseResponse) UnmarshalJSON(b []byte) error {
	type plain BillingOverLimitChooseResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status"})
}

type BillingAutoReloadUpdateInput struct {
	Enabled           bool                       `json:"enabled"`
	Threshold         any                        `json:"threshold"`
	AmountDollars     any                        `json:"amount_dollars"`
	MonthlyCapDollars Optional[any]              `json:"monthly_cap_dollars,omitzero"`
	Extra             map[string]json.RawMessage `json:"-"`
}

func (v BillingAutoReloadUpdateInput) MarshalJSON() ([]byte, error) {
	type plain BillingAutoReloadUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingAutoReloadUpdateInput) UnmarshalJSON(b []byte) error {
	type plain BillingAutoReloadUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"enabled", "threshold", "amount_dollars", "monthly_cap_dollars"})
}

type BillingAutoReloadUpdateResponse struct {
	Object                string                     `json:"object"`
	Enabled               bool                       `json:"enabled"`
	Capped                bool                       `json:"capped"`
	Threshold             any                        `json:"threshold"`
	AmountDollars         any                        `json:"amount_dollars"`
	MonthlyCapDollars     any                        `json:"monthly_cap_dollars"`
	SpentThisMonthDollars int64                      `json:"spent_this_month_dollars"`
	Extra                 map[string]json.RawMessage `json:"-"`
}

func (v BillingAutoReloadUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain BillingAutoReloadUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingAutoReloadUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain BillingAutoReloadUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "enabled", "capped", "threshold", "amount_dollars", "monthly_cap_dollars", "spent_this_month_dollars"})
}

type BillingCreditsBuyInput struct {
	AmountDollars int64                      `json:"amount_dollars"`
	Extra         map[string]json.RawMessage `json:"-"`
}

func (v BillingCreditsBuyInput) MarshalJSON() ([]byte, error) {
	type plain BillingCreditsBuyInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCreditsBuyInput) UnmarshalJSON(b []byte) error {
	type plain BillingCreditsBuyInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"amount_dollars"})
}

type BillingCreditsBuyResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	Kind           string                     `json:"kind"`
	Status         string                     `json:"status"`
	PaymentPageURL any                        `json:"payment_page_url"`
	Credits        any                        `json:"credits"`
	BonusCredits   any                        `json:"bonus_credits"`
	Scheduled      bool                       `json:"scheduled"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BillingCreditsBuyResponse) MarshalJSON() ([]byte, error) {
	type plain BillingCreditsBuyResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCreditsBuyResponse) UnmarshalJSON(b []byte) error {
	type plain BillingCreditsBuyResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "kind", "status", "payment_page_url", "credits", "bonus_credits", "scheduled"})
}

type BillingCreditsRetryInput struct {
	PaymentMethodID Optional[string]           `json:"payment_method_id,omitzero"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v BillingCreditsRetryInput) MarshalJSON() ([]byte, error) {
	type plain BillingCreditsRetryInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCreditsRetryInput) UnmarshalJSON(b []byte) error {
	type plain BillingCreditsRetryInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"payment_method_id"})
}

type BillingCreditsRetryResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	Kind           string                     `json:"kind"`
	Status         string                     `json:"status"`
	PaymentPageURL any                        `json:"payment_page_url"`
	Credits        any                        `json:"credits"`
	BonusCredits   any                        `json:"bonus_credits"`
	Scheduled      bool                       `json:"scheduled"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BillingCreditsRetryResponse) MarshalJSON() ([]byte, error) {
	type plain BillingCreditsRetryResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCreditsRetryResponse) UnmarshalJSON(b []byte) error {
	type plain BillingCreditsRetryResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "kind", "status", "payment_page_url", "credits", "bonus_credits", "scheduled"})
}

type BillingPlanChangeInput struct {
	PaymentMethodID Optional[string]           `json:"payment_method_id,omitzero"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v BillingPlanChangeInput) MarshalJSON() ([]byte, error) {
	type plain BillingPlanChangeInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPlanChangeInput) UnmarshalJSON(b []byte) error {
	type plain BillingPlanChangeInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"payment_method_id"})
}

type BillingPlanChangeResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	Kind           string                     `json:"kind"`
	Status         string                     `json:"status"`
	PaymentPageURL any                        `json:"payment_page_url"`
	Credits        any                        `json:"credits"`
	BonusCredits   any                        `json:"bonus_credits"`
	Scheduled      bool                       `json:"scheduled"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BillingPlanChangeResponse) MarshalJSON() ([]byte, error) {
	type plain BillingPlanChangeResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPlanChangeResponse) UnmarshalJSON(b []byte) error {
	type plain BillingPlanChangeResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "kind", "status", "payment_page_url", "credits", "bonus_credits", "scheduled"})
}

type BillingDowngradeUndoResponse struct {
	Object string                     `json:"object"`
	Status string                     `json:"status"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BillingDowngradeUndoResponse) MarshalJSON() ([]byte, error) {
	type plain BillingDowngradeUndoResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingDowngradeUndoResponse) UnmarshalJSON(b []byte) error {
	type plain BillingDowngradeUndoResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status"})
}

type BillingSubscriptionResumeResponse struct {
	Object string                     `json:"object"`
	Status string                     `json:"status"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BillingSubscriptionResumeResponse) MarshalJSON() ([]byte, error) {
	type plain BillingSubscriptionResumeResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingSubscriptionResumeResponse) UnmarshalJSON(b []byte) error {
	type plain BillingSubscriptionResumeResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status"})
}

type BillingOverduePayInput struct {
	InvoiceNumber   Optional[string]           `json:"invoice_number,omitzero"`
	PaymentMethodID Optional[string]           `json:"payment_method_id,omitzero"`
	Extra           map[string]json.RawMessage `json:"-"`
}

func (v BillingOverduePayInput) MarshalJSON() ([]byte, error) {
	type plain BillingOverduePayInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingOverduePayInput) UnmarshalJSON(b []byte) error {
	type plain BillingOverduePayInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"invoice_number", "payment_method_id"})
}

type BillingOverduePayResponse struct {
	Object         string                     `json:"object"`
	ID             any                        `json:"id"`
	Kind           string                     `json:"kind"`
	Status         string                     `json:"status"`
	PaymentPageURL any                        `json:"payment_page_url"`
	Credits        any                        `json:"credits"`
	BonusCredits   any                        `json:"bonus_credits"`
	Scheduled      bool                       `json:"scheduled"`
	Extra          map[string]json.RawMessage `json:"-"`
}

func (v BillingOverduePayResponse) MarshalJSON() ([]byte, error) {
	type plain BillingOverduePayResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingOverduePayResponse) UnmarshalJSON(b []byte) error {
	type plain BillingOverduePayResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "kind", "status", "payment_page_url", "credits", "bonus_credits", "scheduled"})
}

type BillingPaymentMethodsDefaultResponseOverdueInvoicesItem struct {
	Number      any                        `json:"number"`
	AmountCents int64                      `json:"amount_cents"`
	Status      string                     `json:"status"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BillingPaymentMethodsDefaultResponseOverdueInvoicesItem) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsDefaultResponseOverdueInvoicesItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsDefaultResponseOverdueInvoicesItem) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsDefaultResponseOverdueInvoicesItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"number", "amount_cents", "status"})
}

type BillingPaymentMethodsDefaultResponse struct {
	Object          string                                                    `json:"object"`
	Status          string                                                    `json:"status"`
	OverdueInvoices []BillingPaymentMethodsDefaultResponseOverdueInvoicesItem `json:"overdue_invoices"`
	Extra           map[string]json.RawMessage                                `json:"-"`
}

func (v BillingPaymentMethodsDefaultResponse) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsDefaultResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsDefaultResponse) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsDefaultResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "overdue_invoices"})
}

type BillingPaymentMethodsMoveInput struct {
	Direction string                     `json:"direction"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v BillingPaymentMethodsMoveInput) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsMoveInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsMoveInput) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsMoveInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"direction"})
}

type BillingPaymentMethodsMoveResponseOverdueInvoicesItem struct {
	Number      any                        `json:"number"`
	AmountCents int64                      `json:"amount_cents"`
	Status      string                     `json:"status"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BillingPaymentMethodsMoveResponseOverdueInvoicesItem) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsMoveResponseOverdueInvoicesItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsMoveResponseOverdueInvoicesItem) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsMoveResponseOverdueInvoicesItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"number", "amount_cents", "status"})
}

type BillingPaymentMethodsMoveResponse struct {
	Object          string                                                 `json:"object"`
	Status          string                                                 `json:"status"`
	OverdueInvoices []BillingPaymentMethodsMoveResponseOverdueInvoicesItem `json:"overdue_invoices"`
	Extra           map[string]json.RawMessage                             `json:"-"`
}

func (v BillingPaymentMethodsMoveResponse) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsMoveResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsMoveResponse) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsMoveResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "overdue_invoices"})
}

type BillingPaymentMethodsRemoveResponseOverdueInvoicesItem struct {
	Number      any                        `json:"number"`
	AmountCents int64                      `json:"amount_cents"`
	Status      string                     `json:"status"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BillingPaymentMethodsRemoveResponseOverdueInvoicesItem) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsRemoveResponseOverdueInvoicesItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsRemoveResponseOverdueInvoicesItem) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsRemoveResponseOverdueInvoicesItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"number", "amount_cents", "status"})
}

type BillingPaymentMethodsRemoveResponse struct {
	Object          string                                                   `json:"object"`
	Status          string                                                   `json:"status"`
	OverdueInvoices []BillingPaymentMethodsRemoveResponseOverdueInvoicesItem `json:"overdue_invoices"`
	Extra           map[string]json.RawMessage                               `json:"-"`
}

func (v BillingPaymentMethodsRemoveResponse) MarshalJSON() ([]byte, error) {
	type plain BillingPaymentMethodsRemoveResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingPaymentMethodsRemoveResponse) UnmarshalJSON(b []byte) error {
	type plain BillingPaymentMethodsRemoveResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "overdue_invoices"})
}

type BillingCodesRedeemInput struct {
	Code  string                     `json:"code"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v BillingCodesRedeemInput) MarshalJSON() ([]byte, error) {
	type plain BillingCodesRedeemInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCodesRedeemInput) UnmarshalJSON(b []byte) error {
	type plain BillingCodesRedeemInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"code"})
}

type BillingCodesRedeemResponse struct {
	Object string                     `json:"object"`
	Status string                     `json:"status"`
	Plan   any                        `json:"plan"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v BillingCodesRedeemResponse) MarshalJSON() ([]byte, error) {
	type plain BillingCodesRedeemResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingCodesRedeemResponse) UnmarshalJSON(b []byte) error {
	type plain BillingCodesRedeemResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "status", "plan"})
}

type BillingApprovalLinksCreateInput struct {
	Kind     string                     `json:"kind"`
	Plan     Optional[string]           `json:"plan,omitzero"`
	Interval Optional[string]           `json:"interval,omitzero"`
	Extra    map[string]json.RawMessage `json:"-"`
}

func (v BillingApprovalLinksCreateInput) MarshalJSON() ([]byte, error) {
	type plain BillingApprovalLinksCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingApprovalLinksCreateInput) UnmarshalJSON(b []byte) error {
	type plain BillingApprovalLinksCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"kind", "plan", "interval"})
}

type BillingApprovalLinksCreateResponse struct {
	Object      string                     `json:"object"`
	ID          string                     `json:"id"`
	Kind        string                     `json:"kind"`
	Plan        any                        `json:"plan"`
	Interval    any                        `json:"interval"`
	Status      string                     `json:"status"`
	ApprovalURL any                        `json:"approval_url"`
	ExpiresAt   any                        `json:"expires_at"`
	CompletedAt any                        `json:"completed_at"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BillingApprovalLinksCreateResponse) MarshalJSON() ([]byte, error) {
	type plain BillingApprovalLinksCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingApprovalLinksCreateResponse) UnmarshalJSON(b []byte) error {
	type plain BillingApprovalLinksCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "kind", "plan", "interval", "status", "approval_url", "expires_at", "completed_at"})
}

type BillingApprovalLinksGetResponse struct {
	Object      string                     `json:"object"`
	ID          string                     `json:"id"`
	Kind        string                     `json:"kind"`
	Plan        any                        `json:"plan"`
	Interval    any                        `json:"interval"`
	Status      string                     `json:"status"`
	ApprovalURL any                        `json:"approval_url"`
	ExpiresAt   any                        `json:"expires_at"`
	CompletedAt any                        `json:"completed_at"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v BillingApprovalLinksGetResponse) MarshalJSON() ([]byte, error) {
	type plain BillingApprovalLinksGetResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *BillingApprovalLinksGetResponse) UnmarshalJSON(b []byte) error {
	type plain BillingApprovalLinksGetResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "kind", "plan", "interval", "status", "approval_url", "expires_at", "completed_at"})
}

type AccountUpdateInput struct {
	Name  string                     `json:"name"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v AccountUpdateInput) MarshalJSON() ([]byte, error) {
	type plain AccountUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountUpdateInput) UnmarshalJSON(b []byte) error {
	type plain AccountUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"name"})
}

type AccountUpdateResponse struct {
	Object string                     `json:"object"`
	Name   string                     `json:"name"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v AccountUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain AccountUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain AccountUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "name"})
}

type AccountEmailPreferencesResponse struct {
	Object                string                     `json:"object"`
	MembersAndInvitations bool                       `json:"members_and_invitations"`
	ImportsAndExports     bool                       `json:"imports_and_exports"`
	IntegrationAlerts     bool                       `json:"integration_alerts"`
	WeeklySummary         bool                       `json:"weekly_summary"`
	Extra                 map[string]json.RawMessage `json:"-"`
}

func (v AccountEmailPreferencesResponse) MarshalJSON() ([]byte, error) {
	type plain AccountEmailPreferencesResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountEmailPreferencesResponse) UnmarshalJSON(b []byte) error {
	type plain AccountEmailPreferencesResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "members_and_invitations", "imports_and_exports", "integration_alerts", "weekly_summary"})
}

type AccountEmailPreferencesUpdateInput struct {
	MembersAndInvitations Optional[bool]             `json:"members_and_invitations,omitzero"`
	ImportsAndExports     Optional[bool]             `json:"imports_and_exports,omitzero"`
	IntegrationAlerts     Optional[bool]             `json:"integration_alerts,omitzero"`
	WeeklySummary         Optional[bool]             `json:"weekly_summary,omitzero"`
	Extra                 map[string]json.RawMessage `json:"-"`
}

func (v AccountEmailPreferencesUpdateInput) MarshalJSON() ([]byte, error) {
	type plain AccountEmailPreferencesUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountEmailPreferencesUpdateInput) UnmarshalJSON(b []byte) error {
	type plain AccountEmailPreferencesUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"members_and_invitations", "imports_and_exports", "integration_alerts", "weekly_summary"})
}

type AccountEmailPreferencesUpdateResponse struct {
	Object                string                     `json:"object"`
	MembersAndInvitations bool                       `json:"members_and_invitations"`
	ImportsAndExports     bool                       `json:"imports_and_exports"`
	IntegrationAlerts     bool                       `json:"integration_alerts"`
	WeeklySummary         bool                       `json:"weekly_summary"`
	Extra                 map[string]json.RawMessage `json:"-"`
}

func (v AccountEmailPreferencesUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain AccountEmailPreferencesUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *AccountEmailPreferencesUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain AccountEmailPreferencesUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "members_and_invitations", "imports_and_exports", "integration_alerts", "weekly_summary"})
}

type NotificationsListInput struct {
	Filter Optional[string]           `json:"filter,omitzero"`
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v NotificationsListInput) MarshalJSON() ([]byte, error) {
	type plain NotificationsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsListInput) UnmarshalJSON(b []byte) error {
	type plain NotificationsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"filter", "cursor"})
}

type NotificationsListResponseDataItemActionsItem struct {
	ID    string                     `json:"id"`
	Label string                     `json:"label"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v NotificationsListResponseDataItemActionsItem) MarshalJSON() ([]byte, error) {
	type plain NotificationsListResponseDataItemActionsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsListResponseDataItemActionsItem) UnmarshalJSON(b []byte) error {
	type plain NotificationsListResponseDataItemActionsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "label"})
}

type NotificationsListResponseDataItem struct {
	Object    string                                         `json:"object"`
	ID        string                                         `json:"id"`
	Workspace any                                            `json:"workspace"`
	Kind      string                                         `json:"kind"`
	Title     string                                         `json:"title"`
	Snippet   string                                         `json:"snippet"`
	Unread    bool                                           `json:"unread"`
	Count     int64                                          `json:"count"`
	UpdatedAt any                                            `json:"updated_at"`
	Actions   []NotificationsListResponseDataItemActionsItem `json:"actions"`
	Resolved  any                                            `json:"resolved"`
	Extra     map[string]json.RawMessage                     `json:"-"`
}

func (v NotificationsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain NotificationsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain NotificationsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "workspace", "kind", "title", "snippet", "unread", "count", "updated_at", "actions", "resolved"})
}

type NotificationsListResponse struct {
	Object      string                              `json:"object"`
	Data        []NotificationsListResponseDataItem `json:"data"`
	NextCursor  any                                 `json:"next_cursor"`
	UnreadCount int64                               `json:"unread_count"`
	Extra       map[string]json.RawMessage          `json:"-"`
}

func (v NotificationsListResponse) MarshalJSON() ([]byte, error) {
	type plain NotificationsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsListResponse) UnmarshalJSON(b []byte) error {
	type plain NotificationsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor", "unread_count"})
}

type NotificationsUpdateInput struct {
	Read  bool                       `json:"read"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v NotificationsUpdateInput) MarshalJSON() ([]byte, error) {
	type plain NotificationsUpdateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsUpdateInput) UnmarshalJSON(b []byte) error {
	type plain NotificationsUpdateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"read"})
}

type NotificationsUpdateResponseActionsItem struct {
	ID    string                     `json:"id"`
	Label string                     `json:"label"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v NotificationsUpdateResponseActionsItem) MarshalJSON() ([]byte, error) {
	type plain NotificationsUpdateResponseActionsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsUpdateResponseActionsItem) UnmarshalJSON(b []byte) error {
	type plain NotificationsUpdateResponseActionsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "label"})
}

type NotificationsUpdateResponse struct {
	Object      string                                   `json:"object"`
	ID          string                                   `json:"id"`
	Workspace   any                                      `json:"workspace"`
	Kind        string                                   `json:"kind"`
	Title       string                                   `json:"title"`
	Snippet     string                                   `json:"snippet"`
	Unread      bool                                     `json:"unread"`
	Count       int64                                    `json:"count"`
	UpdatedAt   any                                      `json:"updated_at"`
	Actions     []NotificationsUpdateResponseActionsItem `json:"actions"`
	Resolved    any                                      `json:"resolved"`
	UnreadCount int64                                    `json:"unread_count"`
	Extra       map[string]json.RawMessage               `json:"-"`
}

func (v NotificationsUpdateResponse) MarshalJSON() ([]byte, error) {
	type plain NotificationsUpdateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsUpdateResponse) UnmarshalJSON(b []byte) error {
	type plain NotificationsUpdateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "workspace", "kind", "title", "snippet", "unread", "count", "updated_at", "actions", "resolved", "unread_count"})
}

type NotificationsReadAllResponse struct {
	Object      string                     `json:"object"`
	UnreadCount int64                      `json:"unread_count"`
	Extra       map[string]json.RawMessage `json:"-"`
}

func (v NotificationsReadAllResponse) MarshalJSON() ([]byte, error) {
	type plain NotificationsReadAllResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsReadAllResponse) UnmarshalJSON(b []byte) error {
	type plain NotificationsReadAllResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "unread_count"})
}

type NotificationsRunActionInput struct {
	Action string                     `json:"action"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v NotificationsRunActionInput) MarshalJSON() ([]byte, error) {
	type plain NotificationsRunActionInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsRunActionInput) UnmarshalJSON(b []byte) error {
	type plain NotificationsRunActionInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"action"})
}

type NotificationsRunActionResponseActionsItem struct {
	ID    string                     `json:"id"`
	Label string                     `json:"label"`
	Extra map[string]json.RawMessage `json:"-"`
}

func (v NotificationsRunActionResponseActionsItem) MarshalJSON() ([]byte, error) {
	type plain NotificationsRunActionResponseActionsItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsRunActionResponseActionsItem) UnmarshalJSON(b []byte) error {
	type plain NotificationsRunActionResponseActionsItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"id", "label"})
}

type NotificationsRunActionResponse struct {
	Object      string                                      `json:"object"`
	ID          string                                      `json:"id"`
	Workspace   any                                         `json:"workspace"`
	Kind        string                                      `json:"kind"`
	Title       string                                      `json:"title"`
	Snippet     string                                      `json:"snippet"`
	Unread      bool                                        `json:"unread"`
	Count       int64                                       `json:"count"`
	UpdatedAt   any                                         `json:"updated_at"`
	Actions     []NotificationsRunActionResponseActionsItem `json:"actions"`
	Resolved    any                                         `json:"resolved"`
	UnreadCount int64                                       `json:"unread_count"`
	Extra       map[string]json.RawMessage                  `json:"-"`
}

func (v NotificationsRunActionResponse) MarshalJSON() ([]byte, error) {
	type plain NotificationsRunActionResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *NotificationsRunActionResponse) UnmarshalJSON(b []byte) error {
	type plain NotificationsRunActionResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "workspace", "kind", "title", "snippet", "unread", "count", "updated_at", "actions", "resolved", "unread_count"})
}

type SupportTicketsCreateInput struct {
	Subject   string                     `json:"subject"`
	Message   string                     `json:"message"`
	Topic     string                     `json:"topic"`
	RequestID Optional[string]           `json:"request_id,omitzero"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v SupportTicketsCreateInput) MarshalJSON() ([]byte, error) {
	type plain SupportTicketsCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *SupportTicketsCreateInput) UnmarshalJSON(b []byte) error {
	type plain SupportTicketsCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"subject", "message", "topic", "request_id"})
}

type SupportTicketsCreateResponse struct {
	Object    string                     `json:"object"`
	ID        string                     `json:"id"`
	Status    string                     `json:"status"`
	CreatedAt any                        `json:"created_at"`
	UpdatedAt any                        `json:"updated_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v SupportTicketsCreateResponse) MarshalJSON() ([]byte, error) {
	type plain SupportTicketsCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *SupportTicketsCreateResponse) UnmarshalJSON(b []byte) error {
	type plain SupportTicketsCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "status", "created_at", "updated_at"})
}

type SupportTicketsListInput struct {
	Cursor Optional[string]           `json:"cursor,omitzero"`
	Limit  Optional[int64]            `json:"limit,omitzero"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v SupportTicketsListInput) MarshalJSON() ([]byte, error) {
	type plain SupportTicketsListInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *SupportTicketsListInput) UnmarshalJSON(b []byte) error {
	type plain SupportTicketsListInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"cursor", "limit"})
}

type SupportTicketsListResponseDataItem struct {
	Object    string                     `json:"object"`
	ID        string                     `json:"id"`
	Status    string                     `json:"status"`
	CreatedAt any                        `json:"created_at"`
	UpdatedAt any                        `json:"updated_at"`
	Extra     map[string]json.RawMessage `json:"-"`
}

func (v SupportTicketsListResponseDataItem) MarshalJSON() ([]byte, error) {
	type plain SupportTicketsListResponseDataItem
	return marshalExtras(plain(v), v.Extra)
}
func (v *SupportTicketsListResponseDataItem) UnmarshalJSON(b []byte) error {
	type plain SupportTicketsListResponseDataItem
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "id", "status", "created_at", "updated_at"})
}

type SupportTicketsListResponse struct {
	Object     string                               `json:"object"`
	Data       []SupportTicketsListResponseDataItem `json:"data"`
	NextCursor any                                  `json:"next_cursor"`
	Extra      map[string]json.RawMessage           `json:"-"`
}

func (v SupportTicketsListResponse) MarshalJSON() ([]byte, error) {
	type plain SupportTicketsListResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *SupportTicketsListResponse) UnmarshalJSON(b []byte) error {
	type plain SupportTicketsListResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "data", "next_cursor"})
}

type FeedbackCreateInput struct {
	Message string                     `json:"message"`
	Page    Optional[string]           `json:"page,omitzero"`
	Extra   map[string]json.RawMessage `json:"-"`
}

func (v FeedbackCreateInput) MarshalJSON() ([]byte, error) {
	type plain FeedbackCreateInput
	return marshalExtras(plain(v), v.Extra)
}
func (v *FeedbackCreateInput) UnmarshalJSON(b []byte) error {
	type plain FeedbackCreateInput
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"message", "page"})
}

type FeedbackCreateResponse struct {
	Object string                     `json:"object"`
	Sent   bool                       `json:"sent"`
	Extra  map[string]json.RawMessage `json:"-"`
}

func (v FeedbackCreateResponse) MarshalJSON() ([]byte, error) {
	type plain FeedbackCreateResponse
	return marshalExtras(plain(v), v.Extra)
}
func (v *FeedbackCreateResponse) UnmarshalJSON(b []byte) error {
	type plain FeedbackCreateResponse
	return unmarshalExtras(b, (*plain)(v), &v.Extra, []string{"object", "sent"})
}

// Pricing calls GET /v1/public/pricing. Read AbuseZero's plans and prices, what each plan includes and what every plan shares.
func (c *PublicClient) Pricing(ctx context.Context, input map[string]any, options *CallOptions) (PublicPricingResponse, error) {
	var result PublicPricingResponse
	data, err := c.call(ctx, "public.pricing", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// PriceQuote calls GET /v1/public/pricing/quote. Get the exact price of a plan for a month or a year, or of a credit pack.
func (c *PublicClient) PriceQuote(ctx context.Context, input PublicPricingQuoteInput, options *CallOptions) (PublicPricingQuoteResponse, error) {
	var result PublicPricingQuoteResponse
	data, err := c.call(ctx, "public.pricing_quote", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Docs calls GET /v1/public/docs. List the developer docs pages, with each page's title, path and last update.
func (c *PublicClient) Docs(ctx context.Context, input PublicDocsInput, options *CallOptions) (PublicDocsResponse, error) {
	var result PublicDocsResponse
	data, err := c.call(ctx, "public.docs", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SearchDocs calls GET /v1/public/docs/search. Search the developer docs.
func (c *PublicClient) SearchDocs(ctx context.Context, input PublicDocsSearchInput, options *CallOptions) (PublicDocsSearchResponse, error) {
	var result PublicDocsSearchResponse
	data, err := c.call(ctx, "public.docs_search", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// PublicDocsPage calls GET /v1/public/docs/{path}. Read a developer docs page as Markdown.
func (c *PublicClient) PublicDocsPage(ctx context.Context, input map[string]any, options *CallOptions) (PublicDocsPageResponse, error) {
	var result PublicDocsPageResponse
	data, err := c.call(ctx, "public.docs_page", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// HelpArticles calls GET /v1/public/help. List the help center's articles, with each article's title, path and last update.
func (c *PublicClient) HelpArticles(ctx context.Context, input PublicHelpInput, options *CallOptions) (PublicHelpResponse, error) {
	var result PublicHelpResponse
	data, err := c.call(ctx, "public.help", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SearchHelp calls GET /v1/public/help/search. Search the help center.
func (c *PublicClient) SearchHelp(ctx context.Context, input PublicHelpSearchInput, options *CallOptions) (PublicHelpSearchResponse, error) {
	var result PublicHelpSearchResponse
	data, err := c.call(ctx, "public.help_search", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// PublicHelpArticle calls GET /v1/public/help/{path}. Read a help center article as Markdown.
func (c *PublicClient) PublicHelpArticle(ctx context.Context, input map[string]any, options *CallOptions) (PublicHelpArticleResponse, error) {
	var result PublicHelpArticleResponse
	data, err := c.call(ctx, "public.help_article", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// OpenAPI calls GET /v1/public/openapi. Read the OpenAPI document for the AbuseZero API.
func (c *PublicClient) OpenAPI(ctx context.Context, input map[string]any, options *CallOptions) (PublicOpenAPIResponse, error) {
	var result PublicOpenAPIResponse
	data, err := c.call(ctx, "public.openapi", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Actions calls GET /v1/public/actions. List every API action with its MCP tool, CLI command and scope.
func (c *PublicClient) Actions(ctx context.Context, input map[string]any, options *CallOptions) (PublicActionsResponse, error) {
	var result PublicActionsResponse
	data, err := c.call(ctx, "public.actions", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ErrorCodes calls GET /v1/public/errors. List the API's error codes: what each means, whether to retry and what to do next.
func (c *PublicClient) ErrorCodes(ctx context.Context, input map[string]any, options *CallOptions) (PublicErrorsResponse, error) {
	var result PublicErrorsResponse
	data, err := c.call(ctx, "public.errors", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ReasonCodes calls GET /v1/public/reasons. List the reason codes evaluations return.
func (c *PublicClient) ReasonCodes(ctx context.Context, input map[string]any, options *CallOptions) (PublicReasonsResponse, error) {
	var result PublicReasonsResponse
	data, err := c.call(ctx, "public.reasons", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WebhookEventTypes calls GET /v1/public/webhook-events. List the webhook event types, each with an example payload.
func (c *PublicClient) WebhookEventTypes(ctx context.Context, input map[string]any, options *CallOptions) (PublicWebhookEventsResponse, error) {
	var result PublicWebhookEventsResponse
	data, err := c.call(ctx, "public.webhook_events", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Conditions calls GET /v1/public/conditions. List the conditions the rules in custom policies can use.
func (c *PublicClient) Conditions(ctx context.Context, input map[string]any, options *CallOptions) (PublicConditionsResponse, error) {
	var result PublicConditionsResponse
	data, err := c.call(ctx, "public.conditions", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Glossary calls GET /v1/public/glossary. Read the glossary and the developer FAQ.
func (c *PublicClient) Glossary(ctx context.Context, input map[string]any, options *CallOptions) (PublicGlossaryResponse, error) {
	var result PublicGlossaryResponse
	data, err := c.call(ctx, "public.glossary", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// LegalDocuments calls GET /v1/public/legal. List AbuseZero's legal documents, with each one's version and effective date.
func (c *PublicClient) LegalDocuments(ctx context.Context, input map[string]any, options *CallOptions) (PublicLegalResponse, error) {
	var result PublicLegalResponse
	data, err := c.call(ctx, "public.legal", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// LegalDocument calls GET /v1/public/legal/{document}. Read a legal document as Markdown.
func (c *PublicClient) LegalDocument(ctx context.Context, input map[string]any, options *CallOptions) (PublicLegalDocumentResponse, error) {
	var result PublicLegalDocumentResponse
	data, err := c.call(ctx, "public.legal_document", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Changelog calls GET /v1/public/changelog. Read the product changelog.
func (c *PublicClient) Changelog(ctx context.Context, input map[string]any, options *CallOptions) (PublicChangelogResponse, error) {
	var result PublicChangelogResponse
	data, err := c.call(ctx, "public.changelog", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// BlogPosts calls GET /v1/public/blog. List blog posts.
func (c *PublicClient) BlogPosts(ctx context.Context, input PublicBlogInput, options *CallOptions) (PublicBlogResponse, error) {
	var result PublicBlogResponse
	data, err := c.call(ctx, "public.blog", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// BlogPost calls GET /v1/public/blog/{slug}. Read a blog post as Markdown.
func (c *PublicClient) BlogPost(ctx context.Context, input map[string]any, options *CallOptions) (PublicBlogPostResponse, error) {
	var result PublicBlogPostResponse
	data, err := c.call(ctx, "public.blog_post", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Versions calls GET /v1/public/versions. Read the latest npm package version and the browser script address.
func (c *PublicClient) Versions(ctx context.Context, input map[string]any, options *CallOptions) (PublicVersionsResponse, error) {
	var result PublicVersionsResponse
	data, err := c.call(ctx, "public.versions", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ServiceStatus calls GET /v1/public/status. Read whether AbuseZero is operational or degraded.
func (c *PublicClient) ServiceStatus(ctx context.Context, input map[string]any, options *CallOptions) (PublicStatusResponse, error) {
	var result PublicStatusResponse
	data, err := c.call(ctx, "public.status", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Contacts calls GET /v1/public/contacts. Read how to reach AbuseZero's support, sales and security.
func (c *PublicClient) Contacts(ctx context.Context, input map[string]any, options *CallOptions) (PublicContactsResponse, error) {
	var result PublicContactsResponse
	data, err := c.call(ctx, "public.contacts", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ContactSales calls POST /v1/public/contact/sales. Send a message to AbuseZero's sales team about an Enterprise plan.
func (c *PublicClient) ContactSales(ctx context.Context, input PublicContactSalesInput, options *CallOptions) (PublicContactSalesResponse, error) {
	var result PublicContactSalesResponse
	data, err := c.call(ctx, "public.contact_sales", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ContactSupport calls POST /v1/public/contact/support. Send a request to AbuseZero support without an account.
func (c *PublicClient) ContactSupport(ctx context.Context, input PublicContactSupportInput, options *CallOptions) (PublicContactSupportResponse, error) {
	var result PublicContactSupportResponse
	data, err := c.call(ctx, "public.contact_support", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// StartOnboarding calls POST /v1/public/onboarding-sessions. Start agent onboarding for a new workspace: returns the approval link for the person, and a secret only this agent keeps.
func (c *PublicClient) StartOnboarding(ctx context.Context, input PublicOnboardingStartInput, options *CallOptions) (PublicOnboardingStartResponse, error) {
	var result PublicOnboardingStartResponse
	data, err := c.call(ctx, "public.onboarding_start", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// OnboardingStatus calls POST /v1/public/onboarding-sessions/{id}/status. Read whether the person finished onboarding, optionally waiting up to 25 seconds, and how to collect the sign-in.
func (c *PublicClient) OnboardingStatus(ctx context.Context, input PublicOnboardingStatusInput, options *CallOptions) (PublicOnboardingStatusResponse, error) {
	var result PublicOnboardingStatusResponse
	data, err := c.call(ctx, "public.onboarding_status", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Workspaces calls GET /v1/workspaces. List workspaces available to this identity.
func (c *Client) Workspaces(ctx context.Context, input map[string]any, options *CallOptions) (WorkspacesListResponse, error) {
	var result WorkspacesListResponse
	data, err := c.call(ctx, "workspaces.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Me calls GET /v1/me. Read who this connection acts for: the person, their role and permissions, the workspace and the granted scopes.
func (c *Client) Me(ctx context.Context, input map[string]any, options *CallOptions) (MeGetResponse, error) {
	var result MeGetResponse
	data, err := c.call(ctx, "me.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// APIKeys calls GET /v1/api-keys. List API keys with their type, scopes, last use and expiry. Secrets are never returned.
func (c *Client) APIKeys(ctx context.Context, input APIKeysListInput, options *CallOptions) (APIKeysListResponse, error) {
	var result APIKeysListResponse
	data, err := c.call(ctx, "api-keys.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CreateAPIKey calls POST /v1/api-keys. Create a secret or restricted key. Its secret is shown once, in the response.
func (c *Client) CreateAPIKey(ctx context.Context, input APIKeysCreateInput, options *CallOptions) (APIKeysCreateResponse, error) {
	var result APIKeysCreateResponse
	data, err := c.call(ctx, "api-keys.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateAPIKey calls PATCH /v1/api-keys/{id}. Change an API key's name, scopes or expiry.
func (c *Client) UpdateAPIKey(ctx context.Context, input APIKeysUpdateInput, options *CallOptions) (APIKeysUpdateResponse, error) {
	var result APIKeysUpdateResponse
	data, err := c.call(ctx, "api-keys.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RollAPIKey calls POST /v1/api-keys/{id}/roll. Roll an API key. The old key keeps working for 1 hour, 24 hours or 7 days, and the new secret is shown once.
func (c *Client) RollAPIKey(ctx context.Context, input APIKeysRollInput, options *CallOptions) (APIKeysRollResponse, error) {
	var result APIKeysRollResponse
	data, err := c.call(ctx, "api-keys.roll", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ShowAPIKeyReplacement calls POST /v1/api-keys/{id}/replacement. Show a key's replacement once, when AbuseZero replaced the key for you.
func (c *Client) ShowAPIKeyReplacement(ctx context.Context, input map[string]any, options *CallOptions) (APIKeysReplacementResponse, error) {
	var result APIKeysReplacementResponse
	data, err := c.call(ctx, "api-keys.replacement", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RevokeAPIKey calls POST /v1/api-keys/{id}/revoke. Revoke an API key. It stops working at once.
func (c *Client) RevokeAPIKey(ctx context.Context, input map[string]any, options *CallOptions) (APIKeysRevokeResponse, error) {
	var result APIKeysRevokeResponse
	data, err := c.call(ctx, "api-keys.revoke", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// MachineClients calls GET /v1/machine-clients. List machine clients.
func (c *Client) MachineClients(ctx context.Context, input MachineClientsListInput, options *CallOptions) (MachineClientsListResponse, error) {
	var result MachineClientsListResponse
	data, err := c.call(ctx, "machine-clients.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CreateMachineClient calls POST /v1/machine-clients. Create a machine client. Its secret is shown once, in the response.
func (c *Client) CreateMachineClient(ctx context.Context, input MachineClientsCreateInput, options *CallOptions) (MachineClientsCreateResponse, error) {
	var result MachineClientsCreateResponse
	data, err := c.call(ctx, "machine-clients.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RotateMachineClientSecret calls POST /v1/machine-clients/{id}/rotate-secret. Rotate a machine client's secret. The old secret stops working at once, and the new one is shown once.
func (c *Client) RotateMachineClientSecret(ctx context.Context, input map[string]any, options *CallOptions) (MachineClientsRotateSecretResponse, error) {
	var result MachineClientsRotateSecretResponse
	data, err := c.call(ctx, "machine-clients.rotate_secret", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RevokeMachineClient calls POST /v1/machine-clients/{id}/revoke. Revoke a machine client. Its secret and access tokens stop working at once.
func (c *Client) RevokeMachineClient(ctx context.Context, input map[string]any, options *CallOptions) (MachineClientsRevokeResponse, error) {
	var result MachineClientsRevokeResponse
	data, err := c.call(ctx, "machine-clients.revoke", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// AuthorizedApps calls GET /v1/authorized-apps. List the apps and agents members authorized with OAuth.
func (c *Client) AuthorizedApps(ctx context.Context, input AuthorizedAppsListInput, options *CallOptions) (AuthorizedAppsListResponse, error) {
	var result AuthorizedAppsListResponse
	data, err := c.call(ctx, "authorized-apps.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DisconnectAuthorizedApp calls POST /v1/authorized-apps/{id}/disconnect. Disconnect one of your authorized apps. It loses access at once.
func (c *Client) DisconnectAuthorizedApp(ctx context.Context, input map[string]any, options *CallOptions) (AuthorizedAppsDisconnectResponse, error) {
	var result AuthorizedAppsDisconnectResponse
	data, err := c.call(ctx, "authorized-apps.disconnect", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Members calls GET /v1/members. List the workspace's members with their roles and permissions.
func (c *Client) Members(ctx context.Context, input MembersListInput, options *CallOptions) (MembersListResponse, error) {
	var result MembersListResponse
	data, err := c.call(ctx, "members.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateMember calls PATCH /v1/members/{id}. Change a member's role or permissions.
func (c *Client) UpdateMember(ctx context.Context, input MembersUpdateInput, options *CallOptions) (MembersUpdateResponse, error) {
	var result MembersUpdateResponse
	data, err := c.call(ctx, "members.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RemoveMember calls DELETE /v1/members/{id}. Remove a member from the workspace. They lose access at once.
func (c *Client) RemoveMember(ctx context.Context, input map[string]any, options *CallOptions) (MembersRemoveResponse, error) {
	var result MembersRemoveResponse
	data, err := c.call(ctx, "members.remove", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Invitations calls GET /v1/invitations. List pending invitations.
func (c *Client) Invitations(ctx context.Context, input InvitationsListInput, options *CallOptions) (InvitationsListResponse, error) {
	var result InvitationsListResponse
	data, err := c.call(ctx, "invitations.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// InviteMember calls POST /v1/invitations. Invite someone to the workspace by email, up to the plan's team seats.
func (c *Client) InviteMember(ctx context.Context, input InvitationsCreateInput, options *CallOptions) (InvitationsCreateResponse, error) {
	var result InvitationsCreateResponse
	data, err := c.call(ctx, "invitations.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ResendInvitation calls POST /v1/invitations/{id}/resend. Send a pending invitation again, as a new invitation that replaces it.
func (c *Client) ResendInvitation(ctx context.Context, input map[string]any, options *CallOptions) (InvitationsResendResponse, error) {
	var result InvitationsResendResponse
	data, err := c.call(ctx, "invitations.resend", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CancelInvitation calls POST /v1/invitations/{id}/cancel. Cancel a pending invitation.
func (c *Client) CancelInvitation(ctx context.Context, input map[string]any, options *CallOptions) (InvitationsCancelResponse, error) {
	var result InvitationsCancelResponse
	data, err := c.call(ctx, "invitations.cancel", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WorkspaceSettings calls GET /v1/workspace. Read the workspace's settings: name, icon, time zone, website, product type and data retention.
func (c *Client) WorkspaceSettings(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceGetResponse, error) {
	var result WorkspaceGetResponse
	data, err := c.call(ctx, "workspace.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateWorkspaceSettings calls PATCH /v1/workspace. Rename the workspace, or change its website, product type or time zone.
func (c *Client) UpdateWorkspaceSettings(ctx context.Context, input WorkspaceUpdateInput, options *CallOptions) (WorkspaceUpdateResponse, error) {
	var result WorkspaceUpdateResponse
	data, err := c.call(ctx, "workspace.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UploadWorkspaceIcon calls PUT /v1/workspace/icon. Upload a new workspace icon. It's checked before it's used.
func (c *Client) UploadWorkspaceIcon(ctx context.Context, input WorkspaceIconUploadInput, options *CallOptions) (WorkspaceIconUploadResponse, error) {
	var result WorkspaceIconUploadResponse
	data, err := c.call(ctx, "workspace.icon_upload", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RemoveWorkspaceIcon calls DELETE /v1/workspace/icon. Remove the workspace icon.
func (c *Client) RemoveWorkspaceIcon(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceIconRemoveResponse, error) {
	var result WorkspaceIconRemoveResponse
	data, err := c.call(ctx, "workspace.icon_remove", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WorkspaceIconUpload calls GET /v1/workspace/icon/uploads/{id}. Read whether an uploaded workspace icon is ready.
func (c *Client) WorkspaceIconUpload(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceIconUploadGetResponse, error) {
	var result WorkspaceIconUploadGetResponse
	data, err := c.call(ctx, "workspace.icon_upload_get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SetDataRetention calls PUT /v1/workspace/data-retention. Change how long evaluation evidence is kept. Shortening it deletes older evidence for good.
func (c *Client) SetDataRetention(ctx context.Context, input WorkspaceRetentionInput, options *CallOptions) (WorkspaceRetentionResponse, error) {
	var result WorkspaceRetentionResponse
	data, err := c.call(ctx, "workspace.retention", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SetupProgress calls GET /v1/workspace/setup. Read setup progress: which Get started steps are done.
func (c *Client) SetupProgress(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceSetupResponse, error) {
	var result WorkspaceSetupResponse
	data, err := c.call(ctx, "workspace.setup", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// FinishSetup calls POST /v1/workspace/setup/finish. Finish setting up the workspace, with its website, product type and time zone. Workspace owner only.
func (c *Client) FinishSetup(ctx context.Context, input WorkspaceSetupFinishInput, options *CallOptions) (WorkspaceSetupFinishResponse, error) {
	var result WorkspaceSetupFinishResponse
	data, err := c.call(ctx, "workspace.setup_finish", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DismissSetupGuide calls POST /v1/workspace/setup/dismiss. Stop Get started opening by itself for you.
func (c *Client) DismissSetupGuide(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceSetupDismissResponse, error) {
	var result WorkspaceSetupDismissResponse
	data, err := c.call(ctx, "workspace.setup_dismiss", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// LeaveWorkspace calls POST /v1/workspace/leave. Leave the workspace. Your access ends, and so does this sign-in's.
func (c *Client) LeaveWorkspace(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceLeaveResponse, error) {
	var result WorkspaceLeaveResponse
	data, err := c.call(ctx, "workspace.leave", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Evaluations calls GET /v1/evaluations. List evaluations.
func (c *Client) Evaluations(ctx context.Context, input EvaluationsListInput, options *CallOptions) (EvaluationsListResponse, error) {
	var result EvaluationsListResponse
	data, err := c.call(ctx, "evaluations.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Evaluate calls POST /v1/evaluations. Evaluate a benefit claim.
func (c *Client) Evaluate(ctx context.Context, input EvaluationsCreateInput, options *CallOptions) (any, error) {
	var result any
	data, err := c.call(ctx, "evaluations.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// EvaluationFilters calls GET /v1/evaluation-filters. Read the benefits, results and reasons found in recent evaluations.
func (c *Client) EvaluationFilters(ctx context.Context, input EvaluationsFiltersInput, options *CallOptions) (EvaluationsFiltersResponse, error) {
	var result EvaluationsFiltersResponse
	data, err := c.call(ctx, "evaluations.filters", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetEvaluation calls GET /v1/evaluations/{id}. Read an evaluation.
func (c *Client) GetEvaluation(ctx context.Context, input map[string]any, options *CallOptions) (EvaluationsGetResponse, error) {
	var result EvaluationsGetResponse
	data, err := c.call(ctx, "evaluations.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// EvaluationGraph calls GET /v1/evaluations/{id}/graph. Read an evaluation association graph.
func (c *Client) EvaluationGraph(ctx context.Context, input map[string]any, options *CallOptions) (EvaluationsGraphResponse, error) {
	var result EvaluationsGraphResponse
	data, err := c.call(ctx, "evaluations.graph", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// EvaluationEvidence calls GET /v1/evaluations/{id}/evidence. Read retained evaluation evidence.
func (c *Client) EvaluationEvidence(ctx context.Context, input map[string]any, options *CallOptions) (EvaluationsEvidenceResponse, error) {
	var result EvaluationsEvidenceResponse
	data, err := c.call(ctx, "evaluations.evidence", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Outcomes calls GET /v1/outcomes. List outcomes.
func (c *Client) Outcomes(ctx context.Context, input OutcomesListInput, options *CallOptions) (OutcomesListResponse, error) {
	var result OutcomesListResponse
	data, err := c.call(ctx, "outcomes.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetOutcome calls GET /v1/outcomes/{id}. Read an outcome.
func (c *Client) GetOutcome(ctx context.Context, input map[string]any, options *CallOptions) (OutcomesGetResponse, error) {
	var result OutcomesGetResponse
	data, err := c.call(ctx, "outcomes.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Outcome calls POST /v1/evaluations/{id}/outcome. Record a known outcome, including a benefit delivered after a deny.
func (c *Client) Outcome(ctx context.Context, input OutcomesReportInput, options *CallOptions) (OutcomesReportResponse, error) {
	var result OutcomesReportResponse
	data, err := c.call(ctx, "outcomes.report", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DryRunEvaluation calls POST /v1/evaluations/dry-run. Run a dry-run evaluation without writing claim history. Completed dry-run evaluations consume usage.
func (c *Client) DryRunEvaluation(ctx context.Context, input EvaluationsDryRunInput, options *CallOptions) (any, error) {
	var result any
	data, err := c.call(ctx, "evaluations.dry_run", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Simulate calls POST /v1/simulations. Simulate a retained sequence.
func (c *Client) Simulate(ctx context.Context, input SimulationsSequenceInput, options *CallOptions) (SimulationsSequenceResponse, error) {
	var result SimulationsSequenceResponse
	data, err := c.call(ctx, "simulations.sequence", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Replay calls POST /v1/replays. Replay retained evaluations against a saved policy.
func (c *Client) Replay(ctx context.Context, input SimulationsReplayInput, options *CallOptions) (SimulationsReplayResponse, error) {
	var result SimulationsReplayResponse
	data, err := c.call(ctx, "simulations.replay", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Benefits calls GET /v1/benefits. List benefits.
func (c *Client) Benefits(ctx context.Context, input BenefitsListInput, options *CallOptions) (BenefitsListResponse, error) {
	var result BenefitsListResponse
	data, err := c.call(ctx, "benefits.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CreateBenefit calls POST /v1/benefits. Create a benefit.
func (c *Client) CreateBenefit(ctx context.Context, input BenefitsCreateInput, options *CallOptions) (BenefitsCreateResponse, error) {
	var result BenefitsCreateResponse
	data, err := c.call(ctx, "benefits.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetBenefit calls GET /v1/benefits/{id}. Read a benefit.
func (c *Client) GetBenefit(ctx context.Context, input map[string]any, options *CallOptions) (BenefitsGetResponse, error) {
	var result BenefitsGetResponse
	data, err := c.call(ctx, "benefits.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateBenefit calls PATCH /v1/benefits/{id}. Update a benefit.
func (c *Client) UpdateBenefit(ctx context.Context, input BenefitsUpdateInput, options *CallOptions) (BenefitsUpdateResponse, error) {
	var result BenefitsUpdateResponse
	data, err := c.call(ctx, "benefits.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DeleteBenefit calls DELETE /v1/benefits/{id}. Permanently delete a benefit. This cannot be reversed.
func (c *Client) DeleteBenefit(ctx context.Context, input map[string]any, options *CallOptions) (BenefitsDeleteResponse, error) {
	var result BenefitsDeleteResponse
	data, err := c.call(ctx, "benefits.delete", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Policies calls GET /v1/policies. List policies.
func (c *Client) Policies(ctx context.Context, input PoliciesListInput, options *CallOptions) (PoliciesListResponse, error) {
	var result PoliciesListResponse
	data, err := c.call(ctx, "policies.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CreatePolicy calls POST /v1/policies. Create a policy.
func (c *Client) CreatePolicy(ctx context.Context, input PoliciesCreateInput, options *CallOptions) (PoliciesCreateResponse, error) {
	var result PoliciesCreateResponse
	data, err := c.call(ctx, "policies.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetPolicy calls GET /v1/policies/{id}. Read a policy.
func (c *Client) GetPolicy(ctx context.Context, input map[string]any, options *CallOptions) (PoliciesGetResponse, error) {
	var result PoliciesGetResponse
	data, err := c.call(ctx, "policies.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdatePolicy calls PATCH /v1/policies/{id}. Update a policy.
func (c *Client) UpdatePolicy(ctx context.Context, input PoliciesUpdateInput, options *CallOptions) (PoliciesUpdateResponse, error) {
	var result PoliciesUpdateResponse
	data, err := c.call(ctx, "policies.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DeletePolicy calls DELETE /v1/policies/{id}. Permanently delete an unused custom policy. This cannot be reversed.
func (c *Client) DeletePolicy(ctx context.Context, input map[string]any, options *CallOptions) (PoliciesDeleteResponse, error) {
	var result PoliciesDeleteResponse
	data, err := c.call(ctx, "policies.delete", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DuplicatePolicy calls POST /v1/policies/{id}/duplicate. Duplicate a policy.
func (c *Client) DuplicatePolicy(ctx context.Context, input PoliciesDuplicateInput, options *CallOptions) (PoliciesDuplicateResponse, error) {
	var result PoliciesDuplicateResponse
	data, err := c.call(ctx, "policies.duplicate", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// PolicyVersions calls GET /v1/policies/{id}/versions. List saved policy versions.
func (c *Client) PolicyVersions(ctx context.Context, input map[string]any, options *CallOptions) (PoliciesVersionsResponse, error) {
	var result PoliciesVersionsResponse
	data, err := c.call(ctx, "policies.versions", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetPolicyVersion calls GET /v1/policies/{id}/versions/{version}. Read a saved policy version.
func (c *Client) GetPolicyVersion(ctx context.Context, input map[string]any, options *CallOptions) (PoliciesVersionResponse, error) {
	var result PoliciesVersionResponse
	data, err := c.call(ctx, "policies.version", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// PolicySignals calls GET /v1/policy-signals. List signals available to policy rules.
func (c *Client) PolicySignals(ctx context.Context, input map[string]any, options *CallOptions) (PoliciesSignalsResponse, error) {
	var result PoliciesSignalsResponse
	data, err := c.call(ctx, "policies.signals", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// AccessLists calls GET /v1/access-lists. List access lists.
func (c *Client) AccessLists(ctx context.Context, input AccessListsListInput, options *CallOptions) (AccessListsListResponse, error) {
	var result AccessListsListResponse
	data, err := c.call(ctx, "access-lists.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// AddAccessListEntries calls POST /v1/access-lists. Create an access list entry.
func (c *Client) AddAccessListEntries(ctx context.Context, input AccessListsCreateInput, options *CallOptions) (AccessListsCreateResponse, error) {
	var result AccessListsCreateResponse
	data, err := c.call(ctx, "access-lists.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetAccessListEntry calls GET /v1/access-lists/{id}. Read an access list entry.
func (c *Client) GetAccessListEntry(ctx context.Context, input map[string]any, options *CallOptions) (AccessListsGetResponse, error) {
	var result AccessListsGetResponse
	data, err := c.call(ctx, "access-lists.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DeleteAccessListEntry calls DELETE /v1/access-lists/{id}. Permanently delete an access list entry. This cannot be reversed.
func (c *Client) DeleteAccessListEntry(ctx context.Context, input map[string]any, options *CallOptions) (AccessListsDeleteResponse, error) {
	var result AccessListsDeleteResponse
	data, err := c.call(ctx, "access-lists.delete", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateAccessListEntry calls PATCH /v1/access-lists/{id}. Update an access list entry's list, benefits, reason or expiry.
func (c *Client) UpdateAccessListEntry(ctx context.Context, input AccessListsUpdateInput, options *CallOptions) (AccessListsUpdateResponse, error) {
	var result AccessListsUpdateResponse
	data, err := c.call(ctx, "access-lists.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// LookUpIdentifier calls POST /v1/identifiers/lookup. Look up an ID, email address or IP address.
func (c *Client) LookUpIdentifier(ctx context.Context, input IdentifiersLookupInput, options *CallOptions) (IdentifiersLookupResponse, error) {
	var result IdentifiersLookupResponse
	data, err := c.call(ctx, "identifiers.lookup", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CustomIdentifiers calls GET /v1/custom-identifiers. List custom identifiers.
func (c *Client) CustomIdentifiers(ctx context.Context, input CustomIdentifiersListInput, options *CallOptions) (CustomIdentifiersListResponse, error) {
	var result CustomIdentifiersListResponse
	data, err := c.call(ctx, "custom-identifiers.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RegisterCustomIdentifier calls POST /v1/custom-identifiers. Create a custom identifier.
func (c *Client) RegisterCustomIdentifier(ctx context.Context, input CustomIdentifiersCreateInput, options *CallOptions) (CustomIdentifiersCreateResponse, error) {
	var result CustomIdentifiersCreateResponse
	data, err := c.call(ctx, "custom-identifiers.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetCustomIdentifier calls GET /v1/custom-identifiers/{id}. Read a custom identifier.
func (c *Client) GetCustomIdentifier(ctx context.Context, input map[string]any, options *CallOptions) (CustomIdentifiersGetResponse, error) {
	var result CustomIdentifiersGetResponse
	data, err := c.call(ctx, "custom-identifiers.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DeleteCustomIdentifier calls DELETE /v1/custom-identifiers/{id}. Permanently delete a custom identifier. This cannot be reversed.
func (c *Client) DeleteCustomIdentifier(ctx context.Context, input map[string]any, options *CallOptions) (CustomIdentifiersDeleteResponse, error) {
	var result CustomIdentifiersDeleteResponse
	data, err := c.call(ctx, "custom-identifiers.delete", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Imports calls GET /v1/imports. List imports.
func (c *Client) Imports(ctx context.Context, input ImportsListInput, options *CallOptions) (ImportsListResponse, error) {
	var result ImportsListResponse
	data, err := c.call(ctx, "imports.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ImportClaimHistory calls POST /v1/imports. Import claim history: benefits you granted before AbuseZero. The rows wait for review unless the request sends confirm.
func (c *Client) ImportClaimHistory(ctx context.Context, input ImportsCreateInput, options *CallOptions) (ImportsCreateResponse, error) {
	var result ImportsCreateResponse
	data, err := c.call(ctx, "imports.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CreateCSVImport calls POST /v1/imports/csv. Start a claim history import from a CSV file, and get its upload link.
func (c *Client) CreateCSVImport(ctx context.Context, input ImportsCSVInput, options *CallOptions) (ImportsCSVResponse, error) {
	var result ImportsCSVResponse
	data, err := c.call(ctx, "imports.csv", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetImport calls GET /v1/imports/{id}. Read an import: its state, its rows' counts and its progress.
func (c *Client) GetImport(ctx context.Context, input map[string]any, options *CallOptions) (ImportsGetResponse, error) {
	var result ImportsGetResponse
	data, err := c.call(ctx, "imports.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ListImportProblems calls GET /v1/imports/{id}/problems. List an import's rows that can't be added, with each row's problem and field.
func (c *Client) ListImportProblems(ctx context.Context, input ImportsProblemsInput, options *CallOptions) (ImportsProblemsResponse, error) {
	var result ImportsProblemsResponse
	data, err := c.call(ctx, "imports.problems", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ImportProblemsDownload calls GET /v1/imports/{id}/problems/download. Create a short-lived download link for an import's rows with problems.
func (c *Client) ImportProblemsDownload(ctx context.Context, input map[string]any, options *CallOptions) (ImportsProblemsDownloadResponse, error) {
	var result ImportsProblemsDownloadResponse
	data, err := c.call(ctx, "imports.problems_download", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ConfirmImport calls POST /v1/imports/{id}/confirm. Add an import that's ready to review. Rows that need attention are skipped.
func (c *Client) ConfirmImport(ctx context.Context, input map[string]any, options *CallOptions) (ImportsConfirmResponse, error) {
	var result ImportsConfirmResponse
	data, err := c.call(ctx, "imports.confirm", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DiscardImport calls POST /v1/imports/{id}/discard. Discard an import before it's confirmed. Nothing from it is added, and its rows are deleted.
func (c *Client) DiscardImport(ctx context.Context, input map[string]any, options *CallOptions) (ImportsDiscardResponse, error) {
	var result ImportsDiscardResponse
	data, err := c.call(ctx, "imports.discard", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// FixImportRow calls POST /v1/imports/{id}/fix. Fix one row that needs attention in an import that's ready to review.
func (c *Client) FixImportRow(ctx context.Context, input ImportsFixRowInput, options *CallOptions) (ImportsFixRowResponse, error) {
	var result ImportsFixRowResponse
	data, err := c.call(ctx, "imports.fix_row", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SkipImportRow calls POST /v1/imports/{id}/skip. Skip one row that needs attention in an import that's ready to review.
func (c *Client) SkipImportRow(ctx context.Context, input ImportsSkipRowInput, options *CallOptions) (ImportsSkipRowResponse, error) {
	var result ImportsSkipRowResponse
	data, err := c.call(ctx, "imports.skip_row", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Exports calls GET /v1/exports. List exports.
func (c *Client) Exports(ctx context.Context, input ExportsListInput, options *CallOptions) (ExportsListResponse, error) {
	var result ExportsListResponse
	data, err := c.call(ctx, "exports.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RequestExport calls POST /v1/exports. Create an evaluation export.
func (c *Client) RequestExport(ctx context.Context, input ExportsCreateInput, options *CallOptions) (ExportsCreateResponse, error) {
	var result ExportsCreateResponse
	data, err := c.call(ctx, "exports.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetExport calls GET /v1/exports/{id}. Read export status.
func (c *Client) GetExport(ctx context.Context, input map[string]any, options *CallOptions) (ExportsGetResponse, error) {
	var result ExportsGetResponse
	data, err := c.call(ctx, "exports.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ExportDownload calls GET /v1/exports/{id}/download. Create a short-lived export download URL.
func (c *Client) ExportDownload(ctx context.Context, input map[string]any, options *CallOptions) (ExportsDownloadResponse, error) {
	var result ExportsDownloadResponse
	data, err := c.call(ctx, "exports.download", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RequestWorkspaceExport calls POST /v1/workspace-exports. Request a workspace data export, with the password for its ZIP files.
func (c *Client) RequestWorkspaceExport(ctx context.Context, input WorkspaceExportsCreateInput, options *CallOptions) (WorkspaceExportsCreateResponse, error) {
	var result WorkspaceExportsCreateResponse
	data, err := c.call(ctx, "workspace-exports.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WorkspaceExport calls GET /v1/workspace-exports. Read the workspace's latest data export, and when the next one can be requested.
func (c *Client) WorkspaceExport(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceExportsStatusResponse, error) {
	var result WorkspaceExportsStatusResponse
	data, err := c.call(ctx, "workspace-exports.status", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WorkspaceExportDownload calls GET /v1/workspace-exports/{id}/download. Get short-lived links to the files of a workspace data export you requested.
func (c *Client) WorkspaceExportDownload(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceExportsDownloadResponse, error) {
	var result WorkspaceExportsDownloadResponse
	data, err := c.call(ctx, "workspace-exports.download", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DeleteWorkspaceExport calls DELETE /v1/workspace-exports/{id}. Delete a workspace data export. Its files are removed for good.
func (c *Client) DeleteWorkspaceExport(ctx context.Context, input map[string]any, options *CallOptions) (WorkspaceExportsDeleteResponse, error) {
	var result WorkspaceExportsDeleteResponse
	data, err := c.call(ctx, "workspace-exports.delete", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DryRunDataDeletion calls POST /v1/data-deletions/dry-run. Run a dry run of user-history deletion.
func (c *Client) DryRunDataDeletion(ctx context.Context, input DeletionsDryRunInput, options *CallOptions) (DeletionsDryRunResponse, error) {
	var result DeletionsDryRunResponse
	data, err := c.call(ctx, "deletions.dry_run", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DataDeletions calls GET /v1/data-deletions. List data deletions.
func (c *Client) DataDeletions(ctx context.Context, input DeletionsListInput, options *CallOptions) (DeletionsListResponse, error) {
	var result DeletionsListResponse
	data, err := c.call(ctx, "deletions.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RequestDataDeletion calls POST /v1/data-deletions. Permanently delete user history. This cannot be reversed.
func (c *Client) RequestDataDeletion(ctx context.Context, input DeletionsCreateInput, options *CallOptions) (DeletionsCreateResponse, error) {
	var result DeletionsCreateResponse
	data, err := c.call(ctx, "deletions.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetDataDeletion calls GET /v1/data-deletions/{id}. Read data-deletion status.
func (c *Client) GetDataDeletion(ctx context.Context, input map[string]any, options *CallOptions) (DeletionsGetResponse, error) {
	var result DeletionsGetResponse
	data, err := c.call(ctx, "deletions.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Analytics calls GET /v1/analytics. Read analytics.
func (c *Client) Analytics(ctx context.Context, input AnalyticsListInput, options *CallOptions) (AnalyticsListResponse, error) {
	var result AnalyticsListResponse
	data, err := c.call(ctx, "analytics.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Overview calls GET /v1/overview. Read the overview of the last 24 hours.
func (c *Client) Overview(ctx context.Context, input map[string]any, options *CallOptions) (OverviewGetResponse, error) {
	var result OverviewGetResponse
	data, err := c.call(ctx, "overview.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// BenefitValue calls GET /v1/analytics/value. Read benefit value: what delivered claims were worth and what AbuseZero saved, in any supported currency.
func (c *Client) BenefitValue(ctx context.Context, input AnalyticsValueInput, options *CallOptions) (AnalyticsValueResponse, error) {
	var result AnalyticsValueResponse
	data, err := c.call(ctx, "analytics.value", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DenialReasons calls GET /v1/analytics/denial-reasons. Read why claims were denied: denied claims by reason, largest first.
func (c *Client) DenialReasons(ctx context.Context, input AnalyticsDenialReasonsInput, options *CallOptions) (AnalyticsDenialReasonsResponse, error) {
	var result AnalyticsDenialReasonsResponse
	data, err := c.call(ctx, "analytics.denial_reasons", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RequestLogs calls GET /v1/request-logs. Read request logs.
func (c *Client) RequestLogs(ctx context.Context, input RequestLogsListInput, options *CallOptions) (RequestLogsListResponse, error) {
	var result RequestLogsListResponse
	data, err := c.call(ctx, "request-logs.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetRequestLog calls GET /v1/request-logs/{id}. Read a request log.
func (c *Client) GetRequestLog(ctx context.Context, input map[string]any, options *CallOptions) (RequestLogsGetResponse, error) {
	var result RequestLogsGetResponse
	data, err := c.call(ctx, "request-logs.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// AuditLogs calls GET /v1/audit-logs. Read audit logs.
func (c *Client) AuditLogs(ctx context.Context, input AuditLogsListInput, options *CallOptions) (AuditLogsListResponse, error) {
	var result AuditLogsListResponse
	data, err := c.call(ctx, "audit-logs.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetAuditLog calls GET /v1/audit-logs/{id}. Read an audit log.
func (c *Client) GetAuditLog(ctx context.Context, input map[string]any, options *CallOptions) (AuditLogsGetResponse, error) {
	var result AuditLogsGetResponse
	data, err := c.call(ctx, "audit-logs.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// AccountStatus calls GET /v1/account-status. Read account status: credits, plan, limits and auto-reload.
func (c *Client) AccountStatus(ctx context.Context, input map[string]any, options *CallOptions) (AccountStatusGetResponse, error) {
	var result AccountStatusGetResponse
	data, err := c.call(ctx, "account_status.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// IntegrationStatus calls GET /v1/integration-status. Read integration status: whether keys, the browser script, evaluations, outcomes, webhooks and credits are working, and what to do next.
func (c *Client) IntegrationStatus(ctx context.Context, input map[string]any, options *CallOptions) (IntegrationStatusGetResponse, error) {
	var result IntegrationStatusGetResponse
	data, err := c.call(ctx, "integration_status.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Events calls GET /v1/events. Read account events after a cursor, oldest first, waiting for the next one if you ask.
func (c *Client) Events(ctx context.Context, input EventsListInput, options *CallOptions) (EventsListResponse, error) {
	var result EventsListResponse
	data, err := c.call(ctx, "events.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// BrowserCollection calls GET /v1/browser-collection. Read browser protection: your sites, where the script was seen, and what recent checks received.
func (c *Client) BrowserCollection(ctx context.Context, input map[string]any, options *CallOptions) (BrowserCollectionGetResponse, error) {
	var result BrowserCollectionGetResponse
	data, err := c.call(ctx, "browser_collection.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateBrowserCollection calls PATCH /v1/browser-collection. Update the sites listed for browser protection.
func (c *Client) UpdateBrowserCollection(ctx context.Context, input BrowserCollectionUpdateInput, options *CallOptions) (BrowserCollectionUpdateResponse, error) {
	var result BrowserCollectionUpdateResponse
	data, err := c.call(ctx, "browser_collection.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// BrowserScript calls GET /v1/browser-collection/script. Get the browser script tag, with your publishable key filled in.
func (c *Client) BrowserScript(ctx context.Context, input map[string]any, options *CallOptions) (BrowserCollectionScriptResponse, error) {
	var result BrowserCollectionScriptResponse
	data, err := c.call(ctx, "browser_collection.script", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WebhookEndpoints calls GET /v1/webhook-endpoints. List webhook endpoints.
func (c *Client) WebhookEndpoints(ctx context.Context, input WebhookEndpointsListInput, options *CallOptions) (WebhookEndpointsListResponse, error) {
	var result WebhookEndpointsListResponse
	data, err := c.call(ctx, "webhook-endpoints.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CreateWebhookEndpoint calls POST /v1/webhook-endpoints. Add a webhook endpoint. Its signing secret isn't in the response.
func (c *Client) CreateWebhookEndpoint(ctx context.Context, input WebhookEndpointsCreateInput, options *CallOptions) (WebhookEndpointsCreateResponse, error) {
	var result WebhookEndpointsCreateResponse
	data, err := c.call(ctx, "webhook-endpoints.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetWebhookEndpoint calls GET /v1/webhook-endpoints/{id}. Read a webhook endpoint.
func (c *Client) GetWebhookEndpoint(ctx context.Context, input map[string]any, options *CallOptions) (WebhookEndpointsGetResponse, error) {
	var result WebhookEndpointsGetResponse
	data, err := c.call(ctx, "webhook-endpoints.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateWebhookEndpoint calls PATCH /v1/webhook-endpoints/{id}. Change a webhook endpoint's URL, events or benefits.
func (c *Client) UpdateWebhookEndpoint(ctx context.Context, input WebhookEndpointsUpdateInput, options *CallOptions) (WebhookEndpointsUpdateResponse, error) {
	var result WebhookEndpointsUpdateResponse
	data, err := c.call(ctx, "webhook-endpoints.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DeleteWebhookEndpoint calls DELETE /v1/webhook-endpoints/{id}. Permanently delete a webhook endpoint and its delivery log. This cannot be reversed.
func (c *Client) DeleteWebhookEndpoint(ctx context.Context, input map[string]any, options *CallOptions) (WebhookEndpointsDeleteResponse, error) {
	var result WebhookEndpointsDeleteResponse
	data, err := c.call(ctx, "webhook-endpoints.delete", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SendWebhookTest calls POST /v1/webhook-endpoints/{id}/test. Send a test event to a webhook endpoint.
func (c *Client) SendWebhookTest(ctx context.Context, input WebhookEndpointsTestInput, options *CallOptions) (WebhookEndpointsTestResponse, error) {
	var result WebhookEndpointsTestResponse
	data, err := c.call(ctx, "webhook-endpoints.test", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// EnableWebhookEndpoint calls POST /v1/webhook-endpoints/{id}/enable. Enable a webhook endpoint.
func (c *Client) EnableWebhookEndpoint(ctx context.Context, input map[string]any, options *CallOptions) (WebhookEndpointsEnableResponse, error) {
	var result WebhookEndpointsEnableResponse
	data, err := c.call(ctx, "webhook-endpoints.enable", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DisableWebhookEndpoint calls POST /v1/webhook-endpoints/{id}/disable. Disable a webhook endpoint. It gets only test events until it's enabled again.
func (c *Client) DisableWebhookEndpoint(ctx context.Context, input map[string]any, options *CallOptions) (WebhookEndpointsDisableResponse, error) {
	var result WebhookEndpointsDisableResponse
	data, err := c.call(ctx, "webhook-endpoints.disable", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WebhookSigningSecret calls GET /v1/webhook-endpoints/{id}/secret. Show a webhook endpoint's signing secret. Every read is recorded in the audit trail.
func (c *Client) WebhookSigningSecret(ctx context.Context, input map[string]any, options *CallOptions) (WebhookEndpointsSecretResponse, error) {
	var result WebhookEndpointsSecretResponse
	data, err := c.call(ctx, "webhook-endpoints.secret", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RotateWebhookSigningSecret calls POST /v1/webhook-endpoints/{id}/rotate-secret. Rotate a webhook endpoint's signing secret. The previous secret keeps signing for 24 hours.
func (c *Client) RotateWebhookSigningSecret(ctx context.Context, input map[string]any, options *CallOptions) (WebhookEndpointsRotateSecretResponse, error) {
	var result WebhookEndpointsRotateSecretResponse
	data, err := c.call(ctx, "webhook-endpoints.rotate_secret", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WebhookUsage calls GET /v1/webhook-usage. Read today's webhook deliveries against the daily limit.
func (c *Client) WebhookUsage(ctx context.Context, input map[string]any, options *CallOptions) (WebhookDeliveriesUsageResponse, error) {
	var result WebhookDeliveriesUsageResponse
	data, err := c.call(ctx, "webhook-deliveries.usage", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// WebhookDeliveries calls GET /v1/webhook-deliveries. List webhook deliveries.
func (c *Client) WebhookDeliveries(ctx context.Context, input WebhookDeliveriesListInput, options *CallOptions) (WebhookDeliveriesListResponse, error) {
	var result WebhookDeliveriesListResponse
	data, err := c.call(ctx, "webhook-deliveries.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetWebhookDelivery calls GET /v1/webhook-deliveries/{id}. Read a webhook delivery with its request body and attempts.
func (c *Client) GetWebhookDelivery(ctx context.Context, input map[string]any, options *CallOptions) (WebhookDeliveriesGetResponse, error) {
	var result WebhookDeliveriesGetResponse
	data, err := c.call(ctx, "webhook-deliveries.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ResendWebhookDelivery calls POST /v1/webhook-deliveries/{id}/resend. Send a webhook delivery's event again, to the endpoint's current URL or the URL it was first sent to.
func (c *Client) ResendWebhookDelivery(ctx context.Context, input WebhookDeliveriesResendInput, options *CallOptions) (WebhookDeliveriesResendResponse, error) {
	var result WebhookDeliveriesResendResponse
	data, err := c.call(ctx, "webhook-deliveries.resend", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Plan calls GET /v1/billing/plan. Read the plan, its billing period and renewal date, and the workspace's limits.
func (c *Client) Plan(ctx context.Context, input map[string]any, options *CallOptions) (BillingPlanResponse, error) {
	var result BillingPlanResponse
	data, err := c.call(ctx, "billing.plan", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Usage calls GET /v1/billing/usage. Read credits and usage: allowance credits, pay-as-you-go credits, protected benefits and team seats.
func (c *Client) Usage(ctx context.Context, input map[string]any, options *CallOptions) (BillingUsageResponse, error) {
	var result BillingUsageResponse
	data, err := c.call(ctx, "billing.usage", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CreditHistory calls GET /v1/billing/credit-history. List credit history: credits added, used and removed.
func (c *Client) CreditHistory(ctx context.Context, input BillingCreditHistoryInput, options *CallOptions) (BillingCreditHistoryResponse, error) {
	var result BillingCreditHistoryResponse
	data, err := c.call(ctx, "billing.credit_history", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// DailyUsage calls GET /v1/billing/daily-usage. Read credits used each day of the last 30 days.
func (c *Client) DailyUsage(ctx context.Context, input map[string]any, options *CallOptions) (BillingDailyUsageResponse, error) {
	var result BillingDailyUsageResponse
	data, err := c.call(ctx, "billing.daily_usage", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// EnterprisePlans calls GET /v1/billing/enterprise-plans. List the workspace's Enterprise plans and what each one adds.
func (c *Client) EnterprisePlans(ctx context.Context, input map[string]any, options *CallOptions) (BillingEnterprisePlansResponse, error) {
	var result BillingEnterprisePlansResponse
	data, err := c.call(ctx, "billing.enterprise_plans", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// PreviewEnterprisePlanCode calls POST /v1/billing/enterprise-plan-codes/preview. Preview what an Enterprise plan code adds, without redeeming it.
func (c *Client) PreviewEnterprisePlanCode(ctx context.Context, input BillingCodesPreviewInput, options *CallOptions) (BillingCodesPreviewResponse, error) {
	var result BillingCodesPreviewResponse
	data, err := c.call(ctx, "billing.codes.preview", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpcomingCharge calls GET /v1/billing/upcoming-charge. Read the next charge: its amount, date and payment method.
func (c *Client) UpcomingCharge(ctx context.Context, input map[string]any, options *CallOptions) (BillingUpcomingResponse, error) {
	var result BillingUpcomingResponse
	data, err := c.call(ctx, "billing.upcoming", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Invoices calls GET /v1/billing/invoices. List invoices, newest first, each with its secure payment page link.
func (c *Client) Invoices(ctx context.Context, input BillingInvoicesInput, options *CallOptions) (BillingInvoicesResponse, error) {
	var result BillingInvoicesResponse
	data, err := c.call(ctx, "billing.invoices", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// PaymentMethods calls GET /v1/billing/payment-methods. List saved payment methods: the default first, then the backup payment methods in order.
func (c *Client) PaymentMethods(ctx context.Context, input map[string]any, options *CallOptions) (BillingPaymentMethodsResponse, error) {
	var result BillingPaymentMethodsResponse
	data, err := c.call(ctx, "billing.payment_methods", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// BillingDetails calls GET /v1/billing/details. Read billing details: name, billing email, address and tax ID.
func (c *Client) BillingDetails(ctx context.Context, input map[string]any, options *CallOptions) (BillingDetailsResponse, error) {
	var result BillingDetailsResponse
	data, err := c.call(ctx, "billing.details", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateBillingDetails calls PUT /v1/billing/details. Save billing details: name, billing email, address and tax ID.
func (c *Client) UpdateBillingDetails(ctx context.Context, input BillingDetailsUpdateInput, options *CallOptions) (BillingDetailsUpdateResponse, error) {
	var result BillingDetailsUpdateResponse
	data, err := c.call(ctx, "billing.details.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// GetPayment calls GET /v1/billing/payments/{id}. Read a payment's progress.
func (c *Client) GetPayment(ctx context.Context, input map[string]any, options *CallOptions) (BillingPaymentsGetResponse, error) {
	var result BillingPaymentsGetResponse
	data, err := c.call(ctx, "billing.payments.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// QuotePlanChange calls POST /v1/billing/quotes. Get a quote for a plan or billing period change: what it costs now and at renewal. Nothing changes.
func (c *Client) QuotePlanChange(ctx context.Context, input BillingQuotesCreateInput, options *CallOptions) (BillingQuotesCreateResponse, error) {
	var result BillingQuotesCreateResponse
	data, err := c.call(ctx, "billing.quotes.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SetLowCreditAlert calls PUT /v1/billing/low-credit-alert. Set the low-credit alert level, or clear it to be alerted by your usage pace.
func (c *Client) SetLowCreditAlert(ctx context.Context, input BillingLowCreditAlertUpdateInput, options *CallOptions) (BillingLowCreditAlertUpdateResponse, error) {
	var result BillingLowCreditAlertUpdateResponse
	data, err := c.call(ctx, "billing.low_credit_alert.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ChooseWhatStaysActive calls POST /v1/billing/over-limit. Choose what stays active over your plan limits: keep one benefit or member active and pause another.
func (c *Client) ChooseWhatStaysActive(ctx context.Context, input BillingOverLimitChooseInput, options *CallOptions) (BillingOverLimitChooseResponse, error) {
	var result BillingOverLimitChooseResponse
	data, err := c.call(ctx, "billing.over_limit.choose", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SetAutoReload calls PUT /v1/billing/auto-reload. Set up, change or turn off auto-reload. When it's on, it charges the default payment method on its own.
func (c *Client) SetAutoReload(ctx context.Context, input BillingAutoReloadUpdateInput, options *CallOptions) (BillingAutoReloadUpdateResponse, error) {
	var result BillingAutoReloadUpdateResponse
	data, err := c.call(ctx, "billing.auto_reload.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// BuyCredits calls POST /v1/billing/credit-purchases. Buy credits with a saved payment method.
func (c *Client) BuyCredits(ctx context.Context, input BillingCreditsBuyInput, options *CallOptions) (BillingCreditsBuyResponse, error) {
	var result BillingCreditsBuyResponse
	data, err := c.call(ctx, "billing.credits.buy", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RetryCreditPurchase calls POST /v1/billing/credit-purchases/{id}/retry. Try a declined credit purchase again, with the default or another saved payment method.
func (c *Client) RetryCreditPurchase(ctx context.Context, input BillingCreditsRetryInput, options *CallOptions) (BillingCreditsRetryResponse, error) {
	var result BillingCreditsRetryResponse
	data, err := c.call(ctx, "billing.credits.retry", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ConfirmPlanChange calls POST /v1/billing/quotes/{id}/confirm. Confirm a quote to move to a higher plan on the same or a longer billing period, Builder to Pro or Scale included, or yearly billing on the same plan, with a saved payment method.
func (c *Client) ConfirmPlanChange(ctx context.Context, input BillingPlanChangeInput, options *CallOptions) (BillingPlanChangeResponse, error) {
	var result BillingPlanChangeResponse
	data, err := c.call(ctx, "billing.plan.change", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UndoScheduledDowngrade calls POST /v1/billing/scheduled-change/undo. Undo a scheduled downgrade, so the subscription renews on the current plan.
func (c *Client) UndoScheduledDowngrade(ctx context.Context, input map[string]any, options *CallOptions) (BillingDowngradeUndoResponse, error) {
	var result BillingDowngradeUndoResponse
	data, err := c.call(ctx, "billing.downgrade.undo", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ResumeSubscription calls POST /v1/billing/subscription/resume. Resume a canceled subscription before it ends, so it renews again.
func (c *Client) ResumeSubscription(ctx context.Context, input map[string]any, options *CallOptions) (BillingSubscriptionResumeResponse, error) {
	var result BillingSubscriptionResumeResponse
	data, err := c.call(ctx, "billing.subscription.resume", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// PayOverdueInvoice calls POST /v1/billing/invoices/pay. Pay an overdue invoice with a saved payment method.
func (c *Client) PayOverdueInvoice(ctx context.Context, input BillingOverduePayInput, options *CallOptions) (BillingOverduePayResponse, error) {
	var result BillingOverduePayResponse
	data, err := c.call(ctx, "billing.overdue.pay", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SetDefaultPaymentMethod calls POST /v1/billing/payment-methods/{id}/default. Make a saved payment method the default. It then pays anything overdue.
func (c *Client) SetDefaultPaymentMethod(ctx context.Context, input map[string]any, options *CallOptions) (BillingPaymentMethodsDefaultResponse, error) {
	var result BillingPaymentMethodsDefaultResponse
	data, err := c.call(ctx, "billing.payment_methods.default", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// MoveBackupPaymentMethod calls POST /v1/billing/payment-methods/{id}/move. Move a backup payment method up or down one place in the order.
func (c *Client) MoveBackupPaymentMethod(ctx context.Context, input BillingPaymentMethodsMoveInput, options *CallOptions) (BillingPaymentMethodsMoveResponse, error) {
	var result BillingPaymentMethodsMoveResponse
	data, err := c.call(ctx, "billing.payment_methods.move", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RemovePaymentMethod calls DELETE /v1/billing/payment-methods/{id}. Remove a saved payment method. The default can't be removed.
func (c *Client) RemovePaymentMethod(ctx context.Context, input map[string]any, options *CallOptions) (BillingPaymentMethodsRemoveResponse, error) {
	var result BillingPaymentMethodsRemoveResponse
	data, err := c.call(ctx, "billing.payment_methods.remove", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RedeemEnterprisePlanCode calls POST /v1/billing/enterprise-plan-codes/redeem. Redeem an Enterprise plan code. It adds the plan to the workspace for good.
func (c *Client) RedeemEnterprisePlanCode(ctx context.Context, input BillingCodesRedeemInput, options *CallOptions) (BillingCodesRedeemResponse, error) {
	var result BillingCodesRedeemResponse
	data, err := c.call(ctx, "billing.codes.redeem", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// CreateBillingApprovalLink calls POST /v1/billing/approval-links. Make a link for the owner to add a payment method, or to add one and pay for an upgrade, when the workspace has none saved.
func (c *Client) CreateBillingApprovalLink(ctx context.Context, input BillingApprovalLinksCreateInput, options *CallOptions) (BillingApprovalLinksCreateResponse, error) {
	var result BillingApprovalLinksCreateResponse
	data, err := c.call(ctx, "billing.approval_links.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// BillingApprovalLink calls GET /v1/billing/approval-links/{id}. Read whether the owner finished a billing approval link, declined it or let it expire.
func (c *Client) BillingApprovalLink(ctx context.Context, input map[string]any, options *CallOptions) (BillingApprovalLinksGetResponse, error) {
	var result BillingApprovalLinksGetResponse
	data, err := c.call(ctx, "billing.approval_links.get", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateAccountName calls PATCH /v1/account. Change your name.
func (c *Client) UpdateAccountName(ctx context.Context, input AccountUpdateInput, options *CallOptions) (AccountUpdateResponse, error) {
	var result AccountUpdateResponse
	data, err := c.call(ctx, "account.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// EmailPreferences calls GET /v1/account/email-preferences. Read which optional emails you get.
func (c *Client) EmailPreferences(ctx context.Context, input map[string]any, options *CallOptions) (AccountEmailPreferencesResponse, error) {
	var result AccountEmailPreferencesResponse
	data, err := c.call(ctx, "account.email_preferences", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// UpdateEmailPreferences calls PATCH /v1/account/email-preferences. Choose which optional emails you get. Sign-in, security and billing emails always go.
func (c *Client) UpdateEmailPreferences(ctx context.Context, input AccountEmailPreferencesUpdateInput, options *CallOptions) (AccountEmailPreferencesUpdateResponse, error) {
	var result AccountEmailPreferencesUpdateResponse
	data, err := c.call(ctx, "account.email_preferences_update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Notifications calls GET /v1/notifications. List your notifications, and how many are unread.
func (c *Client) Notifications(ctx context.Context, input NotificationsListInput, options *CallOptions) (NotificationsListResponse, error) {
	var result NotificationsListResponse
	data, err := c.call(ctx, "notifications.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// MarkNotification calls PATCH /v1/notifications/{id}. Mark a notification as read or unread.
func (c *Client) MarkNotification(ctx context.Context, input NotificationsUpdateInput, options *CallOptions) (NotificationsUpdateResponse, error) {
	var result NotificationsUpdateResponse
	data, err := c.call(ctx, "notifications.update", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// MarkAllNotificationsRead calls POST /v1/notifications/read-all. Mark all your notifications as read.
func (c *Client) MarkAllNotificationsRead(ctx context.Context, input map[string]any, options *CallOptions) (NotificationsReadAllResponse, error) {
	var result NotificationsReadAllResponse
	data, err := c.call(ctx, "notifications.read_all", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// RunNotificationAction calls POST /v1/notifications/{id}/run. Run one of a notification's actions. It also needs the scope of the action it runs.
func (c *Client) RunNotificationAction(ctx context.Context, input NotificationsRunActionInput, options *CallOptions) (NotificationsRunActionResponse, error) {
	var result NotificationsRunActionResponse
	data, err := c.call(ctx, "notifications.run_action", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ContactSupport calls POST /v1/support-tickets. Contact AbuseZero support about this workspace. Support replies to your account email.
func (c *Client) ContactSupport(ctx context.Context, input SupportTicketsCreateInput, options *CallOptions) (SupportTicketsCreateResponse, error) {
	var result SupportTicketsCreateResponse
	data, err := c.call(ctx, "support-tickets.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SupportTickets calls GET /v1/support-tickets. List the support tickets you sent about this workspace.
func (c *Client) SupportTickets(ctx context.Context, input SupportTicketsListInput, options *CallOptions) (SupportTicketsListResponse, error) {
	var result SupportTicketsListResponse
	data, err := c.call(ctx, "support-tickets.list", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// SendFeedback calls POST /v1/feedback. Send product feedback to the AbuseZero team.
func (c *Client) SendFeedback(ctx context.Context, input FeedbackCreateInput, options *CallOptions) (FeedbackCreateResponse, error) {
	var result FeedbackCreateResponse
	data, err := c.call(ctx, "feedback.create", input, options)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}
