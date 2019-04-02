package model

import (
	"database/sql"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/satori/go.uuid"
	"reflect"
	"testing"
)

func TestPayment_SetToOneReferenceID(t *testing.T) {
	type fields struct {
		Model                Model
		OrganisationID       uuid.UUID
		Organisation         Organisation
		Amount               string
		Currency             string
		EndToEndReference    string
		NumericReference     string
		PaymentId            string
		PaymentPurpose       string
		PaymentScheme        string
		PaymentType          string
		ProcessingDate       string
		Reference            string
		SchemePaymentSubType string
		SchemePaymentType    string
		BeneficiaryPartyID   sql.NullInt64
		BeneficiaryParty     *Party
		DebtorPartyID        sql.NullInt64
		DebtorParty          *Party
		SponsorPartyID       sql.NullInt64
		SponsorParty         *Party
		ChargesInformationID sql.NullInt64
		ChargesInformation   *Charge
		FXID                 sql.NullInt64
		FX                   *FX
	}
	type args struct {
		name string
		ID   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base", fields{}, args{"organisation", "01234567-0123-0123-0123-0123456789ab"}, false},
		{"existing-relationship", fields{OrganisationID: uuid.NewV4()}, args{"organisation", "01234567-0123-0123-0123-0123456789ab"}, false},
		{"invalid-id", fields{}, args{"organisation", "01234567-0123-0123-0123-0123456789abc"}, true},
		{"invalid-name", fields{}, args{"not-a-relationship", "01234567-0123-0123-0123-0123456789ab"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payment := &Payment{
				Model:                tt.fields.Model,
				OrganisationID:       tt.fields.OrganisationID,
				Organisation:         tt.fields.Organisation,
				Amount:               tt.fields.Amount,
				Currency:             tt.fields.Currency,
				EndToEndReference:    tt.fields.EndToEndReference,
				NumericReference:     tt.fields.NumericReference,
				PaymentId:            tt.fields.PaymentId,
				PaymentPurpose:       tt.fields.PaymentPurpose,
				PaymentScheme:        tt.fields.PaymentScheme,
				PaymentType:          tt.fields.PaymentType,
				ProcessingDate:       tt.fields.ProcessingDate,
				Reference:            tt.fields.Reference,
				SchemePaymentSubType: tt.fields.SchemePaymentSubType,
				SchemePaymentType:    tt.fields.SchemePaymentType,
				BeneficiaryPartyID:   tt.fields.BeneficiaryPartyID,
				BeneficiaryParty:     tt.fields.BeneficiaryParty,
				DebtorPartyID:        tt.fields.DebtorPartyID,
				DebtorParty:          tt.fields.DebtorParty,
				SponsorPartyID:       tt.fields.SponsorPartyID,
				SponsorParty:         tt.fields.SponsorParty,
				ChargesInformationID: tt.fields.ChargesInformationID,
				ChargesInformation:   tt.fields.ChargesInformation,
				FXID:                 tt.fields.FXID,
				FX:                   tt.fields.FX,
			}
			if err := payment.SetToOneReferenceID(tt.args.name, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Payment.SetToOneReferenceID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPayment_GetReferences(t *testing.T) {
	baseID := uuid.NewV4()
	baseOrgID := uuid.NewV4()
	baseRef := []jsonapi.Reference{
		{
			"organisations",
			"organisation",
			false,
			jsonapi.ToOneRelationship,
		},
	}

	type fields struct {
		Model                Model
		OrganisationID       uuid.UUID
		Organisation         Organisation
		Amount               string
		Currency             string
		EndToEndReference    string
		NumericReference     string
		PaymentId            string
		PaymentPurpose       string
		PaymentScheme        string
		PaymentType          string
		ProcessingDate       string
		Reference            string
		SchemePaymentSubType string
		SchemePaymentType    string
		BeneficiaryPartyID   sql.NullInt64
		BeneficiaryParty     *Party
		DebtorPartyID        sql.NullInt64
		DebtorParty          *Party
		SponsorPartyID       sql.NullInt64
		SponsorParty         *Party
		ChargesInformationID sql.NullInt64
		ChargesInformation   *Charge
		FXID                 sql.NullInt64
		FX                   *FX
	}
	tests := []struct {
		name   string
		fields fields
		want   []jsonapi.Reference
	}{
		{"base", fields{Model: Model{ID: baseID}, OrganisationID: baseOrgID}, baseRef},
		{"empty", fields{}, baseRef},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payment := &Payment{
				Model:                tt.fields.Model,
				OrganisationID:       tt.fields.OrganisationID,
				Organisation:         tt.fields.Organisation,
				Amount:               tt.fields.Amount,
				Currency:             tt.fields.Currency,
				EndToEndReference:    tt.fields.EndToEndReference,
				NumericReference:     tt.fields.NumericReference,
				PaymentId:            tt.fields.PaymentId,
				PaymentPurpose:       tt.fields.PaymentPurpose,
				PaymentScheme:        tt.fields.PaymentScheme,
				PaymentType:          tt.fields.PaymentType,
				ProcessingDate:       tt.fields.ProcessingDate,
				Reference:            tt.fields.Reference,
				SchemePaymentSubType: tt.fields.SchemePaymentSubType,
				SchemePaymentType:    tt.fields.SchemePaymentType,
				BeneficiaryPartyID:   tt.fields.BeneficiaryPartyID,
				BeneficiaryParty:     tt.fields.BeneficiaryParty,
				DebtorPartyID:        tt.fields.DebtorPartyID,
				DebtorParty:          tt.fields.DebtorParty,
				SponsorPartyID:       tt.fields.SponsorPartyID,
				SponsorParty:         tt.fields.SponsorParty,
				ChargesInformationID: tt.fields.ChargesInformationID,
				ChargesInformation:   tt.fields.ChargesInformation,
				FXID:                 tt.fields.FXID,
				FX:                   tt.fields.FX,
			}
			if got := payment.GetReferences(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payment.GetReferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayment_GetReferencedIDs(t *testing.T) {
	baseID := uuid.NewV4()
	baseOrgID := uuid.NewV4()
	baseRef := []jsonapi.ReferenceID{
		{
			baseOrgID.String(),
			"organisations",
			"organisation",
			jsonapi.ToOneRelationship,
		},
	}
	emptyRef := []jsonapi.ReferenceID{}

	type fields struct {
		Model                Model
		OrganisationID       uuid.UUID
		Organisation         Organisation
		Amount               string
		Currency             string
		EndToEndReference    string
		NumericReference     string
		PaymentId            string
		PaymentPurpose       string
		PaymentScheme        string
		PaymentType          string
		ProcessingDate       string
		Reference            string
		SchemePaymentSubType string
		SchemePaymentType    string
		BeneficiaryPartyID   sql.NullInt64
		BeneficiaryParty     *Party
		DebtorPartyID        sql.NullInt64
		DebtorParty          *Party
		SponsorPartyID       sql.NullInt64
		SponsorParty         *Party
		ChargesInformationID sql.NullInt64
		ChargesInformation   *Charge
		FXID                 sql.NullInt64
		FX                   *FX
	}
	tests := []struct {
		name   string
		fields fields
		want   []jsonapi.ReferenceID
	}{
		{"base", fields{Model: Model{ID: baseID}, OrganisationID: baseOrgID}, baseRef},
		{"empty", fields{}, emptyRef},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payment := &Payment{
				Model:                tt.fields.Model,
				OrganisationID:       tt.fields.OrganisationID,
				Organisation:         tt.fields.Organisation,
				Amount:               tt.fields.Amount,
				Currency:             tt.fields.Currency,
				EndToEndReference:    tt.fields.EndToEndReference,
				NumericReference:     tt.fields.NumericReference,
				PaymentId:            tt.fields.PaymentId,
				PaymentPurpose:       tt.fields.PaymentPurpose,
				PaymentScheme:        tt.fields.PaymentScheme,
				PaymentType:          tt.fields.PaymentType,
				ProcessingDate:       tt.fields.ProcessingDate,
				Reference:            tt.fields.Reference,
				SchemePaymentSubType: tt.fields.SchemePaymentSubType,
				SchemePaymentType:    tt.fields.SchemePaymentType,
				BeneficiaryPartyID:   tt.fields.BeneficiaryPartyID,
				BeneficiaryParty:     tt.fields.BeneficiaryParty,
				DebtorPartyID:        tt.fields.DebtorPartyID,
				DebtorParty:          tt.fields.DebtorParty,
				SponsorPartyID:       tt.fields.SponsorPartyID,
				SponsorParty:         tt.fields.SponsorParty,
				ChargesInformationID: tt.fields.ChargesInformationID,
				ChargesInformation:   tt.fields.ChargesInformation,
				FXID:                 tt.fields.FXID,
				FX:                   tt.fields.FX,
			}
			if got := payment.GetReferencedIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payment.GetReferencedIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
