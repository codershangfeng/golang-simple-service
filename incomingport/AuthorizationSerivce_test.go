package incomingport

import "testing"

func TestResultCodeString(t *testing.T) {
	cases := []struct {
		enumName   string
		resultCode ResultCode
		expected   string
	}{
		{"Approved", Approved, "APPROVED"},
		{"Error", Error, "ERROR"},
	}
	for _, tc := range cases {
		t.Run(tc.enumName, func(t *testing.T) {
			if got := tc.resultCode.String(); got != tc.expected {
				t.Errorf("%v.String() = %q, but expected = %v", tc.enumName, got, tc.expected)
			}
		})
	}
}
