package model

type(
	 Product struct{
		ID		string  `json:"id"`
		Price		float32 `json:"price"`
		ProductName	string  `json:"product_name"`
		QteStock	int     `json:"qte_stock"`
		ProductGambar	string  `json:"product_gambar"`	
	}
)