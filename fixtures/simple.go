package user

type User struct {
	Name string
}

func (u *User) Save() *User {
	// normal comment
	//: start saving user
	Db.save(u)
	//: end saving user
	//: log something else
}
