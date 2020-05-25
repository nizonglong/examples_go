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

	var users []psql_example.User

	db.Table("user").Where("nick_name = ?", "test7").Or("nick_name = ?", "test8").Find(&users)
	// SELECT * FROM user WHERE nick_name = 'test7' OR nick_name = 'test8';
	fmt.Println(users)

	// Struct
	db.Table("user").Where("nick_name = 'test6'").Or(psql_example.User{NickName: "test2"}).Find(&users)
	// SELECT * FROM user WHERE nick_name = 'test6' OR nick_name = 'test2';
	fmt.Println(users)

	// Map
	db.Table("user").Where("nick_name = 'test01'").Or(map[string]interface{}{"nick_name": "test9"}).Find(&users)
	fmt.Println(users)

	/**
	 *  ORDER
	 */
	db.Table("user").Order("age desc, nick_name").Find(&users)
	// SELECT * FROM user ORDER BY age desc, nick_name;
	fmt.Println(users)

	// Multiple orders
	db.Table("user").Order("age desc").Order("nick_name").Find(&users)
	// SELECT * FROM user ORDER BY age desc, nick_name;
	fmt.Println(users)

	var users1, users2 []psql_example.User
	// ReOrder
	db.Table("user").Order("age desc").Find(&users1).Order("age", true).Find(&users2)
	// SELECT * FROM user ORDER BY age desc; (users1)
	// SELECT * FROM user ORDER BY age; (users2)
	fmt.Println(users1)
	fmt.Println(users2)
}
