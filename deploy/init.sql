-- ----------------------------组服务中心----------------------------
-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group`
(
    `id`             bigint unsigned NOT NULL COMMENT '自增主键',
    `name`           varchar(20) NOT NULL COMMENT '组名称',
    `create_user`    bigint NOT NULL COMMENT '创建者id',
    `ico`            varchar(20) NOT NULL COMMENT '组图标',
    `remark`         varchar(250) NOT NULL COMMENT '备注',
    `parent_id`      bigint NOT NULL COMMENT '父级id',
    `group_type`     tinyint(100) NOT NULL DEFAULT '1' COMMENT '类型: 1部门、2用户组、3群组、4圈子、5话题',
    `rank`           tinyint(100) NOT NULL DEFAULT '1' COMMENT '排序',
    `state`         tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_create_user_name` (`create_user`,`name`),
    KEY `indexFilter_create_time` (`create_time`, `state`, `group_type`, `parent_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='组织管理中心';

-- ----------------------------
-- Table structure for user_group
-- ----------------------------
DROP TABLE IF EXISTS `user_group`;
CREATE TABLE `user_group`
(
    `id`             bigint unsigned NOT NULL COMMENT '自增主键',
    `user_id`        bigint NOT NULL COMMENT '用户id',
    `group_id`       bigint NOT NULL COMMENT '组id',
    `state`         tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_group_user_id` (`user_id`,`group_id`),
    KEY `indexFilter_create_time` (`create_time`, `state`, `group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户、组、角色关系表';

-- ------------project 表，类似于机构、saas里面的公司，作为数据隔离的作用----------------
#

-- ------------application 表----------------
# app英文名称，相当于程序名称
# app中文名称
# app简介
# app需求，需求组id
# app文档，可以是单独文档，也可以是文档组id
# app创建人
# app参与人
# app参与组
# 项目id

-- ------------application-config 表----------------
-- 配置github/gitlab代码仓库
-- 配置webhook
-- 配置其他

-- ------------doc 表----------------
-- ------------docGroup 表----------------

-- ------------testCase 表----------------


-- ------------demandDoc 表----------------
-- ------------demandDocGroup 表----------------

-- ------------bug 表----------------
-- ------------bugGroup 表----------------