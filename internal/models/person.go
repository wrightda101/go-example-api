package models

import "encoding/json"

func (db *DBManager) CreatePerson(p Person) (uint, error) {
	tr := db.db.Create(&p)
	return p.ID, tr.Error
}

func (db *DBManager) GetPerson(id uint) (Person, error) {
	var person Person
	tr := db.db.First(&person, id)
	return person, tr.Error
}

func (db *DBManager) ListPerson() ([]Person, error) {
	persons := make([]Person, 0)
	tr := db.db.Find(&persons)
	return persons, tr.Error
}

func (p *Person) JSON() []byte {
	data, _ := json.Marshal(p)
	return data
}

func PersonFromJSON(data []byte) Person {
	var p Person
	json.Unmarshal(data, &p)
	return p
}
