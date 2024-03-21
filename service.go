package main

import (
	"database/sql"
	"github.com/benebobaa/simple-native-crud/helper"
	"net/http"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateProduct(writer http.ResponseWriter, request *http.Request) {

	//Read request body
	productRequest := Product{}
	err := helper.ReadFromRequestBody(request, &productRequest)
	if err != nil {
		helper.WriteErrToResponseBody(writer, err)
		return
	}

	//Query to insert data
	SQL := `INSERT INTO "products" (name, price, stock) VALUES ($1, $2, $3) RETURNING id`
	err = s.db.QueryRow(SQL, productRequest.Name, productRequest.Price, productRequest.Stock).Scan(&productRequest.ID)
	if err != nil {
		helper.WriteErrToResponseBody(writer, err)
		return
	}

	//Write response
	helper.WriteToResponseBody(writer, productRequest)
}

func (s *Service) GetProducts(writer http.ResponseWriter, request *http.Request) {

	// Query to get all products
	SQL := `SELECT id, name, price, stock FROM "products"`
	rows, err := s.db.Query(SQL)
	if err != nil {
		helper.WriteErrToResponseBody(writer, err)
		return
	}
	defer rows.Close()

	// Iterate over the rows, appending each product to a slice
	products := []Product{}
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			helper.WriteErrToResponseBody(writer, err)
			return
		}
		products = append(products, product)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		helper.WriteErrToResponseBody(writer, err)
		return
	}

	// Write the products to the response body
	helper.WriteToResponseBody(writer, products)
}
