package controllers

import (
	"net/http"
	"os"
	"time"
	"todo_api/initializers"
	"todo_api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){
 // Get the email/pass off req Body
	var body struct{
		FirstName string
		LastName string	
		Username string	
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body",})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "Failed to hash password.",})
		return
	}

	// Create the user
	user := models.User{
		FirstName: body.FirstName,
		LastName: body.LastName,
		Username: body.Username,
		Email: body.Email, 
		Password: string(hash),
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "Failed to create user.",})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{"message": "Sign in successfull"})
}

func Login(c *gin.Context){
	// Get email & pass from  body
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Look for requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare sent in password with saved users password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password",})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
 
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token",})
		return
	}
	
	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login in successfull"})
}

func Validate(c *gin.Context){
 	user,_ := c.Get("user")
	// user.(models.User).Email    -->   to access specific data
	c.JSON(http.StatusOK, gin.H{ "message": user,})
}

func UpdateUser(c *gin.Context) {
    // Find user using JWT token
    userID := c.MustGet("userID").(uint)

    var body struct{
		FirstName string
		LastName string	
		Username string	
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body",})
		return
	}

    // Find the existing user by ID
    var user models.User
    result := initializers.DB.First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Update user user 
    user.FirstName = body.FirstName
    user.LastName = body.LastName
    user.Username = body.Username

    // Save changes to the database
    initializers.DB.Save(&user)

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
    // Find user using JWT token
    userID := c.MustGet("userID").(uint)

    // Find the existing user by ID
    var user models.User
    result := initializers.DB.First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Verify that the user owns this account
    if user.ID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
        return
    }

    // Delete the user from the database
    initializers.DB.Delete(&user)


    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}