package main

import (
	_ "workout-manager-server/common"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbCon, err := gorm.Open("postgres", AppConfig.ConnectionString)
}
