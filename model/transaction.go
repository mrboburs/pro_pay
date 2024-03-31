package model

import "github.com/google/uuid"

type Transaction struct {
	ID       uuid.UUID `json:"id"  db:"id"`
	Amount     float64    `json:"amount" db:"amount"`
	Account    float64        `json:"account" db:"account"`
	TerminalId     int        `json:"terminal_id" db:"terminal_id"`	
	StoreId     int        `json:"store_id" db:"store_id"`
}

type CreateTransaction struct {
	Amount     float64    `json:"amount" db:"amount"`
	Account    string       `json:"account" db:"account"`
	TerminalId     string       `json:"-" db:"terminal_id"`	
	StoreId     string       `json:"store_id" db:"store_id"`

}

type UpdateTransaction struct {
	ID       uuid.UUID `json:"id"  db:"id"`
	Amount     float64    `json:"amount" db:"amount"`
	Account    float64        `json:"account" db:"account"`
	TerminalId     int        `json:"terminal_id" db:"terminal_id"`	
	StoreId     int        `json:"store_id" db:"store_id"`
}


// Store represents the store object
type Store struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Logo  string `json:"logo"`
	GaID  string `json:"ga_id"`
}

// Label represents the label object
type Label struct {
	TypeID   int    `json:"type_id"`
	LabelRU  string `json:"label_ru"`
	LabelUZ  string `json:"label_uz"`
	LabelEN  string `json:"label_en"`
}

// StoreTransaction represents the store transaction object
type StoreTransaction struct {
	SuccessTransID   interface{} `json:"success_trans_id"`
	TransID          int         `json:"trans_id"`
	Store            Store       `json:"store"`
	TerminalID       string      `json:"terminal_id"`
	Account          string      `json:"account"`
	Amount           int         `json:"amount"`
	Confirmed        bool        `json:"confirmed"`
	PrepayTime       interface{} `json:"prepay_time"`
	ConfirmTime      interface{} `json:"confirm_time"`
	Label            Label       `json:"label"`
	Details          interface{} `json:"details"`
	CommissionValue  string      `json:"commission_value"`
	CommissionType   string      `json:"commission_type"`
	Total            int         `json:"total"`
	CardID           interface{} `json:"card_id"`
	StatusCode       interface{} `json:"status_code"`
	StatusMessage    interface{} `json:"status_message"`
}

// Result represents the result object
type Result struct {
	ID uuid.UUID `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Response represents the response object
type Response struct {
	Result          Result          `json:"result"`
	TransactionID   int             `json:"transaction_id"`
	StoreTransaction StoreTransaction `json:"store_transaction"`
}
