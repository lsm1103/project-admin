package uniqueid

import (
	"testing"
)

func TestGenId(t *testing.T) {
	r := GenUid()
	t.Logf("r:%+v\n", r)
}