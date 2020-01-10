package service

import (
	"github.com/Adona12/GoGroup/Movie-and-events/hall"
	"github.com/Adona12/GoGroup/Movie-and-events/model"
)

type HallService struct {
	hallRepo hall.HallRepository
}

func NewHallService(HallRepos hall.HallRepository) hall.HallService {
	return &HallService{hallRepo: HallRepos}
}

// HALLs returns all stored comments
func (hs *HallService) Halls() ([]model.Hall, []error) {
	hll, errs := hs.hallRepo.Halls()
	if len(errs) > 0 {
		return nil, errs
	}
	return hll, errs
}

// func (hs *HallService) Hall(id uint) ([]*model.Hall, []error) {
// 	cmnts, errs := hs.hallRepo.Hall()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return cmnts, errs
// }
// func (hs *HallService) DeleteHall(id uint) (*model.Hall, []error) {
// 	cmnts, errs := hs.hallRepo.Hall()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return cmnts, errs
// }
func (hs *HallService) StoreHall(hall *model.Hall) (*model.Hall, []error) {
	cmnts, errs := hs.hallRepo.StoreHall(hall)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// func (hs *HallService) UpdateHall(hall *model.Hall) (*model.Hall, []error) {
// 	cmnts, errs := hs.hallRepo.StoreHall(hall)
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return cmnts, errs
// }
