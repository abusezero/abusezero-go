# Go HTTP handlers

Create the client once when your server starts, and call `Close` at shutdown. Pass the handler's context to each method, so a disconnect or cancellation stops the call.

```go
func EvaluateClaim(r *http.Request, client *abusezero.Client, claimID, verifiedEmail, deviceToken, clientIP string) (any, error) {
	incoming, err := abusezero.IncomingRequest(r.Method, "https://app.example.com"+r.URL.RequestURI(), r.Header)
	if err != nil {
		return nil, err
	}
	input := abusezero.EvaluationsCreateInput{
		Benefit: "pro_trial", ClaimID: claimID,
		Subject:     abusezero.EvaluationsCreateInputSubject{Email: verifiedEmail},
		Context:     abusezero.Some[any](map[string]any{"ip": clientIP}),
		Identifiers: abusezero.Some(abusezero.EvaluationsCreateInputIdentifiers{DeviceToken: abusezero.Some(deviceToken)}),
		Request:     abusezero.Some(incoming),
	}
	return client.Evaluate(r.Context(), input, &abusezero.CallOptions{IdempotencyKey: "evaluate:" + claimID})
}
```

Build the absolute request URL from your configured application origin. Pass your own stored claim ID, the authenticated user's email, the device token your page sent and the client IP your trusted proxy configuration reports. Check the result as the README shows before you grant the benefit.

For webhooks, read the bounded original request body with `io.ReadAll(http.MaxBytesReader(w, r.Body, maximumBodyBytes))` and check the read error. Then call `VerifyWebhook` with those bytes and `r.Header`, and act on the returned event only after verification succeeds.
