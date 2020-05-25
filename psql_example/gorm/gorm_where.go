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
	// 获取第一个匹配记录
	db.Table("user").Where("nick_name = ?", "test5").First(&user)
	// SELECT * FROM user WHERE nick_name = 'test5' limit 1;
	fmt.Println(user)

	var users []psql_example.User
	// 获取所有匹配记录
	db.Table("user").Where("id > ?", 7).Find(&users)
	// SELECT * FROM user WHERE id > 7;
	fmt.Println(users)

	// IN
	db.Table("user").Where("nick_name in (?)", []string{"test5", "test6", "test7"}).Find(&users)
	fmt.Println(users)

	// LIKE
	db.Table("user").Where("nick_name LIKE ?", "test%").Find(&users)
	fmt.Println(users)

	// AND
	db.Table("user").Where("nick_name = ? AND age >= ?", "test9", "18").Find(&users)
	fmt.Println(users)

	// Struct
	db.Table("user").Where(&psql_example.User{NickName: "test7", Age: 25}).First(&user)
	// SELECT * FROM user WHERE nick_name = "test7" AND age = 25 LIMIT 1;
	fmt.Println(user)

	// Raw SQL 将结果扫描到另一个结构中
	db.Raw("SELECT * FROM user WHERE age = ?", 29).Scan(&user)
	fmt.Println("age=29", user)

	// Map
	db.Table("user").Where(map[string]interface{}{"nick_name": "test8", "age": 26}).Find(&users)
	// SELECT * FROM user WHERE nick_name = "test8" AND age = 26;
	fmt.Println(users)

	// 主键的Slice
	db.Table("user").Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM user WHERE id IN (20, 21, 22);
	fmt.Println(users)

	db.Table("user").Where("nick_name <> ?", "test3").Where("age >= ?", 20).Find(&users)
	//// SELECT * FROM user WHERE nick_name <> 'test3' AND age >= 20 ;
	fmt.Println(users)

	db.Table("user").Where("nick_name = ?", "test6").Or("age = ?", "24").Not("nick_name = ?", "test7").Find(&users)
	fmt.Println(users)
}
