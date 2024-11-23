package models

import (
	"github.com/gofrs/uuid"

	proto "github.com/garden-raccoon/orders/protocols/orders"
)

type Orders struct {
	OrderUuid uuid.UUID `json:"order_uuid"`
	UserUuid  uuid.UUID `json:"user_uuid"`
	Order     []*Order
}

type DummyOrder struct {
	OrderUuid uuid.UUID `json:"order_uuid"`
	UserUuid  uuid.UUID `json:"user_uuid"`
}

// Order is
type Order struct {
	Name     string  `json:"title"`
	Price    float64 `json:"price"`
	MealType string  `json:"mealType"`
	Quantity int     `json:"quantity"`
	Day      string  `json:"day"`
}

func NewOrder(name, mealType, day string, price float64, quantity int) *Order {
	return &Order{
		Name:     name,
		Price:    price,
		MealType: mealType,
		Quantity: quantity,
		Day:      day,
	}
}
func ProtoDummy(o *DummyOrder) *proto.DummyOrder {
	return &proto.DummyOrder{
		OrderUuid: o.OrderUuid.Bytes(),
		UserUuid:  o.UserUuid.Bytes(),
	}
}
func DummyFromProto(pb *proto.DummyOrder) *DummyOrder {
	return &DummyOrder{
		OrderUuid: uuid.FromBytesOrNil(pb.OrderUuid),
		UserUuid:  uuid.FromBytesOrNil(pb.UserUuid),
	}
}

// Proto is
func Proto(o *Order) *proto.Order {
	order := &proto.Order{
		Name:     o.Name,
		MealType: o.MealType,
		Price:    float32(o.Price),
		Quantity: int32(o.Quantity),
		Day:      o.Day,
	}
	return order
}

func OrderFromProto(pb *proto.Order) *Order {
	order := &Order{
		Name:     pb.Name,
		Price:    float64(pb.Price),
		MealType: pb.MealType,

		Quantity: int(pb.Quantity),
		Day:      pb.Day,
	}
	return order
}

// OrdersToProto is
func OrdersToProto(orders *Orders) *proto.Orders {
	pb := &proto.Orders{}
	pb.OrderUuid = orders.OrderUuid.Bytes()
	pb.UserUuid = orders.UserUuid.Bytes()
	for _, b := range orders.Order {
		pb.Orders = append(pb.Orders, Proto(b))
	}
	return pb
}

// OrdersFromProto is
func OrdersFromProto(pb *proto.Orders) *Orders {
	orders := &Orders{}
	orders.OrderUuid = uuid.FromBytesOrNil(pb.OrderUuid)
	orders.UserUuid = uuid.FromBytesOrNil(pb.UserUuid)
	for _, b := range pb.Orders {
		orders.Order = append(orders.Order, OrderFromProto(b))
	}
	return orders
}
