package controllers

import (
	"fmt"
	"github.com/bookstore_oauth-go/oauth"
	"github.com/nhsh1997/bookstore_items-api/domain/items"
	"github.com/nhsh1997/bookstore_items-api/services"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: return error to the caller
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemService.Create(item)

	if err != nil {
		//TODO: Return error json to the user
	}

	fmt.Println(result)
	//TODO: Return created item as json with HTTP status 201 - Created
}

func Get(w http.ResponseWriter, r *http.Request) {

}
