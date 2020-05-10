package user

//User 用户
type User struct {
	Name   string
	Passwd string
	State  int
}

const (
	//Online 在线
	Online = iota
	//Offline 离线
	Offline
)
