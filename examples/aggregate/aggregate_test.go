package aggregate

import (
	"testing"
)

func TestLookup(t *testing.T) {
	if err:=lookup();err!=nil{
		panic(err)
	}
	//require.Nil(t, lookup())
}
