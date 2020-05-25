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

	var users1 []psql_example.User
	var users2 []psql_example.User

	db.Table("user").Limit(3).Find(&users1)
	//// SELECT * FROM user LIMIT 3;

	// Cancel limit condition with -1
	db.Table("user").Limit(10).Find(&users1).Limit(-1).Find(&users2)
	//// SELECT * FROM users LIMIT 10; (users1)
	//// SELECT * FROM users; (users2)
}
