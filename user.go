package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"-"`
	Username string `json:"name"`
	Password string `json:"password"`
}
