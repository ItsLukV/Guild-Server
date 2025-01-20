package middleware

import (
	"net/http"
	"os"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(appData *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the query string (e.g., ?token=<value>)
		service := c.Param("service")

		// Use ValidateToken to check if the token is valid
		token, valid := ValidateToken(c, appData, service)
		if !valid {
			c.Abort() // Stop processing the request if validation fails
			return
		}

		c.Set("validatedToken", token)
		c.Next()
	}
}

func ValidateToken(ctx *gin.Context, appData *app.App, service string) (string, bool) {

	defaultAdminToken := os.Getenv("DEFAULT_ADMIN_TOKEN")

	// Retrieve the token from the query string (e.g., ?token=<value>)
	token := ctx.DefaultQuery("token", "")

	// If the token is empty, return an error
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return "", false
	}

	// Check if the token is the default admin token
	if token == defaultAdminToken {
		// You can skip the database check or apply additional logic for admin access
		return token, true
	}

	// Retrieve the stored token for the given service
	storedToken, err := appData.GetToken(service)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Token not found"})
		return "", false
	}

	// Check if the query token matches the stored token
	if token != storedToken {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return "", false
	}

	return token, true
}
