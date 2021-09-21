package main

import (
	"log"
	"server_elearn/auth"
	"server_elearn/handler"
	"server_elearn/models/chapters"
	"server_elearn/models/courses"
	"server_elearn/models/lessons"
	"server_elearn/models/mentors"
	"server_elearn/models/users"
	"server_elearn/repository"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/server_elearn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&users.User{}, &mentors.Mentor{}, &lessons.Lesson{}, &mentors.Mentor{}, &courses.Course{},  &chapters.Chapter{}, &lessons.Lesson{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	mentorRepository := repository.NewMentorRepository(db)
	courseRepository := repository.NewCourseRepository(db)

	userService := service.NewServiceUser(userRepository)
	authService := auth.NewService()
	authMiddleware := auth.AuthMiddleware(authService, userService)
	mentorService := service.NewServiceMentor(mentorRepository)
	courseService := service.NewServiceCourse(courseRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	mentorHandler := handler.NewMentorHandler(mentorService)
	courseHandler := handler.NewCourseHandler(courseService, mentorService)

	router := gin.Default()

	api := router.Group("/api/v1")
	
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars",authMiddleware, userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware, userHandler.FetchUser)

	api.POST("/mentors", authMiddleware, mentorHandler.AddMentor)
	api.GET("/mentors/:id", mentorHandler.GetMentor)
	api.GET("/mentors", authMiddleware, mentorHandler.GetListMentor)
	api.PUT("/mentors/:id",authMiddleware, mentorHandler.UpdateMentor)
	api.DELETE("/mentors/:id",authMiddleware, mentorHandler.DeleteMentor)

	api.POST("/courses",authMiddleware, courseHandler.CreateCourse)
	api.GET("/courses/:id", courseHandler.GetCourse)
	api.GET("/courses", courseHandler.GetCourses)
	api.PUT("/courses/:id",authMiddleware, courseHandler.UpdateCourse)
	api.DELETE("/courses/:id", authMiddleware, courseHandler.DeleteCourse)

	
	router.Run()

}

