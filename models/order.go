package models

import (
	"github.com/gofrs/uuid"

	proto "github.com/garden-raccoon/orders/protocols/orders"
)

// Order is
type Order struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Uuid        uuid.UUID `json:"uuid"`
	UserUuid    uuid.UUID `json:"user-uuid"`
	Contain     []string  `json:"contain"`
	Quantity    int       `json:"quantity"`
	Day         string    `json:"day"`
}

func NewOrder(title, description, day string, price float64, quantity int, uuid, userUUID uuid.UUID, contain []string) *Order {
	return &Order{
		Title:       title,
		Description: description,
		Price:       price,
		Uuid:        uuid,
		UserUuid:    userUUID,
		Contain:     contain,
		Quantity:    quantity,
		Day:         day,
	}
}

// Proto is
func (o Order) Proto() *proto.Order {
	order := &proto.Order{
		Uuid:        o.Uuid.Bytes(),
		Description: o.Description,
		Contain:     o.Contain,
		Price:       float32(o.Price),
		Title:       o.Title,
	}
	return order
}

func OrderFromProto(pb *proto.Order) *Order {
	order := &Order{
		Title:       pb.Title,
		Description: pb.Description,
		Price:       float64(pb.Price),
		Contain:     pb.Contain,
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
