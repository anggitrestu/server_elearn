package main

import (
	"log"
	"server_elearn/auth"
	"server_elearn/handler"
	"server_elearn/models/chapters"
	"server_elearn/models/courses"
	imagecourses "server_elearn/models/image_courses"
	"server_elearn/models/lessons"
	"server_elearn/models/mentors"
	"server_elearn/models/mycourses"
	"server_elearn/models/orders"
	paymentlogs "server_elearn/models/payment_logs"
	"server_elearn/models/reviews"
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

	db.AutoMigrate(&users.User{}, &mentors.Mentor{}, &lessons.Lesson{}, &mentors.Mentor{}, &courses.Course{},  &chapters.Chapter{}, &lessons.Lesson{}, mycourses.MyCourse{}, &imagecourses.ImageCourse{}, &reviews.Review{}, &orders.Order{}, &paymentlogs.PaymentLog{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	mentorRepository := repository.NewMentorRepository(db)
	courseRepository := repository.NewCourseRepository(db)
	chapterRepository := repository.NewChapterRepository(db)
	lessonRepository := repository.NewLessonRepository(db)
	imageCourseRepository := repository.NewImageCourseRepository(db)

	userService := service.NewServiceUser(userRepository)
	authService := auth.NewService()
	authMiddleware := auth.AuthMiddleware(authService, userService)
	mentorService := service.NewServiceMentor(mentorRepository)
	courseService := service.NewServiceCourse(courseRepository)
	chapterService := service.NewServiceChapter(chapterRepository)
	lessonService := service.NewServiceLesson(lessonRepository)
	imageCourseService := service.NewServiceImageCourse(imageCourseRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	mentorHandler := handler.NewMentorHandler(mentorService)
	courseHandler := handler.NewCourseHandler(courseService, mentorService)
	chapterHandler := handler.NewChapterHandler(chapterService, courseService)
	lessonHandler := handler.NewLessonHandler(lessonService)
	imageCourseHandler := handler.NewImageCourseHandler(imageCourseService, courseService)

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

	api.POST("/chapters", chapterHandler.CreateChapter)
	api.GET("/chapters/:id", chapterHandler.GetChapter)
	api.GET("/chapters", chapterHandler.GetChapters)
	api.PUT("/chapters/:id", chapterHandler.UpdateChapter)
	api.DELETE("/chapters/:id", chapterHandler.DeleteChapter)

	api.POST("/lessons", lessonHandler.CreateLesson)
	api.GET("/lessons/:id", lessonHandler.GetLesson)
	api.GET("/lessons", lessonHandler.GetLessons)
	api.PUT("/lessons/:id", lessonHandler.UpdateLesson)
	api.DELETE("/lessons/:id", lessonHandler.DeleteLesson)

	api.POST("/image-courses", imageCourseHandler.CreateImageCourse)
	api.DELETE("/image-courses/:id", imageCourseHandler.DeleteImageCourse)

	
	
	router.Run()

}

