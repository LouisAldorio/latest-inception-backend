package tool

import (
	"fmt"
	"myapp/graph/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type UserClaim struct {
	Username string
	Role     string
	jwt.StandardClaims
}

func (u *UserClaim) HasRole(role string) bool {
	fmt.Println(u)
	return strings.ToUpper(u.Role) == strings.ToUpper(role)
}

var jwtKey = []byte("secret")

func CreateToken(user model.User) (string, error) {
	var signingMethod = jwt.SigningMethodHS256
	var expiredTime = time.Now().AddDate(0, 1, 0).UnixNano() / int64(time.Millisecond)

	customClaim := UserClaim{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime,
		},
	}

	token := jwt.NewWithClaims(signingMethod, customClaim)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", gqlerror.Errorf(fmt.Sprintf("%s", err))
	}

	return signedToken, nil
}

func ValidateToken(t string) (*jwt.Token, error) {
	token, _ := jwt.ParseWithClaims(t, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return jwtKey, nil
	})

	return token, nil
}
