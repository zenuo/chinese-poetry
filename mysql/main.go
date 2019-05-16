package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// PoetInJSON 诗的Json实体
type PoetInJSON struct {
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Strains    []string `json:"strains"`
	Title      string   `json:"title"`
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

// InsertPoet 读取诗并INSERT到数据库
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

	//遍历文件夹
	files, err := ioutil.ReadDir("/home/yz/project/chinese-poetry/json")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	//正则模式
	validPattern := regexp.MustCompile(`^poet\.(.+?)\.[0-9]+\.json`)
	for _, file := range files {
		//若匹配
		if validPattern.MatchString(file.Name()) {
			//捕获朝代
			dynasty := validPattern.FindStringSubmatch(file.Name())[1]
			//读取文件
			jsonFile, err := os.Open(filepath.Join("..", "json", file.Name()))
			if err != nil {
				log.Panic(err)
				panic(err)
			}
			byteValue, _ := ioutil.ReadAll(jsonFile)
			defer jsonFile.Close()

			//解组
			var poets []PoetInJSON
			err1 := json.Unmarshal(byteValue, &poets)
			if err1 != nil {
				panic(err1)
			}

			//遍历
			for _, poet := range poets {
				//执行INSERT
				if err := tx.Exec("INSERT  INTO `poet` (`author`,`paragraph`,`strains`,`title`,`dynasty`) VALUES (?,?,?,?,?)", poet.Author, strings.Join(poet.Paragraphs, ""), strings.Join(poet.Strains, ""), poet.Title, dynasty).Error; err != nil {
					log.Panicf("%s, %s", poet, err)
					//回滚
					tx.Rollback()
					return err
				}
			}
		}
	}

	return tx.Commit().Error
}
