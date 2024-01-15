package security

import (
	"time"

	"github.com/go-chi/jwtauth"
)

type GetJwtInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtHelper struct {
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewJwtHelper(jwt *jwtauth.JWTAuth, jwtExpiresIn int) *JwtHelper {
	return &JwtHelper{
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (helper *JwtHelper) GenerateJwt(id string) (string, error) {
	_, tokenString, err := helper.Jwt.Encode(map[string]interface{}{
		"sub": id,
		"exp": time.Now().Add(time.Second * time.Duration(helper.JwtExpiresIn)).Unix(),
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
