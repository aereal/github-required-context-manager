package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type testCaseDispatchCmd struct {
	opts          *CLIOpts
	expectedError error
}

type testCaseParseArgs struct {
	argv []string
	err  error
	opts *CLIOpts
}

func TestParseArgs(t *testing.T) {
	testCases := []*testCaseParseArgs{
		&testCaseParseArgs{
			argv: strings.Split("github-required-context-manager -unknown", " "),
			opts: nil,
			err:  fmt.Errorf("flag provided but not defined: -unknown"),
		},
		&testCaseParseArgs{
			argv: strings.Split("github-required-context-manager", " "),
			opts: nil,
			err:  fmt.Errorf("owner cannnot be empty"),
		},
		&testCaseParseArgs{
			argv: strings.Split("github-required-context-manager -owner aereal", " "),
			opts: nil,
			err:  fmt.Errorf("repo cannot be empty"),
		},
		&testCaseParseArgs{
			argv: strings.Split("github-required-context-manager -owner aereal -repo github-required-context-manager", " "),
			opts: nil,
			err:  fmt.Errorf("branch cannot be empty"),
		},
		&testCaseParseArgs{
			argv: strings.Split("github-required-context-manager -owner aereal -repo github-required-context-manager -branch master", " "),
			opts: &CLIOpts{
				owner:              "aereal",
				repo:               "github-required-context-manager",
				branch:             "master",
				insecureSkipVerify: false,
				baseURL:            "https://api.github.com",
				list:               false,
				add:                false,
				delete:             false,
				contextName:        "",
			},
			err: nil,
		},
	}
	for _, testCase := range testCases {
		opts, err := parseArgs(testCase.argv, new(bytes.Buffer))
		if testCase.err == nil {
			if err != nil {
				t.Errorf("Expect parseArgs(%#v) to return no errors but got: %s", testCase.argv, err)
			} else if err = eqOpts(opts, testCase.opts); err != nil {
				t.Errorf("Expect parseArgs(%#v) to return: %#v\nbut got: %#v",
					testCase.argv,
					opts,
					testCase.opts,
				)
			}
		} else if testCase.err != nil {
			if err == nil {
				t.Errorf("Expect parseArgs(%#v) to return some errors but got nothing", testCase.argv)
			} else if testCase.err.Error() != err.Error() {
				t.Errorf("Expect parseArgs(%#v) to return error: %s\nbut got: %s", testCase.argv, testCase.err, err)
			}
		}
	}
}

func TestDispatchCmd(t *testing.T) {
	testCases := []*testCaseDispatchCmd{
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list:   false,
				add:    false,
				delete: false,
			},
			expectedError: fmt.Errorf("Either -list, -add or -delete given"),
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list:   true,
				add:    false,
				delete: false,
			},
			expectedError: nil,
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list:        false,
				add:         true,
				delete:      false,
				contextName: "example-check",
			},
			expectedError: nil,
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list:        false,
				add:         true,
				delete:      false,
				contextName: "",
			},
			expectedError: fmt.Errorf("context required"),
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list:        false,
				add:         false,
				delete:      true,
				contextName: "",
			},
			expectedError: fmt.Errorf("context required"),
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list: true,
				add:  true,
			},
			expectedError: fmt.Errorf("Either -list, -add or -delete given"),
		},
	}
	for _, testCase := range testCases {
		cmd, err := dispatchCmd(testCase.opts)
		if testCase.expectedError == nil && err == nil {
			if cmd == nil {
				t.Errorf("dispatchCmd(%#v) expected return cmd but got nil", testCase.opts)
			}
		} else if testCase.expectedError == nil && err != nil {
			t.Errorf("dispatchCmd(%#v) expected no errors but got %s", testCase.opts, err)
		} else if testCase.expectedError != nil && err == nil {
			t.Errorf("dispatchCmd(%#v) expected some error (%s) but got no errors", testCase.opts, testCase.expectedError)
		} else if testCase.expectedError.Error() != err.Error() {
			t.Errorf("dispatchCmd(%#v) expected error (%s) but got error %s", testCase.opts, testCase.expectedError, err)
		}
	}
}

func eqOpts(got *CLIOpts, expected *CLIOpts) error {
	if expected == nil && got != nil {
		return fmt.Errorf("Expected nothing but got: %#v", got)
	}
	// expected something

	if got == nil {
		return fmt.Errorf("Expected something but got nothing")
	}

	if expected.owner != got.owner {
		return fmt.Errorf("Expected %s to be %v but got %v", "owner", expected.owner, got.owner)
	}
	if expected.repo != got.repo {
		return fmt.Errorf("Expected %s to be %v but got %v", "repo", expected.repo, got.repo)
	}
	if expected.branch != got.branch {
		return fmt.Errorf("Expected %s to be %v but got %v", "branch", expected.branch, got.branch)
	}
	if expected.insecureSkipVerify != got.insecureSkipVerify {
		return fmt.Errorf("Expected %s to be %v but got %v", "insecureSkipVerify", expected.insecureSkipVerify, got.insecureSkipVerify)
	}
	if expected.baseURL != got.baseURL {
		return fmt.Errorf("Expected %s to be %v but got %v", "baseURL", expected.baseURL, got.baseURL)
	}
	if expected.list != got.list {
		return fmt.Errorf("Expected %s to be %v but got %v", "list", expected.list, got.list)
	}
	if expected.add != got.add {
		return fmt.Errorf("Expected %s to be %v but got %v", "add", expected.add, got.add)
	}
	if expected.delete != got.delete {
		return fmt.Errorf("Expected %s to be %v but got %v", "delete", expected.delete, got.delete)
	}
	if expected.contextName != got.contextName {
		return fmt.Errorf("Expected %s to be %v but got %v", "contextName", expected.contextName, got.contextName)
	}

	return nil
}
