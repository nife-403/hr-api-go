package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWTKey = []byte("e6a0ef84c39aad457edd6ef927922ad9") //128bit, could make 256 but eh idc

func GenerateToken() (string, error) {
	//c := jwt.MapClaims{
	//	"exp": time.Now().Add(time.Hour * 24 * 7 * 4).Unix(), // 1 hour expiry
	//}
	c := jwt.MapClaims{}
	c["iat"] = time.Now().Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	new, err := t.SignedString(JWTKey)
	if err != nil {
		return "", err
	}
	err = VerifyToken(new)
	if err != nil {
		return new, err
	}
	return new, nil
}

func VerifyToken(ts string) error {
	//fmt.Println("(VerifyToken()): " + ts)
	claims := jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(ts, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("(verifyToken()) error in parsing token/unexpected signing method with token %s", ts)
		}
		return JWTKey, nil
	})
	if err != nil {
		return err
	}
	if t.Valid {
		return nil
	}
	return fmt.Errorf("invalid token")
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Authorization header missing or malformed"})
			c.Abort()
			return
		}
		t := strings.TrimPrefix(authHeader, "Bearer ")
		err := VerifyToken(t)
		if err != nil {
			c.SecureJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "message": err})
			c.Abort()
			return
		}
		c.Next()
	}
}
