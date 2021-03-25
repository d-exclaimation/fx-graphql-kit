//
//  logger.middleware.go
//  middleware
//
//  Created by d-exclaimation on 9:57 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package middleware

import (
	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

func EndpointLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return em.LoggerWithConfig(em.LoggerConfig{
		Format: "${time_rfc3339_nano} |" +
			color.NewRGBStyle(color.RGB(200, 200, 200), color.HEX("#20bcaf", true)).
			Sprint(" ${status} ${method} ") + "| ${latency_human} | >> ${uri}\n",
	})(next)
}
