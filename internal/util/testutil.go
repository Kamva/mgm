package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// AssertErrIsNil function assert that passed error be nil.
func AssertErrIsNil(t *testing.T, err error) {
	// Inserted model's id should not be nil:
	require.Nil(t, err, "Assertion err: %v", err)
}
