package gorm

import (
	"examples_go/define"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // 这里很容易忘记加导包
)

func MainNot() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=example sslmode=disable password=nizonglong")
	defer db.Close()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	var user define.User
	var users []define.User
	// 下面的输出结果很可能都一样，因为每次取First都是相同的值，除非排除第一个

	db.Table("user").Not("nick_name", "test01").First(&user)
	// SELECT * FROM user WHERE nick_name <> "test6" LIMIT 1;
	fmt.Println(user)

	// Not In
	db.Table("user").Not("nick_name", []string{"test7", "test8"}).Find(&users)
	// SELECT * FROM user WHERE nick_name NOT IN ("test7", "test8");
	fmt.Println(users)

	// Not In slice of primary keys
	db.Table("user").Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM user WHERE id NOT IN (1,2,3);
	fmt.Println(user)

	db.Table("user").Not([]int64{}).First(&user)
	// SELECT * FROM user;
	fmt.Println(user)

	// Plain SQL
	db.Table("user").Not("nick_name = ?", "test9").First(&user)
	// SELECT * FROM user WHERE NOT(nick_name = "test9");
	fmt.Println(user)

	// Struct
	db.Table("user").Not(define.User{NickName: "test10"}).First(&user)
	// SELECT * FROM user WHERE nick_name <> "test10";
	fmt.Println(user)
}
