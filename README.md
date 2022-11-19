
## 业务功能
### 组管理
- 组的增删改查
### 项目管理
- 项目的增删改查？
### application 管理
- application的增删改查
- application的数据库表设计文件 application.sql
- application的接口设计文件 application.api
- 自动生成文档，且可以修改，导出，全文搜索
- 记录测试用例，且可以修改，导出
- 增加新需求时，需要更新api文件
- application开发进度监控功能, 在github/gitlab通过定时获取commit或者webhook更新commit
- application持续集成，建议通过webhook推送后，拉代码，然后测试、打包；
- application部署；通过点击按钮一键部署，通过配置ci后自动部署；


## 核心功能
- 项目设计评审、数据库设计之后，先实现接口设计；
- 由数据库表自动生成基础curd函数；
- 由数据库设计，自动生成相关表的基础后台管理系统；
>w
1、通过数据库表sql，自动生成增删改查的curl代码； @已实现
2、通过数据库表sql，自动生成xxx.api文件；
3、然后由该api文件自动生成一一对接由1生成的curl代码的接口服务； @已实现
4、部署运行（go run/docker）

- 由接口设计，自动生成接口mock服务，这时前端可以借助该服务进行调试开发； @已实现, mock字段的丰富性还需增加
>w
1、改造goctl，在生成接口基础服务的基础上，默认为每个接口接入通过配置的出参类型和特定的mock（如电话、邮箱、人名）的mock；
2、部署运行（go run/docker）

- 生成的mock服务也是该项目的网关，集成认证、鉴权；后端只需要编写相关的业务逻辑无需管理网关；
>i
mock服务的生成，在logic.tpl引入一个mock类方法，该方法可以通过传入的出参（特殊字段要求可以写在字段tag里面）用反射，
获取每个字段的类型和特殊要求进行随机生成相应的值； @已实现

>i TODO 通过中间价集成认证、鉴权到网关；
1、通过 etc/config.yml动态修改认证、鉴权服务的地址、appKey、secretKey;
2、代码集成该认证鉴权sdk

>w
1、开发可以下载上面生成的mock服务代码，删除mock代码，再加入自己的代码即可；

- 实现生成python代码，而不仅仅是go

- 数据库表自动生成方案要点
  - 再读取数据库表时，需要对数据库表进行检查，涉及到用户相关等敏感信息，需要进行提醒、处理；
  - 是否软删除，如果选择软删除则必须要添加 state 字段

## 持续集成
- 实现一个镜像管理服务，由ci执行自动测试自动打包，然后出现在管理平台的版本列表里面，自动通知测试人员进行测试然后发行正式版本；

## 问题
### 数据库相关
- 需要对查询的接口参数，在拼接sql的时候进行处理，把直接写的符号换成eq这些，防止xss攻击；
- 对一些随便写的参数进行过滤
  gt：表示大于>。即greater than
  ge：表示大于等于>=。即greater than or equals to
  lt：表示小于<。即less than
  le：表示小于等于<=。即less than or equals to
  eq：表示等于=。即equals
  ne：表示不等于!=。即not equals

## 新增功能
- 支持关系型、非关系型、大数据数据库
- 支持集群部署，拥有优异的访问性能
- API 支持设置筛选条件、分页等选项
- 支持通过界面创建 API
- 自带请求校验、防SQL注入、HTTPS
- 支持不同集群配置不同的连接信息
- 生成规范的 API 文档以及调用代码
- 支持关系型、非关系型、大数据数据库
  支持 Mysql、Oracle、MariaDB、 SQL Server、PostgreSQL、MongoDB、Redis、Impala 等常见关系型、非关系型、大数据数据库中间件。对于新的的数据库类型可以动态扩充。

## 命令
- goctl model mysql ddl -src=test1.sql  -dir="dataModel/." -c --home libs/template && cp -r common/sqlUtils/* dataModel/

- goctl api go -style goZero --home libs/template -dir test1/. -api test1.api

- mockgen库运行示例
mockgen -destination="mocks/mock_test2.go" -package="mocks" project-admin/test2/mocks Test2Mock

- mockgen代码运行示例
go run . -source /Users/xm/Desktop/go_package/project-admin/common/mocks/tag_v2.go -destination /Users/xm/Desktop/go_package/project-admin/t2.go

- goctl-swagger
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api xxx.api -dir .

- go run goctl.go api go
go run goctl.go api go -style goZero --home ../template -dir ../../projectBuilds/t2 -api ../../adminManage/service1.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../projectBuilds/t2 -api ../../adminManage/service1.api
  
go run goctl.go api go -style goZero --home ../template -dir ../../projectBuilds/t3 -api ../../adminManage/service3.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../projectBuilds/t3 -api ../../adminManage/service3.api

- go run goctl.go mysql ddl  TODO还需要整理一番
go run goctl.go model mysql ddl -src=/Users/xm/Desktop/go_package/project-admin/deploy/init.sql  -dir="../../dataModel_/." -c --home libs/template

[comment]: <> (go run goctl.go api go -style goZero --home ../template -dir ../../projectBuilds/t3 -api ../../adminManage/service3.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../projectBuilds/t3 -api ../../adminManage/service3.api)

- 打包成linux
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o three-tuoming three-tuoming.go

## 生成项目
- 新建app, 创建app文件夹 projectBuilds/项目名
### 生成项目mock服务完整流程
- 设计接口，编写 .api文件
- 生成mock服务，生成代码命令：
  go run goctl.go api go -style goZero --home ../template -dir ../../projectBuilds/项目名 -api ../../projectBuilds/项目名/service.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../projectBuilds/项目名 -api ../../projectBuilds/项目名/service.api
- 更新到数据库
- 打包&运行

### 生成项目后台服务完整流程
- 设计数据库表，编写 .sql文件
- 生成数据库操作模块，生成代码命令：
  go run goctl.go model mysql ddl -src=../../deploy/init.sql  -dir="../../dataModel/项目名Model/." -c --home ../template
- 由 .sql文件生成 .api文件，生成代码：调用adminManage.StartBuild()方法
- 由 .api文件生成后台管理服务代码，生成代码命令：
  go run goctl.go api go -style goZero --home ../template -dir ../../projectBuilds/项目名 -api ../../projectBuilds/项目名/service.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../projectBuilds/项目名 -api ../../projectBuilds/项目名/service.api
- 更新到数据库
- 打包&运行


#### 测试 project1
- go run goctl.go model mysql ddl -src=../../deploy/init.sql  -dir="../../dataModel/project1/." -c --home ../template
- 调用 adminManage.StartBuild(ServiceInfo{
      Title:   "项目1",
      Desc:    "测试项目1；通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
      Author:  "lsm",
      Email:   "18370872400@163.com",
      Version: "v0.1.1",
      ProjectName: "project1",
  	  ServiceType: "admin",
      Host:        "0.0.0.0",
      Port:        "811",
      DataSource: "root:pujian123@tcp(172.16.10.183:4306)/project-admin",
      CacheHost: "172.16.10.183:6379",
      //DataSource: "root:lsm.1018@tcp(127.0.0.1:3306)/project-admin",
      //CacheHost:  "127.0.0.1:6379",
  }, SqlParseCfg{
      filename: "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
      database: "",
      strict:   false,
  })
- go run goctl.go api go -style goZero --home ../template -dir ../../projectBuilds/project1 -api ../../projectBuilds/project1/service.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../projectBuilds/project1 -api ../../projectBuilds/project1/service.api

#### 测试 project2
- go run goctl.go model mysql ddl -src=../../deploy/init.sql  -dir="../../dataModel/project2/." -c --home ../template
- 调用 adminManage.StartBuild(ServiceInfo{
  Title:   "项目2",
  Desc:    "测试项目2；通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
  Author:  "lsm",
  Email:   "18370872400@163.com",
  Version: "v0.1.1",
  ProjectName: "project2",
  ServiceType: "mock",
  Host:        "0.0.0.0",
  Port:        "812",
  DataSource: "root:pujian123@tcp(172.16.10.183:4306)/project-admin",
  CacheHost: "172.16.10.183:6379",
  //DataSource: "root:lsm.1018@tcp(127.0.0.1:3306)/project-admin",
  //CacheHost:  "127.0.0.1:6379",
  }, SqlParseCfg{
  filename: "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
  database: "",
  strict:   false,
})
- go run goctl.go api go -style goZero --home ../template -dir ../../projectBuilds/project2 -api ../../projectBuilds/project2/service.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../projectBuilds/project2 -api ../../projectBuilds/project2/service.api

#### 3
- go run goctl.go model mysql ddl -src=../../deploy/init.sql  -dir="../../dataModel/project3/." -c --home ../template
- go run goctl.go api go -style goZero --home ../template -dir ../../projectBuilds/project3 -api ../../projectBuilds/project3/service.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../projectBuilds/project3 -api ../../projectBuilds/project3/service.api


go run goctl.go api go -style goZero --home ../template -dir ../../project-admin -api ../../project-admin/service.api && goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -dir ../../project-admin -api ../../project-admin/service.api

- 新增样例
```json
{
  "create_user": 111212,
  "demand_ids": 234444,
  "doc_ids": "111，",
  "en_name": "project1",
  "ico": "ssfwfwf",
  "info": "但是更多的废话",
  "join_groups": "134214,",
  "join_users": "134124,",
  "project_id": "34534523",
  "rank": 1,
  "remark": "gsdgds",
  "zn_name": "5675756"
}
```

- TODO swagger.json处理
```json
{
  "info": {
    "contact": {
      "email": "18370872400@163.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "termsOfService": "http://swagger.io/terms/"
  },
  "host": "127.0.0.1:813",
  "externalDocs": {
    "description": "文档",
    "url": "http://doc.pj-ai.com/share_doc/?token=ce67940ae226fe53a640b7750e0bdf73"
  }
}
```

```go
ServiceInfo{
    Title:   "项目3",
    Desc:    "测试项目3；通过api设计文档自动生成服务，并根据api文件配置的字段mock规则进行mock生成结果",
    Author:  "lsm",
    Email:   "18370872400@163.com",
    Version: "v0.1.1",
    ProjectName: "project3",
    ServiceType: "mock",
    Host:        "0.0.0.0",
    Port:        "813",
    DataSource: "root:pujian123@tcp(172.16.10.183:4306)/project-admin",
    CacheHost: "172.16.10.183:6379",
    //DataSource: "root:lsm.1018@tcp(127.0.0.1:3306)/project-admin",
    //CacheHost:  "127.0.0.1:6379",
}, SqlParseCfg{
    filename: "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
    database: "",
    strict:   false,
}

{
    "title": "项目22",
    "author": "lsm",
    "desc": "对研发项目进行管理，包括代码生成、mock服务生成、cicd等；",
    "email": "18370872400@163.com",
    "host": "0.0.0.0",
    "port": "822",
    "version": "v0.1.1"
	
	"projectName": "project22",
    "service_type": "admin",
    "style": "goZero",
    "templatePath": "",
	"cacheHost": "172.16.10.183:6379",
	"dataSource": "root:pujian123@tcp(172.16.10.183:4306)/project-admin",
    "database": "",
    "ddlArg": {
        "cache": true,
        "src": "",
        "strict": false
    },
}


@doc (
summary:获取构建情况（进度、log）
handlerType:getBuild
serviceType:business
)
@handler GetBuild
post /getBuild(GetBuildReq) returns(Build)
```


```text
goctl error: 4:75: expected operand, found ',' (and 1 more errors)\ngoTL_OS=darwin\nGOCTL_ARCH=amd64\nGOCTL_HOME=/Users/xm/.goctl\nGOCTL_DEBUG=False\nGOCTL_CACHE=/Users/xm/.goctl/cache\nGOCTL_VERSION=1.4.2\nPROTOC_VERSION=3.17.3\nPROTOC_GEN_GO_VERSION=\nPROTO_GEN_GO_GRPC_VERSION=\nmessage: go format error:\n

var (
    tgroupFieldNames          = builder.RawFieldNames(\u0026Group{})
    tgroupRows                = strings.Join(groupFieldNames, \",\")
    groupRowsExpectAutoSet   = strings.Join(stringx.Remove(groupFieldNames,  , \"`state`\"), \",\")
    tgroupRowsWithPlaceHolder = strings.Join(stringx.Remove(groupFieldNames, \"`id`\", ), \"=?,\") + \"=?\"
    groupListRows   =  strings.Join( builder.RawFieldNames( \u0026Group{} ), \",\")
    cacheGroupIdPrefix = \"cache:group:id:\"
    cacheGroupCreateUserNamePrefix = \"cache:group:createUser:name:\"
)
```