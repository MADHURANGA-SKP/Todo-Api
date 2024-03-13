package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
	"todo_api/initializers"
	"todo_api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context){
	//get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode/validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
			})
			if err != nil {
				log.Fatal(err)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				//check the exp
				if float64(time.Now().Unix()) > claims["exp"].(float64){
					c.AbortWithStatus(http.StatusUnauthorized)
				}
				//find the user with token aub
				var user models.User
				initializers.DB.First(&user,claims["sub"])

				if user.ID == 0 {
					c.AbortWithStatus(http.StatusUnauthorized)
				}
				//
				c.Set("user", user)
				//continue
				c.Next()
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
		}
}
