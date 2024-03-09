package memstore

import "shoplist/pkg/models"

type MemStore struct {
	data map[int]models.ListItem
	curId int
}

func NewMemStore() *MemStore {
	return &MemStore{
		data: make(map[int]models.ListItem),
		curId: 0,
	}
}

func (m *MemStore) Insert(title, desc string) int {
	m.curId++
	m.data[m.curId] = models.ListItem{
		Id: m.curId,
		Title: title,
		Description: desc,
	}
	return m.curId
}

func (m *MemStore) Delete(id int) {

}

func (m *MemStore) Update(data any, id int) {

}

func (m *MemStore) GetAll() []models.ListItem {
	arr := make([]models.ListItem, m.curId)

	for i := 1; i <= m.curId; i++ {
		arr[i-1] = m.data[i]
	}

	return arr
}

func (m *MemStore) GetById(id int) {

}