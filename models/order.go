package models

import (
	"github.com/gofrs/uuid"
	"github.com/misnaged/scriptorium/logger"

	proto "github.com/garden-raccoon/orders/protocols/orders"
)

// Order is
type Order struct {
	UUID        uuid.UUID
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Contain     []string `json:"contain"`
}

func NewOrder(title, description string, price float64, uuid uuid.UUID, contain []string) *Order {
	return &Order{
		UUID:        uuid,
		Title:       title,
		Description: description,
		Price:       price,
		Contain:     contain,
	}
}

// Proto is
func (b Order) Proto() *proto.Order {
	order := &proto.Order{
		Uuid:        b.UUID.Bytes(),
		Description: b.Description,
		Contain:     b.Contain,
		Price:       float32(b.Price),
		Title:       b.Title,
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
	order.UUID = uuid.FromBytesOrNil(pb.Uuid)
	if order.UUID == uuid.Nil {
		logger.Log().Println("NIL!")
		order.UUID, _ = uuid.NewV4() //todo: REFACTOR!
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
