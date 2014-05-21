package models

import (
		"database/sql"
		"github.com/astaxie/beedb"
		_ "github.com/go-sql-driver/mysql"
		"time"
		"fmt"
)
type Blog struct {
		Id int `PK`
		Title string
		Content string
		Created time.Time
}
func GetLink() beedb.Model {
		db, err := sql.Open("mysql","root:Whwbn123@/test?charset=utf8") 
		if err != nil {
				panic(err)
		}
		orm := beedb.New(db)
		return orm
}
func GetAll() (blogs []Blog) {
		db := GetLink()
		db.FindAll(&blogs)
		return
}
func GetBlog(id int) (blog Blog) {
		db := GetLink()
		db.Where("id=?", id).Find(&blog)
		return
}
func SaveBlog(blog Blog) (bg Blog) {
		db := GetLink()
		err := db.Save(&blog)
		if err != nil{
				fmt.Println("db error")
				fmt.Println(err)
		}
		return bg
}

func DelBlog(blog Blog) {
		db := GetLink()
	    fmt.Println(blog)
		_, err := db.Delete(&blog)
		if err != nil{
				fmt.Println("db error")
				fmt.Println(err)
		}
		return
}

