package controllers

import (
	"net/http"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/gin-gonic/gin"
)

// ValidateToken checks if the token is provided and if it matches the stored token.
func ValidateToken(ctx *gin.Context, appData app.App, service string) (string, bool) {
	// Retrieve the token from the query string (e.g., ?token=<value>)
	token := ctx.DefaultQuery("token", "")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return "", false
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
