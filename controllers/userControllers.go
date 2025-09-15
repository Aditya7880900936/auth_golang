package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Aditya7880900936/auth_golang.git/database"
	helper "github.com/Aditya7880900936/auth_golang.git/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client , "user")

var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
       userId := c.Param("user_id")
	   
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return 
	   } 
	}
}
