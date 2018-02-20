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
			},
			expectedError: fmt.Errorf("Either -list given"),
		},
		&testCaseDispatchCmd{
			opts: &CLIOpts{
				list: true,
			},
			expectedError: nil,
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
