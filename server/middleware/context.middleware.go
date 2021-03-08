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
	"github.com/gin-gonic/gin"
)

// Gin Context Middleware
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}


// Function to extract context from middleware
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
