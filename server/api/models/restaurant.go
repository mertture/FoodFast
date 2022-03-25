package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)


type Order struct {} 

type Comment struct {
	ID				uint64		`gorm:"primary_key;auto_increment" json:"id"`
	Username		string		`gorm:"size:255;not null;" json:"username"`
	RestaurantID  	uint64	`gorm:"not null"; json:"restaurantID"`
	Rate			uint32		`gorm:size:10; not null; json:"rate"`
	CreatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}


type Restaurant struct {
	ID        		uint64      	`gorm:"primary_key;auto_increment" json:"id"`
	Name     		string      	`gorm:"size:255;not null;unique" json:"title"`
	Kitchen  		[]string    	`gorm:"not null;" json:"kitchen"`
	Address   		string	    	`gorm:"size:400;not null;" json:"address"`
	OpensCloses		string	    	`gorm:"size:50;not null;"	 json:"openscloses"`
	Comments		[]Comment		`json:"comments"`
	Orders			[]Order			`json:"orders"`
	CreatedAt 		time.Time   	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time   	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

