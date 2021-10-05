package auth

import (
	"net/http"
	"server_elearn/helper"
	"server_elearn/models/users"
	"strings"

	"github.com/gin-gonic/gin"
)

type Role struct {
	Roles []string
}

func Permission(r *Role) gin.HandlerFunc {
	return func(c *gin.Context) {

		currentUser := c.MustGet("currentUser").(users.User)
		currentRole := currentUser.Role
		success := checkRole(r.Roles, currentRole)
		if !success {
			response := helper.APIResponse("You dont have permission", http.StatusMethodNotAllowed, "error", nil)
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, response)
			return
		}

	}
}


func checkRole(roless []string, currentRole string)(bool){	
	role := strings.Join(roless[:], " ")
    return strings.Contains(role, currentRole) 
}