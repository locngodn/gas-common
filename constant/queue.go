package constant

const (
	CreateStaffQueueName   = "account.create_staff"
	UpdateStaffQueueName   = "account.update_staff"
	ActiveStaffQueueName   = "account.active_staff"
	DeactiveStaffQueueName = "account.deactive_staff"

	CreateAgencyQueueName   = "account.create_agency"
	UpdateAgencyQueueName   = "account.update_agency"
	ActiveAgencyQueueName   = "account.active_agency"
	DeactiveAgencyQueueName = "account.deactive_agency"

	CreateCustomerQueueName       = "account.create_customer"
	UpdateCustomerQueueName       = "account.update_customer"
	ActiveCustomerQueueName       = "account.active_customer"
	DeactiveCustomerQueueName     = "account.deactive_customer"
	AuthenticateCustomerQueueName = "account.authenticate_customer"
	GetAccountRoleQueueName       = "account.get_account_role"

	CreateStoreQueueName   = "account.create_store"
	UpdateStoreQueueName   = "account.update_store"
	ActiveStoreQueueName   = "account.active_store"
	DeactiveStoreQueueName = "account.deactive_store"

	UpdateRoleStaffsQueueName            = "operation.update_role_staffs"
	UpdateAgencyAccountQueueName         = "operation.update_agency_account"
	UpdateBranchAccountQueueName         = "operation.update_branch_account"
	UpdateStaffAccountQueueName          = "operation.update_staff_account"
	UpdateCustomerAccountQueueName       = "operation.update_customer_account"
	UpdateStoreAccountQueueName          = "operation.update_store_account"
	UpdatePointReferralCustomerQueueName = "operation.update_point_referral_customer"

	GetAgencyOrderInfoQueueName       = "operation.get_agency_order_info"
	GetOrderInfoQueueName             = "operation.get_order_info"
	GetGoodsReceiptIssueInfoQueueName = "operation.get_goods_receipt_issue_info"
	GetStoresQueueName                = "operation.get_stores"
	GetUsersQueueName                 = "operation.get_users"
	GetRetailInfoQueueName            = "operation.get_retail_info"
	SyncStoreRatingQueueName          = "operation.sync.store_rating"

	SendNotificationsByUsersQueueName     = "notification.send_notifications_by_users"
	SendNotificationsByFcmTokensQueueName = "notification.send_notifications_by_fcm_tokens"
	SendRetailNotificationQueueName       = "notification.send_retail_notifications"
	SendOrderNotificationQueueName        = "notification.send_order_notifications"
	SendAgencyOrderNotificationQueueName  = "notification.send_order_notifications"
	SendSmsueueName                       = "notification.send_sms"
	UpdateUserSocketStatisticQueueName    = "notification.update_user_socket_statistic"

	SyncStoreDeliveryQueueName = "core.sync.store_delivery"

	DEAD_LETTER         = "_DEAD_LETTER"
	ExchangeTypeFanout  = "fanout"
	ExchangeTypeDirect  = "direct"
	ExchangeTypeHeaders = "headers"
	ExchangeTypeTopic   = "topic"

	SyncStaffStationWarehouseExchangeName = "sync.staff_station_warehouse"
	SyncStaffExchangeName                 = "sync.staff"
	SyncStoreExchangeName                 = "sync.store"
	SyncCustomerExchangeName              = "sync.customer"
	SyncAgencyExchangeName                = "sync.agency"
	SyncBranchExchangeName                = "sync.branch"

	BroadcastSocketRetailMessageExchangeName      = "socket.broadcast.retail_message"
	BroadcastSocketOrderMessageExchangeName       = "socket.broadcast.order_message"
	BroadcastSocketAgencyOrderMessageExchangeName = "socket.broadcast.agency_order_message"

	RandomQueuePrefix = "amq.gen-"
)
