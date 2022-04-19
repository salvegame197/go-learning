package controllers

import (
	"learning/models"
	"learning/services"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	secretkey string = "secretkeyjwt"
)

func GetUser(c *gin.Context) {

	var user []models.User
	result := dbmap.Limit(200).Find(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, user)
}

func GetUserADetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.Find(&user, id)
	if err.Error != nil {
		c.JSON(500, gin.H{"error": err.Error.Error()})
		return
	}
	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	password := user.Password
	email := user.Email
	err := dbmap.Where(map[string]interface{}{"email": user.Email}).First(&user)
	verify := CheckPasswordHash(password, user.Password)

	if err.Error != nil || !verify {
		c.JSON(500, gin.H{"error": err.Error.Error()})
		return
	}
	validToken := services.JWTAuthService().GenerateToken(email, true)
	if validToken == "" {
		c.JSON(500, gin.H{"error": err.Error.Error()})
		return
	}
	c.JSON(200, validToken)
}

// func GenerateJWT(email string) (string, error) {
// 	var mySigningKey = []byte(secretkey)
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["email"] = email
// 	claims["exp"] = time.Now().Add(time.Minute * 90).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}
// 	return tokenString, nil
// }

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
