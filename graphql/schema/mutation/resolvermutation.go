package mutation

import (
	"hello/graphql/config"

	"github.com/graphql-go/graphql"
)

func CreateProdcutMutation(param graphql.ResolveParams) (interface{}, error) {
	id := param.Args["id"].(string)
	price := param.Args["price"].(float32)
	productname := param.Args["product_name"].(string)
	qty := param.Args["qte_stock"].(int)
	productgambar := param.Args["product_gambar"].(string)

	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Query("INSERT INTO toko.product values (?,?,?,?,?)", id, price, productname, qty, productgambar)
	if err != nil {
		panic(err.Error())
	}

	return nil, err
}