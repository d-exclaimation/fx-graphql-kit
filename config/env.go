//
//  env.go
//  config
//
//  Created by d-exclaimation on 7:19 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import "os"

func GetDatabaseURL() string {
	dbURL := os.Getenv("DATABASE_URL")
	if len(dbURL) < 1 {
		dbURL = "postgres://127.0.0.1:5432/fxkit?sslmode=disable"
	}
	return dbURL
}

const defaultPort = "4000"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}