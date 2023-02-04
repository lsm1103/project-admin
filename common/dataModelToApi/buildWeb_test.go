package dataModelToApi

import "testing"

func TestBuildWebCode(t *testing.T) {
	BuildWeb{
		RootPkgPath: "",
		TableList:   nil,
		ServiceCfg:  ServiceInfo{},
	}.BuildWebCode()
	
	t.Logf("r")
}