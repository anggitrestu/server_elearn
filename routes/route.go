package routes

// import (
// 	"net/http"
// 	"server_elearn/auth"
// 	"server_elearn/helper"
// 	"server_elearn/service"
// 	"strings"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func NewRoute(userHandler handler) {
// 	// userRepository := repository.NewUserRepository(db)

// 	// userService := service.NewUserService(userRepository)
// 	// authService := auth.NewService()

// 	// userHandler := handler.NewUserHandler(userService, authService)

// 	router := gin.Default()

// 	api := router.Group("/api/v1")

// 	api.POST("/users", userHandler.RegisterUser)
// 	api.POST("/sessions", userHandler.Login)
// 	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
// 	api.POST("/avatars",authMiddleware(authService, userService), userHandler.UploadAvatar)
// 	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

// 	router.Run()
// }

// func authMiddleware(authService auth.Service, userService service.UserService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if !strings.Contains(authHeader, "Bearer") {
// 			response := helper.APIResponse("Unauthorized1", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		tokenString := ""
// 		arrayToken := strings.Split(authHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		token, err := authService.ValidateToken(tokenString)
// 		if err!=nil {
// 			response := helper.APIResponse("Unauthorized2", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		claim, ok := token.Claims.(jwt.MapClaims)

// 		if !ok || !token.Valid  {
// 			response := helper.APIResponse("Unauthorized3", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		userID := int(claim["user_id"].(float64))
// 		user, err := userService.GetUserByID(userID)
// 		if err!=nil {
// 			response := helper.APIResponse("Unauthorized4", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		c.Set("currentUser", user)

// 	}
// }
