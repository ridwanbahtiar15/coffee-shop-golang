package repositories

import (
	"coffee-shop-golang/internal/models"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type ProductsRepository struct {
	*sqlx.DB
}

func InitializeRepoProducts(db *sqlx.DB) *ProductsRepository {
	cr := ProductsRepository{db}
	return &cr
}

func (r *ProductsRepository) RepsitoryGetAllProducts() ([]models.ProductsModel, error) {
	result := []models.ProductsModel{}
	query := `SELECT p.products_name, p.products_price, p.products_desc,
						p.products_stock, p.products_image, p.categories_id, c.categories_name FROM products p JOIN categories c on p.categories_id = c.categories_id ORDER BY p.products_id ASC`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ProductsRepository) RepsitoryProductsById(id string) ([]models.ProductsModel, error) {
	result := []models.ProductsModel{}
	query := `SELECT p.products_name, p.products_price, p.products_desc, 
						p.products_stock, p.products_image, p.categories_id, c.categories_name FROM products p join categories c on p.categories_id = c.categories_id WHERE p.products_id = $1`
	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ProductsRepository) RepsitoryCreateProducts(body *models.ProductsModel) (sql.Result, error) {
	query := `INSERT INTO products 
						(products_name, products_price, products_desc, products_stock, products_image, categories_id) 
						VALUES (:products_name, :products_price, :products_desc, :products_stock, :products_image, :categories_id)`
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err 
	}
	return result, nil
}

func (r *ProductsRepository) RepsitoryUpdateProducts(body *models.ProductsModel, id string) (sql.Result, error) {
	query := `UPDATE products SET products_name=:products_name,
						products_price=:products_price, products_desc=:products_desc, products_stock=:products_stock, products_image=:products_image, categories_id=:categories_id WHERE products_id =` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err 
	}
	return result, nil
}

func (r *ProductsRepository) RepositoryDeleteProducts(id string) (sql.Result, error) {
	query := `DELETE FROM products WHERE products_id = $1`
	result, err := r.Exec(query, id)
	if err != nil {
		return result, err 
	}
	return result, nil
}

func (r *ProductsRepository) RepsitoryGetFilterProducts(name string, category string, minrange string, maxrange string, page string, limit string) ([]models.ProductsModel, error) {

	newPage, _ := strconv.Atoi("1")
	newLimit, _ := strconv.Atoi("99")

	if minrange == "" { minrange = "10000" }
	if maxrange == "" { maxrange = "120000" }
	if page != "" {
		newPage, _ = strconv.Atoi(page) 
	}
	if limit != "" {
		newLimit, _ = strconv.Atoi(limit) 
	}
	fmt.Println(name)

	result := []models.ProductsModel{}
	query := `SELECT p.products_id, p.products_name, p.products_price, p.products_desc, p.products_stock, p.products_image, c.categories_name
  FROM products p`;

	if name != "" && category != "" {
		query += ` JOIN categories c on p.categories_id = c.categories_id
		WHERE p.products_name like $1
		AND products_price >= $3 and products_price <= $4
		AND c.categories_name = $2
		ORDER BY p.created_at ASC
		LIMIT $5 OFFSET $6`;
		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, "%" + name + "%", category, minrange, maxrange, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	if (category != "") {
		query += ` JOIN categories c on p.categories_id = c.categories_id
		WHERE c.categories_name = $1
		AND products_price >= $2 and products_price <= $3
		ORDER BY p.created_at ASC
		LIMIT $4 OFFSET $5`;
		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, category, minrange, maxrange, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	if (name != "") {
		query += ` JOIN categories c on p.categories_id = c.categories_id
		WHERE p.products_name like $1
		AND products_price >= $2 and products_price <= $3
		ORDER BY p.created_at ASC
		LIMIT $4 OFFSET $5`;
		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, "%" + name + "%", minrange, maxrange, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	return result, nil
}