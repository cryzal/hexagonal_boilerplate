package rest

import (
	"fmt"
	"time"

	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/utils/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoHTTPHandler struct {
	GracefullyShutdown
	Framework *echo.Echo
}

func NewEchoHTTPHandlerDefault(cfg *config.Config) EchoHTTPHandler {
	return NewEchoHTTPHandler(fmt.Sprintf(":%d", cfg.Port))
}

func NewEchoHTTPHandler(address string) EchoHTTPHandler {
	e := echo.New()

	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: logger.MiddlewareLog,
		}),
		middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
	}))

	// Handler for putting app request and response timestamp. This used for get elapsed time
	e.Use(ServiceRequestTime)
	return EchoHTTPHandler{
		GracefullyShutdown: NewGracefullyShutdown(e, address),
		Framework:          e,
	}

}

// RunApplication is implementation of RegistryContract.RunApplication()
func (r *EchoHTTPHandler) RunApplication() {
	r.RunWithGracefullyShutdown()
}

// ServiceRequestTime middleware adds a `Server` header to the response.
func ServiceRequestTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("X-App-RequestTime", time.Now().Format(time.RFC3339))
		return next(c)
	}
}
