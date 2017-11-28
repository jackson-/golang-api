package main

import (
	"database/sql"
	"errors"
)

type item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *item) getItem(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *item) updateItem(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *item) deleteItem(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *item) createItem(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getItems(db *sql.DB, start, count int) ([]item, error) {
	rows, err := db.Query(
		"SELECT id, name, price FROM items LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []item{}

	for rows.Next() {
		var i item
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return items, nil
}
