package models

import "github.com/msazad/go-Ecommerce/pkg/domain"

type Wallet struct{
	Balance int `json:"balance"`
	History []domain.WalletHistory `json:"history"`
}