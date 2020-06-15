package auth

type User struct {
	ID       int
	Name     string   `binding:"required" form:"name" json:"name"  example:"hakutyou"`
	Password password `binding:"required" form:"password" json:"password"  example:"myPassword"`
}

type UserToken struct {
	Token string `binding:"required" form:"token"`
}

type password string

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`"x"`), nil
}
