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
	if strings.Split(c.Request.RequestURI, "?")[0] == "/api/v1/sign-in" {
		return
	}
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
	if requestURI == "/api/v1/timetable" && (claims.RoleID == 1 || claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}
	if requestURI == "/api/v1/students" && (claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}
	if requestURI == "/api/v1/check-in" && (claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}

	if requestURI == "/api/v1/faculties" && (claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}

	if requestURI == "/api/v1/departments" && (claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}

	if requestURI == "/api/v1/groups" && (claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}

	if requestURI == "/api/v1/lessons" && (claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}

	if requestURI == "/api/v1/types" && (claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}

	if requestURI == "/api/v1/times" && (claims.RoleID == 2 || claims.RoleID == 3) {
		return
	}

	if requestURI == "/api/v1/positions" && (claims.RoleID == 2 || claims.RoleID == 3) {
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
