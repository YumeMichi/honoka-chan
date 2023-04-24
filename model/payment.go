package model

// RestrictionInfo ...
type RestrictionInfo struct {
	Restricted bool `json:"restricted"`
}

// UnderAgeInfo ...
type UnderAgeInfo struct {
	BirthSet    bool        `json:"birth_set"`
	HasLimit    bool        `json:"has_limit"`
	LimitAmount interface{} `json:"limit_amount"`
	MonthUsed   int         `json:"month_used"`
}

// SnsProductItemList ...
type SnsProductItemList struct {
	ItemID    int  `json:"item_id"`
	AddType   int  `json:"add_type"`
	Amount    int  `json:"amount"`
	IsFreebie bool `json:"is_freebie"`
}

// SnsProductList ...
type SnsProductList struct {
	ProductID   string               `json:"product_id"`
	Name        string               `json:"name"`
	Price       int                  `json:"price"`
	CanBuy      bool                 `json:"can_buy"`
	ProductType int                  `json:"product_type"`
	ItemList    []SnsProductItemList `json:"item_list"`
}

// ProductItemList ...
type ProductItemList struct {
	ItemID    int  `json:"item_id"`
	AddType   int  `json:"add_type"`
	Amount    int  `json:"amount"`
	IsFreebie bool `json:"is_freebie"`
	IsRankMax bool `json:"is_rank_max,omitempty"`
}

// LimitStatus ...
type LimitStatus struct {
	TermStartDate  string `json:"term_start_date"`
	RemainingTime  string `json:"remaining_time"`
	RemainingCount int    `json:"remaining_count"`
}

// ProductList ...
type ProductList struct {
	ProductID      string            `json:"product_id"`
	Name           string            `json:"name"`
	BannerImgAsset string            `json:"banner_img_asset"`
	Price          int               `json:"price"`
	CanBuy         bool              `json:"can_buy"`
	ProductType    int               `json:"product_type"`
	AnnounceURL    string            `json:"announce_url"`
	ConfirmURL     string            `json:"confirm_url"`
	ItemList       []ProductItemList `json:"item_list"`
	LimitStatus    LimitStatus       `json:"limit_status"`
}

// SubscriptionItemList ...
type SubscriptionItemList struct {
	ItemID    int  `json:"item_id"`
	AddType   int  `json:"add_type"`
	Amount    int  `json:"amount"`
	IsFreebie bool `json:"is_freebie"`
}

// RewardList ...
type RewardList struct {
	ItemID  int `json:"item_id"`
	AddType int `json:"add_type"`
	Amount  int `json:"amount"`
}

// Items ...
type Items struct {
	Seq        int          `json:"seq"`
	RewardList []RewardList `json:"reward_list"`
}

// LicenseInfo ...
type LicenseInfo struct {
	Name  string  `json:"name"`
	Items []Items `json:"items"`
}

// UserStatus ...
type UserStatus struct {
	IsLicensed bool `json:"is_licensed"`
}

// SubscriptionStatus ...
type SubscriptionStatus struct {
	LicenseID     int         `json:"license_id"`
	LicenseType   int         `json:"license_type"`
	LicenseInfo   LicenseInfo `json:"license_info,omitempty"`
	UserStatus    UserStatus  `json:"user_status"`
	PurchaseCount int         `json:"purchase_count"`
	BadgeFlag     bool        `json:"badge_flag"`
}

// SubscriptionList ...
type SubscriptionList struct {
	ProductID          string                 `json:"product_id"`
	Name               string                 `json:"name"`
	BannerImgAsset     string                 `json:"banner_img_asset"`
	Price              int                    `json:"price"`
	CanBuy             bool                   `json:"can_buy"`
	ProductType        int                    `json:"product_type"`
	ProductURL         string                 `json:"product_url"`
	ItemList           []SubscriptionItemList `json:"item_list"`
	LimitStatus        LimitStatus            `json:"limit_status"`
	SubscriptionStatus SubscriptionStatus     `json:"subscription_status"`
}

// ProductListRes ...
type ProductListRes struct {
	RestrictionInfo  RestrictionInfo    `json:"restriction_info"`
	UnderAgeInfo     UnderAgeInfo       `json:"under_age_info"`
	SnsProductList   []SnsProductList   `json:"sns_product_list"`
	ProductList      []ProductList      `json:"product_list"`
	SubscriptionList []SubscriptionList `json:"subscription_list"`
	ShowPointShop    bool               `json:"show_point_shop"`
}

// ProductListResp ...
type ProductListResp struct {
	Result     ProductListRes `json:"result"`
	Status     int            `json:"status"`
	CommandNum bool           `json:"commandNum"`
	TimeStamp  int64          `json:"timeStamp"`
}
