package constant

const (
	EmptyString = ""

	StatusCode = "status_code"
	StatusText = "status_text"

	Language = "language"

	ASC  = "ASC"
	DESC = "DESC"

	TYP = "type"

	PROD  = "production"
	DEV   = "development"
	LOCAL = "local"

	MonthFormat = "01/2006"

	DateFormat    = "02/01/2006"
	SqlDateFormat = "2006-01-02"

	TimeStampFormat    = "02/01/2006 15:04:05"
	SqlTimeStampFormat = "2006-01-02 15:04:05.999999Z07"

	UpdateAtRegex  = `\[\"(.+)\",\)`
	SysPeriodRegex = `\[\"(.+)\",\"(.+)\"\)`

	StationWarehouse     = "stationWarehouse"
	StationWarehouseType = "stationWarehouseType"
	Agency               = "agency"
	Branch               = "branch"
	Search               = "search"
	Role                 = "role"
	Type                 = "type"
	Status               = "status"
	Id                   = "id"
	Province             = "province"
	District             = "district"
	Ward                 = "ward"
	Phone                = "phone"
	RouteType            = "routeType"
	Driver               = "driver"
	LicensePlate         = "licensePlate"
	Vehicle              = "vehicle"
	Route                = "route"
	Location             = "location"
	SerialNo             = "serialNo"
	Barcode              = "barcode"
	Supplier             = "supplier"
	MfgFrom              = "mfgFrom"
	MfgTo                = "mfgTo"
	Mfg                  = "mfg"
	DateFrom             = "dateFrom"
	DateTo               = "dateTo"
	Date                 = "date"
	NumOfLifeCycleFrom   = "numOfLifeCycleFrom"
	NumOfLifeCycleTo     = "numOfLifeCycleTo"
	Brand                = "brand"
	ConsumerType         = "consumerType"
	Category             = "category"
	Market               = "market"
	DiscountMin          = "discountMin"
	DiscountMax          = "discountMax"
	GiftName             = "giftName"
	Name                 = "name"
	ApplyDateFrom        = "applyDateFrom"
	ApplyDateTo          = "applyDateTo"
	Manufactor           = "manufactor"
	DeliveryDateFrom     = "deliveryDateFrom"
	DeliveryDateTo       = "deliveryDateTo"
	GoodsReceiptIssue    = "goodsReceiptIssue"
	Source               = "source"
	Store                = "store"
	PaymentMethod        = "paymentMethod"
	PaymentStatus        = "paymentStatus"
	ShippingType         = "shippingType"
	IssueDate            = "issueDate"
	IssueDateFrom        = "issueDateFrom"
	IssueDateTo          = "issueDateTo"
	Staff                = "staff"
	Customer             = "customer"
	Code                 = "code"
	ProductSeq           = "productSeq"
	OrderSeq             = "orderSeq"
	Product              = "product"
	Voucher              = "voucher"
	Supporter            = "supporter"
	Delivery             = "delivery"
	Quantity             = "quantity"
	OrderBy              = "orderBy"
	Sort                 = "sort"
	HasReferralCode      = "hasReferralCode"

	MaxPagingSize   = 100
	MaxPagingOffset = 1000000

	MaxPagingSizeQueue = 5000

	RequestTimeout = "Request timeout"
)
