package repositories

import (
	"coffee-shop-golang/internal/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type OrdersRepository struct {
	*sqlx.DB
}

func InitializeRepoOrders(db *sqlx.DB) *OrdersRepository {
	cr := OrdersRepository{db}
	return &cr
}

func (r *OrdersRepository) RepsitoryGetAllOrders() ([]models.OrdersModel, error) {
	result := []models.OrdersModel{}
	query := `SELECT * FROM orders`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *OrdersRepository) RepsitoryCreateOrders(body *models.OrdersModel) (error) {
	var err error
	var orderId string

	tx := r.MustBegin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	queryOrder := `INSERT INTO orders (users_id, payment_methods_id, deliveries_id, promos_id) VALUES (:users_id, :payment_methods_id, :deliveries_id, :promos_id) RETURNING orders_id`

	result, execErr := tx.NamedQuery(queryOrder, body)
	if result.Next() {
		result.Scan(&orderId)
		fmt.Println(orderId)
	}

	if execErr != nil {
		tx.Rollback()
		log.Fatal(execErr)
		return execErr
	}

	result.Close()

	queryOrderProduct := `insert into orders_products (orders_id, products_id, sizes_id, orders_products_qty, orders_products_subtotal, hot_or_ice) values (:orders_id, :products_id, :sizes_id, :orders_products_qty, :orders_products_subtotal, :hot_or_ice)`

	_, execErrOp := tx.NamedExec(queryOrderProduct, body)

	if execErrOp != nil {
		tx.Rollback()
		log.Fatal(execErrOp)
		return execErr
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		fmt.Println(err)
	}

	return nil
}