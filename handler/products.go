package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/priyajithc/product-svc/data"
)

type Products struct {
	l *log.Logger
}

func ProductHandler() *Products {
	l := log.New(os.Stdout, "products-svc ", log.LstdFlags)
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Product from ProductHandler...")
	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Product by Category from ProductHandler...")
	// fetch the products from the datastore

	catId := mux.Vars(r)["id"]
	lp := data.GetProductsByCategory(catId)

	p.l.Println(catId)
	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
