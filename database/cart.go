package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("cant find product")
	ErrCantDecodeProducts = errors.New("cant decode products")
	ErrCantUpdateUser     = errors.New("cant update user")
	ErrUserIdNotValid     = errors.New("user id not valid")
	ErrCantFindUser       = errors.New("cant find user")
	ErrCantGetItem        = errors.New("cant get items")
	ErrCantBuyItemFromCart = errors.New("cant buy item from cart")
)

func AddProductToCart()   {}

func RemoveItemFromCart() {}

func BuyItemFromCart() {}

func InstantBuyer() {}
