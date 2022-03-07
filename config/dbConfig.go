package config

import (
	"fmt"
	_ "github.com/lib/pq"
)

const (
	PostgresDriver = "postgres"
	User           = "postgres"
	Host           = "localhost"
	Password       = "postgres"
	Database       = "postgres"
	Port           = "5432"
)

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Database)
