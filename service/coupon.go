package service

import "context"

type coupon struct{}

var Coupon coupon

func (c coupon) Bind(ctx context.Context) error {
	return nil
}

func (c coupon) Unbind(ctx context.Context) error {
	return nil
}

func (c coupon) Freeze(ctx context.Context, coupon int64) error {
	return nil
}

func (c coupon) Unfreeze(ctx context.Context, coupon int64) error {
	return nil
}

func (c coupon) Detail() {

}

func (c coupon) List() {

}
