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
