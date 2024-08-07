package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountIDFromAdminRoleID(t *testing.T) {
	t.Parallel()
	tt := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name:     "one hyphen",
			Input:    "example-ab1cd",
			Expected: "ab1cd",
		},
		{
			Name:     "multiple hyphens",
			Input:    "example--inc-ab1cd",
			Expected: "ab1cd",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			actual := AccountIDFromAdminRoleID(tc.Input)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}
