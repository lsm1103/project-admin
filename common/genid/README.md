# 说明 @v1.0.1

本组件是基于github.com/srlemon/gen-id的核心功能进行修改，调整底层对数据随机的方案

# gen-id
一个身份证、名字、邮箱、地址、手机号码等随机生成的sdk

# Installation
`go get project-admin/common/genid`

如果网速过慢:
```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
go get github.com/srlemon/gen-id
```

# Example

```golang
package main

import (
	"fmt"
	"project-admin/common/genid"
	"project-admin/common/genid/generator"
)

func main()  {
	// 生成总的信息
	fmt.Println(gen_id.NewGeneratorData())
	// 分个单独获取
	g:=new(generator.GeneratorData)
	fmt.Println(g.GeneratorPhone())
	fmt.Println(g.GeneratorName())
	fmt.Println(g.GeneratorIDCart())
	fmt.Println(g.GeneratorEmail())
	fmt.Println(g.GeneratorBankID())
	fmt.Println(g.GeneratorAddress())
}

```

# Statement
本项目用于开发环境,涉及商业用途用本人无关
