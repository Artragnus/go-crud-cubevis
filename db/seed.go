package db

import (
	"context"
	"database/sql"
)

func Seed(dataSourceName string) error {
	var products = []CreateProductParams{
		{Name: "Smartphone X", Value: 2999},
		{Name: "Notebook Pro", Value: 5499},
		{Name: "Fone Bluetooth", Value: 249},
		{Name: "Smartwatch Fit", Value: 799},
		{Name: "Tablet Max", Value: 1599},
		{Name: "Câmera Pro", Value: 3499},
		{Name: "Monitor 4K", Value: 1899},
		{Name: "Teclado Gamer", Value: 499},
		{Name: "Mouse Sem Fio", Value: 199},
		{Name: "Caixa de Som Bluetooth", Value: 349},
	}

	conn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	defer conn.Close()

	q := New(conn)

	result, err := q.GetProducts(context.Background())

	if err != nil {
		return err
	}

	if len(result) == 0 {
		for _, p := range products {
			err := q.CreateProduct(context.Background(), p)
			if err != nil {
				return err
			}
		}
	}

	return nil

}
