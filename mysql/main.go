package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Poet 诗
type Poet struct {
	ID        uint64 `gorm:"Column:id;auto_increment"`
	Author    string `gorm:"Column:author"`
	Paragraph string `gorm:"Column:paragraph"`
	Strains   string `gorm:"Column:strains"`
	Title     string `gorm:"Column:title"`
	Dynasty   string `gorm:"Column:dynasty"`
}

func (Poet) TableName() string {
	return "poet"
}

func main() {
	// 打开数据库连接
	db, _ := gorm.Open("mysql", "app:123456@tcp(localhost:3306)/chinese-poetry?charset=utf8&parseTime=True")
	err := InsertPoet(db)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func InsertPoet(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&Poet{Dynasty: "tang", Author: "李白", Paragraph: "意识"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
