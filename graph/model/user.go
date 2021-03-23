//
//  user.go
//  model
//
//  Created by d-exclaimation on 6:53 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
