create database configure;
use configure;

#环境表
create table environment(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `keyword` char(32) binary not null comment '环境标识',
    `name` char(32) not null comment '环境名称',
    `description` varchar(256) not null  comment '环境描述',
    `token` char(32) not null comment '环境token',
    `status` bool not null comment '启用状态',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    `updated_at`bigint unsigned default null  comment '修改时间',
    index(`created_at`),
    unique index(`keyword`),
    unique index(`token`)
)engine innodb charset utf8;


#服务表
create table `server`(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `keyword` char(32) binary not null comment '服务标识',
    `name` char(32) not null comment '服务名称',
    `description` varchar(256) not null  comment '服务描述',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    `updated_at`bigint unsigned default null  comment '修改时间',
    index(`created_at`),
    unique index(`keyword`)
)engine innodb charset utf8;


#资源表
create table `resource`(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `keyword` char(32) binary not null comment '资源标识',
    `fields` text not null comment '资源字段',
    `tag` char(32) not null comment '资源标签;mysql/mongo/consul..',
    `description` varchar(256) not null  comment '资源描述',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    `updated_at`bigint unsigned default null  comment '修改时间',
    index(`created_at`),
    unique index(`keyword`)
)engine innodb charset utf8;


# 资源服务表，限定资源可以被哪些服务使用
create table resource_server(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `resource_id` bigint unsigned not null comment '资源id',
    `server_id` bigint unsigned not null comment '服务id',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    unique index(`resource_id`,`server_id`)
)engine innodb charset utf8;
# resource_value 资源值
create table resource_value(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `environment_id` bigint unsigned not null comment '环境id',
    `resource_id` bigint unsigned not null comment '资源id',
    `values` text not null comment '资源值',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    `updated_at`bigint unsigned default null  comment '修改时间',
    unique index(`environment_id`,`resource_id`)
)engine innodb charset utf8;


#业务表
create table `business`(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `server_id` bigint unsigned not null comment '服务id',
    `keyword` char(32) binary not null comment '业务标识',
    `description` varchar(256) not null  comment '业务描述',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    `updated_at`bigint unsigned default null  comment '修改时间',
    index(`created_at`),
    unique index(`server_id`,`keyword`)
)engine innodb charset utf8;


# 业务字段值
create table business_value(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `environment_id` bigint unsigned not null comment '环境id',
    `business_id` bigint unsigned not null comment '业务id',
    `value` text not null comment '业务值',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    `updated_at`bigint unsigned default null  comment '修改时间',
    unique index(`environment_id`,`business_id`)
)engine innodb charset utf8;


create table template(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `server_id` bigint unsigned not null comment '服务id',
    `content` text not null comment '模板内容',
    `description` varchar(256) not null comment '版本描述',
    `version` varchar(64) not null comment '模板版本',
    `is_use` bool not null comment  '是否使用',
    `format` char(32) not null comment '模板格式',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    `updated_at`bigint unsigned default null  comment '修改时间',
    unique index(`server_id`,`version`)
)engine innodb charset utf8;


create table configure(
    `id` bigint unsigned not null primary key auto_increment comment '自增id',
    `server_id` bigint unsigned not null comment '服务id',
    `environment_id` bigint unsigned not null comment '环境id',
    `content` text not null comment '模板内容',
    `description` varchar(256) not null comment '版本描述',
    `version` varchar(64) not null comment '模板版本',
    `format` char(32) not null comment '模板格式',
    `operator` varchar(32) not null comment '操作人',
    `operator_id` bigint unsigned not null comment '操作人id',
    `created_at` bigint unsigned default null  comment '创建时间',
    `updated_at`bigint unsigned default null  comment '修改时间',
    unique index(`server_id`,`environment_id`)
)engine innodb charset utf8;


# create table operate_log(
#     `id` bigint unsigned not null primary key auto_increment comment '自增id',
#     `server_id` bigint unsigned not null comment '服务id',
#     `environment_id` bigint unsigned not null comment '环境id',
#     `compare` text not null comment '差异内容',
#     `created_at` bigint unsigned default null  comment '创建时间',
#     unique index(`server_id`,`version`)
# )engine innodb charset utf8;

