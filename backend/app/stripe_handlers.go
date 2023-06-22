package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// CheckoutInput struct for data needed when checkout
type CheckoutInput struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Amount    int64 `json:"amount" binding:"required"`

	SuccessUrl string `json:"success_url" binding:"required"`
	FailedUrl  string `json:"failed_url" binding:"required"`
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

	prod, err := a.db.GetProduct(input.ProductID)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFound(errors.New("product is not found"))
	}

	if err != nil {
		log.Error().Err(err).Send()
		return nil, InternalServerError(errors.New(internalServerErrorMsg))
	}

	paramsCheckout := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(prod.PriceID),
				Quantity: stripe.Int64(input.Amount),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(input.SuccessUrl),
		CancelURL:  stripe.String(input.FailedUrl),
	}

	s, err := session.New(paramsCheckout)

	if err != nil {
		log.Error().Err(err).Send()
		return nil, InternalServerError(errors.New(internalServerErrorMsg))
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)

	return ResponseMsg{
		Message: "Redirected",
		Data:    nil,
	}, Ok()
}
