package seeders

import (
	"echo-demo-project/models"

	"github.com/jinzhu/gorm"
)

type UserSeeder struct {
	DB *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{DB: db}
}

func (userSeeder *UserSeeder) SetUsers() {
	users := map[int]map[string]string{
		1: {
			"email":    "user1@example.com",
			"name":     "user1",
			"password": "$2y$12$5cg0So96dfKll14wrivrp.ccl/zcbtnx9Cge4UFqkvw5H89dfpSCe ", //password1
		},
		2: {
			"email":    "user2@example.com",
			"name":     "user2",
			"password": "$2y$12$NU7se45IiQoostWNprx4Iu0nC4I5CCHDt74JMT3Lsr0FRN7wl3UZu ", //password2
		},
	}

	for key, value := range users {
		user := models.User{}
		userSeeder.DB.First(&user, key)

		if user.ID == 0 {
			user.ID = uint(key)
			user.Email = value["email"]
			user.Name = value["name"]
			user.Password = value["password"]
			userSeeder.DB.Create(&user)
		}
	}
}
