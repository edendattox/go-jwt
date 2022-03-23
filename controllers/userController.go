package controllers

import (
	"github.com/edendattox/go-jwt/database"
	"github.com/go-playground/validator/v10"
	// "context"
	// "fmt"
	// "go-jwt/helpers"
	// helper "go-jwt/helpers"
	// "go-jwt/models"
	// "log"
	// "net/http"
	// "strconv"
	// "time"
	// "github.com/edendattox/go-jwt/database"
	// "github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v18"
	// "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

func Signup() {

}

func Login() {

}

func GetUsers() {

}

func GetUser() {

}
