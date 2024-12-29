package service

import (
	"github.com/souhailBektachi/hexa-go-crud/internal/core/ports"

	. "github.com/souhailBektachi/hexa-go-crud/internal/core/domain"
)

type SomthingService struct {
	repo ports.SomthingRepository
}

func NewSomthingService(repo ports.SomthingRepository) *SomthingService {
	return &SomthingService{repo}
}

func (s *SomthingService) Save(somthin Somthing) error {
	return s.repo.Save(somthin)
}
func (s *SomthingService) FindById(id int) (Somthing, error) {
	return s.repo.FindById(id)
}

func (s *SomthingService) FindAll() ([]Somthing, error) {
	return s.repo.FindAll()
}
func (s *SomthingService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}

func (s *SomthingService) Update(somthin Somthing) error {
	return s.repo.Update(somthin)
}
