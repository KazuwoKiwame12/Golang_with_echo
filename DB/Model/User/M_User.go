package user

import db "github.com/KazuwoKiwame12/test-echo-with-postgres/DB"

//User ...table: Usersのモデル
type User struct {
	ID   int
	Name string
}

//Create ...Usersモデルの保存
func Create(name string) *User {
	db := db.Connect()
	defer db.Close()

	user := User{Name: name}
	db.Create(&user)

	return &user
}

//Update ...Usersモデルの更新
func Update(id int, newName string) *User {
	db := db.Connect()
	defer db.Close()

	user := User{}
	user.ID = id
	db.Model(&user).Update("name", newName)

	return &user
}

//GetUser ...Usersモデルの取得
func Get(id int) *User {
	db := db.Connect()
	defer db.Close()

	user := User{}
	db.First(&user, id)

	return &user
}
