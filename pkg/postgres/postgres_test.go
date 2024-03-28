package postgres

import (
	"log"
	"shoplist/pkg/models"
	"testing"
)

var db *Database

func TestMain(m *testing.M) {
	var err error
	db, err = NewDb("postgres://postgres:password@localhost:5432/test")
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func TestDatabase_GetById(t *testing.T) {
	m, err := db.GetById(1)
	if err != nil {
		t.Error(err)
	}
	if m.Title != "abcd" {
		t.Error("value returned doesn't match expected value")
	}
}

func TestDatabase_Insert(t *testing.T) {
	id, err := db.Insert("testing", "inserted from TestDataBase_Insert", 1)
	if err != nil {
		t.Error(err)
	}
	m, err := db.GetById(id)
	if err != nil {
		t.Error(err)
	}
	if m.Title != "testing" {
		t.Error("value returned doesn't match expected value")
	}
}

func TestDatabase_Delete(t *testing.T) {
	err := db.Delete(1)
	if err != nil {
		t.Error(err)
	}
	_, err = db.GetById(1)
	if err == nil {
		t.Error("should have been deleted")
	}
}

func TestDatabase_GetAll(t *testing.T) {
	data, err := db.GetAll(1)
	if err != nil {
		t.Error(err)
	}
	expected := []models.ListItem{
		{
			Id:          2,
			Title:       "efgh",
			Description: "item 2",
			UserId:      1,
		},
		{
			Id:          4,
			Title:       "testing",
			Description: "inserted from TestDataBase_Insert",
			UserId:      1,
		},
	}
	for i, v := range data {
		e := expected[i]
		if e.Id != v.Id || e.Title != v.Title || e.Description != v.Description || e.UserId != v.UserId {
			t.Error(e, v)
		}
	}
}

func TestDatabase_Update(t *testing.T) {
	m := models.ListItem{
		Id:          4,
		Title:       "testing",
		Description: "updated?",
		UserId:      1,
	}
	_, err := db.Update(m)
	if err != nil {
		t.Error(err)
	}

	ret, err := db.GetById(m.Id)
	if err != nil {
		t.Error(err)
	}

	if ret.Description != m.Description {
		t.Fail()
	}
}
