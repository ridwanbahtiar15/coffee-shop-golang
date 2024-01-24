package repositories

import (
	"coffee-shop-golang/internal/models"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type IProductsRepository interface {
	// RepositoryGetFilterProducts(name string, category string, minrange string, maxrange string, page string, limit string, sort string) ([]models.ProductsResponseModel, error)
	RepositoryCountProducts(name string, category string, minrange string, maxrange string) ([]string, error)
	RepositoryGetAllProducts(name string, category string, minrange string, maxrange string, page string, limit string, sort string) ([]models.ProductsResponseModel, error)
	RepositoryProductsById(id string) ([]models.ProductsResponseModel, error)
	RepositoryCreateProducts(body *models.ProductsModel) (int, error)
	RepositoryUpdateProducts(body *models.UpdateProductsModel, id string) (error)
	RepositoryUpdateImgProducts(productImg string, id string) (error)
	RepositoryDeleteProducts(id string) (int64, error)
}

type ProductsRepository struct {
	*sqlx.DB
}

func InitializeRepoProducts(db *sqlx.DB) *ProductsRepository {
	cr := ProductsRepository{db}
	return &cr
}

func (r *ProductsRepository) RepositoryGetAllProducts(name string, category string, minrange string, maxrange string, page string, limit string, sort string) ([]models.ProductsResponseModel, error) {
	// result := []models.ProductsResponseModel{}
	// query := `SELECT p.products_id, p.products_name, p.products_price, p.products_desc,
	// 					p.products_stock, p.products_image, p.categories_id, c.categories_name FROM products p JOIN categories c on p.categories_id = c.categories_id ORDER BY p.products_id ASC`
	// err := r.Select(&result, query)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil

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

	result := []models.ProductsResponseModel{}
	query := `SELECT p.products_id, p.products_name, p.products_price, p.products_desc, p.products_stock, p.products_image, p.categories_id, c.categories_name
  FROM products p`;

	if name != "" && category != "" {
		query += ` JOIN categories c on p.categories_id = c.categories_id WHERE p.products_name like $1 AND products_price >= $3 and products_price <= $4 AND c.categories_name = $2`;
		
		switch sort {
		case "minprice":
			query += ` ORDER BY p.products_price ASC LIMIT $5 OFFSET $6`
		case "maxprice":
			query += ` ORDER BY p.products_price DESC LIMIT $5 OFFSET $6`
		default:
			query += ` ORDER BY p.products_price ASC LIMIT $5 OFFSET $6`
		}

		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, fmt.Sprintf("%%%s%%", name), category, minrange, maxrange, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	if category != "" {
		query += ` JOIN categories c on p.categories_id = c.categories_id
		WHERE c.categories_name = $1
		AND products_price >= $2 and products_price <= $3`;

		switch sort {
		case "minprice":
			query += ` ORDER BY p.products_price ASC LIMIT $4 OFFSET $5`
		case "maxprice":
			query += ` ORDER BY p.products_price DESC LIMIT $4 OFFSET $5`
		default:
			query += ` ORDER BY p.products_price ASC LIMIT $4 OFFSET $5`
		}

		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, category, minrange, maxrange, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	if name != "" {
		query += ` JOIN categories c on p.categories_id = c.categories_id
		WHERE p.products_name like $1
		AND products_price >= $2 and products_price <= $3`;

		switch sort {
		case "minprice":
			query += ` ORDER BY p.products_price ASC LIMIT $4 OFFSET $5`
		case "maxprice":
			query += ` ORDER BY p.products_price DESC LIMIT $4 OFFSET $5`
		default:
			query += ` ORDER BY p.products_price ASC LIMIT $4 OFFSET $5`
		}

		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, fmt.Sprintf("%%%s%%", name), minrange, maxrange, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	if name == "" && category == "" {
		query += ` JOIN categories c on p.categories_id = c.categories_id 
		WHERE products_price >= $1 AND products_price <= $2`

		switch sort {
		case "minprice":
			query += ` ORDER BY p.products_price ASC LIMIT $3 OFFSET $4`
		case "maxprice":
			query += ` ORDER BY p.products_price DESC LIMIT $3 OFFSET $4`
		default:
			query += ` ORDER BY p.products_price ASC LIMIT $3 OFFSET $4`
		}

		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, minrange, maxrange, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	offset := newPage * newLimit - newLimit;
	err := r.Select(&result, query, minrange, maxrange, newLimit, strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ProductsRepository) RepositoryProductsById(id string) ([]models.ProductsResponseModel, error) {
	result := []models.ProductsResponseModel{}
	query := `SELECT p.products_name, p.products_price, p.products_desc, 
						p.products_stock, p.products_image, p.categories_id, c.categories_name FROM products p join categories c on p.categories_id = c.categories_id WHERE p.products_id = $1`
	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ProductsRepository) RepositoryCreateProducts(body *models.ProductsModel) (int, error) {
	var id int
	query := `INSERT INTO products 
						(products_name, products_price, products_desc, products_stock, categories_id) 
						VALUES 
						(:products_name, :products_price, :products_desc, :products_stock, :categories_id) 
						RETURNING products_id`
	result, err := r.NamedQuery(query, body)
	if err != nil {
		return 0, err 
	}
	if result.Next() {
    result.Scan(&id)
	}
	return id, nil
}

func (r *ProductsRepository) RepositoryUpdateImgProducts(productImg string, id string) (error) {
	query := `UPDATE products SET products_image = $1 WHERE products_id = $2`
	_, err := r.Exec(query, productImg,  id)
	if err != nil {
		return err 
	}
	return nil
}

func (r *ProductsRepository) RepositoryUpdateProducts(body *models.UpdateProductsModel, id string) (error) {
	query := `UPDATE products SET products_name=:products_name,
						products_price=:products_price, products_desc=:products_desc, products_stock=:products_stock, products_image=:products_image, categories_id=:categories_id WHERE products_id =` + id
	_, err := r.NamedExec(query, body)
	if err != nil {
		return err 
	}
	return nil
}

func (r *ProductsRepository) RepositoryDeleteProducts(id string) (int64, error) {
	var res int64 = 1
	query := `DELETE FROM products WHERE products_id = $1`
	result, err := r.Exec(query, id)
	if err != nil {
		return 0, err 
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		res = 0
		return res, nil
	}
	return res, nil
	
}

// func (r *ProductsRepository) RepositoryGetFilterProducts(name string, category string, minrange string, maxrange string, page string, limit string, sort string) ([]models.ProductsResponseModel, error) {

// 	newPage, _ := strconv.Atoi("1")
// 	newLimit, _ := strconv.Atoi("99")

// 	if minrange == "" { minrange = "10000" }
// 	if maxrange == "" { maxrange = "120000" }
// 	if page != "" {
// 		newPage, _ = strconv.Atoi(page) 
// 	}
// 	if limit != "" {
// 		newLimit, _ = strconv.Atoi(limit) 
// 	}

// 	result := []models.ProductsResponseModel{}
// 	query := `SELECT p.products_id, p.products_name, p.products_price, p.products_desc, p.products_stock, p.products_image, p.categories_id, c.categories_name
//   FROM products p`;

// 	if name != "" && category != "" {
// 		query += ` JOIN categories c on p.categories_id = c.categories_id WHERE p.products_name like $1 AND products_price >= $3 and products_price <= $4 AND c.categories_name = $2`;
		
// 		switch sort {
// 		case "minprice":
// 			query += ` ORDER BY p.products_price ASC LIMIT $5 OFFSET $6`
// 		case "maxprice":
// 			query += ` ORDER BY p.products_price DESC LIMIT $5 OFFSET $6`
// 		default:
// 			query += ` ORDER BY p.products_price ASC LIMIT $5 OFFSET $6`
// 		}

// 		offset := newPage * newLimit - newLimit;
// 		err := r.Select(&result, query, fmt.Sprintf("%%%s%%", name), category, minrange, maxrange, newLimit, strconv.Itoa(offset))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return result, nil
// 	}

// 	if category != "" {
// 		query += ` JOIN categories c on p.categories_id = c.categories_id
// 		WHERE c.categories_name = $1
// 		AND products_price >= $2 and products_price <= $3`;

// 		switch sort {
// 		case "minprice":
// 			query += ` ORDER BY p.products_price ASC LIMIT $4 OFFSET $5`
// 		case "maxprice":
// 			query += ` ORDER BY p.products_price DESC LIMIT $4 OFFSET $5`
// 		default:
// 			query += ` ORDER BY p.products_price ASC LIMIT $4 OFFSET $5`
// 		}

// 		offset := newPage * newLimit - newLimit;
// 		err := r.Select(&result, query, category, minrange, maxrange, newLimit, strconv.Itoa(offset))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return result, nil
// 	}

// 	if name != "" {
// 		query += ` JOIN categories c on p.categories_id = c.categories_id
// 		WHERE p.products_name like $1
// 		AND products_price >= $2 and products_price <= $3`;

// 		switch sort {
// 		case "minprice":
// 			query += ` ORDER BY p.products_price ASC LIMIT $4 OFFSET $5`
// 		case "maxprice":
// 			query += ` ORDER BY p.products_price DESC LIMIT $4 OFFSET $5`
// 		default:
// 			query += ` ORDER BY p.products_price ASC LIMIT $4 OFFSET $5`
// 		}

// 		offset := newPage * newLimit - newLimit;
// 		err := r.Select(&result, query, fmt.Sprintf("%%%s%%", name), minrange, maxrange, newLimit, strconv.Itoa(offset))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return result, nil
// 	}

// 	if name == "" && category == "" {
// 		query += ` JOIN categories c on p.categories_id = c.categories_id 
// 		WHERE products_price >= $1 AND products_price <= $2`

// 		switch sort {
// 		case "minprice":
// 			query += ` ORDER BY p.products_price ASC LIMIT $3 OFFSET $4`
// 		case "maxprice":
// 			query += ` ORDER BY p.products_price DESC LIMIT $3 OFFSET $4`
// 		default:
// 			query += ` ORDER BY p.products_price ASC LIMIT $3 OFFSET $4`
// 		}

// 		offset := newPage * newLimit - newLimit;
// 		err := r.Select(&result, query, minrange, maxrange, newLimit, strconv.Itoa(offset))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return result, nil
// 	}

// 	return result, nil
// }

func (r *ProductsRepository) RepositoryCountProducts(name string, category string, minrange string, maxrange string) ([]string, error) {
	if minrange == "" { minrange = "10000" }
	if maxrange == "" { maxrange = "120000" }
	count := []string{}

	query := `SELECT COUNT(*) FROM products p`

	if name != "" && category != "" {
		query += ` JOIN categories c ON p.categories_id = c.categories_id 
								WHERE p.products_name LIKE $1 
								AND c.categories_name = $2 
								AND p.products_price >= $3 AND p.products_price <= $4`

		err := r.Select(&count, query, fmt.Sprintf("%%%s%%", name), category, minrange, maxrange)
		if err != nil {
			return nil, err
		}
		return count, nil
	}

	if name != "" {
		query += ` WHERE p.products_name LIKE $1 
								AND p.products_price >= $2 AND p.products_price <= $3`;
		err := r.Select(&count, query, fmt.Sprintf("%%%s%%", name), minrange, maxrange)
		if err != nil {
			return nil, err
		}
		return count, nil
	}

	if category != "" {
		query += ` JOIN categories c ON p.categories_id = c.categories_id 
								WHERE c.categories_name = $1 
								AND p.products_price >= $2 AND p.products_price <= $3`;
		err := r.Select(&count, query, category, minrange, maxrange)
		if err != nil {
			return nil, err
		}
		return count, nil
	}

	if name == "" && category == "" {
		query += ` WHERE p.products_price >= $1 AND p.products_price <= $2`
		err := r.Select(&count, query, minrange, maxrange)
		if err != nil {
			return nil, err
		}
		return count, nil
	}
	return count, nil
}
