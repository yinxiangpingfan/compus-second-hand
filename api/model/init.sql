-- Active: 1762331896607@@127.0.0.1@3306@compus_second_hand

CREATE TABLE campus (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    is_able BOOLEAN NOT NULL DEFAULT TRUE-- 校区是否可用
)

insert into campus (name, created_at, updated_at)
 values 
 ('凌水校区', NOW(), NOW()),
 ('开发区校区', NOW(), NOW()),
 ('盘锦校区', NOW(), NOW());


CREATE TABLE user (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    gender int,-- 0是男，1是女
    campus_id INT NOT NULL,-- 是哪个校区的用户
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_able BOOLEAN NOT NULL DEFAULT TRUE,-- 账号是否可用
    Foreign Key (campus_id) REFERENCES campus(id)
)

alter table user add column pic VARCHAR(255) NOT NULL DEFAULT '';-- 用户头像

CREATE TABLE category (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,-- 分类名称
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
)

CREATE TABLE goods (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,-- 哪个用户发布的
    title VARCHAR(255) NOT NULL,-- 商品标题
    description TEXT NOT NULL,-- 商品描述
    price DECIMAL(10,2) NOT NULL,-- 商品价格
    category_id BIGINT NOT NULL,-- 商品分类
    collect_count INT NOT NULL DEFAULT 0,-- 收藏数
    campus_id INT NOT NULL,-- 是哪个校区的商品
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    state INT NOT NULL DEFAULT 0,-- 0正在售卖 1 已经交易完成 2 商品被拉黑
    Foreign Key (user_id) REFERENCES user(id),
    Foreign Key (campus_id) REFERENCES campus(id),
    Foreign Key (category_id) REFERENCES category(id)
)

CREATE TABLE goods_image (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    goods_id BIGINT NOT NULL,-- 哪个商品的图片
    image_url VARCHAR(255) NOT NULL,-- 图片url
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    Foreign Key (goods_id) REFERENCES goods(id)
)

CREATE TABLE goods_comment (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    goods_id BIGINT NOT NULL,-- 哪个商品的评论
    user_id BIGINT NOT NULL,-- 哪个用户评论的
    comment TEXT NOT NULL,-- 评论内容
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    Foreign Key (goods_id) REFERENCES goods(id),
    Foreign Key (user_id) REFERENCES user(id)
)

CREATE TABLE collect(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    goods_id BIGINT NOT NULL,-- 商品id
    user_id BIGINT NOT NULL,-- 收藏的用户
    Foreign Key (goods_id) REFERENCES goods(id),
    Foreign Key (user_id) REFERENCES user(id)
)

CREATE TABLE `order` (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    goods_id BIGINT NOT NULL,-- 商品id
    buy_id BIGINT NOT NULL,-- 买家id
    sell_id BIGINT NOT NULL,-- 卖家id
    order_amount DECIMAL(10,2) NOT NULL,-- 订单金额
    order_time DATETIME NOT NULL,-- 订单时间
    Foreign Key (goods_id) REFERENCES goods(id),
    Foreign Key (buy_id) REFERENCES user(id),
    Foreign Key (sell_id) REFERENCES user(id)
)
