package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"student-management/internal/models"
)

func AuthorizationMiddleware(c *gin.Context) {
	tokenStr, err := getFromHeader(c, "Authorization")
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	claims, err := parseToken(tokenStr)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	handlePermissions(c, claims)

	c.Set(`claims`, claims)

	c.Next()

}

func handlePermissions(c *gin.Context, claims *models.UserClaims) {

	requestURI := strings.Split(c.Request.RequestURI, "?")[0]
	if requestURI[:6] == "/admin" && claims.RoleID == 3 {
		return
	}
	if requestURI == "/api/v1/timetable" && (claims.RoleID == 1 || claims.RoleID == 2) {
		return
	}
	if requestURI == "/api/v1/students" && (claims.RoleID == 2) {
		return
	}
	if requestURI == "/api/v1/check-in" && (claims.RoleID == 2) {
		return
	}

	errorResponse(c, http.StatusForbidden, map[string]interface{}{
		"status":  "error",
		"message": `you have no permission to view resource`,
	})

}

func parseToken(tokenStr string) (*models.UserClaims, error) {
	claims := &models.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("bad request")
	}
	return claims, nil
}

func getFromHeader(c *gin.Context, headerName string) (string, error) {
	value := c.GetHeader(headerName)
	if value == "" {
		return "", errors.New(headerName + " not found in header")
	}
	return value, nil
}

func errorResponse(c *gin.Context, statusCode int, resp interface{}) {
	c.AbortWithStatusJSON(statusCode, resp)
}
