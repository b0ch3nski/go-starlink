package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/b0ch3nski/go-starlink/model/device"
)

const DefaultDishyAddr = "192.168.100.1:9200"

type Client interface {
	Status(context.Context) (*device.DishGetStatusResponse, error)
	Unstow(context.Context) error
	Stow(context.Context) error
	Reboot(context.Context) error
}

var _ Client = (*client)(nil)

type client struct {
	conn    *grpc.ClientConn
	dcl     device.DeviceClient
	timeout time.Duration
}

// NewClient creates new Starlink client.
func NewClient(ctx context.Context, addr string) (*client, error) {
	defaultTimeout := 5 * time.Second

	ctxDial, cancelDial := context.WithTimeout(ctx, defaultTimeout)
	defer cancelDial()

	conn, errDial := grpc.DialContext(ctxDial, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if errDial != nil {
		return nil, fmt.Errorf("failed creating gRPC connection to Starlink Dish: %w", errDial)
	}
	return &client{conn: conn, dcl: device.NewDeviceClient(conn), timeout: defaultTimeout}, nil
}

func (c *client) WithTimeout(timeout time.Duration) *client {
	c.timeout = timeout
	return c
}

func (c *client) do(ctx context.Context, req *device.Request) (*device.Response, error) {
	ctxReq, cancelReq := context.WithTimeout(ctx, c.timeout)
	defer cancelReq()

	resp, err := c.dcl.Handle(ctxReq, req)
	if err != nil {
		return resp, fmt.Errorf("failed executing request[%+v] to Starlink Dish: %w", req.Request, err)
	}
	return resp, nil
}

func (c *client) Status(ctx context.Context) (*device.DishGetStatusResponse, error) {
	req := &device.Request{Request: &device.Request_GetStatus{}}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetDishGetStatus(), nil
}

func (c *client) unstow(ctx context.Context, unstow bool) error {
	req := &device.Request{
		Request: &device.Request_DishStow{
			DishStow: &device.DishStowRequest{
				Unstow: unstow,
			},
		},
	}
	_, err := c.do(ctx, req)
	return err
}

func (c *client) Unstow(ctx context.Context) error {
	return c.unstow(ctx, true)
}

func (c *client) Stow(ctx context.Context) error {
	return c.unstow(ctx, false)
}

func (c *client) Reboot(ctx context.Context) error {
	req := &device.Request{Request: &device.Request_Reboot{}}
	_, err := c.do(ctx, req)
	return err
}
