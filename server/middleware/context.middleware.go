//
//  context.middleware.go
//  middleware
//
//  Created by d-exclaimation on 8:11 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package middleware

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

const echoContext = "EchoContextKey"

// Echo Context Middleware
func EchoContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		newContext := context.WithValue(ctx.Request().Context(), echoContext, ctx)
		ctx.SetRequest(ctx.Request().WithContext(newContext))
		return next(ctx)
	}
}

// Extract Context
func EchoFromContext(ctx context.Context) (echo.Context, error) {
	con := ctx.Value(echoContext)
	if con == nil {
		err := fmt.Errorf("could not retrieve context")
		return nil, err
	}

	ec, ok := con.(echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context failed to be casted")
		return nil, err
	}
	return ec, nil
}