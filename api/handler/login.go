package handler

import (
	"api_gateway/generated/user"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Maxfiy kalit
var jwtKey = []byte("ecoTrack3050720241605ecoTrack3")

// Foydalanuvchi ma'lumotlari
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT da'volar (claims)
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}


// JWT token yaratish funksiyasi
func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// JWT tokenni tekshirish middleware
func(h *Handler) ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			c.Abort()
			return
		}
		tokenStr := cookie.Value
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Login handler
func(h *Handler) Login(c *gin.Context) {
	u := &user.LoginUser{}
	if err := c.BindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	fmt.Println(u)
	// Aslida foydalanuvchi ma'lumotlari ma'lumotlar bazasida tekshiriladi
	resp, err := h.UserService.Login(c, u)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if resp.Password != u.Password || resp.Email != u.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tokenString, err := GenerateJWT(u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}



	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(5 * time.Minute),
	})

	c.JSON(http.StatusOK, gin.H{"Username": resp.Username, "Email": resp.Email, "Token": tokenString})
}