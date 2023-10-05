package storage

import (
	"database/sql"
	"del/models"
	"log"

	_ "github.com/lib/pq"
)

type IStorage interface {
	GetOrder(orderId int) *models.Order
	AddOrder(o *models.Order) int
}

type Storage struct {
	Db *sql.DB
}

func NewStorage(db *sql.DB) IStorage {
	return &Storage{
		Db: db,
	}
}

func (s *Storage) GetOrder(orderId int) *models.Order {
	var order models.Order
	query := `SELECT * FROM uzumdeliverytable WHERE id = $1`
	err := s.Db.QueryRow(query, orderId).Scan(&order.ID, &order.Name, &order.Phone,
		&order.Address, &order.Meta, &order.Status, &order.Deliviry_time)
	if err != nil {
		log.Println(err, "select row")
	}
	return &order
}

func (s *Storage) AddOrder(o *models.Order) int {
	var id int
	query := `INSERT INTO uzumdeliverytable (name, phone,address,meta,status,delivery_time) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := s.Db.QueryRow(query, o.Name, o.Phone, o.Address, o.Meta, o.Status, o.Deliviry_time).Scan(&id)
	if err != nil {
		log.Println(err, "insert order")
		return 0
	}

	return id
}
