package main

import (
	"examples/psql_example"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // 这里很容易忘记加导包
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=example sslmode=disable password=nizonglong")
	defer db.Close()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	var user psql_example.User
	// Save将包括执行更新SQL时的所有字段，即使它没有更改
	db.First(&user)
	user.Id = 12
	user.NickName = "test5"
	user.Age = 23
	db.Save(&user)
}
