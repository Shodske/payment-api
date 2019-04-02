package model

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/satori/go.uuid"
	"google.golang.org/genproto/googleapis/type/money"
)

// AccountType type and constants.
type AccountType int

const (
	BasicAccount AccountType = iota
	PremiumAccount
)

// Struct representation of a payment. Instances of this struct can be marshaled to a json resource according to the
// json:api specification.
type Payment struct {
	Model `json:"-"`

	OrganisationID uuid.UUID    `json:"-" gorm:"type:uuid REFERENCES organisations(id)"`
	Organisation   Organisation `json:"-" gorm:"association_autoupdate:false"`

	Amount               string `json:"amount,omitempty"`
	Currency             string `json:"currency,omitempty"`
	EndToEndReference    string `json:"end_to_end_reference,omitempty"`
	NumericReference     string `json:"numeric_reference,omitempty"`
	PaymentId            string `json:"payment_id,omitempty"`
	PaymentPurpose       string `json:"payment_purpose,omitempty"`
	PaymentScheme        string `json:"payment_scheme,omitempty"`
	PaymentType          string `json:"payment_type,omitempty"`
	ProcessingDate       string `json:"processing_date,omitempty"`
	Reference            string `json:"reference,omitempty"`
	SchemePaymentSubType string `json:"scheme_payment_sub_type,omitempty"`
	SchemePaymentType    string `json:"scheme_payment_type,omitempty"`

	BeneficiaryPartyID sql.NullInt64 `json:"-" gorm:"type:integer REFERENCES parties(id)"`
	BeneficiaryParty   *Party        `json:"beneficiary_party,omitempty"`

	DebtorPartyID sql.NullInt64 `json:"-" gorm:"type:integer REFERENCES parties(id)"`
	DebtorParty   *Party        `json:"debtor_party,omitempty"`

	SponsorPartyID sql.NullInt64 `json:"-" gorm:"type:integer REFERENCES parties(id)"`
	SponsorParty   *Party        `json:"sponsor_party,omitempty"`

	ChargesInformationID sql.NullInt64 `json:"-" gorm:"type:integer REFERENCES charges(id)"`
	ChargesInformation   *Charge       `json:"charges_information,omitempty"`

	FXID sql.NullInt64 `json:"-" sql:"type:integer REFERENCES fxes(id)"`
	FX   *FX           `json:"fx,omitempty"`
}

type Party struct {
	ID                uint        `json:"-" gorm:"primary_key"`
	AccountName       string      `json:"account_name,omitempty"`
	AccountNumber     string      `json:"account_number,omitempty"`
	AccountNumberCode string      `json:"account_number_code,omitempty"`
	AccountType       AccountType `json:"account_type,omitempty"`
	Address           string      `json:"address,omitempty"`
	BankID            string      `json:"bank_id,omitempty"`
	BankIDCode        string      `json:"bank_id_code,omitempty"`
	Name              string      `json:"name,omitempty"`
}

type Charge struct {
	ID                      uint   `json:"-" gorm:"primary_key"`
	BearerCode              string `json:"bearer_code,omitempty"`
	ReceiverChargesAmount   string `json:"receiver_charges_amount,omitempty"`
	ReceiverChargesCurrency string `json:"receiver_charges_currency,omitempty"`
	SenderCharges           []struct {
		Amount   string `json:"amount,omitempty"`
		Currency string `json:"currency,omitempty"`
	} `json:"sender_charges,omitempty" gorm:"type:json"`
}

type FX struct {
	ID                uint        `json:"-" gorm:"primary_key"`
	ContractReference string      `json:"contract_reference,omitempty"`
	ExchangeRate      string      `json:"exchange_rate,omitempty"`
	OriginalValue     money.Money `json:"original_value,omitempty"`
}

// Method required to implement `jsonapi.UnmarshalToOneRelations`, which we need to set the organisation relationship.
func (payment *Payment) SetToOneReferenceID(name, ID string) error {
	id, err := uuid.FromString(ID)
	if err != nil {
		return err
	}

	switch name {
	case "organisation":
		payment.OrganisationID = id
	default:
		return errors.New(fmt.Sprintf("invalid relationship name `%s`", name))
	}

	return nil
}

// Method required to implement `jsonapi.MarshalReferences`.
func (payment *Payment) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Name:         "organisation",
			Type:         "organisations",
			IsNotLoaded:  false,
			Relationship: jsonapi.ToOneRelationship,
		},
	}
}

// Method required to implement `jsonapi.MarshalLinkedRelations`
func (payment *Payment) GetReferencedIDs() []jsonapi.ReferenceID {
	return []jsonapi.ReferenceID{
		{
			Name:         "organisation",
			Type:         "organisations",
			Relationship: jsonapi.ToOneRelationship,
			ID:           payment.OrganisationID.String(),
		},
	}
}
