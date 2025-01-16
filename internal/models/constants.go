package models

// Constants để sử dụng trong code
const (
	ROLE_ADMIN  = "ADMIN"
	ROLE_SELLER = "SELLER"
	ROLE_BUYER  = "BUYER"

	DEFAULT_ROLE = ROLE_BUYER
)

// Permissions cho từng role
var (
	AdminPermissions = []string{
		"manage_users",
		"manage_stores",
		"manage_products",
		"manage_orders",
		"view_analytics",
		"manage_roles",
	}

	SellerPermissions = []string{
		"manage_own_store",
		"manage_own_products",
		"view_own_orders",
		"view_own_analytics",
	}

	BuyerPermissions = []string{
		"view_products",
		"manage_own_cart",
		"place_orders",
		"manage_own_profile",
	}
)
