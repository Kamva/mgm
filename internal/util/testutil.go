package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// AssertErrIsNil function assert that the passed-in error is nil.
func AssertErrIsNil(t *testing.T, err error) {
	// The inserted model's id should not be nil:
	require.Nil(t, err, "Assertion err: %v", err)
}
