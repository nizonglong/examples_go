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

	/**
	 * 如果只想更新更改的字段，可以使用Update, Updates
	 */
	var user = psql_example.User{
		Id:       22,
		NickName: "test",
		Age:      16,
	}

	// 更新单个属性（如果更改）
	fmt.Println("before:", user)
	db.Table("user").Model(&user).Update("nick_name", "test-22")
	// UPDATE user SET nick_name='test-22' WHERE id=22;
	fmt.Println("after:", user)

	// 使用组合条件更新单个属性
	db.Table("user").Model(&user).Where("age = ?", 30).Update("nick_name", "test12-md")

	// 使用`map`更新多个属性，只会更新这些更改的字段
	db.Table("user").Model(&user).Updates(map[string]interface{}{"nick_name": "test-nick", "age": 18})

	// 使用`struct`更新多个属性，只会更新这些更改的和非空白字段
	db.Table("user").Model(&user).Updates(psql_example.User{NickName: "hello", Age: 18})
	// UPDATE users SET nick_name='hello', age=18 WHERE id = 18;

	// 使用`RowsAffected`获取更新记录计数
	count := db.Model(psql_example.User{}).Updates(psql_example.User{NickName: "hello", Age: 18}).RowsAffected
	fmt.Println(count)

}
