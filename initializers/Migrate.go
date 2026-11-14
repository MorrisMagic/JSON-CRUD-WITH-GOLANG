package initializers

import (
	"crud/models"
)

func Migrate() {
	DB.AutoMigrate(&models.Post{})
}
