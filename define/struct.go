package define

type User struct {
    Id       int32 `gorm:"primary_key"`
    Uuid     string
    NickName string
    Email    string
    Password string
    Age      int32
}

