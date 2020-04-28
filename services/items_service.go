package services

import "github.com/nhsh1997/bookstore_items-api/domain/items"
import "github.com/nhsh1997/bookstore_utils-go/rest_errors"

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestError)
}
