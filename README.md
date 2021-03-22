package main


type CASRateReq struct {
	State       string    `json:"STATE,omitempty"`
	Country     string    `json:"COUNTRY,omitempty"`
	Role        int       `json:"ROLE,omitempty"`
	CASID       string    `json:"CASID,omitempty"`
	ProduceName string    `json:"PRODUCE,omitempty"`
	PriceRate   []CASRate `json:"CASRATES,omitempty"`
}
type CASRateAsset struct {
	DocType     string    `json:"docType,omitempty"`
	Country     string    `json:"COUNTRY,omitempty"`
	State       string    `json:"STATE,omitempty"`
	CASID       string    `json:"CASID,omitempty"`
	ProduceName string    `json:"PRODUCE,omitempty"`
	PriceRate   []CASRate `json:"CASRATES,omitempty"`
}

type CASRate struct {
	CurrencyUnit string      `json:"CURRENCY,omitempty"`
	QuantityUnit interface{} `json:"QUANTITYUNIT,omitempty"`
	DurationUnit string      `json:"DURATIONUNIT,omitempty"`
	Value        float64     `json:"VALUE,omitempty"`
}

type DCRateReq struct {
	State       string  `json:"STATE,omitempty"`
	Country     string  `json:"COUNTRY,omitempty"`
	ProduceName string  `json:"PRODUCE,omitempty"`
	Role        int     `json:"ROLE,omitempty"`
	PriceRate   float32 `json:"PRICERATE,omitempty"`
}
type DCRateAsset struct {
	DocType     string  `json:"docType,omitempty"`
	State       string  `json:"STATE,omitempty"`
	Country     string  `json:"COUNTRY,omitempty"`
	ProduceName string  `json:"PRODUCE,omitempty"`
	Role        int     `json:"ROLE,omitempty"`
	PriceRate   float32 `json:"PRICERATE,omitempty"`
}


type DeliveryRateReq struct {
	State       string         `json:"STATE,omitempty"`
	Country     string         `json:"COUNTRY,omitempty"`
	ProduceName string         `json:"PRODUCE,omitempty"`
	Role        int            `json:"ROLE,omitempty"`
	PriceRate   []DeliveryRate `json:"PRICERATE,omitempty"`
}
type DeliveryRateAsset struct {
	DocType     string         `json:"docType,omitempty"`
	State       string         `json:"STATE,omitempty"`
	Country     string         `json:"COUNTRY,omitempty"`
	ProduceName string         `json:"PRODUCE,omitempty"`
	PriceRate   []DeliveryRate `json:"PRICERATE,omitempty"`
}

type DeliveryRate struct {
	CurrencyUnit string      `json:"CURRENCY,omitempty"`
	QuantityUnit interface{} `json:"QUANTITYUNIT,omitempty"`
	DistanceUnit string      `json:"DISTANCEUNIT,omitempty"`
	Value        float64     `json:"VALUE,omitempty"`
}

type DfarmRateReq struct {
	State       string   `json:"STATE,omitempty"`
	Country     string   `json:"COUNTRY,omitempty"`
	ProduceName string   `json:"PRODUCE,omitempty"`
	Role        int      `json:"ROLE,omitempty"`
	PriceRate   DfarmPer `json:"PRICERATE,omitempty"`
}
type DfarmRateAsset struct {
	DocType     string   `json:"docType,omitempty"`
	State       string   `json:"STATE,omitempty"`
	Country     string   `json:"COUNTRY,omitempty"`
	ProduceName string   `json:"PRODUCE,omitempty"`
	PriceRate   DfarmPer `json:"PRICERATE,omitempty"`
}

type DfarmPer struct {
	Totalper     float64 `json:"DFARMPER,omitempty"`
	MaxAgentPer  float64 `json:"MAXAGENTPER,omitempty"`
	DemurragePer float64 `json:"DEMURRAGEPER,omitempty"`
}
type Payable struct {
	Verity   string  `json:"verity,omitempty"`
	Currency string  `json:"currency,omitempty"` //usd ,Rs
	Amount   float64 `json:"amount,omitempty"`
	Status   string  `json:"status,omitempty"` //Paid,Unpaid

}
//FarmerPayableAsset struct
type FarmerPayableAsset struct {
	FarmerID           string    `json:"farmerID,omitempty"`
	ProduceID          string    `json:"produceID,omitempty"`
	Payables           []Payable `json:"payables,omitempty"`
	ExpectedPayable    float64   `json:"expectedPayables,omitempty"`
	TotalPayable       float64   `json:"TotalFarmerPayable,omitempty"`
	TotalPayableStatus string    `json:"TotalPayableAssetstatus,omitempty"`
}

type GenAsset struct {
	AssetName   string        `json:"assetName"`             //Name of Asset i.e IOT Device Info
	Keys        []string      `json:"keys"`                  //Primary Key to add this Asset to Ledger
	EntityCount int32         `json:"entityCount,omitempty"` //No. of Asset data
	AssetDatas  []interface{} `json:"assetDatas,omitempty"`  //List of asset data
	QueryString string        `json:"queryString,omitempty"` //Generic Couch query string ,this value must be not set if ledger world state is stateDB
	Bookmark    string        `json:"bookmark,omitempty"`
}

// GenAssetResult :This go structure  will contain information about result of Asset creation request
type GenAssetResult struct {
	Keys   []string `json:"keys,omitempty"`
	Result string   `json:"result,omitempty"`
}

type ParticipantRate struct {
	DeliveryRateVal DeliveryRate
	DCpercentage    float32
	Dfarmpercentage DfarmPer
}

type QualityPaymentInfo struct {
	QualityType string  `json:"QUALITYTYPE,omitempty"` //Grade A /Grade B
	Qty         uint64  `json:"Qty,omitempty"`
	Amount      float64 `json:"AMOUNT,omitempty"`
	Currency    string  `json:"CURRENCY,omitempty"`
}
type ParticipantInfo struct {
	Type    int    `json:"ROLE,omitempty"`
	ID      string `json:"MAINID,omitempty"`
	OrderID string `json:"ORDERID,omitempty"`
	Country string `json:"COUNTRY,omitempty"`
	State   string `json:"STATE,omitempty"`
}

type ParticipantPaymentInvoiceAsset struct {
	DocType                string               `json:"docType,omitempty"`
	PaymentInvoiceID       string               `json:"PAYMENTINVOICEID,omitempty"`
	OrderID                string               `json:"ORDERID,omitempty"`
	ProduceID              string               `json:"PRODUCEID,omitempty"`
	Variety                string               `json:"VARIETY,omitempty"`
	Qty                    uint64               `json:"QTY,omitempty"`
	BuyerID                string               `json:"buyerID,omitempty"`
	BuyerOrderID           string               `json:"BUYERORDERID,omitempty"`
	ID                     string               `json:"ID,omitempty"`
	Type                   int                  `json:"TYPE,omitempty"`
	Currency               string               `json:"CURRENCY,omitempty"`
	PayableAmount          float64              `json:"PAYABLEAMOUNT,omitempty"`
	QualityPaymentInfos    []QualityPaymentInfo `json:"QPINFOS,omitempty"`
	AccountPaymentResponse string               `json:"ACCOUNTPAYMENTRESPONSE,omitempty"`
	Status                 string               `json:"STATUS,omitempty"`
}

type PCPackagingRateAsset struct {
	DocType          string             `json:"docType,omitempty"`
	ProduceName      string             `json:"PRODUCE,omitempty"`
	Country          string             `json:"COUNTRY,omitempty"`
	State            string             `json:"STATE,omitempty"`
	Variety          string             `json:"VARIETY,omitempty"`
	PCID             string             `json:"PCID,omitempty"`
	TableVerityPrice []QualityPriceRate `json:"TABLEVARIETYPRICE,omitempty"`
	Selectedunit     Unit               `json:"SELECTED_UNIT,omitempty"`
	Currecny         string             `json:"CURRENCY,omitempty"`
}

type Rate struct {
	QuantityUnit string  `json:"UNIT,omitempty"`     //pound,Ton,KG etc
	CurrencyUnit string  `json:"CURRENCY,omitempty"` //usd,euro,inr etc
	Value        float64 `json:"VALUE,omitempty"`
}
type FarmerPriceRateAsset struct {
	DocType          string `json:"docType,omitempty"`
	FarmerID         string `json:"FARMERID,omitempty"`
	ProduceID        string `json:"PRODUCEID,omitempty"` //primary key
	Variety          string `json:"VARIETY,omitempty"`   //primery key
	State            string `json:"STATE,omitempty"`
	Country          string `json:"COUNTRY,omitempty"`
	Price            []Rate `json:"PRICE,omitempty"`
	BaseUnit         string `json:"BASE_UNIT,omitempty"`
	Selectunit       Unit   `json:"SELECTED_UNIT,omitempty"`
	IsaskingPriceset bool   `json:"ISASKINGPRICESET,omitempty"`
	AskingPrice      []Rate `json:"ASKINGPRICE,omitempty"` // farmer can set this price
}

type VarietyBasedRate struct {
	Variety    string `json:"VARIETY,omitempty"`
	MinRate    Rate   `json:"MinRate,omitempty"`
	MarketRate Rate   `json:"MarketRate,omitempty"`
}
type ProduceRateAsset struct {
	DocType     string             `json:"docType,omitempty"`
	ProduceName string             `json:"PRODUCE,omitempty"`
	State       string             `json:"STATE,omitempty"`
	Country     string             `json:"COUNTRY,omitempty"`
	BaseUnit    string             `json:"BASE_UNIT,omitempty"`
	SelectUnit  Unit               `json:"SELECTED_UNIT,omitempty"`
	CustomTime  string             `json:"customTime,omitempty"`
	Rates       []VarietyBasedRate `json:"RATES,omitempty"`
}

type ProduceRateResponse struct {
	ProduceName string             `json:"PRODUCE,omitempty"`
	State       string             `json:"STATE,omitempty"`
	Country     string             `json:"COUNTRY,omitempty"`
	BaseUnit    string             `json:"BASE_UNIT,omitempty"`
	SelectUnit  Unit               `json:"SELECTED_UNIT,omitempty"`
	Rates       []VarietyBasedRate `json:"RATES,omitempty"` //Variety as a string key
}

//[] {QUALITYTYPE,MODE ,VALUE}
type QualityPerInfo struct {
	Type     string  `json:"TYPE,omitempty"`
	OpMode   bool    `json:"OPMODE,omitempty"` //0 means add/more 1 means less/subscract
	PerValue float32 `json:"PERVALUE,omitempty"`
}

//New Structure for India Price calculation for Produce

type TableVarietyPerRateAsset struct {
	DocType         string           `json:"docType,omitempty"`
	ProduceName     string           `json:"PRODUCE,omitempty"`
	State           string           `json:"STATE,omitempty"`
	Country         string           `json:"COUNTRY,omitempty"`
	Variety         string           `json:"VARIETY,omitempty"`
	QualityPerInfos []QualityPerInfo `json:"QUALITYPERINFOS,omitempty"`
}
type PriceInqueryRequest struct {
	ProduceName         string            `json:"PRODUCE,omitempty"`
	Variety             string            `json:"VARIETY,omitempty"`
	BaseUnit            string            `json:"BASE_UNIT,omitempty"`
	SelectUnit          Unit              `json:"SELECTED_UNIT,omitempty"`
	ParticipantTypeList []ParticipantInfo `json:"PARTICIPANTTYPELIST,omitempty"`
}
type PriceInqueryResponse struct {
	ProduceName       string             `json:"PRODUCE,omitempty"`
	Variety           string             `json:"VARIETY,omitempty"`
	BaseUnit          string             `json:"BASE_UNIT,omitempty"`
	SelectUnit        Unit               `json:"SELECTED_UNIT,omitempty"`
	IsPCIncluded      bool               `json:"ISPCINCLUDED,omitempty"`
	VarietyRates      []Rate             `json:"VARIETYRATES,omitempty"`
	QualityPriceRates []QualityPriceRate `json:"QUALITYPRICERATES,omitempty"` //Variety as a string key
}

//0,1-Farmer  incase array is empty or have value 0/1 we will return Farmer Marklet pirce
//2-CAS
//3-PC
//4-Transport
//5-DC
//6-Dfarm
//7-Buyer
//13 Tenants like YNG and others
type PriceRateKey struct {
	ProduceID   string `json:"PRODUCEID,omitempty"` //primary key
	Variety     string `json:"VARIETY,omitempty"`   //primery key
	ProduceName string `json:"PRODUCE,omitempty"`
	Country     string `json:"COUNTRY,omitempty"`
	State       string `json:"STATE,omitempty"`
}
type Gap struct {
	GapID          string `json:"GAPID,omitempty"`
	Status         string `json:"STATUS,omitempty"`
	ExpirationDate string `json:"EXPIRATIONDATTE,omitempty"`
}

type ProduceQuantity struct {
	VarietyType            string `json:"name,omitempty"`
	EstimatedTotalQuantity uint64 `json:"qty,omitempty"`
	// FinalQuantity          uint64 `json:"FINALQUANTITY,omitempty"`
	//UsableQuantity         uint64         `json:"USABLEQUANTITY,omitempty"`
	//AvailableQuantity      uint64         `json:"AVAILABLEQUANTITY,omitempty"`
	//DiscardedQuantity      uint64         `json:"DISCARDEDQUANTITY,omitempty"`
	//ProcessingQuantity uint64 `json:"PROCESSINGQUANTITY,omitempty"`
	//TableVerities          []TableVariety `json:"TABLEVARIETY,omitempty"`
	//Payment Related information
	//EstimatedTotalPayment uint64        `json:"EstimatedTotalPayment,omitempty"`
	//FinalTotalPayment float64 `json:"FinalTotalPayment,omitempty"`
	//PaymentInfos          []PaymentInfo `json:"PaymentInfos,omitempty"`
}

// ProduceAsset with key as ProduceID
type ProduceAsset struct {
	DocType           string            `json:"docType,omitempty"`
	ProduceID         string            `json:"PRODUCEID,omitempty"`
	ProduceName       string            `json:"PRODUCE,omitempty"`
	ProduceQuantities []ProduceQuantity `json:"PRODUCEQUANTITES,omitempty"` //key as variety
	FarmLocation      string            `json:"FARMLOCATION,omitempty"`
	PlantingDate      string            `json:"PLANTINGDATE,omitempty"`
	GAPInfo           Gap               `json:"GAPINFO,omitempty"`
	FarmerID          string            `json:"FARMERID,omitempty"`
	Status            string            `json:"STATUS,omitempty"` //Registered,Approved,In-transit,Cleaning done,Inspection done,Delivered to Buyer,Financial settlement done from buyer Side,
	BaseUnit          string            `json:"BASE_UNIT,omitempty"`
	Unit              Unit              `json:"SELECTED_UNIT,omitempty"`
	//financial settlement done from dFarm Side,Tracking closed
	//StatusHistories []StatusHistory `json:"statusHistories,omitempty"`
}



type TransportInfoAsset struct {
	DocType       string       `json:"docType,omitempty"`
	SourceID      string       `json:"SORUCEID,omitempty"`      //Primary Key
	DestinationID string       `json:"DESTINATIONID,omitempty"` //Primary Key
	Distance      DistanceUnit `json:"DISTANCEUNIT,omitempty"`
}

type TruckorQuantity struct {
	Verity     string `json:"verity,omitempty"`
	Weight     uint64 `json:"weight,omitempty"`
	WeightUnit string `json:"weightUnit,omitempty"`
}
type TruckerPayableAsset struct {
	GID         string            `json:"gid,omitempty"`
	TruckerID   string            `json:"TruckerId,omitempty"`
	Origin      string            `json:"origin,omitempty"`
	Destination string            `json:"destination,omitempty"`
	Quantities  []TruckorQuantity `json:"quantities,omitempty"`
	Payable     float64           `json:"payable,omitempty"`
	Currency    string            `json:"currency,omitempty"`
	Status      string            `json:"status,omitempty"`
}

type UserPaymentStatRequest struct {
	ProduceID        string `json:"PRODUCEID,omitempty"`
	Variety          string `json:"VARIETY,omitempty"`
	UserType         int    `json:"USERTYPE,omitempty"`
	UserID           string `json:"USERID,omitempty"`
	BuyerOrderID     string `json:"BUYERORDERID,omitempty"`
	UserOrderID      string `json:"USERORDERID,omitempty"`
	PayableInvoiceID string `json:"PAYABLEINVOICEID,omitempty"`
}

type PaymentStat struct {
	Variety                      string               `json:"VARIETY,omitempty"`
	OrderID                      string               `json:"ORDERID,omitempty"`
	Currency                     string               `json:"CURRENCY,omitempty"`
	AvailableQtyforSell          uint64               `json:"AVAILABLEQTYFORSELL,omitempty"`
	AvailableQtyBreakdownforSell []QualityPaymentInfo `json:"AVAILABLEQTYBREAKDOWNFORSELL,omitempty"`
	QtyforReceivedPayment        uint64               `json:"QTYFORRECEIVEDPAYMENT,omitempty"`
	ReceivedPayment              float64              `json:"RECEIVEDPAYMENT,omitempty"`
	BreakdownReceivedPayment     []QualityPaymentInfo `json:"BREAKDOWNRECEIVEDPAYMENT,omitempty"`
	QtyforPendingPayment         uint64               `json:"QTYFORPENDINGPAYMENT,omitempty"`
	BreakdownPendingPayment      []QualityPaymentInfo `json:"BREAKDOWNPENDINGPAYMENT,omitempty"`
	PendingPayment               float64              `json:"PENDINGPAYMENT,omitempty"`
}
type UserPaymentStatResponse struct {
	ProduceID    string        `json:"PRODUCEID,omitempty"`
	PaymentStats []PaymentStat `json:"PAYMENTSTATS,omitempty"`
}

type WholeSellerReq struct {
	State         string  `json:"STATE,omitempty"`
	Country       string  `json:"COUNTRY,omitempty"`
	ProduceName   string  `json:"PRODUCE,omitempty"`
	Variety       string  `json:"VARIETY,omitempty"`
	Role          int     `json:"ROLE,omitempty"`
	SourceID      string  `json:"SOURCEID,omitempty"`
	DestinationID string  `json:"DESTINATIONID,omitempty"`
	PerValue      float64 `json:"PERVALUE,omitempty"`
}

//WholeSellerRatePerAsset asset to provide Markup and discount on final Buyer payment for a produce Variety

type WholeSellerRatePerAsset struct {
	DocType       string  `json:"docType,omitempty"`
	SourceID      string  `json:"SOURCEID,omitempty"`
	DestinationID string  `json:"DESTINATIONID,omitempty"`
	Variety       string  `json:"VARIETY,omitempty"`
	ProduceName   string  `json:"PRODUCE,omitempty"`
	QualityType   string  `json:"QUALITYTYPE,omitempty"`
	PerValue      float64 `json:"PERVALUE,omitempty"`
}


type AgentRateReq struct {
	State         string  `json:"STATE,omitempty"`
	Country       string  `json:"COUNTRY,omitempty"`
	ProduceName   string  `json:"PRODUCE,omitempty"`
	Variety       string  `json:"VARIETY,omitempty"`
	Role          int     `json:"ROLE,omitempty"`
	SourceID      string  `json:"SOURCEID,omitempty"`
	DestinationID string  `json:"DESTINATIONID,omitempty"`
	PerValue      float64 `json:"PERVALUE,omitempty"`
}

//AgentRatePerAsset asset to provide Markup and discount on final Buyer payment for a produce Variety

type AgentRatePerAsset struct {
	DocType       string  `json:"docType,omitempty"`
	SourceID      string  `json:"SOURCEID,omitempty"`
	DestinationID string  `json:"DESTINATIONID,omitempty"`
	Variety       string  `json:"VARIETY,omitempty"`
	ProduceName   string  `json:"PRODUCE,omitempty"`
	QualityType   string  `json:"QUALITYTYPE,omitempty"`
	PerValue      float64 `json:"PERVALUE,omitempty"`
}



type AssetData struct {
	OrderID string `json:"ORDERID,omitempty"`
}
type AssetDataChild struct {
	ChildOrderID string `json:"ChildOrderID,omitempty"`
}
type Trucker struct {
	FullName      string `json:"fullName,omitempty"`
	LicensePlate  string `json:"licensePlate,omitempty"`
	DriverLicense string `json:"driverLicense,omitempty"`
	Type          string `json:"type,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
	Size          uint64 `json:"size,omitempty"`
	Color         string `json:"color,omitempty"`
}

type ReceiveData struct {
	TruckID          uint64   `json:"truckId,omitempty"`
	ActualQty        uint64   `json:"actualQty,omitempty"`
	AgreeToActualQty bool     `json:"agreeToActualQty,omitempty"`
	CSQty            uint64   `json:"csQty,omitempty"`
	Room             uint64   `json:"room,omitempty"`
	Bins             []string `json:"bins,omitempty"`
}

type TransportaionoFCAS struct {
	TruckID         uint64  `json:"truckId,omitempty"`
	OrderBy         string  `json:"orderBy,omitempty"`
	PickUpDate      string  `json:"pickUpDate,omitempty"`
	PickUpLocation  string  `json:"pickUpLocation,omitempty"`
	DropOffLocation string  `json:"dropOffLocation,omitempty"`
	Price           float64 `json:"price,omitempty"`
	Insurance       string  `json:"insurance,omitempty"`
	Truck           Trucker `json:"trucker,omitempty"`
	Status          string  `json:"STATUS,omitempty"`
	ActualQty       uint64  `json:"actualQty,omitempty"`
	PickUpTime      string  `json:"pickUpTime,omitempty"`
}

type Unit struct {
	Name   string  `json:"NAME,omitempty"`
	Weight float64 `json:"WEIGHT,omitempty"`
}
type DistanceUnit struct {
	Name  string `json:"NAME,omitempty"`
	Value uint64 `json:"VALUE,omitempty"`
}

//Need to Discuss about Flow
// Coldstorage
type CSOrderAsset struct {
	DocType                  string               `json:"docType,omitempty"`
	OrderType                uint32               `json:"ORDERTYPE,omitempty"`
	OrderID                  string               `json:"ORDERID,omitempty"`
	OrderUnixTime            int64                `json:"orderUnixTime,omitempty"`
	SourceID                 string               `json:"SOURCEID,omitempty"`
	ProduceID                string               `json:"PRID,omitempty"`
	Variety                  string               `json:"variety,omitempty"`
	HarvestDate              string               `json:"HARVESTDATE,omitempty"`
	Qty                      uint64               `json:"QTY,omitempty"`
	DestinationID            string               `json:"DESTINATIONID,omitempty"`
	StorageFacility          string               `json:"STORAGE_FACILITY,omitempty"`
	QtyforSale               uint64               `json:"Quantity_of_sale,omitempty"`
	Status                   string               `json:"STATUS,omitempty"` //Order Requested from Farmer,Order approved from CAS,Picked up by trucker,in Transit,Delivered to CAS
	Transports               []TransportaionoFCAS `json:"Transportaion,omitempty"`
	Receives                 []ReceiveData        `json:"receive,omitempty"`
	ActualQty                uint64               `json:"TotalActualQuantity,omitempty"`
	TotaltransportationPrice float64              `json:"TotalTransportationCost,omitempty"`
	BaseUnit                 string               `json:"BASE_UNIT,omitempty"`
	SenderUnit               Unit                 `json:"SENDER_UNIT,omitempty"`
	Country                  string               `json:"COUNTRY,omitempty"`
}

//DC

type Transportaion struct {
	TruckID         uint64  `json:"truckId,omitempty"`
	OrderBy         string  `json:"orderBy,omitempty"`
	PickUpDate      string  `json:"pickUpDate,omitempty"`
	PickUpLocation  string  `json:"pickUpLocation,omitempty"`
	DropOffLocation string  `json:"dropOffLocation,omitempty"`
	Price           float64 `json:"price,omitempty"`
	Insurance       string  `json:"insurance,omitempty"`
	Truck           Trucker `json:"trucker,omitempty"`
	Status          string  `json:"STATUS,omitempty"`
	PickUpTime      string  `json:"pickUpTime,omitempty"`
	ActualQty       uint64  `json:"actualQty,omitempty"`
	ReceiveDate     string  `json:"receiveDate,omitempty"`
}

type TableVariety struct {
	Name            string  `json:"name,omitempty"`
	Value           float64 `json:"quantity,omitempty"`
	SelecteUnitdata Unit    `json:"selected_unit,omitempty"`
}

type ParentIDInfo struct {
	ProduceID      string         `json:"PRID,omitempty"`
	TableVarieties []TableVariety `json:"QTY,omitempty"`
	ParentOrderId  string         `json:"ParentOrderId,omitempty"`
	ChildOrderID   string         `json:"ChildOrderID,omitempty"`
}

type DCOrder struct {
	SourceID                 string          `json:"SOURCEID,omitempty"`
	DestinationID            string          `json:"DESTINATIONID,omitempty"`
	OrderType                uint32          `json:"ORDERTYPE,omitempty"`
	Produce                  string          `json:"Produce,omitempty"`
	Variety                  string          `json:"Variety,omitempty"`
	BaseUnit                 string          `json:"BASE_UNIT,omitempty"`
	RequiredDate             string          `json:"requiredDate,omitempty"`
	Qty                      uint64          `json:"QTY,omitempty"`
	OrderID                  string          `json:"ORDERID,omitempty"`
	Status                   string          `json:"STATUS,omitempty"`
	Transports               []Transportaion `json:"Transportaion,omitempty"`
	ParentInfo               []ParentIDInfo  `json:"bcParentorderinfo,omitempty"`
	TotaltransportationPrice float64         `json:"totalTransportationcost,omitempty"`
}

type DCOrderAsset struct {
	DocType                  string          `json:"docType,omitempty"`
	OrderType                uint32          `json:"ORDERTYPE,omitempty"`
	SourceID                 string          `json:"SOURCEID,omitempty"`
	DestinationID            string          `json:"DESTINATIONID,omitempty"`
	OrderUnixTime            int64           `json:"orderUnixTime,omitempty"`
	Produce                  string          `json:"Produce,omitempty"`
	Variety                  string          `json:"variety,omitempty"`
	RequiredDate             string          `json:"requiredDate,omitempty"`
	Qty                      uint64          `json:"QTY,omitempty"`
	ProduceID                string          `json:"PRID,omitempty"`
	TableVarieties           []TableVariety  `json:"tableVarietyDC,omitempty"`
	ParentOrderID            string          `json:"ParentOrderId,omitempty"`
	OrderID                  string          `json:"ORDERID,omitempty"`
	ChildOrderID             string          `json:"ChildOrderID,omitempty"`
	Status                   string          `json:"STATUS,omitempty"`
	Transports               []Transportaion `json:"Transportaion,omitempty"`
	TotaltransportationPrice float64         `json:"totalTransportationCost,omitempty"`
}

//PC
type TransportaionoFPC struct {
	TruckID    uint64 `json:"truckId,omitempty"`
	OrderBy    string `json:"orderBy,omitempty"`
	PickUpDate string `json:"pickUpDate,omitempty"`
	// PickUpLocation  string  `json:"pickUpLocation,omitempty"`
	DropOffLocation string  `json:"dropOffLocation,omitempty"`
	Price           float64 `json:"price,omitempty"`
	Insurance       string  `json:"insurance,omitempty"`
	Truck           Trucker `json:"trucker,omitempty"`
	Status          string  `json:"STATUS,omitempty"`
	ActualQty       uint64  `json:"actualQty,omitempty"` // need to set data type
	PickUpTime      string  `json:"pickUpTime,omitempty"`
	ReceiveDate     string  `json:"receiveDate,omitempty"`
}

type Processe struct {
	Qty uint64 `json:"qty,omitempty"`
}

type Wastage struct {
	Qty uint64 `json:"qty,omitempty"`
}

//change in mongodb
//add  TableVariety to Order
type ProcessData struct {
	TableVarieties []TableVariety `json:"tableVariety,omitempty"`
	Processed      []TableVariety `json:"processVariety,omitempty"`
	Wastages       Wastage        `json:"wastages,omitempty"`
}

/* type PCOrderAsset struct {
	DocType                  string              `json:"docType,omitempty"`
	OrderType                uint32              `json:"ORDERTYPE,omitempty"`
	SourceID                 string              `json:"SOURCEID,omitempty"`
	DestinationID            string              `json:"DESTINATIONID,omitempty"`
	OrderUnixTime            int64               `json:"orderUnixTime,omitempty"`
	ProduceID                string              `json:"PRID,omitempty"`
	Variety                  string              `json:"variety,omitempty"`
	Qty                      uint64              `json:"QTY,omitempty"`
	OrderID                  string              `json:"ORDERID,omitempty"`
	Status                   string              `json:"STATUS,omitempty"`
	Transports               []TransportaionoFPC `json:"Transportaion,omitempty"`
	ProcessProduce           ProcessData         `json:"processData,omitempty"`
	ParentID                 string              `json:"parentorderId,omitempty"`
	TotaltransportationPrice float64             `json:"totalTransportationcost,omitempty"`
	TotalavailableQty        []TableVariety      `json:"TotalavailableQty,omitempty"` // added new
	Baseunit                 string              `json:"BASE_UNIT,omitempty"`
} */

type BuyerOrderAsset struct {
	DocType          string            `json:"docType,omitempty"`
	OrderType        uint32            `json:"ORDERTYPE,omitempty"`
	DCID             string            `json:"SOURCEID,omitempty"`
	OrderUnixTime    int64             `json:"orderUnixTime,omitempty"`
	ProduceID        string            `json:"PRID,omitempty"`
	Variety          string            `json:"variety,omitempty"`
	Qty              uint64            `json:"totalQTY,omitempty"`
	BuyerID          string            `json:"DESTINATIONID,omitempty"`
	TotalPrice       float64           `json:"price,omitempty"` // total price
	QtyBreakdowns    []TableVariety    `json:"QTY,omitempty"`
	// ParticipantInfos []ParticipantInfo `json:"PARTICIPANTINFOS,omitempty"`
	OrderID          string            `json:"ORDERID,omitempty"`
	Status           string            `json:"STATUS,omitempty"`
	ParentID         string            `json:"parentorderId,omitempty"` // NEED To discuss with prasanna
	//TransportRate    Rate              `json:"TRANSPORTRATE,omitempty"`
	PayableAssetID string `json:"payableAssetID,omitempty"`
}

//Rate for any kind of produce varity quality rate.

type QualityPriceRate struct {
	QualityName string `json:"QUALITYNAME,omitempty"`
	// Rates       []Rate `json:"RATES,omitempty"`
}

// From here CC order Datastructure added
type ProducelistData struct {
	ItemId          int     `json:"itemId,omitempty"`
	ID              string  `json:"MAINID,omitempty"`
	OrderBy         string  `json:"orderBy,omitempty"`
	PickUpLocation  string  `json:"pickUpLocation,omitempty"`
	DropOffLocation string  `json:"dropOffLocation,omitempty"`
	Price           string  `json:"price,omitempty"`
	Insurance       string  `json:"insurance,omitempty"`
	BARCODEID       string  `json:"BARCODEID,omitempty"`
	ActualQty       string  `json:"actualQty,omitempty"`
	RejectedQty     string  `json:"rejectedQty,omitempty"`
	RecevieDate     string  `json:"RECEVIE_DATE,omitempty"`
	ReceviveTime    string  `json:"RECEVIVE_TIME,omitempty"`
	ReceiveStatus   string  `json:"RECEIVE_STATUS,omitempty"`
	Truck           Trucker `json:"trucker,omitempty"`
}

type CCOrderAsset struct {
	DocType             string            `json:"docType,omitempty"`
	OrderType           uint32            `json:"ORDERTYPE,omitempty"`
	OrderID             string            `json:"ORDERID,omitempty"`
	SourceID            string            `json:"SOURCEID,omitempty"`
	ProduceID           string            `json:"PRID,omitempty"`
	OrderUnixTime       int64             `json:"orderUnixTime,omitempty"`
	Produce             string            `json:"produce,omitempty"`
	Variety             string            `json:"variety,omitempty"`
	HarvestDate         string            `json:"HARVESTDATE,omitempty"`
	Qty                 uint64            `json:"QTY,omitempty"`
	DestinationID       string            `json:"DESTINATIONID,omitempty"`
	Status              string            `json:"STATUS,omitempty"` //Order Requested from Farmer,Order approved from CAS,Picked up by trucker,in Transit,Delivered to CAS
	ReceiveType         string            `json:"RECEIVE_TYPE,omitempty"`
	BaseUnit            string            `json:"BASE_UNIT,omitempty"`
	SenderUnit          Unit              `json:"SENDER_UNIT,omitempty"`
	Country             string            `json:"COUNTRY,omitempty"`
	State               string            `json:"SOURCE_STATE,omitempty"`
	TotalavailableQty   uint32            `json:"TotalavailableQty,omitempty"` // added new
	ActualHarvest       string            `json:"ACTUAL_HARVEST,omitempty"`
	TotalActualQuantity uint32            `json:"TotalActualQuantity,omitempty"` // added new
	// ParticipantInfos    []ParticipantInfo `json:"PARTICIPANTINFOS,omitempty"`
	Producelist         []ProducelistData `json:"PRODUCELIST,omitempty"`
}

//PHOrderAsset Order info for pack house
type PHOrderAsset struct {
	DocType       string `json:"docType,omitempty"`
	OrderUnixTime int64  `json:"orderUnixTime,omitempty"`
	SourceID      string `json:"SOURCEID,omitempty"`
	DestinationID string `json:"DESTINATIONID,omitempty"`
	OrderID       string `json:"ORDERID,omitempty"`
	OrderType     uint64 `json:"ORDERTYPE,omitempty"`
	ProduceID     string `json:"PRID,omitempty"`
	Variety       string `json:"variety,omitempty"`
	Qty           uint64 `json:"QTY,omitempty"`
	BaseUnit      string `json:"BASE_UNIT,omitempty"`
	SenderUnit    Unit   `json:"SENDER_UNIT,omitempty"`
	Status        string `json:"STATUS,omitempty"`
}

// GRNAsset is used for GRN asset
type GRNAsset struct {
	DocType          string  `json:"docType,omitempty"`
	GrnNo            string  `json:"GRNNO,omitempty"`
	ProduceID        string  `json:"PRID,omitempty"`
	Produce          string  `json:"PRODUCE,omitempty"`
	Variety          string  `json:"VARIETY,omitempty"`
	HarvestDate      string  `json:"HARVESTDATE,omitempty"`
	MainID           string  `json:"MAINID,omitempty"`
	Name             string  `json:"NAME,omitempty"`
	BaseUnit         string  `json:"BASE_UNIT,omitempty"`
	Currency         string  `json:"CURRENCY,omitempty"`
	State            string  `json:"STATE,omitempty"`
	Country          string  `json:"COUNTRY,omitempty"`
	Status           string  `json:"STATUS,omitempty"` //Pending,Approved and Processed
	Rate             float64 `json:"RATE,omitempty"`   // can change
	Qty              uint64  `json:"QTY,omitempty"`    // can change
	BeforeProcessQty uint64  `json:"BEFORE_PROCESS_QTY,omitempty"`
	GeneratedAt      string  `json:"GENERATED_AT,omitempty"`
	RecepientID      string  `json:"RECEPIENT_ID,omitempty"`
	Comment          string  `json:"COMMENT,omitempty"`
}

// Participanttypelist Declare
type Participanttypelist struct {
	OrderID string `json:"ORDERID,omitempty"`
	MainID  string `json:"MAINID,omitempty"`
	Role    string `json:"ROLE,omitempty"`
	State   string `json:"STATE,omitempty"`
	Country uint64 `json:"COUNTRY,omitempty"`
}

type Producelist struct {
	ItemID        string  `json:"itemId,omitempty"`
	Truck         Trucker `json:"trucker,omitempty"`
	ActualQty     uint64  `json:"actualQty,omitempty"`
	LoadingStatus string  `json:"LOADING_STATUS,omitempty"`
	ReceiveDate   string  `json:"RECEIVE_DATE,omitempty"`
	LoadDate      string  `json:"LOAD_DATE,omitempty"`
	LoadTime      string  `json:"LOAD_TIME,omitempty"`
	DestinationID string  `json:"DESTINATIONID,omitempty"`
	PcArrival     string  `json:"PC_ARRIVAL,omitempty"`
	Comments      string  `json:"COMMENTS,omitempty"`
}

// Receivelist Declare
type Receivelist struct {
	ItemID        string `json:"itemId,omitempty"`
	BarcodeID     string `json:"BARCODEID,omitempty"`
	CCReceiveDate string `json:"CCRECEIVE_DATE,omitempty"`
	Type          string `json:"TYPE,omitempty"`
	Qty           uint64 `json:"QTY,omitempty"`
	PCReceiveDate string `json:"PCRECEIVE_DATE,omitempty"`
	PCReceiveTime string `json:"PCRECEIVE_TIME,omitempty"`
	STATUS        string `json:"STATUS,omitempty"`
}

// ProcessingCenterOrders Declare
type PCOrderAsset struct {
	DocType             string                `json:"docType,omitempty"`
	OrderID             string                `json:"ORDERID,omitempty"`
	OrderUnixTime       int64                 `json:"orderUnixTime,omitempty"`
	SourceID            string                `json:"SOURCEID,omitempty"`
	DestinationID       string                `json:"DESTINATIONID,omitempty"`
	SourceRoleID        string                `json:"SOURCE_ROLEID,omitempty"`
	DestinationRoleID   string                `json:"DESTINATION_ROLEID,omitempty"`
	ProduceID           string                `json:"PRID,omitempty"`
	Variety             string                `json:"variety,omitempty"`
	Produce             string                `json:"PRODUCE,omitempty"`
	Qty                 uint64                `json:"QTY,omitempty"`
	Status              string                `json:"STATUS,omitempty"`
	SourceState         string                `json:"SOURCE_STATE,omitempty"`
	SourceCountry       string                `json:"SOURCE_COUNTRY,omitempty"`
	DestinationState    string                `json:"DESTINATION_STATE,omitempty"`
	DestinationCountry  string                `json:"DESTINATION_COUNTRY,omitempty"`
	BaseUnit            string                `json:"BASE_UNIT,omitempty"`
	ParticipantTypeList []Participanttypelist `json:"PARTICIPANTTYPELIST,omitempty"`
	ProduceList         []Producelist         `json:"PRODUCELIST,omitempty"`
	ActualHarvest       string                `json:"ACTUAL_HARVEST,omitempty"`
	TotalavailableQty   uint64                `json:"TotalavailableQty,omitempty"`
	ReceiveType         string                `json:"RECEIVE_TYPE,omitempty"`
	FarmerID            string                `json:"FARMERID,omitempty"`
	SourceReceiveDate   string                `json:"SOURCE_RECEIVE_DATE,omitempty"`
	GrnNo               string                `json:"GRNNO,omitempty"`
	ReceiveList         []Receivelist         `json:"RECEIVELIST,omitempty"`
}
