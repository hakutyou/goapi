package Account

import "github.com/dgrijalva/jwt-go"

var (
	JwtCfg struct {
		JwtSecret []byte `json:"JWT_SECRET"`
	}
)

type Claims struct {
	UserID uint `json:"userid"`
	jwt.StandardClaims
}
