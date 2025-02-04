package infraestructure

//LSP Liskov Substitution Principle (Principio de sustitución de Liskov)
//Se puede sustituir un tipo por otro si cumple con la misma interfaz
import (
	"database/sql"
	"errors"
	entities "main/src/Domain-negocio/Entities"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Save(product entities.Product) error {
	_, err := r.db.Exec("INSERT INTO products (name,quantity,codigo_barras) VALUES ($1, $2, $3)", product.Name, product.Quantity, product.CodigoBarras)
	return err
}

func (r *PostgresProductRepository) GetAll() ([]entities.Product, error) {
	rows, err := r.db.Query("SELECT id, name, quantity, codigo_barras FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.CodigoBarras); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *PostgresProductRepository) GetOne(id int) (entities.Product, error) {
	row := r.db.QueryRow("SELECT id,name, price, quantity FROM products WHERE id = $1", id)
	var product entities.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Quantity, &product.CodigoBarras); err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

// duda esto puede retornar un texto?
func (r *PostgresProductRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id =$1", id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("producto no encontrado")
	}
	return nil
}

func (r *PostgresProductRepository) Put(id int, product entities.Product) error {
	result, err := r.db.Exec("UPDATE products SET name = $1 , quantity = $2, codigo_barras = $3 WHERE id = $4", product.Name, product.Quantity, product.CodigoBarras, id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("producto no encontrado")
	}
	return nil
}
