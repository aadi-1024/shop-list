package storage

import "shoplist/pkg/models"

// Storage defines a uniform interface so any form of storage may be plugged in
type Storage interface {
	Insert(title, desc string, uid int) (int, error)
	Delete(id int) error
	Update(data models.ListItem) (models.ListItem, error)
	GetAll(uid int) ([]models.ListItem, error)
	GetById(id int) (*models.ListItem, error)
}
