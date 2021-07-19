package test_context

import (
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	Ctx()
	time.Sleep(time.Minute * 2)
}

type OrderItem struct {
	Code int64           `json:"code"`
	Data []OrderItemInfo `json:"data"`
	Msg  string          `json:"msg"`
}

type OrderItemSerializerData struct {
	ID                int64       `json:"id"`
	ItemID            int64       `json:"item_id"`
	Price             int64       `json:"price"`
	DiscountPrice     string      `json:"discount_price"`
	CommodityPrice    interface{} `json:"commodity_price"`
	Number            int64       `json:"number"`
	ServiceItemID     int64       `json:"service_item_id"`
	Unit              string      `json:"unit"`
	UnitID            string      `json:"unit_id"`
	Memo              interface{} `json:"memo"`
	CustomizeName     string      `json:"customize_name"`
	CostPrice         float64     `json:"cost_price"`
	UseCardID         interface{} `json:"use_card_id"`
	UseCardNum        int64       `json:"use_card_num"`
	UseCardPrice      int64       `json:"use_card_price"`
	IsAllowedPurchase bool        `json:"is_allowed_purchase"`
	OrderID           int64       `json:"order_id"`
	ProjectID         int64       `json:"project_id"`
	OriginTotalPrice  int64       `json:"origin_total_price"`

	Name          string      `json:"name"`
	FirstType     int64       `json:"first_type"`
	FirstTypeName string      `json:"first_type_name"`
	Attribute     string      `json:"attribute"`
	Brand         string      `json:"brand"`
	CouponName    interface{} `json:"coupon_name"`

	InventoryAmount   int64   `json:"inventory_amount"`
	InventoryPrice    float64 `json:"inventory_price"`
	InventoryUnitID   string  `json:"inventory_unit_id"`
	InventoryUnitName string  `json:"inventory_unit_name"`
	ItemName          string  `json:"item_name"`
	MaterialPicker    string  `json:"material_picker"`
	Model             string  `json:"model"`

	NoTaxInventoryPrice float64 `json:"no_tax_inventory_price"`
	NoTaxPrice          int64   `json:"no_tax_price"`
	OeCode              string  `json:"oe_code"`

	PickNum        int64             `json:"pick_num"`
	Position       string            `json:"position"`
	ProductCode    interface{}       `json:"product_code"`
	ProjectName    string            `json:"project_name"`
	Quality        ToolStruct_sub1   `json:"quality"`
	ReturnNumber   int64             `json:"return_number"`
	SBarcode       string            `json:"s_barcode"`
	SecondType     int64             `json:"second_type"`
	SecondTypeName string            `json:"second_type_name"`
	Spec           string            `json:"spec"`
	Technicians    []ToolStruct_sub2 `json:"technicians"`
	ThirdType      int64             `json:"third_type"`
	ThirdTypeName  string            `json:"third_type_name"`
	UnitInfo       ToolStruct_sub4   `json:"unit_info"`

	UseCardName interface{} `json:"use_card_name"`
}

type ToolStruct_sub4 struct {
	BaseUnitID string            `json:"base_unit_id"`
	Category   int64             `json:"category"`
	CorpID     int64             `json:"corp_id"`
	Entry      []ToolStruct_sub3 `json:"entry"`
	UnitID     string            `json:"unit_id"`
	UnitName   string            `json:"unit_name"`
}

type OrderItemInfo struct {
	ClaimParty           interface{}           `json:"claim_party"`
	ClaimPartyDisplay    string                `json:"claim_party_display"`
	InsuranceCompanyID   interface{}           `json:"insurance_company_id"`
	InsuranceCompanyName string                `json:"insurance_company_name"`
	ProjectID            int64                 `json:"project_id"`
	ProjectName          string                `json:"project_name"`
	Data                 []OrderItemSerializer `json:"data"`
}

type ToolStruct_sub2 struct {
	Commission            string `json:"commission"`
	ID                    int64  `json:"id"`
	IsLeader              bool   `json:"is_leader"`
	Name                  string `json:"name"`
	Position              int64  `json:"position"`
	PositionDisplay       string `json:"position_display"`
	TechnicianType        string `json:"technician_type"`
	TechnicianTypeDisplay string `json:"technician_type_display"`
}

type OrderItemSerializer struct {
	Data              []OrderItemSerializerData `json:"data"`
	FirstTypeID       int64                     `json:"first_type_id"`
	FirstTypeName     string                    `json:"first_type_name"`
	ProjectTotalPrice int64                     `json:"project_total_price"`
}

type ToolStruct_sub3 struct {
	Default  bool   `json:"default"`
	Rate     int64  `json:"rate"`
	UnitID   string `json:"unit_id"`
	UnitName string `json:"unit_name"`
}

type ToolStruct_sub1 struct {
	ID   interface{} `json:"id"`
	Name string      `json:"name"`
}
