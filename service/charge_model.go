package service

import (
	"time"
)

type statusCharge string


type taxID struct {
	taxID string
	Type string `json:"type"`
}
type customer struct {
	name string
	email string
	phone string
	taxID taxID
	correlationID string
}

type params struct {
	start *time.Time
	end *time.Time
	status *string
}

type charge struct {
	value *uint32
	customer *customer
	comment *string
	brCode *string
	status *string
	correlationID *string
	paymentLinkID *string
	paymentLinkUrl *interface{}
	globalID *interface{}
	transactionID *interface{}
	identifier *string
	qrCodeImage *interface{}
	additionalInfo *[]map[string]interface{}
	pixKey *string
	createdAt *string
	updatedAt *string
	expiresIn *string
	
}