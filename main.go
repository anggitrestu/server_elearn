package main

import (
	"log"
	"net/http"
	"server_elearn/auth"
	"server_elearn/handler"
	"server_elearn/helper"
	"server_elearn/models/mentors"
	"server_elearn/models/users"
	"server_elearn/repository"
	"server_elearn/service"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/server_elearn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&users.User{}, &mentors.Mentor{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	mentorRepository := repository.NewMentorRepository(db)
	
	userService := service.NewServiceUser(userRepository)
	authService := auth.NewService()
	mentorService := service.NewServiceMentor(mentorRepository)
	
	userHandler := handler.NewUserHandler(userService, authService)
	mentorHandler := handler.NewMentorHandler(mentorService)


	router := gin.Default()

	api := router.Group("/api/v1")
	
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars",authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.POST("/mentors", mentorHandler.AddMentor)
	api.GET("/mentors/:id", mentorHandler.GetMentor)
	api.GET("/mentors", mentorHandler.GetListMentor)
	api.PUT("/mentors/:id", mentorHandler.UpdateMentor)
	api.DELETE("/mentors/:id", mentorHandler.DeleteMentor)


	
	router.Run()

}





func authMiddleware(authService auth.Service, userService service.ServiceUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized1", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err!=nil {
			response := helper.APIResponse("Unauthorized2", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid  {
			response := helper.APIResponse("Unauthorized3", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err!=nil {
			response := helper.APIResponse("Unauthorized4", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}
