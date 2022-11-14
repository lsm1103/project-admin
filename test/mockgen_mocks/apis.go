package mockgen_mocks

import "project-admin/test2/internal/types"

type Test2Mock interface {
	CreateSingleMsg(req types.CreateSingleMsg) error
	DeleteSingleMsg(req types.DeleteSingleMsg) error
	UpdateSingleMsg(req types.UpdateSingleMsg) error
	GetSingleMsg(req types.GetSingleMsg) (resp *types.SingleMsg, err error)
	GetsSingleMsg(req types.GetsReq) (resp *types.SingleMsgList, err error)
}
