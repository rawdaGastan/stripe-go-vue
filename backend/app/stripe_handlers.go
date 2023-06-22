package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rawdaGastan/stripe-go-vue/models"
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// CheckoutInput struct for data needed when checkout
type CheckoutInput struct {
	Cart []SellInput `json:"cart" binding:"required"`

	SuccessUrl string `json:"success_url" binding:"required"`
	FailedUrl  string `json:"failure_url" binding:"required"`
}

// SellInput struct for data needed when selling products
type SellInput struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Amount    int64 `json:"amount" binding:"required"`
}

// CreateInput struct for data needed when creating a new product
type CreateInput struct {
	Name     string          `json:"name" binding:"required"`
	Price    int64           `json:"price" binding:"required"`
	Currency models.Currency `json:"currency" binding:"required"`
	Amount   int64           `json:"amount" binding:"required"`
}

func (a *App) createCheckoutSession(r *http.Request, w http.ResponseWriter) (interface{}, Response) {
	var input CheckoutInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, BadRequest(errors.New("failed to read input data"))
	}

	err = validator.Validate(input)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, BadRequest(errors.New("invalid input data"))
	}

	var stripeItems []*stripe.CheckoutSessionLineItemParams
	for _, item := range input.Cart {
		prod, err := a.db.GetProduct(item.ProductID)
		if err == gorm.ErrRecordNotFound {
			return nil, NotFound(errors.New("product is not found"))
		}

		if err != nil {
			log.Error().Err(err).Send()
			return nil, InternalServerError(errors.New(internalServerErrorMsg))
		}

		err = a.db.SellProduct(item.ProductID, item.Amount)
		if err != nil {
			log.Error().Err(err).Send()
			return nil, InternalServerError(errors.New(internalServerErrorMsg))
		}

		stripeItem := stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(prod.PriceID),
			Quantity: stripe.Int64(item.Amount),
		}

		stripeItems = append(stripeItems,
			&stripeItem,
		)
	}

	paramsCheckout := &stripe.CheckoutSessionParams{
		LineItems:  stripeItems,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(input.SuccessUrl),
		CancelURL:  stripe.String(input.FailedUrl),
	}

	s, err := session.New(paramsCheckout)

	if err != nil {
		log.Error().Err(err).Send()
		return nil, InternalServerError(errors.New(internalServerErrorMsg))
	}

	return ResponseMsg{
		Message: "Redirect",
		Data:    s.URL,
	}, Ok()
}

func (a *App) getProducts(r *http.Request, w http.ResponseWriter) (interface{}, Response) {
	prods, err := a.db.ListProducts()
	if err == gorm.ErrRecordNotFound {
		return ResponseMsg{
			Message: "No products found",
			Data:    prods,
		}, Ok()
	}

	if err != nil {
		log.Error().Err(err).Send()
		return nil, InternalServerError(errors.New(internalServerErrorMsg))
	}

	return ResponseMsg{
		Message: "Products found",
		Data:    prods,
	}, Ok()
}

func (a *App) createProduct(r *http.Request, w http.ResponseWriter) (interface{}, Response) {
	var input CreateInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Error().Err(err).Send()
		return nil, BadRequest(errors.New("failed to read input data"))
	}

	err = validator.Validate(input)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, BadRequest(errors.New("invalid input data"))
	}

	prod := models.Product{
		Name:     input.Name,
		Price:    input.Price,
		Currency: input.Currency,
		Amount:   input.Amount,
	}

	err = a.db.CreateProduct(&prod)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, InternalServerError(errors.New(internalServerErrorMsg))
	}

	return ResponseMsg{
		Message: "Product is created successfully",
		Data:    nil,
	}, Created()
}

func (a *App) sellProduct(r *http.Request, w http.ResponseWriter) (interface{}, Response) {
	var input SellInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Error().Err(err).Send()
		return nil, BadRequest(errors.New("failed to read input data"))
	}

	err = validator.Validate(input)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, BadRequest(errors.New("invalid input data"))
	}

	err = a.db.SellProduct(input.ProductID, input.Amount)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, InternalServerError(errors.New(internalServerErrorMsg))
	}

	return ResponseMsg{
		Message: "Product is sold successfully",
		Data:    nil,
	}, Ok()
}
