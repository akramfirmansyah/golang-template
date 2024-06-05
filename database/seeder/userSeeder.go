package seeder

import (
	"errors"
	"log"

	"github.com/akramfirmansyah/golang-template/database"
	"github.com/akramfirmansyah/golang-template/database/model"
	"github.com/akramfirmansyah/golang-template/utils"
	"gorm.io/gorm"
)

func UserSeeder() {
	var user model.User

	if err := database.DB.Where("username = ?", "superAdmin").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pass := "admin"

			hash, err := utils.HashPassword(pass)
			if err != nil {
				log.Fatal(err)
			}

			newUser := model.User{
				Username: "superAdmin",
				Password: hash,
			}

			database.DB.Create(&newUser)
		}
	}
}
