package crud

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLookup(t *testing.T) {
	if err:=lookup();err!=nil{
		panic(err)
	}
	require.Nil(t, lookup())
}
