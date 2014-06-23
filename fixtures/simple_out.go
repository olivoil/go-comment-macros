package user

type User struct {
	Name string
}

func (u *User) Save() *User {
	// normal comment
	fmt.Println("start saving user")
	metrics.Start("saving user")
	Db.save(u)
	fmt.Println("end saving user")
	metrics.End("saving user")
	fmt.Println("log something else")
}
