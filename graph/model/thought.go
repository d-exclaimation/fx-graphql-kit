//
//  thought.go
//  model
//
//  Created by d-exclaimation on 7:10 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

type Thought struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageURL *string `json:"imageURL"`
	UserID   string  `json:"user"`
}
