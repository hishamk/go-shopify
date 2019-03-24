package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const draftOrdersBasePath = "admin/draft_orders"
const draftOrdersResourceName = "draft_orders"

// DraftOrderService is an interface for interfacing with the draftOrders endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/orders/draftorder
type DraftOrderService interface {
	List(interface{}) ([]DraftOrder, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*DraftOrder, error)
	Create(DraftOrder) (*DraftOrder, error)
	Update(DraftOrder) (*DraftOrder, error)

	// MetafieldsService used for DraftOrder resource to communicate with Metafields resource
	MetafieldsService

	// FulfillmentsService used for DraftOrder resource to communicate with Fulfillments resource
	FulfillmentsService
}

// DraftOrderServiceOp handles communication with the draft order related methods of the
// Shopify API.
type DraftOrderServiceOp struct {
	client *Client
}

// A struct for all available draft order count options
type DraftOrderCountOptions struct {
	Page              int       `url:"page,omitempty"`
	Limit             int       `url:"limit,omitempty"`
	SinceID           int       `url:"since_id,omitempty"`
	CreatedAtMin      time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax      time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin      time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax      time.Time `url:"updated_at_max,omitempty"`
	Order             string    `url:"order,omitempty"`
	Fields            string    `url:"fields,omitempty"`
	Status            string    `url:"status,omitempty"`
	FinancialStatus   string    `url:"financial_status,omitempty"`
	FulfillmentStatus string    `url:"fulfillment_status,omitempty"`
}

// A struct for all available draft order list options.
// See: https://help.shopify.com/api/reference/orders/draftorder#index
type DraftOrderListOptions struct {
	Page              int       `url:"page,omitempty"`
	Limit             int       `url:"limit,omitempty"`
	SinceID           int       `url:"since_id,omitempty"`
	Status            string    `url:"status,omitempty"`
	FinancialStatus   string    `url:"financial_status,omitempty"`
	FulfillmentStatus string    `url:"fulfillment_status,omitempty"`
	CreatedAtMin      time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax      time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin      time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax      time.Time `url:"updated_at_max,omitempty"`
	ProcessedAtMin    time.Time `url:"processed_at_min,omitempty"`
	ProcessedAtMax    time.Time `url:"processed_at_max,omitempty"`
	Fields            string    `url:"fields,omitempty"`
	Order             string    `url:"order,omitempty"`
}

// DraftOrder represents a Shopify order
type DraftOrder struct {
	ID              int                  `json:"id,omitempty"`
	Name            string               `json:"name,omitempty"`
	Email           string               `json:"email,omitempty"`
	CreatedAt       *time.Time           `json:"created_at,omitempty"`
	UpdatedAt       *time.Time           `json:"updated_at,omitempty"`
	CompletedAt     *time.Time           `json:"completed_at,omitempty"`
	OrderID         int                  `json:"order_id,omitempty"`
	Customer        *Customer            `json:"customer,omitempty"`
	BillingAddress  *Address             `json:"billing_address,omitempty"`
	ShippingAddress *Address             `json:"shipping_address,omitempty"`
	Currency        string               `json:"currency,omitempty"`
	TotalPrice      *decimal.Decimal     `json:"total_price,omitempty"`
	SubtotalPrice   *decimal.Decimal     `json:"subtotal_price,omitempty"`
	Note            string               `json:"note,omitempty"`
	NoteAttributes  []NoteAttribute      `json:"note_attributes,omitempty"`
	InvoiceSentAt   *time.Time           `json:"invoice_sent_at,omitempty"`
	InvoiceURL      string               `json:"invoice_url,omitempty"`
	LineItems       []DraftOrderLineItem `json:"line_items,omitempty"`
	ShippingLine    ShippingLine         `json:"shipping_line,omitempty"`
	Tags            string               `json:"tags,omitempty"`
	TaxExempt       bool                 `json:"tax_exempt,omitempty"`
	TaxLines        []TaxLine            `json:"tax_lines,omitempty"`
	AppliedDiscount AppliedDiscount      `json:applied_discount,omitempty`
	TaxesIncluded   bool                 `json:"taxes_included,omitempty"`
	TotalTax        *decimal.Decimal     `json:"total_tax,omitempty"`
	Status          string               `json:"status,omitempty"`

	//Fulfillments      []Fulfillment    `json:"fulfillments,omitempty"`
	//FulfillmentStatus string           `json:"fulfillment_status,omitempty"`
	//CartToken         string           `json:"cart_token,omitempty"`
	//Transactions      []Transaction    `json:"transactions,omitempty"`
	//Confirmed         bool             `json:"confirmed,omitempty"`
	//TotalPriceUSD     *decimal.Decimal `json:"total_price_usd,omitempty"`
	//CheckoutToken     string           `json:"checkout_token,omitempty"`
	//Reference         string           `json:"reference,omitempty"`
	//SourceIdentifier  string           `json:"source_identifier,omitempty"`
	//SourceURL         string           `json:"source_url,omitempty"`
	//DeviceID          int              `json:"device_id,omitempty"`
	//LandingSiteRef    string           `json:"landing_site_ref,omitempty"`
	//CheckoutID        int              `json:"checkout_id,omitempty"`
	//ContactEmail      string           `json:"contact_email,omitempty"`

	Metafields []Metafield `json:"metafields,omitempty"`
}

type ShippingLine struct {
	Handle string           `json:"handle,omitempty"`
	Price  *decimal.Decimal `json:"price,omitempty"`
	Title  string           `json:"title,omitempty"`
}

type AppliedDiscount struct {
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Value       *decimal.Decimal `json:"value,omitempty"`
	ValueType   string           `json:"value_type,omitempty"`
	Amount      *decimal.Decimal `json:"amount,omitempty"`
}

type DraftOrderLineItem struct {
	ID                  int              `json:"id,omitempty"`
	ProductID           int              `json:"product_id,omitempty"`
	VariantID           int              `json:"variant_id,omitempty"`
	Quantity            int              `json:"quantity,omitempty"`
	Price               *decimal.Decimal `json:"price,omitempty"`
	AppliedDiscount     AppliedDiscount  `json:"applied_discount,omitempty"`
	TaxLines            []TaxLine        `json:"tax_lines,omitempty"`
	Taxable             bool             `json:"taxable,omitempty"`
	Properties          []NoteAttribute  `json:"properties,omitempty"`
	GiftCard            bool             `json:"gift_card,omitempty"`
	Name                string           `json:"name,omitempty"`
	Vendor              string           `json:"vendor,omitempty"`
	VariantTitle        string           `json:"variant_title,omitempty"`
	Title               string           `json:"title,omitempty"`
	SKU                 string           `json:"sku,omitempty"`
	RequiresShipping    bool             `json:"requires_shipping,omitempty"`
	Grams               int              `json:"grams,omitempty"`
	FulfillmentService  string           `json:"fulfillment_service,omitempty"`
	FulfillableQuantity int              `json:"fulfillable_quantity,omitempty"`
	Custom              bool             `json:"custom,omitempty"`
}

// Represents the result from the draftOrders/X.json endpoint
type DraftOrderResource struct {
	Order *DraftOrder `json:"draft_order"`
}

// Represents the result from the draftOrders.json endpoint
type DraftOrdersResource struct {
	Orders []DraftOrder `json:"draft_orders"`
}

// List draftOrders
func (s *DraftOrderServiceOp) List(options interface{}) ([]DraftOrder, error) {
	path := fmt.Sprintf("%s.json", draftOrdersBasePath)
	resource := new(DraftOrdersResource)
	err := s.client.Get(path, resource, options)
	return resource.Orders, err
}

// Count draftOrders
func (s *DraftOrderServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", draftOrdersBasePath)
	return s.client.Count(path, options)
}

// Get individual order
func (s *DraftOrderServiceOp) Get(orderID int, options interface{}) (*DraftOrder, error) {
	path := fmt.Sprintf("%s/%d.json", draftOrdersBasePath, orderID)
	fmt.Println(path)
	resource := new(DraftOrderResource)
	err := s.client.Get(path, resource, options)
	return resource.Order, err
}

// Create order
func (s *DraftOrderServiceOp) Create(order DraftOrder) (*DraftOrder, error) {
	path := fmt.Sprintf("%s.json", draftOrdersBasePath)
	wrappedData := DraftOrderResource{Order: &order}
	resource := new(DraftOrderResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Order, err
}

// Update order
func (s *DraftOrderServiceOp) Update(order DraftOrder) (*DraftOrder, error) {
	path := fmt.Sprintf("%s/%d.json", draftOrdersBasePath, order.ID)
	wrappedData := DraftOrderResource{Order: &order}
	resource := new(DraftOrderResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Order, err
}

// List metafields for a draft order
func (s *DraftOrderServiceOp) ListMetafields(orderID int, options interface{}) ([]Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return metafieldService.List(options)
}

// Count metafields for a draft order
func (s *DraftOrderServiceOp) CountMetafields(orderID int, options interface{}) (int, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return metafieldService.Count(options)
}

// Get individual metafield for a draft order
func (s *DraftOrderServiceOp) GetMetafield(orderID int, metafieldID int, options interface{}) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return metafieldService.Get(metafieldID, options)
}

// Create a new metafield for a draft order
func (s *DraftOrderServiceOp) CreateMetafield(orderID int, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return metafieldService.Create(metafield)
}

// Update an existing metafield for a draft order
func (s *DraftOrderServiceOp) UpdateMetafield(orderID int, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return metafieldService.Update(metafield)
}

// Delete an existing metafield for a draft order
func (s *DraftOrderServiceOp) DeleteMetafield(orderID int, metafieldID int) error {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return metafieldService.Delete(metafieldID)
}

// List fulfillments for a draft order
func (s *DraftOrderServiceOp) ListFulfillments(orderID int, options interface{}) ([]Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return fulfillmentService.List(options)
}

// Count fulfillments for a draft order
func (s *DraftOrderServiceOp) CountFulfillments(orderID int, options interface{}) (int, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return fulfillmentService.Count(options)
}

// Get individual fulfillment for a draft order
func (s *DraftOrderServiceOp) GetFulfillment(orderID int, fulfillmentID int, options interface{}) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return fulfillmentService.Get(fulfillmentID, options)
}

// Create a new fulfillment for a draft order
func (s *DraftOrderServiceOp) CreateFulfillment(orderID int, fulfillment Fulfillment) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return fulfillmentService.Create(fulfillment)
}

// Update an existing fulfillment for a draft order
func (s *DraftOrderServiceOp) UpdateFulfillment(orderID int, fulfillment Fulfillment) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return fulfillmentService.Update(fulfillment)
}

// Complete an existing fulfillment for a draft order
func (s *DraftOrderServiceOp) CompleteFulfillment(orderID int, fulfillmentID int) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return fulfillmentService.Complete(fulfillmentID)
}

// Transition an existing fulfillment for a draft order
func (s *DraftOrderServiceOp) TransitionFulfillment(orderID int, fulfillmentID int) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return fulfillmentService.Transition(fulfillmentID)
}

// Cancel an existing fulfillment for a draft order
func (s *DraftOrderServiceOp) CancelFulfillment(orderID int, fulfillmentID int) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: orderID}
	return fulfillmentService.Cancel(fulfillmentID)
}
