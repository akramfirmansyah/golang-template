package migration

import (
	"fmt"

	"github.com/akramfirmansyah/golang-template/database"
	"github.com/akramfirmansyah/golang-template/database/model"
)

func MigrateModel() {
	_ = database.DB.AutoMigrate(&model.User{})
	fmt.Println("Success migrate database!")
}
