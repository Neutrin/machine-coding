package user

type User struct {
	name     string
	password string
	mark     rune
}

func (u *User) GetMark() rune {
	return u.mark
}

func GetUser(name string, password string, mark rune) User {
	return User{
		name:     name,
		password: password,
		mark:     mark,
	}
}

func (u *User) GetName() string {
	return u.name
}
