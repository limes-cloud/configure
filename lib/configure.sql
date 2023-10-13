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
    unique index(`keyword`)
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


