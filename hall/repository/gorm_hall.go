package repository

import (
	"github.com/Adona12/GoGroup/Movie-and-events/hall"
	"github.com/Adona12/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
)

// CommentGormRepo implements menu.CommentRepository interface
type HallGormRepo struct {
	conn *gorm.DB
}

// NewHALLGormRepo returns new object of CommentGormRepo
func NewHallGormRepo(db *gorm.DB) hall.HallRepository {
	return &HallGormRepo{conn: db}
}
func (hllRepo *HallGormRepo) Halls() ([]model.Hall, []error) {
	hll := []model.Hall{}
	errs := hllRepo.conn.Find(&hll).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return hll, errs
}

// Comment retrieves a customer comment from the database by its id
// func (hllRepo *HallGormRepo) Hall(id uint) (*model.Hall, []error) {
// 	hll := model.Hall{}
// 	errs := hllRepo.conn.First(&hll, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return &hll, errs
// }

// // UpdateComment updates a given customer comment in the database
// func (hllRepo *HallGormRepo) UpdateHall(hall *model.Hall) (*model.Hall, []error) {
// 	hll := hall
// 	errs := hllRepo.conn.Save(hll).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return hll, errs
// }

// // DeleteComment deletes a given customer comment from the database
// func (hllRepo *HallGormRepo) DeleteHall(id uint) (*model.Hall, []error) {
// 	hll, errs := hllRepo.Hall(id)

// 	if len(errs) > 0 {
// 		return nil, errs
// 	}

// 	errs = hllRepo.conn.Delete(hll, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return hll, errs
// }

// StoreComment stores a given customer comment in the database
func (hllRepo *HallGormRepo) StoreHall(hall *model.Hall) (*model.Hall, []error) {
	hll := hall
	errs := hllRepo.conn.Create(hll).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return hll, errs
}
