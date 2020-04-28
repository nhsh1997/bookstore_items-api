package services

import (
	"github.com/nhsh1997/bookstore_items-api/domain/items"
)
import "github.com/nhsh1997/bookstore_utils-go/rest_errors"

var(
	ItemService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestError)
	Get(string) (*items.Item, *rest_errors.RestError)
}

type itemService struct{}

func (s *itemService) Create(items.Item) (*items.Item, *rest_errors.RestError) {
	return nil, rest_errors.NewNotImplementedError("implement me")
}

func (s *itemService) Get(string) (*items.Item, *rest_errors.RestError) {
	return nil, rest_errors.NewNotImplementedError("implement me")
}
