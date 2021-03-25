//
//  services.go
//  services
//
//  Created by d-exclaimation on 5:28 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package errors

import "fmt"

type ServiceError struct {
	Type 	 uint
	Response string
}

func NewServiceError(httpStatus uint, message string) *ServiceError {
	return &ServiceError{
		Type:     httpStatus,
		Response: message,
	}
}

func (err *ServiceError) Error() string {
	return fmt.Sprintf("[%d]: %s", err.Type, err.Response)
}
