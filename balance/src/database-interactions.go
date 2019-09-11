package main

import (
	"context"
	"math/rand"
	"time"

	"./proto"
)

type AccountsModel struct {
	ID string `gorm:"type:varchar(20);PRIMARY_KEY;UNIQUE;NOT NULL"`
	CreatedAt time.Time `gorm:"type:timestamp;DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	ModifiedAt time.Time `gorm:"type:timestamp;DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	BalanceAmount int64 `gorm:"NOT NULL"`
	Name string `gorm:"type:varchar(64);NOT NULL"`
}

/*type ScyllaAccountsModel struct {
	ID string
	CreatedAt time.Time
	ModifiedAt time.Time
	BalanceAmount int64
	Name string
}*/

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateID(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func (s *server) RecordTransaction(ctx context.Context, request *proto.Transaction) (*proto.TransactionResponse, error) {
	from := AccountsModel{}
	if s.Database.Where("ID = ?", request.SenderID).First(&from).Error != nil {
		return &proto.TransactionResponse{
			Status:               false,
			Message:              "Account " + request.SenderID + " does not exist.",
			Transaction:          nil,
		}, nil
	}
	to := AccountsModel{}
	if s.Database.Where("ID = ?", request.ReceiverID).First(&to).Error != nil {
		return &proto.TransactionResponse{
			Status:               false,
			Message:              "Account " + request.ReceiverID + " does not exist.",
			Transaction:          nil,
		}, nil
	}
	newFromBalance := from.BalanceAmount - request.Amount
	newToBalance := to.BalanceAmount + request.Amount
	s.Database.Model(&from).Where("ID = ?", request.SenderID).Updates(map[string]interface{}{
		"balance_amount": 		newFromBalance,
		"modified_at": 			time.Now(),
	})
	s.Database.Model(&to).Where("ID = ?", request.ReceiverID).Updates(map[string]interface{}{
		"balance_amount": newToBalance,
		"modified_at": time.Now(),
	})
	return &proto.TransactionResponse{
		Status:               true,
		Message:              "Transaction successful.",
		Transaction:          &proto.Transaction{
			ID:                   request.ID,
			SenderID:             request.SenderID,
			ReceiverID:           request.ReceiverID,
			CreatedAt:            request.CreatedAt,
			ModifiedAt:           request.ModifiedAt,
			Description:          request.Description,
			Amount:               request.Amount,
			Currency:             request.Currency,
			Note:                 request.Note,
		},
	}, nil
}

func (s *server) CreateAccount(ctx context.Context, request *proto.Account) (*proto.AccountResponse, error) {
	account := AccountsModel{
		ID:            GenerateID(20),
		CreatedAt:     time.Now(),
		ModifiedAt:    time.Now(),
		BalanceAmount: request.BalanceAmount,
		Name:          request.Name,
	}
	for s.Database.NewRecord(&account) {
		account.ID = GenerateID(20)
	}
	s.Database.Create(&account)
	return &proto.AccountResponse{
		Status: 		true,
		Message: 		"Account successfully created.",
		Account: 		&proto.Account{
			ID:                   account.ID,
			BalanceAmount:        account.BalanceAmount,
			Name:                 account.Name,
			CreatedAt:            account.CreatedAt.Format("01-02-2006 15:04:05"),
			ModifiedAt:           account.ModifiedAt.Format("01-02-2006 15:04:05"),
		},
	}, nil
}

func (s *server) DebitAccount(ctx context.Context, request *proto.AccountAction) (*proto.AccountResponse, error) {
	account := AccountsModel{}
	if s.Database.Where("ID = ?", request.AccountID).First(&account).Error != nil {
		return &proto.AccountResponse{
			Status: 		false,
			Message: 		"Account " + request.AccountID + " does not exist.",
			Account: 		nil,
		}, nil
	}
	s.Database.Model(&AccountsModel{}).Where("ID = ?", request.AccountID).Update(map[string]interface{}{
		"balance_amount": account.BalanceAmount - request.Amount,
		"modified_at": time.Now(),
	})
	return &proto.AccountResponse{
		Status: 		true,
		Message: 		"Account successfully debited.",
		Account: 		&proto.Account{
			ID:                   account.ID,
			BalanceAmount:        account.BalanceAmount - request.Amount,
			Name:                 account.Name,
			CreatedAt:            account.CreatedAt.Format("01-02-2006 15:04:05"),
			ModifiedAt:           account.ModifiedAt.Format("01-02-2006 15:04:05"),
		},
	}, nil
}

func (s *server) CreditAccount(ctx context.Context, request *proto.AccountAction) (*proto.AccountResponse, error) {
	account := AccountsModel{}
	if s.Database.Where("ID = ?", request.AccountID).First(&account).Error != nil {
		return &proto.AccountResponse{
			Status: 		false,
			Message: 		"Account " + request.AccountID + " does not exist.",
			Account: 		nil,
		}, nil
	}
	s.Database.Model(&AccountsModel{}).Where("ID = ?", request.AccountID).Update(map[string]interface{}{
		"balance_amount": account.BalanceAmount + request.Amount,
		"modified_at": time.Now(),
	})
	return &proto.AccountResponse{
		Status: 		true,
		Message: 		"Account successfully credited.",
		Account: 		&proto.Account{
			ID:                   account.ID,
			BalanceAmount:        account.BalanceAmount + request.Amount,
			Name:                 account.Name,
			CreatedAt:            account.CreatedAt.Format("01-02-2006 15:04:05"),
			ModifiedAt:           account.ModifiedAt.Format("01-02-2006 15:04:05"),
		},
	}, nil
}

func (s *server) GetAccount(ctx context.Context, request *proto.AccountAction) (*proto.AccountResponse, error) {
	account := AccountsModel{}
	if s.Database.Where("ID = ?", request.AccountID).First(&account).Error != nil {
		return &proto.AccountResponse{
			Status: 		false,
			Message: 		"Account " + request.AccountID + " does not exist.",
			Account: 		nil,
		}, nil
	}
	return &proto.AccountResponse{
		Status: 		true,
		Message: 		"Account successfully retrieved.",
		Account: 		&proto.Account{
			ID:                   account.ID,
			BalanceAmount:        account.BalanceAmount,
			Name:                 account.Name,
			CreatedAt:            account.CreatedAt.Format("01-02-2006 15:04:05"),
			ModifiedAt:           account.ModifiedAt.Format("01-02-2006 15:04:05"),
		},
	}, nil
}