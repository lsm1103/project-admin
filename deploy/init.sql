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
  COLLATE = utf8mb4_bin COMMENT ='用户、组关系表';

-- ------------group_group_relation 表----------------
DROP TABLE IF EXISTS `group_group_relation`;
CREATE TABLE `group_group_relation` (
    `id`              bigint unsigned NOT NULL COMMENT '自增主键',
    `create_user`     bigint NOT NULL COMMENT '创建者id',
    `master_group_id` bigint NOT NULL COMMENT '主组id',
    `from_group_id`   bigint NOT NULL COMMENT '从组id',
    `relation`        tinyint(4) NOT NULL DEFAULT '1' COMMENT '关系，-1不允许，1允许读写(非自己)，2只允许读(非自己)',
    `state`           tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`     datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`create_user`,`master_group_id`,`from_group_id`,`relation`),
    KEY `index_filter` (`create_time`, `state`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='组与组关系描述表';

-- ------------config 表----------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config`
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
    `ico`            char(200) NOT NULL DEFAULT '' COMMENT '图标',
    `info`           char(200) NOT NULL COMMENT '简介',
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

-- ------------application-config 表----------------
# 配置github/gitlab代码仓库, 配置webhook
DROP TABLE IF EXISTS `application_config`;
CREATE TABLE `application_config`
(
    `id`             bigint unsigned NOT NULL COMMENT '主键',
    `create_user`   int(20) NOT NULL COMMENT '所属用户',
    `application_id` bigint NOT NULL COMMENT '应用id',
    `config_id`      bigint NOT NULL COMMENT '配置id',
    `state`          tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`create_user`,`application_id`,`config_id`),
    KEY `index_filter` (`create_time`, `state`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='应用配置表';

-- ------------doc 表----------------
DROP TABLE IF EXISTS `doc`;
CREATE TABLE `doc` (
    `id`            bigint unsigned NOT NULL COMMENT '主键',
    `name`          char(50) NOT NULL COMMENT '文档标题',
    `create_user`   int(20) NOT NULL COMMENT '所属用户',
    `pre_content`   text NOT NULL COMMENT '编辑内容',
    `content`       text NOT NULL COMMENT '文档内容',
    `parent_doc`    int(20) NOT NULL COMMENT '上级文档',
    `group_id`       int(20) NOT NULL COMMENT '所属文档组',
    `sort`          int(20) NOT NULL COMMENT '排序',
    `editor_mode`   tinyint(4) NOT NULL COMMENT '编辑器模式,1表示Editormd编辑器，2表示Vditor编辑器，3表示iceEditor编辑器',
    `open_children` tinyint(4) NOT NULL COMMENT '展开下级目录',
    `show_children` tinyint(4) NOT NULL COMMENT '显示下级文档',
    `state`         tinyint(4) NOT NULL DEFAULT '1' COMMENT '文档状态，-2删除，-1禁用，待审核-草稿0，启用1',
    `create_time`   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique` (`name`,`create_user`),
    KEY `index_filter` (`create_time`, `state`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='文档表';
-- ------------dochistory 文档历史表----------------
DROP TABLE IF EXISTS `doc_history`;
CREATE TABLE `doc_history` (
    `id`            bigint unsigned NOT NULL COMMENT '主键',
    `pre_content`   text NOT NULL COMMENT '编辑内容',
    `create_user`   int(20) NOT NULL COMMENT '所属用户',
    `doc_id`        int(20) NOT NULL COMMENT '文档id',
    `create_time`   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `index_filter` (`create_time`,`create_user`, `doc_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='文档历史表';


-- ------------demand 表----------------
-- ------------demandGroup 表----------------


-- ------------testCase 表----------------


-- ------------bug 表----------------
-- ------------bugGroup 表----------------