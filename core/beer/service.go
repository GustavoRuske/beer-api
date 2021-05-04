package beer

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Store(beer *Beer) error
	Update(beer *Beer) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (service *Service) GetAll() ([]*Beer, error) {
	var result []*Beer

	rows, err := service.DB.Query("select id, name, type, style from beer")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var beer Beer
		err = rows.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)
		if err != nil {
			return nil, err
		}

		result = append(result, &beer)
	}
	return result, nil
}

func (service *Service) Get(ID int64) (*Beer, error) {
	var beer Beer

	statement, err := service.DB.Prepare("select id, name, type, style from beer where id = ?")

	if err != nil {
		return nil, err
	}

	defer statement.Close()
	err = statement.QueryRow(ID).Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)

	if err != nil {
		return nil, err
	}

	return &beer, nil
}

func (service *Service) Store(beer *Beer) error {
	tx, err := service.DB.Begin()

	if err != nil {
		return err
	}

	statement, err := tx.Prepare("insert into beer(id, name, type, style) values(?,?,?,?)")

	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(&beer.ID, &beer.Name, &beer.Type, &beer.Style)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (service *Service) Update(beer *Beer) error {
	if beer.ID == 0 {
		return fmt.Errorf("Invalid Id")
	}
	tx, err := service.DB.Begin()

	if err != nil {
		return err
	}

	statement, err := tx.Prepare("update beer set name=?, type=?, style=? where id=?")

	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(&beer.Name, &beer.Type, &beer.Style, &beer.ID)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (service *Service) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("Invalid Id")
	}
	tx, err := service.DB.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from beer where id=?", ID)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
