package v1

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = 2 * time.Hour

// 中华人民共和国万岁
var MySecret = []byte("zhrmghgws")

func (r *Controller) Test(c *gin.Context) {
	c.String(http.StatusOK, "Test")
}

func (r *Controller) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func (r *Controller) JwtEncode(c *gin.Context) {
	jwt, _ := EnCodeToken("jieqiang")
	jwtDecode, _ := DecodeToken(jwt)
	res := fmt.Sprintf("jwt(jieqiang)=%v, %v", jwt, jwtDecode)
	c.String(http.StatusOK, res)
}

func EnCodeToken(username string) (string, error) {
	c := MyClaims{
		// Username: "jieqiang",
		Username: username,
		// username,
		StandardClaims: jwt.StandardClaims{
			// jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "jwtJieqiang",
			// Audience:  "",
			// Id:        "",
			// IssuedAt:  0,
			// NotBefore: 0,
			// Subject:   "gin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

// EnCodeToken 生成JWT
// func EnCodeToken(username string) (string, error) {
// 	// 创建一个我们自己的声明
// 	c := MyClaims{
// 		username, // 自定义字段
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
// 			Issuer:    "my-project",                               // 签发人
// 		},
// 	}
// 	// 使用指定的签名方法创建签名对象
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
// 	// 使用指定的secret签名并获得完整的编码后的字符串token
// 	return token.SignedString(MySecret)
// }

func DecodeToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{
		Username:       "",
		StandardClaims: jwt.StandardClaims{},
	}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
