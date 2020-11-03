package data

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

// ProductList is a structure that holds a collection of Product
type ProductList struct {
	Products []Product `json:"products"`
}

// in memory cache of products
var productList ProductList

// GetProducts returns a list of products
func GetProducts() ProductList {
	if len(productList.Products) == 0 {
		productList = getProductList()
	}
	return productList
}

// Reads the product details from a JSON file
// GetProducts returns a list of products
func GetProductsByCategory(catId string) ProductList {
	var catpl ProductList
	pl := GetProducts()
	for _, p := range pl.Products {
		if p.Category == catId {
			catpl.Products = append(catpl.Products, p)
		}
	}

	return catpl
}

func getProductList() ProductList {
	dataPath := os.Getenv("PRODUCTS_JSON_PATH")
	fmt.Println(dataPath)
	file, err := ioutil.ReadFile(dataPath)
	if err != nil {
		fmt.Println(err)
	}
	data := ProductList{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

// Encode the products to a JSON format
func (p ProductList) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
