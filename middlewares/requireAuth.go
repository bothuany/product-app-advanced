package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"product-app/controller/response"
	"product-app/domain"
	"product-app/initializers"
	"time"
)

func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("Authorization")
			if err != nil || cookie.Value == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: No Authorization cookie provided")
			}

			token, jwtErr := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Invalid token")
				}

				return []byte(os.Getenv("SECRET")), nil
			})

			if jwtErr != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Invalid token")
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if time.Now().Unix() > int64(claims["exp"].(float64)) {
					return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Token expired")
				}

				var user domain.User
				initializers.DB.First(&user, uint(claims["sub"].(float64)))

				if user.ID == 0 {
					return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Invalid token")
				}

				var userResponse response.GetUserByIdResponse
				userResponse = userResponse.ToResponse(&user)
				c.Set("user", userResponse)
			}

			return next(c)
		}
	}
}
