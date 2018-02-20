package main

import (
	"fmt"
	"testing"
)

type testCaseDispatchCmd struct {
	opts          *CLIOpts
	expectedError error
}

func TestDispatchCmd(t *testing.T) {
	testCases := []*testCaseDispatchCmd{
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list: false,
				add:  false,
			},
			expectedError: fmt.Errorf("Either -list or -add given"),
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list: true,
				add:  false,
			},
			expectedError: nil,
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list:        false,
				add:         true,
				contextName: "example-check",
			},
			expectedError: nil,
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list:        false,
				add:         true,
				contextName: "",
			},
			expectedError: fmt.Errorf("context required"),
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list: true,
				add:  true,
			},
			expectedError: fmt.Errorf("Either -list or -add given"),
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
