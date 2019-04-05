package source

import (
	"github.com/Shodske/payment-api/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

var organisationFixtures = []*model.Organisation{
	{
		Model: model.Model{ID: uuid.NewV4()},
		Name:  "Test Organisation A",
	},
	{
		Model: model.Model{ID: uuid.NewV4()},
		Name:  "Test Organisation B",
	},
	{
		Model: model.Model{ID: uuid.NewV4(), DeletedAt: &time.Time{}},
		Name:  "Deleted Test Organisation",
	},
}

// TODO Find out why the nested structs cause errors while setting up the test database. This may have something to do
//      with sqlite as these errors don't seem to appear with postgres.
var paymentFixtures = []*model.Payment{
	{
		Model: model.Model{ID: uuid.NewV4()},

		OrganisationID: organisationFixtures[0].ID,

		Amount:               "13.37",
		Currency:             "GBP",
		EndToEndReference:    "Some reference A",
		NumericReference:     "90011337420",
		PaymentID:            "01234659876",
		PaymentPurpose:       "Just paying",
		PaymentScheme:        "FPS",
		PaymentType:          "Credit",
		ProcessingDate:       "2017-01-18",
		Reference:            "Another reference",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",

		//BeneficiaryParty: &model.Party{
		//	AccountName:       "B Gates",
		//	AccountNumber:     "333333333",
		//	AccountNumberCode: "BBAN",
		//	AccountType:       model.PremiumAccount,
		//	Address:           "1 Road Somewhere",
		//	BankID:            "5000000",
		//	BankIDCode:        "GBDSC",
		//	Name:              "Bill Gates",
		//},
		//
		//DebtorParty: &model.Party{
		//	AccountName:       "S Balmer",
		//	AccountNumber:     "NL99ABNA0123456789",
		//	AccountNumberCode: "IBAN",
		//	AccountType:       model.PremiumAccount,
		//	Address:           "4 Sure Over Here",
		//	BankID:            "700500",
		//	BankIDCode:        "GBDDD",
		//	Name:              "Steve Balmer",
		//},
		//
		//SponsorParty: &model.Party{
		//	AccountNumber: "321654",
		//	BankID:        "456798",
		//	BankIDCode:    "GBDAS",
		//},
		//
		//ChargesInformation: &model.Charge{
		//	BearerCode:              "4532452",
		//	ReceiverChargesAmount:   "1.12",
		//	ReceiverChargesCurrency: "USD",
		//	SenderCharges: []*model.CurrencyAmount{
		//		{
		//			Amount:   "987.32",
		//			Currency: "USD",
		//		},
		//		{
		//			Amount:   "9.32",
		//			Currency: "USD",
		//		},
		//	},
		//},
		//
		//FX: &model.FX{
		//	ContractReference: "FX123",
		//	ExchangeRate:      "2.00000",
		//	OriginalAmount:    "200000.01",
		//	OriginalCurrency:  "USD",
		//},
	},
	{
		Model:          model.Model{ID: uuid.NewV4()},
		OrganisationID: organisationFixtures[0].ID,

		Amount:               "5.55",
		Currency:             "GBP",
		EndToEndReference:    "Some reference B",
		NumericReference:     "987984654",
		PaymentID:            "116544",
		PaymentPurpose:       "Some purpose",
		PaymentScheme:        "FPS",
		PaymentType:          "Debit",
		ProcessingDate:       "2017-01-18",
		Reference:            "Another reference",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",

		//BeneficiaryParty: &model.Party{
		//	AccountName:       "B Gates",
		//	AccountNumber:     "333333333",
		//	AccountNumberCode: "BBAN",
		//	AccountType:       model.BasicAccount,
		//	Address:           "1 Road Somewhere",
		//	BankID:            "5000000",
		//	BankIDCode:        "GBDSC",
		//	Name:              "Bill Gates",
		//},
		//
		//DebtorParty: &model.Party{
		//	AccountName:       "S Balmer",
		//	AccountNumber:     "NL99ABNA0123456789",
		//	AccountNumberCode: "IBAN",
		//	AccountType:       model.PremiumAccount,
		//	Address:           "4 Sure Over Here",
		//	BankID:            "700500",
		//	BankIDCode:        "GBDDD",
		//	Name:              "Steve Balmer",
		//},
		//
		//SponsorParty: &model.Party{
		//	AccountNumber: "98987987987",
		//	BankID:        "400450",
		//	BankIDCode:    "GBDAS",
		//},
		//
		//ChargesInformation: &model.Charge{
		//	BearerCode:              "987987",
		//	ReceiverChargesAmount:   "11.99",
		//	ReceiverChargesCurrency: "USD",
		//	SenderCharges: []*model.CurrencyAmount{
		//		{
		//			Amount:   "900000000.12",
		//			Currency: "USD",
		//		},
		//		{
		//			Amount:   "2.01",
		//			Currency: "USD",
		//		},
		//		{
		//			Amount:   "7.77",
		//			Currency: "USD",
		//		},
		//	},
		//},
		//
		//FX: &model.FX{
		//	ContractReference: "FX321",
		//	ExchangeRate:      "1.77700",
		//	OriginalAmount:    "90.01",
		//	OriginalCurrency:  "USD",
		//},
	},
	{
		Model:          model.Model{ID: uuid.NewV4()},
		OrganisationID: organisationFixtures[1].ID,

		Amount:               "987.54",
		Currency:             "USD",
		EndToEndReference:    "Some reference C",
		NumericReference:     "90011337420",
		PaymentID:            "01234659876",
		PaymentPurpose:       "Just paying",
		PaymentScheme:        "FPS",
		PaymentType:          "Credit",
		ProcessingDate:       "2017-01-18",
		Reference:            "Another reference",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",

		//BeneficiaryParty: &model.Party{
		//	AccountName:       "B Gates",
		//	AccountNumber:     "333333333",
		//	AccountNumberCode: "BBAN",
		//	AccountType:       model.PremiumAccount,
		//	Address:           "1 Road Somewhere",
		//	BankID:            "5000000",
		//	BankIDCode:        "GBDSC",
		//	Name:              "Bill Gates",
		//},
		//
		//DebtorParty: &model.Party{
		//	AccountName:       "S Balmer",
		//	AccountNumber:     "NL99ABNA0123456789",
		//	AccountNumberCode: "IBAN",
		//	AccountType:       model.PremiumAccount,
		//	Address:           "4 Sure Over Here",
		//	BankID:            "700500",
		//	BankIDCode:        "GBDDD",
		//	Name:              "Steve Balmer",
		//},
		//
		//SponsorParty: &model.Party{
		//	AccountNumber: "32134",
		//	BankID:        "65476958789",
		//	BankIDCode:    "GBDAS",
		//},
		//
		//ChargesInformation: &model.Charge{
		//	BearerCode:              "46876987",
		//	ReceiverChargesAmount:   "4.20",
		//	ReceiverChargesCurrency: "USD",
		//	SenderCharges: []*model.CurrencyAmount{
		//		{
		//			Amount:   "8.20",
		//			Currency: "USD",
		//		},
		//	},
		//},
		//
		//FX: &model.FX{
		//	ContractReference: "FX123",
		//	ExchangeRate:      "7.01010",
		//	OriginalAmount:    "20.87",
		//	OriginalCurrency:  "USD",
		//},
	},
	{
		Model:          model.Model{ID: uuid.NewV4(), DeletedAt: &time.Time{}},
		OrganisationID: organisationFixtures[2].ID,

		Amount:               "13.37",
		Currency:             "GBP",
		EndToEndReference:    "Some reference D",
		NumericReference:     "4242",
		PaymentID:            "01234452452659876",
		PaymentPurpose:       "Just paying again",
		PaymentScheme:        "FPS",
		PaymentType:          "Credit",
		ProcessingDate:       "2017-01-18",
		Reference:            "Another reference",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",

		//BeneficiaryParty: &model.Party{
		//	AccountName:       "B Gates",
		//	AccountNumber:     "333333333",
		//	AccountNumberCode: "BBAN",
		//	AccountType:       model.PremiumAccount,
		//	Address:           "1 Road Somewhere",
		//	BankID:            "5000000",
		//	BankIDCode:        "GBDSC",
		//	Name:              "Bill Gates",
		//},
		//
		//DebtorParty: &model.Party{
		//	AccountName:       "S Balmer",
		//	AccountNumber:     "NL99ABNA0123456789",
		//	AccountNumberCode: "IBAN",
		//	AccountType:       model.PremiumAccount,
		//	Address:           "4 Sure Over Here",
		//	BankID:            "700500",
		//	BankIDCode:        "GBDDD",
		//	Name:              "Steve Balmer",
		//},
		//
		//SponsorParty: &model.Party{
		//	AccountNumber: "321654",
		//	BankID:        "456798",
		//	BankIDCode:    "GBDAS",
		//},
		//
		//ChargesInformation: &model.Charge{
		//	BearerCode:              "9877698354",
		//	ReceiverChargesAmount:   "1.12",
		//	ReceiverChargesCurrency: "USD",
		//	SenderCharges: []*model.CurrencyAmount{
		//		{
		//			Amount:   "987.32",
		//			Currency: "USD",
		//		},
		//		{
		//			Amount:   "9.32",
		//			Currency: "USD",
		//		},
		//	},
		//},
		//
		//FX: &model.FX{
		//	ContractReference: "FX123",
		//	ExchangeRate:      "0.00001",
		//	OriginalAmount:    "200000000000.01",
		//	OriginalCurrency:  "USD",
		//},
	},
}

// GetOrganisationFixtures method returns all Organisation fixtures used to populate the test database. When `deleted`
// bool is set to true, it returns all Organisations marked as deleted, all other Organisations otherwise.
func GetOrganisationFixtures(deleted bool) []*model.Organisation {
	orgs := []*model.Organisation{}

	for _, org := range organisationFixtures {
		if (org.DeletedAt != nil) == deleted {
			orgs = append(orgs, org)
		}
	}

	return orgs
}

// GetPaymentFixtures method returns all Payment fixtures used to populate the test database. When `deleted`
// bool is set to true, it returns all Payments marked as deleted, all other Payments otherwise.
func GetPaymentFixtures(deleted bool) []*model.Payment {
	payments := []*model.Payment{}

	for _, payment := range paymentFixtures {
		if (payment.DeletedAt != nil) == deleted {
			payments = append(payments, payment)
		}
	}

	return payments
}

// NewTestDatabase method creates a database connection for testing purposes, using sqlite.
func NewTestDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/api.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = migrate(db); err != nil {
		return nil, err
	}

	if err = populate(db); err != nil {
		return nil, err
	}

	return db, nil
}

// Create all database tables.
func migrate(db *gorm.DB) error {
	// First drop all tables, so we don't have residual data that can cause errors.
	db.DropTableIfExists(
		&model.Payment{},
		&model.FX{},
		&model.Charge{},
		&model.CurrencyAmount{},
		&model.Party{},
		&model.Organisation{},
	)

	return db.AutoMigrate(
		&model.Organisation{},
		&model.Party{},
		&model.Charge{},
		&model.CurrencyAmount{},
		&model.FX{},
		&model.Payment{},
	).Error
}

// Populate the database with test data.
func populate(db *gorm.DB) error {
	for _, org := range organisationFixtures {
		deleted := org.DeletedAt
		org.DeletedAt = nil
		if err := db.Create(org).Error; err != nil {
			return err
		}
		// Make sure the fixtures are updated to correctly represent the database.
		db.Where(org).First(org)

		if deleted == nil {
			continue
		}

		// Make sure the organisation is deleted.
		if err := db.Delete(org).Error; err != nil {
			return err
		}
		org.DeletedAt = deleted
	}

	for _, payment := range paymentFixtures {
		deleted := payment.DeletedAt
		payment.DeletedAt = nil
		if err := db.Create(payment).Error; err != nil {
			return err
		}

		// Make sure the fixtures are updated to correctly represent the database.
		db.Where(payment).First(payment)
		if deleted == nil {
			continue
		}

		// Make sure the organisation is deleted.
		if err := db.Delete(payment).Error; err != nil {
			return err
		}
		payment.DeletedAt = deleted
	}

	return nil
}
