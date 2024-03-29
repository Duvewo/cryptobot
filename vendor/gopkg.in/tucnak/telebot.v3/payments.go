package telebot

import (
	"encoding/json"
	"math"
)

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	Sender  *User           `json:"from"`
	ID      string          `json:"id"`
	Payload string          `json:"invoice_payload"`
	Address ShippingAddress `json:"shipping_address"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// ShippingOption represents one shipping option.
type ShippingOption struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Prices []Price `json:"prices"`
}

// Payment contains basic information about a successful payment.
type Payment struct {
	Currency         string `json:"currency"`
	Total            int    `json:"total_amount"`
	Payload          string `json:"invoice_payload"`
	OptionID         string `json:"shipping_option_id"`
	Order            Order  `json:"order_info"`
	TelegramChargeID string `json:"telegram_payment_charge_id"`
	ProviderChargeID string `json:"provider_payment_charge_id"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	Sender   *User  `json:"from"`
	ID       string `json:"id"`
	Currency string `json:"currency"`
	Payload  string `json:"invoice_payload"`
	Total    int    `json:"total_amount"`
	OptionID string `json:"shipping_option_id"`
	Order    Order  `json:"order_info"`
}

// Order represents information about an order.
type Order struct {
	Name        string          `json:"name"`
	PhoneNumber string          `json:"phone_number"`
	Email       string          `json:"email"`
	Address     ShippingAddress `json:"shipping_address"`
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Payload     string  `json:"payload"`
	Currency    string  `json:"currency"`
	Prices      []Price `json:"prices"`
	Token       string  `json:"provider_token"`
	Data        string  `json:"provider_data"`

	Photo     *Photo `json:"photo"`
	PhotoSize int    `json:"photo_size"`

	// Unique deep-linking parameter that can be used to
	// generate this invoice when used as a start parameter.
	Start string `json:"start_parameter"`

	// Shows the total price in the smallest units of the currency.
	// For example, for a price of US$ 1.45 pass amount = 145.
	Total int `json:"total_amount"`

	NeedName            bool `json:"need_name"`
	NeedPhoneNumber     bool `json:"need_phone_number"`
	NeedEmail           bool `json:"need_email"`
	NeedShippingAddress bool `json:"need_shipping_address"`
	SendPhoneNumber     bool `json:"send_phone_number_to_provider"`
	SendEmail           bool `json:"send_email_to_provider"`
	Flexible            bool `json:"is_flexible"`
}

// Price represents a portion of the price for goods or services.
type Price struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// Currency contains information about supported currency for payments.
type Currency struct {
	Code         string      `json:"code"`
	Title        string      `json:"title"`
	Symbol       string      `json:"symbol"`
	Native       string      `json:"native"`
	ThousandsSep string      `json:"thousands_sep"`
	DecimalSep   string      `json:"decimal_sep"`
	SymbolLeft   bool        `json:"symbol_left"`
	SpaceBetween bool        `json:"space_between"`
	Exp          int         `json:"exp"`
	MinAmount    interface{} `json:"min_amount"`
	MaxAmount    interface{} `json:"max_amount"`
}

func (c Currency) FromTotal(total int) float64 {
	return float64(total) / math.Pow(10, float64(c.Exp))
}

func (c Currency) ToTotal(total float64) int {
	return int(total) * int(math.Pow(10, float64(c.Exp)))
}

var SupportedCurrencies = make(map[string]Currency)

func init() {
	err := json.Unmarshal([]byte(dataCurrencies), &SupportedCurrencies)
	if err != nil {
		panic(err)
	}
}
