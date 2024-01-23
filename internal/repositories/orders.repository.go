package repositories

import (
	"coffee-shop-golang/internal/models"
	"fmt"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type IOrderRepository interface {
	RepositoryGetAllOrders(orderNumber string, page string, limit string, sort string) ([]models.OrdersResponseModel, error)
	RepositoryGetOrdersById(id string) ([]models.OrdersResponseModel, error)
	RepositoryCreateOrders(body *models.OrdersModel) (error)
	RepositoryCountOrders(orderNumber string) ([]string, error)
	RepositoryUpdateOrders(body *models.OrderUpdateModel, id string) (error)
}

type OrdersRepository struct {
	*sqlx.DB
}

func InitializeRepoOrders(db *sqlx.DB) *OrdersRepository {
	cr := OrdersRepository{db}
	return &cr
}

// func (r *OrdersRepository) RepositoryGetAllOrders() ([]models.OrdersResponseModel, error) {
// 	result := []models.OrdersResponseModel{}
// 	query := `SELECT o.orders_id, o.users_id, o.deliveries_id, p.promos_name, 
// 						o.payment_methods_id, o.orders_status, o.orders_total
// 	 					FROM orders o
// 						JOIN promos p ON o.promos_id = p.promos_id`
// 	err := r.Select(&result, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

func (r *OrdersRepository) RepositoryGetAllOrders(orderNumber string, page string, limit string, sort string) ([]models.OrdersResponseModel, error) {
	newPage, _ := strconv.Atoi("1")
	newLimit, _ := strconv.Atoi("99")

	if page != "" {
		newPage, _ = strconv.Atoi(page) 
	}
	if limit != "" {
		newLimit, _ = strconv.Atoi(limit) 
	}

	result := []models.OrdersResponseModel{}
	query := `SELECT o.orders_id, o.users_id, o.deliveries_id, p.promos_name, 
						o.payment_methods_id, o.orders_status, o.orders_total
	 					FROM orders o
						JOIN promos p ON o.promos_id = p.promos_id`
	if orderNumber != "" {
		query += ` WHERE o.orders_id = $1`
		switch sort {
		case "asc":
			query += ` ORDER BY o.orders_id ASC LIMIT $2 OFFSET $3`
		case "desc":
			query += ` ORDER BY o.orders_id DESC LIMIT $2 OFFSET $3`
		default:
			query += ` ORDER BY o.orders_id ASC LIMIT $2 OFFSET $3`
		}
		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, orderNumber, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	switch sort {
	case "asc":
		query += ` ORDER BY o.orders_id ASC LIMIT $1 OFFSET $2`
	case "desc":
		query += ` ORDER BY o.orders_id DESC LIMIT $1 OFFSET $2`
	default:
		query += ` ORDER BY o.orders_id ASC LIMIT $1 OFFSET $2`
	}
	offset := newPage * newLimit - newLimit;
	err := r.Select(&result, query, newLimit, strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *OrdersRepository) RepositoryGetOrdersById(id string) ([]models.OrdersResponseModel, error) {
	result := []models.OrdersResponseModel{}
	query := `SELECT o.orders_id, o.users_id, o.deliveries_id, p.promos_name, 
						o.payment_methods_id, o.orders_status, o.orders_total
						FROM orders o
						JOIN promos p ON o.promos_id = p.promos_id WHERE orders_id = $1`
	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *OrdersRepository) RepositoryCreateOrders(body *models.OrdersModel) (error) {
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

	queryOrderProduct := `insert into orders_products (orders_id, products_id, sizes_id, orders_products_qty, orders_products_subtotal, hot_or_ice) values ($1, $2, $3, $4, $5, $6)`
	values := []any{orderId, body.Products_id, body.Sizes_id, body.Orders_products_qty, body.Orders_products_subtotal, body.Hot_or_ice}

	_, execErrOp := tx.Exec(queryOrderProduct, values...)

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

// func (r *OrdersRepository) RepositoryGetFilterOrders(orderNumber string, page string, limit string, sort string) ([]models.OrdersResponseModel, error) {
// 	newPage, _ := strconv.Atoi("1")
// 	newLimit, _ := strconv.Atoi("99")

// 	if page != "" {
// 		newPage, _ = strconv.Atoi(page) 
// 	}
// 	if limit != "" {
// 		newLimit, _ = strconv.Atoi(limit) 
// 	}

// 	result := []models.OrdersResponseModel{}
// 	query := `SELECT o.orders_id, o.users_id, o.deliveries_id, p.promos_name, 
// 						o.payment_methods_id, o.orders_status, o.orders_total
// 	 					FROM orders o
// 						JOIN promos p ON o.promos_id = p.promos_id`
// 	if orderNumber != "" {
// 		query += ` WHERE o.orders_id = $1`
// 		switch sort {
// 		case "asc":
// 			query += ` ORDER BY o.orders_id ASC LIMIT $2 OFFSET $3`
// 		case "desc":
// 			query += ` ORDER BY o.orders_id DESC LIMIT $2 OFFSET $3`
// 		default:
// 			query += ` ORDER BY o.orders_id ASC LIMIT $2 OFFSET $3`
// 		}
// 		offset := newPage * newLimit - newLimit;
// 		err := r.Select(&result, query, orderNumber, newLimit, strconv.Itoa(offset))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return result, nil
// 	}

// 	switch sort {
// 	case "asc":
// 		query += ` ORDER BY o.orders_id ASC LIMIT $1 OFFSET $2`
// 	case "desc":
// 		query += ` ORDER BY o.orders_id DESC LIMIT $1 OFFSET $2`
// 	default:
// 		query += ` ORDER BY o.orders_id ASC LIMIT $1 OFFSET $2`
// 	}
// 	offset := newPage * newLimit - newLimit;
// 	err := r.Select(&result, query, newLimit, strconv.Itoa(offset))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

func (r *OrdersRepository) RepositoryCountOrders(orderNumber string) ([]string, error) {
	count := []string{}
	query := `SELECT COUNT(*) FROM orders o`

	if orderNumber != "" {
		query += ` WHERE o.orders_id = $1`
		err := r.Select(&count, query, orderNumber)
		if err != nil {
			return nil, err
		}
		return count, nil
	}

	err := r.Select(&count, query)
		if err != nil {
			return nil, err
		}
		return count, nil
}

func (r *OrdersRepository) RepositoryUpdateOrders(body *models.OrderUpdateModel, id string) (error) {
	query := `UPDATE orders SET orders_status=:orders_status
						WHERE orders_id =` + id
	_, err := r.NamedExec(query, body)
	if err != nil {
		return err 
	}
	return nil
}