type (
	create{{.tableName}}Req {
	    {{.createFields}}
	}

	update{{.tableName}}Req {
	    {{.updateFields}}
	}
	{{.tableName}} {
	    {{.allFields}}
        /*
            Id               int64  `json:"id"`                                     // 主键
            Seq              int64  `json:"seq"`                                    // 消息主键,每个单聊都维护一个消息主键
            SenderType       string `json:"sender_type"`                            // 发送者类型：1发消息，2打招呼，3转发
            SenderId         int64  `json:"sender_id"`                              // 发送者id
            SenderDeviceId   string `json:"sender_device_id"`                       // 发送设备id
            ReceiverId       int64  `json:"receiver_id"`                            // 接收者id
            ReceiverDeviceId string `json:"receiver_device_id"`                     // 接收设备id：多个设备id"，"隔开，*表示全部设备
            MsgType          string `json:"msg_type"`                               // 消息类型
            Content          string `json:"content"`                                // 消息内容
            ParentId         int64  `json:"parent_id"`                              // 父级id，引用功能
            SendTime         string `json:"send_time"`                              // 消息发送时间
            State            int64  `json:"state"`                                  // 消息状态：-1撤回，0未处理，1已读
            CreateTime       string `json:"createTime"`                             // 创建时间
            UpdateTime       string `json:"updateTime"`                             // 更新时间
        */
	}

	// 查询列表结果
	{{.tableName}}List {
		List     []*{{.tableName}} `json:"list"tag:"list"`                       // 数据列表
		Current  int64        `json:"current"tag:"uint"min:"1"max:"10"`     // 当前页
		PageSize int64        `json:"pageSize"tag:"uint"content:"10|20"`    // 页面大小
		IsNext   bool         `json:"isNext"tag:"char"content:"true|false"` // 是否有下一页
		// total           int64        `json:"total"`  // 总数
	}
)
