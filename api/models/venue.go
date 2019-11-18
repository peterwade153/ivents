package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type Venue struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Description string `gorm:"not null"                 json:"description"`
	Location    string `gorm:"size:100;not null"        json:"location"`
	Capacity    int    `gorm:"not null"                 json:"capacity"`
	Category    string `gorm:"size:100;not null"        json:"category"`
	CreatedBy   User   `gorm:"foreignKey:UserID;"       json:"-"`
	UserID      uint   `gorm:"not null"                 json:"user_id"`
}

func (v *Venue) Prepare() {
	v.Name = strings.TrimSpace(v.Name)
	v.Description = strings.TrimSpace(v.Description)
	v.Location = strings.TrimSpace(v.Location)
	v.Category = strings.TrimSpace(v.Category)
	v.CreatedBy = User{}
}

func (v *Venue) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		{
			return nil
		}
	default:
		if v.Name == "" {
			return errors.New("Name is required")
		}
		if v.Description == "" {
			return errors.New("Description about venue is required")
		}
		if v.Location == "" {
			return errors.New("Location of venue is required")
		}
		if v.Category == "" {
			return errors.New("Category of venue is required")
		}
		if v.Capacity < 0 {
			return errors.New("Capacity of venue is invalid")
		}
		return nil
	}
}

func (v *Venue) Save(db *gorm.DB) (*Venue, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Debug().Create(&v).Error
	if err != nil {
		return &Venue{}, err
	}
	return v, nil
}

func (v *Venue) GetVenue(db *gorm.DB) (*Venue, error) {
	venue := &Venue{}
	if err := db.Debug().Table("venues").Where("name = ?", v.Name).First(venue).Error; err != nil {
		return nil, err
	}
	return venue, nil
}

func GetVenues(db *gorm.DB) (*[]Venue, error) {
	venues := []Venue{}
	if err := db.Debug().Table("venues").Find(&venues).Error; err != nil {
		return &[]Venue{}, err
	}
	return &venues, nil
}

func GetVenueById(id int, db *gorm.DB) (*Venue, error) {
	venue := &Venue{}
	if err := db.Debug().Table("venues").Where("id = ?", id).First(venue).Error; err != nil {
		return nil, err
	}
	return venue, nil
}

func (v *Venue) UpdateVenue(id int, db *gorm.DB) (*Venue, error) {
	if err := db.Debug().Table("venues").Where("id = ?", id).Updates(Venue{
		Name: v.Name, 
		Description: v.Description, 
		Location: v.Location, 
		Capacity: v.Capacity, 
		Category: v.Category}).Error; err != nil {
		return &Venue{}, err
	}
	return v, nil
}

func DeleteVenue(id int, db *gorm.DB) error {
	if err := db.Debug().Table("venues").Where("id = ?", id).Delete(&Venue{}).Error; err != nil {
		return err
	}
	return nil
}
