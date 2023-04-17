package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go_restful_api/helper"
	"go_restful_api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
func main() {

	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
