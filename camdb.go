package camdb

//import "time"

// Tx .
type Tx struct {
	ID         int64     `xorm:"pk autoincr"`
	TX         string    `xorm:"notnull index"`
	From       string    `xorm:"index(from_to)"`
	To         string    `xorm:"index(from_to)"`
	Asset      string    `xorm:"notnull"`
	Value      string    `xorm:"notnull"`
	Block      uint64    `xorm:"notnull index"`
	CreateTime int64     `xorm:"notnull"`
}

// TableName .
func (tx *Tx) TableName() string {
	return "cam_tx"
}

// Block .
type Block struct {
	ID         int64     `xorm:"pk autoincr"`
	Block      int64     `xorm:"notnull index"`
	SysFee     float64   `xorm:"notnull"`
	NetFee     float64   `xorm:"notnull"`
	CreateTime int64     `xorm:"notnull"`
}

// TableName .
func (tx *Block) TableName() string {
	return "cam_block"
}

// UTXO .
type UTXO struct {
	ID          int64      `xorm:"pk autoincr"`
	TX          string     `xorm:"notnull index(vout)"`
	N           int        `xorm:"notnull index(vout)"`
	Address     string     `xorm:"notnull index(unclaimed)"`
	CreateBlock int64      `xorm:"notnull index"`
	SpentBlock  int64      `xorm:"notnull index default (-1)"`
	Asset       string     `xorm:"notnull index(unclaimed)"`
	Value       string     `xorm:"notnull"`
	CreateTime  int64      `xorm:"notnull"`
	SpentTime   int64      `xorm:"TIMESTAMP"`
	Claimed     bool       `xorm:"index(unclaimed)"`
}

// TableName .
func (table *UTXO) TableName() string {
	return "cam_utxo"
}

// Order .
type Order struct {
	ID          int64      `json:"-" xorm:"pk autoincr"`
	TX          string     `json:"tx" xorm:"index notnull"`
	From        string     `json:"from" xorm:"index(from_to)"`
	To          string     `json:"to" xorm:"index(from_to)"`
	Asset       string     `json:"asset" xorm:"notnull"`
	Value       string     `json:"value" xorm:"notnull"`
	Block       int64      `json:"blocks" xorm:"default (-1)"`
	CreateTime  int64      `json:"createTime" xorm:"notnull"`
	ConfirmTime int64      `json:"confirmTime" xorm:"notnull"`
	Status      int64      `json:"status" xorm:"default (1)"`
}

// TableName xorm table name
func (table *Order) TableName() string {
	return "cam_order"
}

// Wallet .
type Wallet struct {
	ID         int64     `xorm:"pk autoincr"`
	Address    string    `xorm:"index(address_userid)"`
	CreateTime int64     `xorm:"notnull"`
}

// TableName xorm table name
func (table *Wallet) TableName() string {
	return "cam_wallet"
}
