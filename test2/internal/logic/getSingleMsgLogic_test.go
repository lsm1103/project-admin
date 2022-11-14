package logic

import (
	"context"
	"testing"
)

func Test_RespMock(t *testing.T) {
	logic := NewGetSingleMsgLogic(context.Background(), nil)
	resp, err := logic.GetSingleMsg(nil)
	t.Logf("resp: %+v, %+v\n", resp, err)
	return
}
