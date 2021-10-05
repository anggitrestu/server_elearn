package main

import (
	"fmt"
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
	"server_elearn/models/reviews"
	"server_elearn/models/users"
	"server_elearn/repository"
	"server_elearn/repository/drivers/mysql"
	"server_elearn/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`configs/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB){
	err := db.AutoMigrate(&users.User{}, &mentors.Mentor{}, &lessons.Lesson{}, &mentors.Mentor{}, &courses.Course{},  &chapters.Chapter{}, &lessons.Lesson{}, mycourses.MyCourse{}, &imagecourses.ImageCourse{}, &reviews.Review{}, &orders.Order{})
	if err != nil {
		panic(err)
	}
}

func main() {

	mysqlConfig := mysql.ConfigDb {
		DbUser:     viper.GetString(`databases.mysql.user`),
		DbPassword: viper.GetString(`databases.mysql.password`),
		DbHost:     viper.GetString(`databases.mysql.host`),
		DbPort:     viper.GetString(`databases.mysql.port`),
		DbName:     viper.GetString(`databases.mysql.dbname`),
	}

	db := mysqlConfig.InitialDb()
	DbMigrate(db)
	
	configJWT := viper.GetString(`jwt.SECRET_KEY`)
	fmt.Println(configJWT)
	midtrans_client_key := viper.GetString(`midtrans.MIDTRANS_CLIENT_KEY`)
	midtrans_server_key := viper.GetString(`midtrans.MIDTRANS_SERVER_KEY`)

	userRepository := repository.NewUserRepository(db)
	mentorRepository := repository.NewMentorRepository(db)
	courseRepository := repository.NewCourseRepository(db)
	chapterRepository := repository.NewChapterRepository(db)
	lessonRepository := repository.NewLessonRepository(db)
	imageCourseRepository := repository.NewImageCourseRepository(db)
	reviewRepository := repository.NewReviewRepository(db)
	myCourseRepository := repository.NewMyCourseRepository(db)
	orderRepository := repository.NewOrderRepository(db)

	userService := service.NewServiceUser(userRepository)
	authService := auth.NewService(configJWT)
	authMiddleware := auth.AuthMiddleware(authService, userService)
	canAdmin := auth.Permission(&auth.Role{Roles: []string{"admin"}})
	canAll := auth.Permission(&auth.Role{Roles: []string{"admin", "student"}})
	mentorService := service.NewServiceMentor(mentorRepository)
	courseService := service.NewServiceCourse(courseRepository)
	chapterService := service.NewServiceChapter(chapterRepository)
	lessonService := service.NewServiceLesson(lessonRepository)
	imageCourseService := service.NewServiceImageCourse(imageCourseRepository)
	reviewService := service.NewServiceReview(reviewRepository)
	myCourseService := service.NewServiceMyCourse(myCourseRepository)
	paymentService := service.NewServicePayment(midtrans_client_key, midtrans_server_key)
	orderService := service.NewServiceOrder(orderRepository, *paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	mentorHandler := handler.NewMentorHandler(mentorService)
	courseHandler := handler.NewCourseHandler(courseService, mentorService)
	chapterHandler := handler.NewChapterHandler(chapterService, courseService)
	lessonHandler := handler.NewLessonHandler(lessonService)
	imageCourseHandler := handler.NewImageCourseHandler(imageCourseService, courseService)
	reviewHandler := handler.NewReviewHandler(reviewService, courseService)
	myCourseHandler := handler.NewMyCourseHandler(myCourseService, courseService, userService, orderService)
	orderHandler := handler.NewOrderHandler(orderService, myCourseService)

	router := gin.Default()

	api := router.Group("/api/v1")
	
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", authMiddleware,  userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware, userHandler.FetchUser)

	api.POST("/mentors", authMiddleware, canAdmin, mentorHandler.AddMentor)
	api.GET("/mentors/:id",  mentorHandler.GetMentor)
	api.GET("/mentors",   mentorHandler.GetListMentor)
	api.PUT("/mentors/:id",authMiddleware, canAdmin, mentorHandler.UpdateMentor)
	api.DELETE("/mentors/:id",authMiddleware, canAdmin, mentorHandler.DeleteMentor)

	api.POST("/courses", authMiddleware, canAdmin, courseHandler.CreateCourse)
	api.GET("/courses/:id",  courseHandler.GetCourse)
	api.GET("/courses", courseHandler.GetCourses)
	api.PUT("/courses/:id",authMiddleware, canAdmin, courseHandler.UpdateCourse)
	api.DELETE("/courses/:id", authMiddleware, canAdmin, courseHandler.DeleteCourse)

	api.POST("/chapters", authMiddleware, canAdmin, chapterHandler.CreateChapter)
	api.GET("/chapters/:id", chapterHandler.GetChapter)
	api.GET("/chapters", chapterHandler.GetChapters)
	api.PUT("/chapters/:id",authMiddleware, canAdmin, chapterHandler.UpdateChapter)
	api.DELETE("/chapters/:id", authMiddleware, canAdmin, chapterHandler.DeleteChapter)

	api.POST("/lessons",authMiddleware, canAdmin, lessonHandler.CreateLesson)
	api.GET("/lessons/:id", lessonHandler.GetLesson)
	api.GET("/lessons", lessonHandler.GetLessons)
	api.PUT("/lessons/:id",authMiddleware, canAdmin, lessonHandler.UpdateLesson)
	api.DELETE("/lessons/:id",authMiddleware,canAdmin, lessonHandler.DeleteLesson)

	api.POST("/image-courses",authMiddleware, canAdmin, imageCourseHandler.CreateImageCourse)
	api.DELETE("/image-courses/:id", authMiddleware, canAdmin, imageCourseHandler.DeleteImageCourse)

	api.POST("/reviews", authMiddleware, canAll, reviewHandler.CreateReview)
	api.PUT("/reviews/:id", authMiddleware, canAll, reviewHandler.UpdateReview)
	api.DELETE("/reviews/:id", authMiddleware, canAll, reviewHandler.DeleteReview)

	api.POST("/my-courses", authMiddleware, canAll, myCourseHandler.CreateMyCourse)
	api.GET("/my-courses", authMiddleware, canAll, myCourseHandler.GetAllMyCourse)
	
	api.GET("/orders", authMiddleware, orderHandler.GetOrders)
	api.POST("/webhook", orderHandler.Webhook)


	router.Run()

}

