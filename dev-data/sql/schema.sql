create database if not exists `iris-backend`;

use `iris-backend`;

-- 用户信息表
create table if not exists `tb_user`
(
    `id`          varchar(32)  not null primary key,
    `username`    varchar(100) not null comment '登录名',
    `nickname`    varchar(100) not null comment '用户昵称',
    `email`       varchar(100) comment '邮箱',
    `password`    varchar(50)  not null comment '密码',
    `avatar`      varchar(100) comment '头像URL路径',
    `create_user` varchar(32) comment '创建用户',
    `create_time` datetime              DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `update_user` varchar(32) comment '更新用户',
    `update_time` datetime on update CURRENT_TIMESTAMP comment '更新时间',
    `is_delete`   varchar(1)   not null default '0' comment '是否删除',
    constraint tb_user_username_uindex unique (`username`)
) engine = InnoDB comment '用户信息表';

-- 文章表
create table if not exists `tb_article`
(
    `id`            varchar(32)  not null primary key,
    `title`         varchar(100) not null comment '文章标题',
    `remark`        varchar(200) not null comment '简介',
    `pic`           varchar(200) comment '背景URL路径',
    `publish_state` varchar(1)            default '0' comment '发布状态',
    `publish_time`  varchar(200) comment '发布时间',
    `create_user`   varchar(32) comment '创建用户',
    `create_time`   datetime              DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `update_user`   varchar(32) comment '更新用户',
    `update_time`   datetime on update CURRENT_TIMESTAMP comment '更新时间',
    `is_delete`     varchar(1)   not null default '0' comment '是否删除'
) engine = InnoDB comment '文章表';

-- 文章表
create table if not exists `tb_article_detail`
(
    `article_id` varchar(32) not null primary key,
    `content`    text comment '文章内容'
) engine = InnoDB comment '文章详情表';


-- 评论表·
create table if not exists `tb_comment`
(
    `id`          varchar(32)  not null primary key,
    `pid`         varchar(32)  not null default '0' comment '父评论id',
    `article_id`  varchar(32)  not null comment '文章id',
    `content`     varchar(200) not null comment '评论内容',
    `create_user` varchar(32) comment '创建用户',
    `create_time` datetime              DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `update_user` varchar(32) comment '更新用户',
    `update_time` datetime on update CURRENT_TIMESTAMP comment '更新时间',
    `is_delete`   varchar(1)   not null default '0' comment '是否删除'
) engine = InnoDB comment '评论表';

-- 查询条件建立索引
create index tb_comment_pid_aid_index on `iris-backend`.`tb_comment` (`pid`, `article_id`);