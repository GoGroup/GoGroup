package service

import (
	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/GoGroup/Movie-and-events/model"
)

type CinemaService struct {
	cinemaRepo cinema.CinemaRepository
}

func NewCinemaService(CinemaRepos cinema.CinemaRepository) cinema.CinemaService {

	return &CinemaService{cinemaRepo: CinemaRepos}
}

// HALLs returns all stored comments
func (cs *CinemaService) Cinemas() ([]model.Cinema, []error) {
	cll, errs := cs.cinemaRepo.Cinemas()
	if len(errs) > 0 {
		return nil, errs
	}
	return cll, errs
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
func (cs *CinemaService) StoreCinema(cinema *model.Cinema) (*model.Cinema, []error) {
	cmnts, errs := cs.cinemaRepo.StoreCinema(cinema)
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
