package storage

import "shoplist/pkg/models"

//Storage defines a uniform interface so any form of storage may be plugged in
type Storage interface {
	Insert(title, desc string) int
	Delete(id int)
	Update(data any, id int)
	GetAll() []models.ListItem
	GetById(id int)
}