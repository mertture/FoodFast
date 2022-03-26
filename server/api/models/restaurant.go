package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)


type Order struct {
	ID				uint64		`gorm:"primary_key;auto_increment" json:"id"`
	UserID			uint64		`gorm:"primary_key;" json:"user_id"`
	Username		string		`gorm:"size:255;not null;" json:"username"`
	RestaurantID  	uint64		`gorm:"not null"; json:"restaurant_id"`
	OrderedFood		[]Food		`gorm:"not null;" json:"ordered_food"`
	OrderedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
	Price			uint64		`gorm:"not null"; json:"price"`
	PaymentType		string		`gorm:"not null;" json:"payment_type"`
} 

type Comment struct {
	ID				uint64		`gorm:"primary_key;auto_increment" json:"id"`
	UserID			uint64		`gorm:"primary_key;" json:"user_id"`
	Username		string		`gorm:"size:255;not null;" json:"username"`
	RestaurantID  	uint64		`gorm:"not null"; json:"restaurant_id"`
	OrderID			uint64		`gorm:"primary_key;auto_increment" json:"id"`
	Rate			uint32		`gorm:size:10; not null; json:"rate"`
	CreatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type Food struct {
	ID				uint64		`gorm:"primary_key;auto_increment" json:"id"`
	Name     		string      `gorm:"size:255;not null;unique" json:"name"`
	Ingredients		[]string	`gorm:"not null;" json:"ingredients"`
	RestaurantID  	uint64		`gorm:"not null"; json:"restaurant_id"`
	Type     		string      `gorm:"size:255;" json:"type"`
	Price			uint64		`gorm:"not null"; json:"price"`
}


type Restaurant struct {
	ID        		uint64      	`gorm:"primary_key;auto_increment" json:"id"`
	Name     		string      	`gorm:"size:255;not null;unique" json:"name"`
	Kitchen  		[]string    	`gorm:"not null;" json:"kitchen"`
	Menu			[]Food			`gorm:"not null;" json:"menu"`
	Address   		string	    	`gorm:"size:400;not null;" json:"address"`
	OpensCloses		string	    	`gorm:"size:50;not null;" json:"opens_closes"`
	Comments		[]Comment		`json:"comments"`
	Orders			[]Order			`json:"orders"`
	CreatedAt 		time.Time   	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time   	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}




func (o *Order) Prepare() {
	o.ID = 0
	o.UserID = 0
	o.Username = html.EscapeString(strings.TrimSpace(o.Username))
	o.RestaurantID = 0

	for i,_ := range o.OrderedFood {
		o.OrderedFood[i] = Food{}
	}
	o.OrderedAt = time.Now()
	o.Price = 0
	o.PaymentType = html.EscapeString(strings.TrimSpace(o.PaymentType))
}


func (c *Comment) Prepare() {
	c.ID = 0
	c.UserID = 0
	c.Username = html.EscapeString(strings.TrimSpace(c.Username))
	c.RestaurantID = 0
	c.OrderID = 0
	c.Rate = 0
	c.CreatedAt = time.Now()
}

func (f *Food) Prepare() {
	f.ID = 0
	f.Name = html.EscapeString(strings.TrimSpace(f.Name))
	for i,_ := range f.Ingredients {
		f.Ingredients[i] = html.EscapeString(strings.TrimSpace(f.Ingredients[i]))
	}
	f.RestaurantID = 0
	f.Type = html.EscapeString(strings.TrimSpace(f.Type))
	f.Price = 0
}

func (r *Restaurant) Prepare() {
	r.ID = 0
	r.Name = html.EscapeString(strings.TrimSpace(r.Name))

	for i,_ := range r.Kitchen {
		r.Kitchen[i] = html.EscapeString(strings.TrimSpace(r.Kitchen[i]))
	}
	for i,_ := range r.Menu {
		r.Menu[i] = Food{}
	}

	r.Address = html.EscapeString(strings.TrimSpace(r.Address))
	r.OpensCloses = html.EscapeString(strings.TrimSpace(r.OpensCloses))
	
	for i,_ := range r.Comments {
		r.Comments[i] = Comment{}
	}
	for i,_ := range r.Orders {
		r.Orders[i] = Order{}
	}

	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()

}


func (r *Restaurant) Validate() error {

	if r.Name == "" {
		return errors.New("Required Name")
	}
	if r.Address == "" {
		return errors.New("Required Address")
	}
	if r.OpensCloses == "" {
		return errors.New("Required Open Closing time")
	}
	if len(r.Kitchen) == 0 {
		return errors.New("Required Kitchen")
	}
	if len(r.Menu) == 0 {
		return errors.New("Required Menu")
	}
	return nil
}
