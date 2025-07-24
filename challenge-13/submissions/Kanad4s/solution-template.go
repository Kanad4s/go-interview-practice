package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Product represents a product in the inventory system
type Product struct {
	ID       int64
	Name     string
	Price    float64
	Quantity int
	Category string
}

// ProductStore manages product operations
type ProductStore struct {
	db *sql.DB
}

// NewProductStore creates a new ProductStore with the given database connection
func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

// InitDB sets up a new SQLite database and creates the products table
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id INTEGER PRIMARY KEY, name TEXT, price REAL, quantity INTEGER, category TEXT);")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateProduct adds a new product to the database
func (ps *ProductStore) CreateProduct(product *Product) error {
	result, err := ps.db.Exec("INSERT INTO products (name, price, quantity, category) VALUES (?, ?, ?, ?)",
		product.Name, product.Price, product.Quantity, product.Category)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = id
	return nil
	// TODO: Insert the product into the database
	// TODO: Update the product.ID with the database-generated ID
}

// GetProduct retrieves a product by ID
func (ps *ProductStore) GetProduct(id int64) (*Product, error) {
	row := ps.db.QueryRow("SELECT * FROM products WHERE id = ?", id)

    p := &Product{}

    err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity, &p.Category)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("product with ID %d not found", id)
        }

        return nil, err
    }

	return p, nil
}

// UpdateProduct updates an existing product
func (ps *ProductStore) UpdateProduct(product *Product) error {
	_, err := ps.db.Exec("UPDATE products SET name = ?, price = ?, quantity = ?, category = ? WHERE id = ?",
		product.Name, product.Price, product.Quantity, product.Category, product.ID)
	if err != nil {
		return err
	}
	// TODO: Update the product in the database
	// TODO: Return an error if the product doesn't exist
	return nil
}

// DeleteProduct removes a product by ID
func (ps *ProductStore) DeleteProduct(id int64) error {
	_, err := ps.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
	// TODO: Delete the product from the database
	// TODO: Return an error if the product doesn't exist
}

// ListProducts returns all products with optional filtering by category
func (ps *ProductStore) ListProducts(category string) ([]*Product, error) {
	var rows *sql.Rows
	var err error
	if category == "" {
		rows, err = ps.db.Query("SELECT * FROM products")
	} else {
		rows, err = ps.db.Query("SELECT * FROM products WHERE category = ?", category)
	}
	if err != nil {
		return nil, err
	}

	var products []*Product
	for rows.Next() {
		product := new(Product)
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil

	// TODO: Query the database for products
	// TODO: If category is not empty, filter by category
	// TODO: Return a slice of Product pointers
}

// BatchUpdateInventory updates the quantity of multiple products in a single transaction
func (ps *ProductStore) BatchUpdateInventory(updates map[int64]int) error {
	tx, err := ps.db.Begin()

	stmt, err := tx.Prepare("UPDATE products SET quantity = ? WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for id, val := range updates {
		res, err := stmt.Exec(val, id)
		if err != nil {
			tx.Rollback()
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			tx.Rollback()
			return err
		}

		if rowsAffected == 0 {
			tx.Rollback()
			return fmt.Errorf("product with ID %d not found", id)
		}
	}

	return tx.Commit()
	// TODO: Start a transaction
	// TODO: For each product ID in the updates map, update its quantity
	// TODO: If any update fails, roll back the transaction
	// TODO: Otherwise, commit the transaction
}

func main() {
	// Optional: you can write code here to test your implementation
}
