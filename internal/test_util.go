package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func AssertErrIsNil(t *testing.T, err error) {
	// Inserted model's id should not be nil:
	require.Nil(t, err, "Assertion err: %v", err)
}
