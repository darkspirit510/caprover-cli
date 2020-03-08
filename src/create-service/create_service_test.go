package create_service

import (
	"strings"
	"testing"
)

func AssertStartsWith(t *testing.T, actual string, prefix string) {
	t.Helper()

	if !strings.HasPrefix(actual, prefix) {
		t.Error()
	}
}

func TestDoAction(t *testing.T) {
	t.Run("should return usage when no subcommand is specified", func(t *testing.T) {
		AssertStartsWith(t, DoAction([]string{}), "Usage: ")
	})
}
