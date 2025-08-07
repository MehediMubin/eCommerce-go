package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgUrl string `json:"imageUrl"`
}

var productList []Product

func NewProduct(id int, title, description string, price float64, imgUrl string) *Product {
	return &Product{
		ID: id,
		Title: title,
		Description: description,
		Price: price,
		ImgUrl: imgUrl,
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from eCommerce Project")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func sendData(w http.ResponseWriter, data any, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(rootHandler))

	mux.Handle("GET /products", http.HandlerFunc(getProducts))

	mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))

	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	product1 := NewProduct(1, "Orange", "Orange is orange", 3.23, "https://www.dole.com/sites/default/files/media/2025-01/oranges.png")

	product2 := NewProduct(2, "Apple", "I love Apple", 2.23, "https://5.imimg.com/data5/AK/RA/MY-68428614/apple.jpg")

	product3 := NewProduct(3, "Banana", "Banana is interesting", 4.23, "https://png.pngtree.com/png-clipart/20220716/ourmid/pngtree-banana-yellow-fruit-banana-skewers-png-image_5944324.png")

	product4 := NewProduct(4, "Grape", "Grape is sour", 10.23, "https://hips.hearstapps.com/hmg-prod/images/766/grapes-main-1506688521.jpg?resize=640:*")
	
	product5 := NewProduct(5, "Mango", "Mango is fruit ka raja", 6.23, "https://5.imimg.com/data5/SELLER/Default/2023/9/344928632/OW/RQ/XC/25352890/yellow-mango-500x500.jpeg")

	productList = append(productList, *product1, *product2, *product3, *product4, *product5)
}