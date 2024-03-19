package models

import (
	"testing"

	"gotest.tools/assert"
)

func testDB() DBManager {
	db := NewMustFromConfig(Config{
		Host:     "localhost",
		Port:     5432,
		DbName:   "test",
		User:     "example",
		Pass:     "example",
		SslMode:  "disable",
		TimeZone: "GMT",
	})
	db.db.Where("1 = 1").Delete(&Person{})
	return db
}

func TestDBManager_CreatePerson(t *testing.T) {
	prefix := "CreatePerson"
	type args struct {
		p Person
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"create a person", args{p: Person{FirstName: prefix + "joe", LastName: prefix + "bloggs", Age: 30}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testDB()
			if _, err := db.CreatePerson(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.CreatePerson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_CreateAndGetPerson(t *testing.T) {
	prefix := "CreateAndGetPerson"
	type args struct {
		p Person
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"create and get person 1", args{p: Person{FirstName: prefix + "first_name_a", LastName: prefix + "last_name_a", Age: 20}}, false},
		{"create and get person 2", args{p: Person{FirstName: prefix + "first_name_b", LastName: prefix + "last_name_b", Age: 30}}, false},
		{"create and get person 3", args{p: Person{FirstName: prefix + "first_name_c", LastName: prefix + "last_name_c", Age: 40}}, false},
		{"create and get person 4", args{p: Person{FirstName: prefix + "first_name_d", LastName: prefix + "last_name_d", Age: 50}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testDB()
			id, err := db.CreatePerson(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.CreatePerson() error = %v, wantErr %v", err, tt.wantErr)
			}
			returnedPerson, err := db.GetPerson(id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.GetPerson() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.args.p.FirstName, returnedPerson.FirstName)
			assert.Equal(t, tt.args.p.LastName, returnedPerson.LastName)
			assert.Equal(t, tt.args.p.Age, returnedPerson.Age)
		})
	}
}

func TestDBManager_ListPerson(t *testing.T) {
	prefix := "ListPerson"
	type args struct {
		p Person
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"list person 1", args{p: Person{FirstName: prefix + "first_name_a", LastName: prefix + "last_name_a", Age: 20}}, false},
		{"list person 2", args{p: Person{FirstName: prefix + "first_name_b", LastName: prefix + "last_name_b", Age: 30}}, false},
		{"list person 3", args{p: Person{FirstName: prefix + "first_name_c", LastName: prefix + "last_name_c", Age: 40}}, false},
		{"list person 4", args{p: Person{FirstName: prefix + "first_name_d", LastName: prefix + "last_name_d", Age: 50}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testDB()
			id, err := db.CreatePerson(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.CreatePerson() error = %v, wantErr %v", err, tt.wantErr)
			}
			returnedPersons, err := db.ListPerson()
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.GetPerson() error = %v, wantErr %v", err, tt.wantErr)
			}

			var foundPerson Person
			for _, p := range returnedPersons {
				if p.ID == id {
					foundPerson = p
				}
			}

			assert.Equal(t, tt.args.p.FirstName, foundPerson.FirstName)
			assert.Equal(t, tt.args.p.LastName, foundPerson.LastName)
			assert.Equal(t, tt.args.p.Age, foundPerson.Age)
		})
	}
}
