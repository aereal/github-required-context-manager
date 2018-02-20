package main

import (
	"context"
	"fmt"
)

type Cmd interface {
	Do(ctx context.Context) error
}

type AddCmd struct {
	opts *CLIOpts
}

func newAddCmd(opts *CLIOpts) (*AddCmd, error) {
	if opts.contextName == "" {
		return nil, fmt.Errorf("context required")
	}
	cmd := &AddCmd{
		opts: opts,
	}
	return cmd, nil
}

func (c *AddCmd) Do(ctx context.Context) error {
	client, err := newGithubClient(ctx, c.opts.baseURL, c.opts.insecureSkipVerify)
	if err != nil {
		return err
	}

	u := fmt.Sprintf(
		"repos/%v/%v/branches/%v/protection/required_status_checks/contexts",
		c.opts.owner,
		c.opts.repo,
		c.opts.branch,
	)
	ctxs := []string{c.opts.contextName}
	req, err := client.NewRequest("POST", u, ctxs)
	if err != nil {
		return err
	}
	if _, err = client.Do(ctx, req, nil); err != nil {
		return err
	}
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

type DeleteCmd struct {
	opts *CLIOpts
}

func newDeleteCmd(opts *CLIOpts) (*DeleteCmd, error) {
	if opts.contextName == "" {
		return nil, fmt.Errorf("context required")
	}
	cmd := &DeleteCmd{
		opts: opts,
	}
	return cmd, nil
}

func (c *DeleteCmd) Do(ctx context.Context) error {
	client, err := newGithubClient(ctx, c.opts.baseURL, c.opts.insecureSkipVerify)
	if err != nil {
		return err
	}

	u := fmt.Sprintf(
		"repos/%v/%v/branches/%v/protection/required_status_checks/contexts",
		c.opts.owner,
		c.opts.repo,
		c.opts.branch,
	)
	ctxs := []string{c.opts.contextName}
	req, err := client.NewRequest("DELETE", u, ctxs)
	if err != nil {
		return err
	}
	if _, err = client.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
