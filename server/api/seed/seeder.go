package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mertture/FoodFast/api/models"
)

var users = []models.User{
	models.User{
		Username: "Steven victor",
		UserType: "Owner",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Username: "Martin Luther",
		UserType: "Owner",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var foods = []models.Food{
	models.Food{
		Name: "Big Mac",
		Type: "Burger",
		Price: 8.99,
	},
	models.Food{
		Name: "Philadelphia Roll",
		Type: "Sushi",
		Price: 11.99,
	},
}


var comments = []models.Comment {
	models.Comment {
		UserID: 1,
		RestaurantID: 1,
		Rate: 4,
		CommentText: "It was okay.",
	},
	models.Comment {
		UserID: 2,
		RestaurantID: 2,
		Rate: 3,
		CommentText: "Sushi was spicy.",
	},
}

var orders = []models.Order {
	models.Order {
		OrderedID: 1,
	},
	models.Order {
		OrderedID: 2,
	},
}

var menus = []models.Menu {
	models.Menu {
		FoodID: 1,
		RestaurantID: 1,
	},
	models.Menu {
		FoodID: 2,
		RestaurantID: 2,
	},
}

var restaurants = []models.Restaurant{
	models.Restaurant{
		Name:   "McDonalds",
		Kitchen: "Burger",
		Address: "New York",
		OpensCloses: "11:30 AM - 02:00 AM",
	},
	models.Restaurant{
		Name:   "Sushi&Spice",
		Kitchen: "Sushi",
		Address: "Istanbul",
		OpensCloses: "11:30 AM - 02:00 AM",
	},
}


func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Comment{}, &models.Order{}, &models.Menu{}, &models.Restaurant{}, &models.User{}, &models.Food{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Restaurant{}, &models.Menu{}, &models.Order{}, &models.Comment{}, &models.Food{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Restaurant{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Menu{}).AddForeignKey("restaurant_id", "restaurants(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Menu{}).AddForeignKey("food_id", "foods(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Order{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Order{}).AddForeignKey("restaurant_id", "restaurants(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Order{}).AddForeignKey("ordered_id", "foods(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Comment{}).AddForeignKey("restaurant_id", "restaurants(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Comment{}).AddForeignKey("order_id", "orders(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}


	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		restaurants[i].UserID = users[i].ID

		err = db.Debug().Model(&models.Restaurant{}).Create(&restaurants[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}