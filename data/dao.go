package data

import (
	"sync"

	err "github.com/badoll/go-chat/error"
	"github.com/badoll/go-chat/user"
)

//Dao 数据源
type Dao struct {
	Users map[string]*user.User
	mutex sync.RWMutex
}

//Init 初始化数据源
func (d *Dao) Init() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.Users = make(map[string]*user.User)
}

//AddUser 添加用户
func (d *Dao) AddUser(u user.User) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if _, ok := d.Users[u.Name]; ok {
		return err.ErrResp(err.UserExisted)
	}
	d.Users[u.Name] = &u
	return nil
}

//CheckPasswd 验证用户密码
func (d *Dao) CheckPasswd(u user.User, pwd string) (bool, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	if v, ok := d.Users[u.Name]; ok {
		return v.Passwd == u.Passwd, nil
	}
	return false, err.ErrResp(err.UserNotExist)
}

//ChangeUserState 更改用户状态
func (d *Dao) ChangeUserState(u user.User, state int) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if _, ok := d.Users[u.Name]; !ok {
		return err.ErrResp(err.UserNotExist)
	}
	d.Users[u.Name].State = state
	return nil
}

//GetOnlineUsers 获取在线用户
func (d *Dao) GetOnlineUsers() []user.User {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	users := make([]user.User, 0)
	for _, v := range d.Users {
		if v.State == user.Online {
			users = append(users, *v)
		}
	}
	return users
}
