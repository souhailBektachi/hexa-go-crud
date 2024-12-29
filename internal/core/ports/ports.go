package ports

import . "github.com/souhailBektachi/hexa-go-crud/internal/core/domain"

type SomthingRepository interface {
	Save(somthin Somthing) error
	FindById(id int) (Somthing, error)
	FindAll() ([]Somthing, error)
	DeleteById(id int) error
	Update(somthin Somthing) error
}

type SomthingService interface {
	Save(somthin Somthing) error
	FindById(id int) (Somthing, error)
	FindAll() ([]Somthing, error)
	DeleteById(id int) error
	Update(somthin Somthing) error
}
