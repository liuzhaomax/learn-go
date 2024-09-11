package pay

import "time"

type Merchant struct {
	MerchantID uint64    `gorm:"primaryKey;autoIncrement"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	Platform   string    `gorm:"type:enum('WECHAT', 'ALIPAY');not null"`
	// wechat
	MchID    string `gorm:"size:32;not null"`
	AppID    string `gorm:"size:32;not null"`
	APIKey   string `gorm:"size:64"`
	CertPath string `gorm:"size:256"`
	KeyPath  string `gorm:"size:256"`
	Sandbox  bool   `gorm:"default:false"`
	// alipay
	AlipayAppID      string `gorm:"size:32"`
	AlipayPrivateKey string `gorm:"type:text"`
	AlipayPublicKey  string `gorm:"type:text"`
}

type Order struct {
	OrderID     uint64    `gorm:"primaryKey;autoIncrement"`
	UserID      uint64    `gorm:"not null"`
	MerchantID  uint64    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	OrderNumber string    `gorm:"size:64;unique;not null"`
	TotalAmount float64   `gorm:"type:decimal(10,2);not null"`
	Currency    string    `gorm:"size:10;default:'CNY'"`
	Description string    `gorm:"size:255"`
	Status      string    `gorm:"type:enum('PENDING', 'PAID', 'FAILED', 'CANCELLED');default:'PENDING'"`
}

type Transaction struct {
	TransactionID     uint64    `gorm:"primaryKey;autoIncrement"`
	OrderID           uint64    `gorm:"not null"`
	MerchantID        uint64    `gorm:"not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
	TransactionNumber string    `gorm:"size:64;unique"`
	Platform          string    `gorm:"type:enum('WECHAT', 'ALIPAY');not null"`
	TradeType         string    `gorm:"type:enum('JSAPI', 'NATIVE', 'APP', 'WEB', 'WAP');not null"`
	PaymentMethod     string    `gorm:"type:enum('WECHAT_PAY', 'ALIPAY');not null"`
	Amount            float64   `gorm:"type:decimal(10,2);not null"`
	Currency          string    `gorm:"size:10;default:'CNY'"`
	Status            string    `gorm:"type:enum('SUCCESS', 'REFUND', 'NOTPAY', 'CLOSED', 'REVOKED', 'USERPAYING', 'PAYERROR');default:'NOTPAY'"`
	CreateTime        time.Time // 实际支付交易创建时间
	PaymentTime       time.Time // 实际支付交易发生时间
	RefundTime        time.Time // 实际支付交易退款时间
	ExtraData         string    `gorm:"type:json"`
	// wechat
	WeChatTransactionNo string `gorm:"size:64"` // 微信支付交易号
	PayerOpenID         string `gorm:"size:64"`
	BankType            string `gorm:"size:32"`
	// alipay
	AlipayTradeNo string `gorm:"size:64"`
	AlipayBuyerID string `gorm:"size:64"`
}

type Refund struct {
	RefundID       uint64    `gorm:"primaryKey;autoIncrement"`
	TransactionID  uint64    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	RefundNumber   string    `gorm:"size:64;unique;not null"`
	Platform       string    `gorm:"type:enum('WECHAT', 'ALIPAY');not null"`
	RefundAmount   float64   `gorm:"type:decimal(10,2);not null"`
	Currency       string    `gorm:"size:10;default:'CNY'"`
	Status         string    `gorm:"type:enum('SUCCESS', 'FAILED', 'PROCESSING');default:'PROCESSING'"`
	Reason         string    `gorm:"size:255"`
	AlipayRefundNo string    `gorm:"size:64;unique"` // alipay独有
	RefundTime     time.Time
}

type PaymentLog struct {
	LogID         uint64 `gorm:"primaryKey;autoIncrement"`
	OrderID       uint64
	TransactionID uint64
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	Platform      string    `gorm:"type:enum('WECHAT', 'ALIPAY');default:'WECHAT'"`
	Level         string    `gorm:"type:enum('INFO', 'WARNING', 'ERROR');default:'INFO'"`
	Message       string
}

type IdempotencyKey struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`                              // 生成时间
	ExpiresAt      time.Time `gorm:"not null"`                                    // 过期时间
	IdempotencyKey string    `gorm:"size:255;not null;unique;index:idx_idem_key"` // 幂等键，必须唯一
	Result         string    `gorm:"type:text"`                                   // 请求处理的结果（可以是 JSON、XML 等格式）
}
