package orders

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/garden-raccoon/orders/models"
	proto "github.com/garden-raccoon/orders/protocols/orders"
)

// TODO need to set timeout via lib initialisation
// timeOut is  hardcoded GRPC requests timeout value
const timeOut = 60

// Debug on/off
var Debug = true

//

// IOrderAPI is
type IOrderAPI interface {
	GetOrders() ([]*models.Order, error)
	AllOrders(pb *proto.Orders) ([]*models.Order, error)

	OrderByTitle(title string) (*models.Order, error)

	CreateOrder(s *models.Order) error
	// Close GRPC Api connection
	Close() error
}

// Api is profile-service GRPC Api
// structure with client Connection
type Api struct {
	addr    string
	timeout time.Duration
	*grpc.ClientConn
	proto.OrderServiceClient
}

// New create new Battles Api instance
func New(addr string) (IOrderAPI, error) {
	api := &Api{timeout: timeOut * time.Second}

	if err := api.initConn(addr); err != nil {
		return nil, fmt.Errorf("create Battles Api:  %w", err)
	}

	api.OrderServiceClient = proto.NewOrderServiceClient(api.ClientConn)
	return api, nil
}
func (api *Api) AllOrders(pb *proto.Orders) ([]*models.Order, error) {
	ppp := models.OrdersFromProto(pb)
	return ppp, nil
}
func (api *Api) GetOrders() ([]*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	var resp *proto.Orders
	empty := new(proto.OrderEmpty)
	resp, err := api.OrderServiceClient.GetOrders(ctx, empty)
	if err != nil {
		return nil, fmt.Errorf("GetOrders api request: %w", err)
	}

	orders := models.OrdersFromProto(resp)

	if Debug {
		fmt.Println(orders)
	}
	return orders, nil
}
func (api *Api) CreateOrder(s *models.Order) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	order := &proto.Order{
		Title:       s.Title,
		Description: s.Description,
		Price:       float32(s.Price),
		Uuid:        s.UUID.Bytes(),
		Contain:     s.Contain,
	}
	// DEBUG Info
	if Debug {
		fmt.Println("order is")
		fmt.Println(order)

		if order == nil {
			fmt.Println("order is nil")
		} else {
			fmt.Println("order is not nil")
		}
		if api.OrderServiceClient == nil {
			fmt.Println("OrderServiceClient is nil")
		} else {
			fmt.Println("OrderServiceClient is not nil")
		}
	}
	//--- **** ---///

	_, err = api.OrderServiceClient.CreateOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("create order api request: %w", err)
	}
	return nil
}

// initConn initialize connection to Grpc servers
func (api *Api) initConn(addr string) (err error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	api.ClientConn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	return
}

// OrderByTitle is
func (api *Api) OrderByTitle(title string) (*models.Order, error) {
	getter := &proto.OrderGetter{
		Getter: &proto.OrderGetter_Title{
			Title: title,
		},
	}
	return api.getOrder(getter)
}

// getOrder is
func (api *Api) getOrder(getter *proto.OrderGetter) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	resp, err := api.OrderServiceClient.OrderByTitle(ctx, getter)
	if err != nil {
		return nil, fmt.Errorf("get battle api request: %w", err)
	}

	return models.OrderFromProto(resp), nil
}
