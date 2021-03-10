//
//  connection.go
//  db
//
//  Created by d-exclaimation on 7:18 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package db

import (
	"github.com/d-exclaimation/fx-graphql-kit/config"
	"github.com/d-exclaimation/fx-graphql-kit/db/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Fx Provider
func PostgresProvider() *gorm.DB {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDatabaseURL(),
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		log.Fatalln(err)
	}

	if err = conn.AutoMigrate(&entities.User{}, &entities.Thought{}); err != nil {
		log.Fatalln(err)
	}
	return conn
}
