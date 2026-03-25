package preserver

import (
	"baobab/internal/place"
	"fmt"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type PlaceDB struct {
	gorm.Model
	Address    string
	Country    string
	Prefecture string
	City       string
	Postal     string
}

func (PlaceDB) TableName() string {
	return "places"
}

type PlaceRepository interface {
	Create(place *place.Place) error
	GetAll() ([]place.Place, error)
}

type gormPlaceRepository struct {
	db *gorm.DB
}

// MapToDB: Place -> PlaceDB への変換
func PlaceRepoMapToDB(u *place.Place) (*PlaceDB, error) {
	dbPlace := &PlaceDB{}
	err := copier.Copy(dbPlace, u)
	return dbPlace, err
}

// MapToDomain: PlaceDB -> Place への変換
func PlaceRepoMapToDomain(dbU *PlaceDB) (*place.Place, error) {
	place := &place.Place{}
	err := copier.Copy(place, dbU)
	return place, err
}

// MapToDBArray: []Place -> []PlaceDB への変換
func PlaceRepoMapToDBArray(us []place.Place) ([]PlaceDB, error) {
	var dbPlaces []PlaceDB
	err := copier.Copy(&dbPlaces, us)
	return dbPlaces, err
}

// MapToDomainArray: []PlaceDB -> []Place への変換
func PlaceRepoMapToDomainArray(dbUs []PlaceDB) ([]place.Place, error) {
	var places []place.Place
	err := copier.Copy(&places, dbUs)
	return places, err
}

func NewPlaceRepository(db *gorm.DB) PlaceRepository {
	return &gormPlaceRepository{db: db}
}

func (r *gormPlaceRepository) Create(place *place.Place) error {
	dbplace, err := PlaceRepoMapToDB(place)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}
	return r.db.Create(dbplace).Error
}

func (r *gormPlaceRepository) GetAll() ([]place.Place, error) {
	var dbplaces []PlaceDB
	err := r.db.Find(&dbplaces).Error
	if err != nil {
		return nil, err
	}

	placesPtr, err := PlaceRepoMapToDomainArray(dbplaces)
	if err != nil {
		return nil, err
	}

	return placesPtr, nil
}
