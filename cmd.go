package main

import (
	"context"
	"fmt"
)

type Cmd interface {
	Do(ctx context.Context) error
}

type ListCmd struct {
	opts *CLIOpts
}

func newListCmd(opts *CLIOpts) (*ListCmd, error) {
	cmd := &ListCmd{opts: opts}
	return cmd, nil
}

func (c *ListCmd) Do(ctx context.Context) error {
	client, err := newGithubClient(ctx, c.opts.baseURL, c.opts.insecureSkipVerify)
	if err != nil {
		return err
	}

	contexts, _, err := client.Repositories.ListRequiredStatusChecksContexts(
		ctx,
		c.opts.owner,
		c.opts.repo,
		c.opts.branch,
	)
	if err != nil {
		return err
	}
	for _, checkCtx := range contexts {
		fmt.Println(checkCtx)
	}
	return nil
}
