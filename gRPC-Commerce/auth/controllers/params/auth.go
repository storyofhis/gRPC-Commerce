package params

type Register struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
