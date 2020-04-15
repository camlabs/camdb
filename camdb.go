package camdb

//import "time"

// Tx .
type Tx struct {
	ID         int     `xorm:"pk autoincr"`
	TX         string    `xorm:"'t_x' notnull index(t_x_from_to_asset)"`
	From       string    `xorm:"'from' notnull index(t_x_from_to_asset)"`
	To         string    `xorm:"'to' notnull index(t_x_from_to_asset)"`
	Asset      string    `xorm:"'asset' notnull index(t_x_from_to_asset)"`
	Value      string    `xorm:"'value' notnull"`
	Block      uint64    `xorm:"'block' notnull"`
	CreateTime int     `xorm:"'create_time' notnull"`
	TxIndex    int     `xorm:"'tx_index' notnull default (0)"`
}

// TableName .
func (tx *Tx) TableName() string {
	return "cam_tx"
}

// Block .
type Block struct {
	ID         int     `xorm:"pk autoincr"`
	Block      int64     `xorm:"'block' notnull index(block_index)"`
	SysFee     float32   `xorm:"'sys_fee' notnull"`
	NetFee     float32   `xorm:"'net_fee' notnull"`
	CreateTime int     `xorm:"'create_time' notnull"`
}

// TableName .
func (tx *Block) TableName() string {
	return "cam_block"
}

// UTXO .
type UTXO struct {
	ID          int      `xorm:"pk autoincr"`
	TX          string     `xorm:"'t_x' notnull index(t_x_n)"`
	N           int        `xorm:"'n' notnull index(t_x_n)"`
	Address     string     `xorm:"'address' notnull"`
	CreateBlock int64      `xorm:"'create_block'"`
	SpentBlock  int64      `xorm:"'spent_block' default (-1)"`
	Asset       string     `xorm:"'asset' notnull"`
	Value       string     `xorm:"'value' notnull"`
	CreateTime  int      `xorm:"'create_time' notnull"`
	SpentTime   int      `xorm:"'spent_time' null"`
	Claimed     bool       `xorm:"'claimed' notnull"`
}

// TableName .
func (table *UTXO) TableName() string {
	return "cam_utxo"
}

// Order .
type Order struct {
	ID          int      `xorm:"pk autoincr"`
	TX          string     `xorm:"'t_x' index(t_x_from_to_asset) notnull"`
	From        string     `xorm:"'from' index(t_x_from_to_asset)"`
	To          string     `xorm:"'to' index(t_x_from_to_asset)"`
	Asset       string     `xorm:"'asset' index(t_x_from_to_asset) notnull"`
	Value       string     `xorm:"'value' notnull"`
	Block       int64      `xorm:"'block' default (-1)"`
	CreateTime  int      `xorm:"'create_time'"`
	ConfirmTime int      `xorm:"'confirm_time'"`
	Status      int      `xorm:"'status' notnull default (1)"`
	TxIndex     int      `xorm:"'tx_index' notnull default (0)"`
}

// TableName xorm table name
func (table *Order) TableName() string {
	return "cam_order"
}

// Wallet .
type Wallet struct {
	ID         int     `xorm:"pk autoincr"`
	Address    string    `xorm:"'address' index(address)"`
	CreateTime int     `xorm:"'create_time' notnull"`
}

// TableName xorm table name
func (table *Wallet) TableName() string {
	return "cam_wallet"
}
