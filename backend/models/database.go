// Package models for database models
package models

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

// DB struct hold db instance
type DB struct {
	db *gorm.DB
}

// NewDB creates new DB
func NewDB() DB {
	return DB{}
}

// Connect connects to database file
func (d *DB) Connect(file string) error {
	gormDB, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = gormDB
	return nil
}

// Migrate migrates db schema
func (d *DB) Migrate() error {
	return d.db.AutoMigrate(&Product{})
}

// CreateProduct creates a new product
func (d *DB) CreateProduct(p *Product) error {
	params := &stripe.ProductParams{Name: stripe.String(p.Name)}
	prod, err := product.New(params)
	if err != nil {
		return err
	}

	paramsPrice := &stripe.PriceParams{
		Product:    stripe.String(prod.ID),
		UnitAmount: stripe.Int64(p.Price),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
	}

	priceObj, err := price.New(paramsPrice)
	if err != nil {
		return err
	}

	p.PriceID = priceObj.ID
	return d.db.Create(&p).Error
}

// GetProduct returns a product by its id
func (d *DB) GetProduct(id int64) (Product, error) {
	var res Product
	query := d.db.First(&res, "id = ?", id)
	return res, query.Error
}

// ListProducts returns all products
func (d *DB) ListProducts() ([]Product, error) {
	var products []Product
	err := d.db.Find(&products).Error
	return products, err
}

// BuyProduct adds amount bought of a product
func (d *DB) BuyProduct(id int64, amount int64) error {
	var res Product
	result := d.db.Model(&res).Where("id = ?", id).Set("amount = amount + ?", amount)
	return result.Error
}

// SellProduct subtracts amount sold from a product
func (d *DB) SellProduct(id int64, amount int64) error {
	var res Product
	result := d.db.Model(&res).Where("id = ?", id).Set("amount = amount - ?", amount)
	return result.Error
}
