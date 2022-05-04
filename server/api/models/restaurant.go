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
	UserID			uint64		`json:"user_id"`
	RestaurantID  	uint64		`gorm:"not null;" json:"restaurant_id"`
	OrderedID		uint64		`gorm:"not null;" json:"ordered_id"`
	OrderedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
}

type Comment struct {
	ID				uint64		`gorm:"primary_key;auto_increment" json:"id"`
	UserID			uint64		`json:"user_id"`
	RestaurantID  	uint64		`gorm:"not null;" json:"restaurant_id"`
	OrderID			uint64		`gorm:"not null;" json:"order_id"`
	Rate			uint32		`gorm:"size:10; not null;" json:"rate"`
	CommentText		string		`gorm:"size:255;" json:"comment_text"`
	CreatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type Food struct {
	ID				uint64		`gorm:"primary_key;auto_increment" json:"id"`
	Name     		string      `gorm:"size:255;not null;" json:"name"`
	Type     		string      `gorm:"size:255;" json:"type"`
	Price			float64		`gorm:"not null;" json:"price"`
}

type Menu struct {
	ID				uint64		`gorm:"primary_key;auto_increment" json:"id"`
	FoodID			uint64		`gorm:"not null;" json:"food_id"`
	RestaurantID  	uint64		`gorm:"not null;" json:"restaurant_id"`
}


type Restaurant struct {
	ID        		uint64      	`gorm:"primary_key;auto_increment" json:"id"`
	UserID			uint64			`gorm:"not null;" json:"user_id"`
	Name     		string      	`gorm:"size:255;not null;" json:"name"`
	Kitchen  		string    		`json:"kitchen"`
	Address   		string	    	`gorm:"size:400;not null;" json:"address"`
	OpensCloses		string	    	`gorm:"size:50;not null;" json:"opens_closes"`
	CreatedAt 		time.Time   	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time   	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}




func (o *Order) Prepare() {
	o.ID = 0
	o.UserID = 0
	o.RestaurantID = 0
	o.OrderedID = 0
	o.OrderedAt = time.Now()
}

func (m *Menu) Prepare() {
	m.ID = 0
	m.FoodID = 0
	m.RestaurantID = 0 
}


func (c *Comment) Prepare() {
	c.ID = 0
	c.UserID = 0
	c.RestaurantID = 0
	c.OrderID = 0
	c.Rate = 0
	c.CreatedAt = time.Now()
}

func (f *Food) Prepare() {
	f.ID = 0
	f.Name = html.EscapeString(strings.TrimSpace(f.Name))
	f.Type = html.EscapeString(strings.TrimSpace(f.Type))
	f.Price = 0
}

func (r *Restaurant) Prepare() {
	r.ID = 0
	r.UserID = 0
	r.Name = html.EscapeString(strings.TrimSpace(r.Name))
	r.Kitchen = html.EscapeString(strings.TrimSpace(r.Kitchen))
	r.Address = html.EscapeString(strings.TrimSpace(r.Address))
	r.OpensCloses = html.EscapeString(strings.TrimSpace(r.OpensCloses))
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

	return nil
}

func (r *Restaurant) SaveRestaurant(db *gorm.DB) (*Restaurant, error) {
	var err error
	err = db.Debug().Model(&Restaurant{}).Create(&r).Error
	if err != nil {
		return &Restaurant{}, err
	}
	if r.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", r.ID).Take(&r.Name).Error
		if err != nil {
			return &Restaurant{}, err
		}
	}
	return r, nil
}

func (r *Restaurant) FindAllRestaurants(db *gorm.DB) (*[]Restaurant, error) {
	var err error
	Restaurants := []Restaurant{}
	err = db.Debug().Model(&Restaurant{}).Limit(100).Find(&Restaurants).Error
	if err != nil {
		return &[]Restaurant{}, err
	}
	if len(Restaurants) > 0 {
		for i, _ := range Restaurants {
			err := db.Debug().Model(&User{}).Where("id = ?", Restaurants[i].ID).Take(&Restaurants[i].Name).Error
			if err != nil {
				return &[]Restaurant{}, err
			}
		}
	}
	return &Restaurants, nil
}

func (r *Restaurant) FindRestaurantByID(db *gorm.DB, pid uint64) (*Restaurant, error) {
	var err error
	err = db.Debug().Model(&Restaurant{}).Where("id = ?", pid).Take(&r).Error
	if err != nil {
		return &Restaurant{}, err
	}
	if r.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", r.ID).Take(&r.Name).Error
		if err != nil {
			return &Restaurant{}, err
		}
	}
	return r, nil
}

func (r *Restaurant) UpdateARestaurant(db *gorm.DB) (*Restaurant, error) {

	var err error

	err = db.Debug().Model(&Restaurant{}).Where("id = ?", r.ID).Updates(Restaurant{Name: r.Name, Kitchen: r.Kitchen, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Restaurant{}, err
	}
	if r.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", r.ID).Take(&r.Name).Error
		if err != nil {
			return &Restaurant{}, err
		}
	}
	return r, nil
}

func (r *Restaurant) DeleteARestaurant(db *gorm.DB, rid uint64) (int64, error) {

	db = db.Debug().Model(&Restaurant{}).Where("id = ?", rid).Take(&Restaurant{}).Delete(&Restaurant{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Restaurant not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}





