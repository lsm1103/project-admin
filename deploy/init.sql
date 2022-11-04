-- ----------------------------组服务中心----------------------------
-- ------------group 表----------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group`
(
    `id`             bigint unsigned NOT NULL COMMENT '自增主键',
    `name`           char(20) NOT NULL COMMENT '组名称',
    `create_user`    bigint NOT NULL COMMENT '创建者id',
    `ico`            char(20) NOT NULL DEFAULT '' COMMENT '组图标',
    `remark`         varchar(250) NOT NULL COMMENT '备注',
    `parent_id`      bigint NOT NULL COMMENT '父级id',
    `group_type`     tinyint(100) NOT NULL DEFAULT '1' COMMENT '类型: 1部门、2用户组、3群组、4圈子、5话题',
    `rank`           tinyint(100) NOT NULL DEFAULT '1' COMMENT '排序',
    `state`          tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`create_user`,`name`),
    KEY `index_filter` (`create_time`, `state`, `group_type`, `parent_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='组织管理中心';

-- ------------user_group 表----------------
DROP TABLE IF EXISTS `user_group`;
CREATE TABLE `user_group`
(
    `id`             bigint unsigned NOT NULL COMMENT '自增主键',
    `user_id`        bigint NOT NULL COMMENT '用户id',
    `group_id`       bigint NOT NULL COMMENT '组id',
    `state`          tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`user_id`,`group_id`),
    KEY `index_filter` (`create_time`, `state`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户、组、角色关系表';

-- ------------project 表，类似于机构、saas里面的公司，作为数据隔离的作用----------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project`
(
    `id`             bigint unsigned NOT NULL COMMENT '主键',
    `name`           char(20) NOT NULL COMMENT '名称',
    `ico`            char(20) NOT NULL DEFAULT '' COMMENT '图标',
    `info`           char(20) NOT NULL COMMENT '简介',
    `project_type`   tinyint(100) NOT NULL DEFAULT '1' COMMENT '类型: 1编程、2其他',
    `create_user`    bigint NOT NULL COMMENT '创建者id',
    `join_users`     char(200) NOT NULL COMMENT '参与者ids',
    `join_groups`    char(200) NOT NULL COMMENT '参与组ids',
    `remark`         varchar(250) NOT NULL COMMENT '备注',
    `rank`           tinyint(100) NOT NULL DEFAULT '1' COMMENT '排序',
    `state`          tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`create_user`,`name`, `project_type`),
    KEY `index_filter` (`create_time`, `state`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='项目表';

-- ------------application 表----------------
DROP TABLE IF EXISTS `application`;
CREATE TABLE `application`
(
    `id`             bigint unsigned NOT NULL COMMENT '主键',
    `zn_name`        char(20) NOT NULL COMMENT '中文名称',
    `en_name`        char(20) NOT NULL COMMENT '英文名称，相当于程序名称',
    `ico`            char(20) NOT NULL DEFAULT '' COMMENT '图标',
    `info`           char(20) NOT NULL COMMENT '简介',
    `create_user`    bigint NOT NULL COMMENT '创建者id',
    `demand_ids`     bigint NOT NULL COMMENT '需求组ids',
    `doc_ids`        char(200) NOT NULL COMMENT '文档组ids',
    `join_users`     char(200) NOT NULL COMMENT '参与者ids',
    `join_groups`    char(200) NOT NULL COMMENT '参与组ids',
    `project_id`     char(200) NOT NULL COMMENT '所属项目id',
    `remark`         varchar(250) NOT NULL COMMENT '备注',
    `rank`           tinyint(100) NOT NULL DEFAULT '1' COMMENT '排序',
    `state`          tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`create_user`, `en_name`, `project_id`),
    KEY `index_filter` (`create_time`, `state`, `zn_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='应用表';

-- ------------config 表----------------
DROP TABLE IF EXISTS `application_config`;
CREATE TABLE `application_config`
(
    `id`             bigint unsigned NOT NULL COMMENT '主键',
    `user_id`        bigint NOT NULL COMMENT '用户id',
    `key`            bigint NOT NULL COMMENT 'key',
    `value`          bigint NOT NULL COMMENT 'value',
    `state`          tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`user_id`,`key`),
    KEY `index_filter` (`create_time`, `state`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='配置表';

-- ------------application-config 表----------------
# 配置github/gitlab代码仓库；配置webhook
DROP TABLE IF EXISTS `application_config`;
CREATE TABLE `application_config`
(
    `id`             bigint unsigned NOT NULL COMMENT '主键',
    `application_id` bigint NOT NULL COMMENT '应用id',
    `config_id`      bigint NOT NULL COMMENT '配置id',
    `state`          tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`application_id`,`config_id`),
    KEY `index_filter` (`create_time`, `state`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='应用配置表';


-- ------------doc 表----------------
-- ------------docGroup 表----------------

-- ------------testCase 表----------------


-- ------------demandDoc 表----------------
-- ------------demandDocGroup 表----------------

-- ------------bug 表----------------
-- ------------bugGroup 表----------------