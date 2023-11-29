package backend

import (
	"testing"
)

func TestCliCent_int(t *testing.T) {
	i := -123
	j := uint32(i)
	t.Log(i)
	t.Log(j)
}
