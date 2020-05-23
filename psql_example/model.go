package psql_example

type User struct {
	Id       int32 `gorm:"primary_key"`
	Uuid     string
	NickName string
	Email    string
	Password string
	Age      int32
}
