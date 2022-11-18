package logic

import (
	"context"
	"encoding/json"
	"testing"
)

func Test_GetsSingleMsg_RespMock(t *testing.T) {
	logic := NewGetsSingleMsgLogic(context.Background(), nil)
	resp, err := logic.GetsSingleMsg(nil)
	marshal, err := json.Marshal(resp)
	if err != nil {
		t.Logf("err:%s", err.Error())
	}
	t.Logf("resp: %+v, %s\n", resp, marshal)
	return
}
