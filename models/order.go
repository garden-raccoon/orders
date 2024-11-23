package models

import (
	"github.com/gofrs/uuid"

	proto "github.com/garden-raccoon/orders/protocols/orders"
)

type Orders struct {
	Order []*Order
}

// Order is
type Order struct {
	Name      string    `json:"title"`
	Price     float64   `json:"price"`
	MealType  string    `json:"mealType"`
	OrderUuid uuid.UUID `json:"orderUuid"`
	UserUuid  uuid.UUID `json:"user_uuid"`
	Quantity  int       `json:"quantity"`
	Day       string    `json:"day"`
}

func NewOrder(name, mealType, day string, price float64, quantity int, orderUUID, userUUID uuid.UUID) *Order {
	return &Order{
		Name:      name,
		Price:     price,
		OrderUuid: orderUUID,
		UserUuid:  userUUID,
		MealType:  mealType,
		Quantity:  quantity,
		Day:       day,
	}
}

// Proto is
func (o Order) Proto() *proto.Order {
	order := &proto.Order{
		OrderUuid: o.OrderUuid.Bytes(),
		UserUuid:  o.UserUuid.Bytes(),
		Name:      o.Name,
		MealType:  o.MealType,
		Price:     float32(o.Price),
		Quantity:  int32(o.Quantity),
		Day:       o.Day,
	}
	return order
}

func OrderFromProto(pb *proto.Order) *Order {
	order := &Order{
		Name:      pb.Name,
		Price:     float64(pb.Price),
		MealType:  pb.MealType,
		OrderUuid: uuid.FromBytesOrNil(pb.OrderUuid),
		UserUuid:  uuid.FromBytesOrNil(pb.UserUuid),
		Quantity:  int(pb.Quantity),
		Day:       pb.Day,
	}
	return order
}

// OrdersToProto is
func OrdersToProto(orders []*Order) (pb *proto.Orders) {
	for _, b := range orders {
		pb.Orders = append(pb.Orders, b.Proto())
	}
	return
}

// OrdersFromProto is
func OrdersFromProto(pb *proto.Orders) (shops []*Order) {
	for _, b := range pb.Orders {
		shops = append(shops, OrderFromProto(b))
	}
	return
}
