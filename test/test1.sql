-- ----------------------------
-- Table structure for offline_msg
-- ----------------------------
DROP TABLE IF EXISTS `offline_msg`;
CREATE TABLE `offline_msg`
(
    `id`          bigint unsigned NOT NULL COMMENT '主键',
    `user_id`     bigint NOT NULL COMMENT '用户id',
    `device_id`   varchar(100) NOT NULL COMMENT '设备id',
    `object_type` tinyint  NOT NULL COMMENT '对象类型,1:friend；2：群组',
    `object_id`   bigint NOT NULL COMMENT '对象id, friendId/groupId',
    `last_ack_seq`         bigint NOT NULL COMMENT '最后确认序列号',
    `newest_seq`         bigint NOT NULL COMMENT '最新的消息序列号',
    `state`        tinyint(4) NOT NULL DEFAULT '0' COMMENT '消息状态：-1撤回，0未处理，1已读',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_object` (`user_id`, `device_id`,`object_type`,`object_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='离线消息';

-- ----------------------------
-- Table structure for single_msg
-- ----------------------------
DROP TABLE IF EXISTS `single_msg`;
CREATE TABLE `single_msg`
(
    `id`            bigint unsigned NOT NULL COMMENT '主键',
    `seq`           bigint NOT NULL COMMENT '消息序列号,每个单聊都维护一个消息序列号',
    `sender_type`   tinyint(3) NOT NULL DEFAULT '1' COMMENT '发送者类型：1朋友，2打招呼，3转发',
    `sender_id`     bigint NOT NULL COMMENT '发送者id',
    `sender_device_id` varchar(100) NOT NULL COMMENT '发送设备id',
    `receiver_id`   bigint NOT NULL COMMENT '接收者id(friendId)',
    `receiver_device_id` varchar(100) NOT NULL COMMENT '接收设备id：多个设备id"，"隔开，*表示全部设备',
    `msg_type`      tinyint(4) NOT NULL DEFAULT '0' COMMENT '消息类型：0文本、1图文、2语音、3视频、4链接',
    `content`       blob NOT NULL COMMENT '消息内容',
    `parent_id`     bigint NOT NULL COMMENT '父级id，引用功能',
    `send_time`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '消息发送时间',
    `state`        tinyint(4) NOT NULL DEFAULT '0' COMMENT '消息状态：-1撤回，0未处理，1已读',
    `create_time`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `k_send_time_state_msg_type_sender_type` (`send_time`,`state`,`msg_type`,`sender_type`),
    UNIQUE KEY `uk_seq_sender_id_receiver_id` (`seq`,`sender_id`, `receiver_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='单聊消息表';