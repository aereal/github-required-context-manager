package main

import (
	"context"
	"fmt"
)

type Cmd interface {
	Do(ctx context.Context) error
}

type AddCmd struct {
	opts        *CLIOpts
	contextName string
}

func newAddCmd(opts *CLIOpts) (*AddCmd, error) {
	ctxName := opts.contextName
	if ctxName == "" {
		return nil, fmt.Errorf("context required")
	}
	cmd := &AddCmd{
		contextName: ctxName,
		opts:        opts,
	}
	return cmd, nil
}

func (c *AddCmd) Do(ctx context.Context) error {
	return nil
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
