package query

import (
	"BELAJAR-GOLANG/graphql/schema/model"
	"hello/graphql/config"

	"github.com/graphql-go/graphql"
)

func ProductResolve(param graphql.ResolveParams) (interface{}, error){
	var a model.Product
	var b []model.Product
	
	db, err := config.GetConnection()
	if err != nil{
		panic(err.Error())
	}
	b = b[:0]
	result, err := db.Query("SELECT id, product_name, qte_stock, ifnull(product_gambar,'') FROM Products")
	if err != nil{
		panic(err.Error())
	}

	for result.Next(){
		err = result.Scan(&a.ID, &a.ProductName, &a.QteStock, &a.ProductGambar)
		if err != nil{
			panic(err.Error())
		}
		b = append(b, a)
	}
	return b, nil
}