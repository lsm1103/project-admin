package dataModel_

import (
	"golang.org/x/xerrors"
)

var (
	//错误
	ErrNotFound = xerrors.New("找不到数据")

	// 状态
	Del     int64 = -2 //删除
	Disable int64 = -1 //禁用
	Audited int64 = 1  //待审核
	Enable  int64 = 2  //启用
)
