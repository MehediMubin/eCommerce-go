package cmd

import (
	"ecommerce-go/global_router"
	"ecommerce-go/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server running on :3000")

	globalRoute := global_router.GlobalRoute(mux)

	err := http.ListenAndServe(":3000", globalRoute)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}