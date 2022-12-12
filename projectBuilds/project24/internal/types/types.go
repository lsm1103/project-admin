// Code generated by goctl. DO NOT EDIT.
package types

type DeleteReq struct {
	Id int64 `json:"id"` // 主键
}

type GetReq struct {
	Id int64 `path:"id"` // 主键
}

type GetsQueryItem struct {
	Key        string `json:"key"`                                   // key
	Val        string `json:"val"`                                   // val int/str('xxx')/list('xx,xx,xx')
	Handle     string `json:"handle"`                                // 操作方法 =/>/</like
	NextHandle string `json:"nextHandle,options=or|and,default=and"` // 与下一个条件的操作方法 or/and
}

type GetsReq struct {
	Query    []*GetsQueryItem `json:"query"`                              //查询
	OrderBy  string           `json:"orderBy,default=id"`                 //分组
	Sort     string           `json:"sort,options=desc|asc,default=desc"` //排序 desc/asc
	Current  int64            `json:"current,default=1"`                  //当前页
	PageSize int64            `json:"pageSize,default=10"`                //页面大小
}

type CreateGroupReq struct {
	Name       string `json:"name"`        // 组名称
	CreateUser int64  `json:"create_user"` // 创建者id
	Ico        string `json:"ico"`         // 组图标
	Remark     string `json:"remark"`      // 备注
	ParentId   int64  `json:"parent_id"`   // 父级id
	GroupType  int64  `json:"group_type"`  // 类型: 1部门、2用户组、3群组、4圈子、5话题
	Rank       int64  `json:"rank"`        // 排序
}

type UpdateGroupReq struct {
	Id         int64  `json:"id"`                   // 自增主键
	Name       string `json:"name,optional"`        // 组名称
	CreateUser int64  `json:"create_user,optional"` // 创建者id
	Ico        string `json:"ico,optional"`         // 组图标
	Remark     string `json:"remark,optional"`      // 备注
	ParentId   int64  `json:"parent_id,optional"`   // 父级id
	GroupType  int64  `json:"group_type,optional"`  // 类型: 1部门、2用户组、3群组、4圈子、5话题
	Rank       int64  `json:"rank,optional"`        // 排序
	State      int64  `json:"state,optional"`       // 状态，-2删除，-1禁用，待审核0，启用1
}

type Group struct {
	Id         int64  `json:"id"`          // 自增主键
	Name       string `json:"name"`        // 组名称
	CreateUser int64  `json:"create_user"` // 创建者id
	Ico        string `json:"ico"`         // 组图标
	Remark     string `json:"remark"`      // 备注
	ParentId   int64  `json:"parent_id"`   // 父级id
	GroupType  int64  `json:"group_type"`  // 类型: 1部门、2用户组、3群组、4圈子、5话题
	Rank       int64  `json:"rank"`        // 排序
	State      int64  `json:"state"`       // 状态，-2删除，-1禁用，待审核0，启用1
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

type GroupList struct {
	List     []*Group `json:"list"tag:"list"`                     // 数据列表
	Current  int64    `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64    `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool     `json:"isNext"tag:"bool"`                   // 是否有下一页
}

type CreateUserGroupReq struct {
	UserId  int64 `json:"user_id"`  // 用户id
	GroupId int64 `json:"group_id"` // 组id
}

type UpdateUserGroupReq struct {
	Id      int64 `json:"id"`                // 自增主键
	UserId  int64 `json:"user_id,optional"`  // 用户id
	GroupId int64 `json:"group_id,optional"` // 组id
	State   int64 `json:"state,optional"`    // 状态，-2删除，-1禁用，待审核0，启用1
}

type UserGroup struct {
	Id         int64  `json:"id"`          // 自增主键
	UserId     int64  `json:"user_id"`     // 用户id
	GroupId    int64  `json:"group_id"`    // 组id
	State      int64  `json:"state"`       // 状态，-2删除，-1禁用，待审核0，启用1
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

type UserGroupList struct {
	List     []*UserGroup `json:"list"tag:"list"`                     // 数据列表
	Current  int64        `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64        `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool         `json:"isNext"tag:"bool"`                   // 是否有下一页
}

type CreateGroupGroupRelationReq struct {
	CreateUser    int64 `json:"create_user"`     // 创建者id
	MasterGroupId int64 `json:"master_group_id"` // 主组id
	FromGroupId   int64 `json:"from_group_id"`   // 从组id
	Relation      int64 `json:"relation"`        // 关系，-1禁止(非自己)，1可读(非自己)，2可读写(非自己)
}

type UpdateGroupGroupRelationReq struct {
	Id            int64 `json:"id"`                       // 自增主键
	CreateUser    int64 `json:"create_user,optional"`     // 创建者id
	MasterGroupId int64 `json:"master_group_id,optional"` // 主组id
	FromGroupId   int64 `json:"from_group_id,optional"`   // 从组id
	Relation      int64 `json:"relation,optional"`        // 关系，-1禁止(非自己)，1可读(非自己)，2可读写(非自己)
	State         int64 `json:"state,optional"`           // 状态，-2删除，-1禁用，待审核0，启用1
}

type GroupGroupRelation struct {
	Id            int64  `json:"id"`              // 自增主键
	CreateUser    int64  `json:"create_user"`     // 创建者id
	MasterGroupId int64  `json:"master_group_id"` // 主组id
	FromGroupId   int64  `json:"from_group_id"`   // 从组id
	Relation      int64  `json:"relation"`        // 关系，-1禁止(非自己)，1可读(非自己)，2可读写(非自己)
	State         int64  `json:"state"`           // 状态，-2删除，-1禁用，待审核0，启用1
	CreateTime    string `json:"create_time"`     // 创建时间
	UpdateTime    string `json:"update_time"`     // 更新时间
}

type GroupGroupRelationList struct {
	List     []*GroupGroupRelation `json:"list"tag:"list"`                     // 数据列表
	Current  int64                 `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64                 `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool                  `json:"isNext"tag:"bool"`                   // 是否有下一页
}

type CreateConfigReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Key    string `json:"key"`     // key
	Value  string `json:"value"`   // value
}

type UpdateConfigReq struct {
	Id     int64  `json:"id"`               // 主键
	UserId int64  `json:"user_id,optional"` // 用户id
	Key    string `json:"key,optional"`     // key
	Value  string `json:"value,optional"`   // value
	State  int64  `json:"state,optional"`   // 状态，-2删除，-1禁用，待审核0，启用1
}

type Config struct {
	Id         int64  `json:"id"`          // 主键
	UserId     int64  `json:"user_id"`     // 用户id
	Key        string `json:"key"`         // key
	Value      string `json:"value"`       // value
	State      int64  `json:"state"`       // 状态，-2删除，-1禁用，待审核0，启用1
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

type ConfigList struct {
	List     []*Config `json:"list"tag:"list"`                     // 数据列表
	Current  int64     `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64     `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool      `json:"isNext"tag:"bool"`                   // 是否有下一页
}

type CreateProjectReq struct {
	Name        string `json:"name"`         // 名称
	Ico         string `json:"ico"`          // 图标
	Info        string `json:"info"`         // 简介
	ProjectType int64  `json:"project_type"` // 类型: 1编程、2其他
	CreateUser  int64  `json:"create_user"`  // 创建者id
	JoinUsers   string `json:"join_users"`   // 参与者ids
	JoinGroups  string `json:"join_groups"`  // 参与组ids
	Remark      string `json:"remark"`       // 备注
	Rank        int64  `json:"rank"`         // 排序
}

type UpdateProjectReq struct {
	Id          int64  `json:"id"`                    // 主键
	Name        string `json:"name,optional"`         // 名称
	Ico         string `json:"ico,optional"`          // 图标
	Info        string `json:"info,optional"`         // 简介
	ProjectType int64  `json:"project_type,optional"` // 类型: 1编程、2其他
	CreateUser  int64  `json:"create_user,optional"`  // 创建者id
	JoinUsers   string `json:"join_users,optional"`   // 参与者ids
	JoinGroups  string `json:"join_groups,optional"`  // 参与组ids
	Remark      string `json:"remark,optional"`       // 备注
	Rank        int64  `json:"rank,optional"`         // 排序
	State       int64  `json:"state,optional"`        // 状态，-2删除，-1禁用，待审核0，启用1
}

type Project struct {
	Id          int64  `json:"id"`           // 主键
	Name        string `json:"name"`         // 名称
	Ico         string `json:"ico"`          // 图标
	Info        string `json:"info"`         // 简介
	ProjectType int64  `json:"project_type"` // 类型: 1编程、2其他
	CreateUser  int64  `json:"create_user"`  // 创建者id
	JoinUsers   string `json:"join_users"`   // 参与者ids
	JoinGroups  string `json:"join_groups"`  // 参与组ids
	Remark      string `json:"remark"`       // 备注
	Rank        int64  `json:"rank"`         // 排序
	State       int64  `json:"state"`        // 状态，-2删除，-1禁用，待审核0，启用1
	CreateTime  string `json:"create_time"`  // 创建时间
	UpdateTime  string `json:"update_time"`  // 更新时间
}

type ProjectList struct {
	List     []*Project `json:"list"tag:"list"`                     // 数据列表
	Current  int64      `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64      `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool       `json:"isNext"tag:"bool"`                   // 是否有下一页
}

type CreateApplicationReq struct {
	ZnName     string `json:"zn_name"`     // 中文名称
	EnName     string `json:"en_name"`     // 英文名称，相当于程序名称
	Ico        string `json:"ico"`         // 图标
	Info       string `json:"info"`        // 简介
	CreateUser int64  `json:"create_user"` // 创建者id
	DemandIds  int64  `json:"demand_ids"`  // 需求组ids
	DocIds     string `json:"doc_ids"`     // 文档组ids
	JoinUsers  string `json:"join_users"`  // 参与者ids
	JoinGroups string `json:"join_groups"` // 参与组ids
	ProjectId  string `json:"project_id"`  // 所属项目id
	Remark     string `json:"remark"`      // 备注
	Rank       int64  `json:"rank"`        // 排序
}

type UpdateApplicationReq struct {
	Id         int64  `json:"id"`                   // 主键
	ZnName     string `json:"zn_name,optional"`     // 中文名称
	EnName     string `json:"en_name,optional"`     // 英文名称，相当于程序名称
	Ico        string `json:"ico,optional"`         // 图标
	Info       string `json:"info,optional"`        // 简介
	CreateUser int64  `json:"create_user,optional"` // 创建者id
	DemandIds  int64  `json:"demand_ids,optional"`  // 需求组ids
	DocIds     string `json:"doc_ids,optional"`     // 文档组ids
	JoinUsers  string `json:"join_users,optional"`  // 参与者ids
	JoinGroups string `json:"join_groups,optional"` // 参与组ids
	ProjectId  string `json:"project_id,optional"`  // 所属项目id
	Remark     string `json:"remark,optional"`      // 备注
	Rank       int64  `json:"rank,optional"`        // 排序
	State      int64  `json:"state,optional"`       // 状态，-2删除，-1禁用，待审核0，启用1
}

type Application struct {
	Id         int64  `json:"id"`          // 主键
	ZnName     string `json:"zn_name"`     // 中文名称
	EnName     string `json:"en_name"`     // 英文名称，相当于程序名称
	Ico        string `json:"ico"`         // 图标
	Info       string `json:"info"`        // 简介
	CreateUser int64  `json:"create_user"` // 创建者id
	DemandIds  int64  `json:"demand_ids"`  // 需求组ids
	DocIds     string `json:"doc_ids"`     // 文档组ids
	JoinUsers  string `json:"join_users"`  // 参与者ids
	JoinGroups string `json:"join_groups"` // 参与组ids
	ProjectId  string `json:"project_id"`  // 所属项目id
	Remark     string `json:"remark"`      // 备注
	Rank       int64  `json:"rank"`        // 排序
	State      int64  `json:"state"`       // 状态，-2删除，-1禁用，待审核0，启用1
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

type ApplicationList struct {
	List     []*Application `json:"list"tag:"list"`                     // 数据列表
	Current  int64          `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64          `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool           `json:"isNext"tag:"bool"`                   // 是否有下一页
}

type CreateApplicationInfoReq struct {
	CreateUser    int64  `json:"create_user"`    // 所属用户
	ApplicationId int64  `json:"application_id"` // 应用id
	Version       string `json:"version"`        // 版本号
	ConfigIds     string `json:"config_ids"`     // 配置ids ","号分隔
}

type UpdateApplicationInfoReq struct {
	Id            int64  `json:"id"`                      // 主键
	CreateUser    int64  `json:"create_user,optional"`    // 所属用户
	ApplicationId int64  `json:"application_id,optional"` // 应用id
	Version       string `json:"version,optional"`        // 版本号
	ConfigIds     string `json:"config_ids,optional"`     // 配置ids ","号分隔
	State         int64  `json:"state,optional"`          // 状态，-2删除，-1禁用，待审核0，启用1
}

type ApplicationInfo struct {
	Id            int64  `json:"id"`             // 主键
	CreateUser    int64  `json:"create_user"`    // 所属用户
	ApplicationId int64  `json:"application_id"` // 应用id
	Version       string `json:"version"`        // 版本号
	ConfigIds     string `json:"config_ids"`     // 配置ids ","号分隔
	State         int64  `json:"state"`          // 状态，-2删除，-1禁用，待审核0，启用1
	CreateTime    string `json:"create_time"`    // 创建时间
	UpdateTime    string `json:"update_time"`    // 更新时间
}

type ApplicationInfoList struct {
	List     []*ApplicationInfo `json:"list"tag:"list"`                     // 数据列表
	Current  int64              `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64              `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool               `json:"isNext"tag:"bool"`                   // 是否有下一页
}

type CreateDocReq struct {
	Name         string `json:"name"`          // 文档标题
	CreateUser   int64  `json:"create_user"`   // 所属用户
	PreContent   string `json:"pre_content"`   // 编辑内容
	Content      string `json:"content"`       // 文档内容
	ParentDoc    int64  `json:"parent_doc"`    // 上级文档
	GroupId      int64  `json:"group_id"`      // 所属文档组
	Sort         int64  `json:"sort"`          // 排序
	EditorMode   int64  `json:"editor_mode"`   // 编辑器模式,1表示Editormd编辑器，2表示Vditor编辑器，3表示iceEditor编辑器
	OpenChildren int64  `json:"open_children"` // 展开下级目录
	ShowChildren int64  `json:"show_children"` // 显示下级文档
}

type UpdateDocReq struct {
	Id           int64  `json:"id"`                     // 主键
	Name         string `json:"name,optional"`          // 文档标题
	CreateUser   int64  `json:"create_user,optional"`   // 所属用户
	PreContent   string `json:"pre_content,optional"`   // 编辑内容
	Content      string `json:"content,optional"`       // 文档内容
	ParentDoc    int64  `json:"parent_doc,optional"`    // 上级文档
	GroupId      int64  `json:"group_id,optional"`      // 所属文档组
	Sort         int64  `json:"sort,optional"`          // 排序
	EditorMode   int64  `json:"editor_mode,optional"`   // 编辑器模式,1表示Editormd编辑器，2表示Vditor编辑器，3表示iceEditor编辑器
	OpenChildren int64  `json:"open_children,optional"` // 展开下级目录
	ShowChildren int64  `json:"show_children,optional"` // 显示下级文档
	State        int64  `json:"state,optional"`         // 文档状态，-2删除，-1禁用，待审核-草稿0，启用1
}

type Doc struct {
	Id           int64  `json:"id"`            // 主键
	Name         string `json:"name"`          // 文档标题
	CreateUser   int64  `json:"create_user"`   // 所属用户
	PreContent   string `json:"pre_content"`   // 编辑内容
	Content      string `json:"content"`       // 文档内容
	ParentDoc    int64  `json:"parent_doc"`    // 上级文档
	GroupId      int64  `json:"group_id"`      // 所属文档组
	Sort         int64  `json:"sort"`          // 排序
	EditorMode   int64  `json:"editor_mode"`   // 编辑器模式,1表示Editormd编辑器，2表示Vditor编辑器，3表示iceEditor编辑器
	OpenChildren int64  `json:"open_children"` // 展开下级目录
	ShowChildren int64  `json:"show_children"` // 显示下级文档
	State        int64  `json:"state"`         // 文档状态，-2删除，-1禁用，待审核-草稿0，启用1
	CreateTime   string `json:"create_time"`   // 创建时间
	UpdateTime   string `json:"update_time"`   // 更新时间
}

type DocList struct {
	List     []*Doc `json:"list"tag:"list"`                     // 数据列表
	Current  int64  `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64  `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool   `json:"isNext"tag:"bool"`                   // 是否有下一页
}

type CreateDocHistoryReq struct {
	PreContent string `json:"pre_content"` // 编辑内容
	CreateUser int64  `json:"create_user"` // 所属用户
	DocId      int64  `json:"doc_id"`      // 文档id
}

type UpdateDocHistoryReq struct {
	Id         int64  `json:"id"`                   // 主键
	PreContent string `json:"pre_content,optional"` // 编辑内容
	CreateUser int64  `json:"create_user,optional"` // 所属用户
	DocId      int64  `json:"doc_id,optional"`      // 文档id
}

type DocHistory struct {
	Id         int64  `json:"id"`          // 主键
	PreContent string `json:"pre_content"` // 编辑内容
	CreateUser int64  `json:"create_user"` // 所属用户
	DocId      int64  `json:"doc_id"`      // 文档id
	CreateTime string `json:"create_time"` // 创建时间
}

type DocHistoryList struct {
	List     []*DocHistory `json:"list"tag:"list"`                     // 数据列表
	Current  int64         `json:"current"tag:"uint"min:"1"max:"1000"` // 当前页
	PageSize int64         `json:"pageSize"tag:"uint"min:"10"max:"20"` // 页面大小
	IsNext   bool          `json:"isNext"tag:"bool"`                   // 是否有下一页
}
