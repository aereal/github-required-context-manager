package main

import (
	"flag"
	"fmt"
)

const (
	defaultBaseURL = "https://api.github.com"
)

type CLIOpts struct {
	owner              string
	repo               string
	branch             string
	insecureSkipVerify bool
	baseURL            string
	list               bool
	add                bool
	contextName        string
}

func dispatchCmd(opts *CLIOpts) (Cmd, error) {
	if opts.list && opts.add {
	} else if opts.list {
		cmd, err := newListCmd(opts)
		if err != nil {
			return nil, err
		}
		return cmd, nil
	} else if opts.add {
		cmd, err := newAddCmd(opts)
		if err != nil {
			return nil, err
		}
		return cmd, nil
	}
	return nil, fmt.Errorf("Either -list or -add given")
}

func parseArgs(argv []string) (*CLIOpts, error) {
	args := &CLIOpts{}
	flgs := flag.NewFlagSet(argv[0], flag.ContinueOnError)
	flgs.StringVar(&args.owner, "owner", "", "owner of repo")
	flgs.StringVar(&args.repo, "repo", "", "repo name")
	flgs.StringVar(&args.branch, "branch", "", "branch name")
	flgs.BoolVar(&args.insecureSkipVerify, "insecure-skip-verify", false, "skip verification of cert")
	flgs.StringVar(&args.baseURL, "base-url", defaultBaseURL, "custom GitHub base URL if you use GitHub Enterprise")
	flgs.BoolVar(&args.list, "list", false, "list required contexts")
	flgs.BoolVar(&args.add, "add", false, "add required context")
	flgs.StringVar(&args.contextName, "context", "", "context name")

	if err := flgs.Parse(argv[1:]); err != nil {
		return nil, err
	}

	if args.owner == "" {
		return nil, fmt.Errorf("owner cannnot be empty")
	}

	if args.repo == "" {
		return nil, fmt.Errorf("repo cannot be empty")
	}

	if args.branch == "" {
		return nil, fmt.Errorf("branch cannot be empty")
	}

	return args, nil
}
