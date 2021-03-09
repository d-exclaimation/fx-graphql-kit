//
//  services.go
//  services
//
//  Created by d-exclaimation on 5:28 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package errors


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
