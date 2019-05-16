package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Poet 诗
type Poet struct {
	gorm.Model
	ID        uint64 `gorm:"Column:id;auto_increment"`
	Author    string `gorm:"Column:author"`
	Paragraph string `gorm:"Column:paragraph"`
	Strains   string `gorm:"Column:strains"`
	Title     string `gorm:"Column:title"`
}

func main() {

}
