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

	var userTmp psql_example.User
	// First查询第一条数据，按照主键排序
	// SELECT * FROM user ORDER BY id LIMIT 1;
	result := db.Table("user").First(&userTmp)
	if err := result.Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(userTmp)

	var userTmp2 psql_example.User
	// Last查询最后一条数据，按照主键排序
	// SELECT * FROM user ORDER BY id DESC LIMIT 1;
	result = db.Table("user").Last(&userTmp2)
	if err := result.Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(userTmp2)

	var userTmp3 psql_example.User
	// First使用主键获取记录
	// SELECT * FROM user WHERE id = 10;
	result = db.Table("user").First(&userTmp3, 10)
	if err := result.Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(userTmp3)

	var userList []psql_example.User
	// 查找所有记录
	// SELECT * FROM user;
	result = db.Table("user").Find(&userList)
	if err := result.Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(userList)

}
