package user

type User struct {
	Username 	string
	Email	 	string
}

var Users []string

func Create(username string) string {
	newUser := User{
		Username: username,
	}
	return newUser.Username
}
