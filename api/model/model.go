package model

//数据库映射的结构体

import (
	"time"
)

// Campus 校区
type Campus struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsAble    bool // 校区是否可用
}

// User 用户
type User struct {
	ID        uint64
	Username  string
	Password  string
	Gender    int // 0 男 1 女
	Email     string
	Pic       string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsAble    bool   // 账号是否可用
	CampusID  uint64 `db:"campus_id"`
}

// Category 商品分类
type Category struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Goods 商品
type Goods struct {
	ID           uint64
	UserID       uint64
	Title        string
	Description  string
	Price        float64
	CategoryID   uint64
	CollectCount uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	State        int
	CampusID     int
}

// GoodsImage 商品图片
type GoodsImage struct {
	ID        uint64
	GoodsID   uint64
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GoodsComment 商品评论
type GoodsComment struct {
	ID        uint64
	GoodsID   uint64
	UserID    uint64
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Collect 收藏
type Collect struct {
	ID      uint64
	GoodsID uint64
	UserID  uint64

	Goods Goods
	User  User
}

// Order 订单
type Order struct {
	ID          uint64
	GoodsID     uint64
	BuyID       uint64
	SellID      uint64
	OrderAmount float64
	OrderTime   time.Time
}
