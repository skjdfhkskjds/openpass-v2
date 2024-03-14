package user

type User struct {
	username string

	paswordHash []byte
}

type requestType struct {
	Username string
	MasterPassword string
}

func NewUser(req requestType) *User {
	user := &User{
		username: req.Username,
		paswordHash: hashPassword(req.MasterPassword),
	}

	db.SetUserData(user)

	return user
}

// 
func (db *DB) SetUserData(user *User) {
	db.Write([]byte(user.username), user.paswordHash)

	db.Get([]byte(user.username))
}

