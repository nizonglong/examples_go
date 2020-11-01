package gorm

import (
	"examples_go/define"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // 这里很容易忘记加导包
)

func MainSelectInner() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=example sslmode=disable password=nizonglong")
	defer db.Close()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	// 使用主键查询时，应仔细检查所传递的值是否为有效主键，以避免SQL注入
	var user define.User
	var users []define.User

	// 按主键获取
	db.Table("user").First(&user, 23)
	// SELECT * FROM user WHERE id = 23 LIMIT 1;
	fmt.Println(user)

	// 简单SQL
	db.Table("user").Find(&user, "nick_name = ?", "test3")
	// SELECT * FROM user WHERE nick_name = "test3";
	fmt.Println(user)

	db.Table("user").Find(&users, "nick_name <> ? AND age > ?", "test3", 21)
	// SELECT * FROM user WHERE nick_name <> "jinzhu" AND age > 20;
	fmt.Println(users)

	// Struct
	db.Table("user").Find(&users, define.User{Age: 20})
	// SELECT * FROM user WHERE age = 20;
	fmt.Println(users)

	// Map
	db.Table("user").Find(&users, map[string]interface{}{"age": 29})
	// SELECT * FROM user WHERE age = 20;
	fmt.Println(users)
}
