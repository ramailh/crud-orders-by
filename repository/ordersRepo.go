package repository

import "crud-gorm-gin/model"

type OrdersRepo interface {
	Create(param model.Order) (model.Order, error)
	Get() ([]model.Order, error)
	Update(param model.Order) (model.Order, error)
	Delete(id int64) error
}
