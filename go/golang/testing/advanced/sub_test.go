package sub_test

import (
	"testing"
)

func Sum(a, b int) int {
	return a + b
}

func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup")
	return func(t *testing.T) {
		t.Log("teardown")
	}
}

// testing.M provides environment for setup/teardown, but it is for all tests in file
// this example defines test suites as collection of subsets whose purpose is same
// https://stackoverflow.com/questions/42310088/setup-and-teardown-for-each-test-using-std-testing-package/42310257#42310257
func TestAdvancedSub(t *testing.T) {
	cases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"basic", 1, 2, 3},
		{"minus", 1, -2, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)

			result := Sum(tc.a, tc.b)
			if result != tc.expected {
				t.Fatalf("expected sum %v, but got %v", tc.expected, result)
			}
		})
	}
}
