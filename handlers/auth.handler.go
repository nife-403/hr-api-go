package handlers

import (
	"fmt"
	"net/http"
	"reflect"

	"hr-api/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Body struct {
	Token string `json:"token" binding:"required"`
}

func CheckTokenHandler(c *gin.Context) {
	var ts Body
	if err := c.BindJSON(&ts); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if ts.Token == "" {
		c.SecureJSON(http.StatusBadRequest, gin.H{"message": "token not found"})
		return
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(ts.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JWTKey, nil
	})

	// ... error handling
	if err != nil {
		c.SecureJSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}
	if token.Valid {
		c.SecureJSON(http.StatusOK, gin.H{"message": "token is valid"})
	}
	c.SecureJSON(http.StatusOK, gin.H{"claims": claims})
}

func GenerateTokenHandler(c *gin.Context) {
	//ahh := c.GetHeader("Authorization")
	//if ahh == "" {
	//	c.SecureJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
	//	return
	//}
	//ts := strings.TrimPrefix(ahh, "Bearer ")
	t, err := auth.GenerateToken()
	if err != nil {
		c.SecureJSON(http.StatusInternalServerError, gin.H{"error": err})
		fmt.Println(reflect.ValueOf(&c).Elem().IsZero(), err)
		c.Abort()
		return
	}
	c.SecureJSON(http.StatusCreated, gin.H{"new_token": t})
	//t, _ := jwt.Parse(ts, func(t *jwt.Token) (any, error) { return auth.JWTKey, nil })
	//if t.Valid {
	//	c.SecureJSON(http.StatusAccepted, gin.H{"new_token": t})
	//} else {
	//	c.SecureJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//}
}
