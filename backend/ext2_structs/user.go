package file

type User struct {
	I_uid int64  // User ID
	I_gid int64  // Group ID
	Name  string // Username
	Pass  string // Password
	Grp   string // Group name
}

func NewUser() *User {
	return &User{
		I_uid: 0,
		I_gid: 0,
		Name:  "",
		Pass:  "",
		Grp:   "",
	}
}
