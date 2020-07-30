package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/MgHtinLynn/golang-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/argon2"
	"net/http"
)

type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

// GET /books
// Get all books
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.CreateUserInput
	fmt.Sprintln(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config := &PasswordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	password, err := GeneratePassword(config, input.Password)

	if err != nil {
		panic(err)
	}

	// Create Book
	user := models.User{Name: input.Name, Password: password, Email: input.Email, Phone: input.Phone, Address: input.Address}
	fmt.Sprintln(user)
	status := db.Create(&user)
	fmt.Sprintln(status)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GeneratePassword is used to generate a new password hash for storing and
// comparing at a later date.
func GeneratePassword(c *PasswordConfig, password string) (string, error) {

	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
	return full, nil
}

// GET /user/:id
// Find a user
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /user/:id
// Update a user
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not found!"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
