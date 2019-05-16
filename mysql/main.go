package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Poet 诗的数据库实体
type Poet struct {
	ID        uint64 `gorm:"Column:id;auto_increment"`
	Author    string `gorm:"Column:author"`
	Paragraph string `gorm:"Column:paragraph"`
	Strains   string `gorm:"Column:strains"`
	Title     string `gorm:"Column:title"`
	Dynasty   string `gorm:"Column:dynasty"`
}

// PoetInJson 诗的Json实体
type PoetInJson struct {
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Strains    []string `json:"strains"`
	Title      string   `json:"title"`
}

func (Poet) TableName() string {
	return "poet"
}

func main() {
	// 打开数据库连接
	db, _ := gorm.Open("mysql", "app:123456@tcp(localhost:3306)/chinese-poetry?charset=utf8mb4&parseTime=True&loc=Local")
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

	//读取文件
	jsonFile, err := os.Open("/home/yz/project/chinese-poetry/json/poet.song.0.json")
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	//解组
	var poets []PoetInJson
	err1 := json.Unmarshal(byteValue, &poets)
	if err1 != nil {
		panic(err1)
	}

	//遍历
	for _, poet := range poets {
		if err := tx.Create(&Poet{Dynasty: "song", Author: poet.Author, Paragraph: strings.Join(poet.Paragraphs, ""), Strains: strings.Join(poet.Strains, ""), Title: poet.Title}).Error; err != nil {
			log.Panicf("%s, %s", poet, err)
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
