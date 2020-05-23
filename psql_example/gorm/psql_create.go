package main

import (
	"bytes"
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

	// 单个创建插入
	user := psql_example.User{
		Uuid:     "uidtest01",
		NickName: "test01",
		Email:    "test01@qq.com",
		Password: "pwd01",
		Age:      18,
	}
	db.Table("user").Create(&user)
	// 若NewRecord插入成功返回false，失败返回true
	if !db.Table("user").NewRecord(&user) {
		fmt.Println("插入user到数据库成功，user=", user)
	} else {
		fmt.Println("插入user到数据库失败！，user=", user)
	}

	db.Table("user").Save(&user)

	users := make([]*psql_example.User, 0)
	for i := 0; i < 14; i++ {
		user := psql_example.User{
			Uuid:     fmt.Sprintf("uuid%v", i),
			NickName: fmt.Sprintf("test%v", i),
			Email:    fmt.Sprintf("test%v@qq.com", i),
			Password: fmt.Sprintf("pwd%v", i),
			Age:      int32(18 + i),
		}

		users = append(users, &user)
	}

	// 批量插入
	_ = BatchSave(db, users)
}

// BatchSave 批量插入数据
func BatchSave(db *gorm.DB, users []*psql_example.User) error {
	var buffer bytes.Buffer
	sql := "insert into User (uuid,nick_name,password,email,age) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, user := range users {
		if i == len(users)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s',%d);", user.Uuid, user.NickName, user.Password, user.Email, user.Age))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s',%d); %s", user.Uuid, user.NickName, user.Password, user.Email, user.Age, sql))
		}
		fmt.Println(buffer.String())
	}
	return db.Exec(buffer.String()).Error
}
