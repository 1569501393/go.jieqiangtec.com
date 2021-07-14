package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Text struct {
	ID      uint   `gorm:"primaryKey;column:id;type:int unsigned;not null"`    // 主键
	Name    string `gorm:"column:name;type:varchar(255);not null;default:''"`  // 名称
	Title   string `gorm:"column:title;type:varchar(255);not null;default:''"` // 标题
	Desc    string `gorm:"column:desc;type:varchar(255);not null;default:''"`  // 描述
	Content string `gorm:"column:content;type:text;not null;default:''"`       // 内容
	// Type    int    `gorm:"column:type;type:tinyint;not null;default:0"`        // 文本类型
	Type uint8 `gorm:"column:type;type:tinyint;not null;default:0"` // 文本类型
	// CreatedAt uint   `gorm:"column:created_at;type:int unsigned;not null;default:0"` // 创建时间
	// UpdatedAt uint   `gorm:"column:updated_at;type:int unsigned;not null;default:0"` // 更新时间
	CreatedAt uint `gorm:"autoCreateTime"`                                         // 创建时间,自动更新
	UpdatedAt uint `gorm:"autoUpdateTime"`                                         // 更新时间,自动更新
	DeletedAt uint `gorm:"column:deleted_at;type:int unsigned;not null;default:0"` // 删除时间
}

// TableName 为gorm中的表设置单数名称
func (Text) TableName() string {
	return "text"
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3308)/shanya?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	// dsn := "root:root@tcp(127.0.0.1:3308)/shanya?charset=utf8mb4&parseTime=true&loc=local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("err:%v \n", err)
	}

	fmt.Printf("db:%v \n", db)

	// 创建记录
	text := Text{
		Name:    "jieqiang",
		Title:   "title",
		Desc:    "desc",
		Content: "content",
		Type:    1,
		// CreatedAt: 0,
		// UpdatedAt: 0,

		DeletedAt: 0,
	}

	result := db.Create(&text)

	fmt.Printf("创建记录===id:%v, result:%v, text:%v, &text:%v \n", text.ID, result, text, &text)

	// 用指定的字段创建记录
	// 创建记录并更新给出的字段。
	// result = db.Select("Name", "Content").Create(&text)
	result = db.Select("Name", "Title", "Desc", "Type").Create(&text)
	if result != nil {
		fmt.Printf("用指定的字段创建记录===id:%v, error:%v \n", text.ID, result)
	} else {
		fmt.Printf("用指定的字段创建记录===id:%v, ok:%v \n", text.ID, result)
	}

	// 创建记录并更新未给出的字段。
	result = db.Omit("id", "Name").Create(&text)
	fmt.Printf("创建记录并更新未给出的字段result:%v", result)

	// 批量插入
	// 要有效地插入大量记录，请将一个 slice 传递给 Create 方法。 将切片数据传递给 Create 方法，GORM 将生成一个单一的 SQL 语句来插入所有数据，并回填主键的值，钩子方法也会被调用。

	// db.Delete(&text)
}
