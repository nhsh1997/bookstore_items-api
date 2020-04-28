package app

import (
	"github.com/nhsh1997/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls()  {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}