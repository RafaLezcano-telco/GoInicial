package handler

import (
	"errors"

	"github.com/RafaLezcano-telco/GoInicial/Testing/clase5/api/model"
)

type StorageMockOk struct{}

func (psm *StorageMockOk) Create(person *model.Person) error {
	return nil
}
func (psm *StorageMockOk) Update(ID int, person *model.Person) error {
	return nil
}
func (psm *StorageMockOk) Delete(ID int) error {
	return nil
}
func (psm *StorageMockOk) GetByID(ID int) (model.Person, error) {
	return model.Person{}, nil
}
func (psm *StorageMockOk) GetAll() (model.Persons, error) {
	return model.Persons{}, nil
}

type StorageMockError struct{}

func (psm *StorageMockError) Create(person *model.Person) error {
	return errors.New("error")
}
func (psm *StorageMockError) Update(ID int, person *model.Person) error {
	return errors.New("error")
}
func (psm *StorageMockError) Delete(ID int) error {
	return errors.New("error")
}
func (psm *StorageMockError) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("error")
}
func (psm *StorageMockError) GetAll() (model.Persons, error) {
	return nil, errors.New("error")
}
