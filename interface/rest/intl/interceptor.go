package intl

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (r *Routes) Auth() echo.MiddlewareFunc {

	// set jwt
	JWTCustomConfig := mid.JWTConfig{
		SigningKey: []byte(r.Config.Jwt.Extl.Secret),
	}
	r.HTTPHandler.Framework.Use(mid.JWTWithConfig(JWTCustomConfig))

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			u := c.Get("user").(*jwt.Token)
			claims := u.Claims.(jwt.MapClaims)
			id := claims["id"]
			ClientID := r.Config.Jwt.Extl.ClientID
			fmt.Println(id)
			fmt.Println(ClientID)
			if id != ClientID {
				return echo.NewHTTPError(http.StatusForbidden, "Client ID unauthorize!")
			}

			c.Set("ClientID", id)

			return next(c)
		}
	}

}
