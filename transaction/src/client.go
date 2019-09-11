package main

import (
	"./proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Describe transaction table for database
type TransactionModel struct {
	ID string `gorm:"type:varchar(20);PRIMARY_KEY;UNIQUE;NOT NULL"`
	SenderID string `gorm:"type:varchar(20);NOT NULL"`
	ReceiverID string `gorm:"type:varchar(20);NOT NULL"`
	CreatedAt time.Time `gorm:"type:timestamp;DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	ModifiedAt time.Time `gorm:"type:timestamp;DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	Description string `gorm:"type:varchar(255)"`
	Amount int64 `gorm:"NOT NULL"`
	Currency string `gorm:"type:varchar(10);NOT NULL"`
	Note string `gorm:"type:varchar(255)"`
}

// Insert a transaction into the database
func CreateTransaction(db *gorm.DB, t *TransactionModel) {
	for db.NewRecord(&t) {
		t.ID = GenerateID(20)
	}
	db.Create(&t)
}

// Generate random ID
func GenerateID(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func main() {
	fmt.Printf("[Client] Transaction service starting...\n")
	connection, err := grpc.Dial("balance:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCommunicationClient(connection)
	defer connection.Close()
	g := gin.Default()
	db, err := gorm.Open("mysql", "root:password@tcp(db:3306)/balance?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if !db.HasTable(&TransactionModel{}) {
		db.CreateTable(&TransactionModel{})
	}

	// Send an account for creation to the balance that will insert it into database
	g.POST("/create-account", func (ctx *gin.Context) {
		balanceAmount, err := strconv.ParseInt(ctx.PostForm("balance_amount"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter balance_amount"})
			return
		}
		name := ctx.PostForm("name")

		req := &proto.Account{Name: name, BalanceAmount: balanceAmount}

		if response, err := client.CreateAccount(context.Background(), req); err == nil {
			if !response.Status {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": response.Message,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"Message": response.Message,
					"AccountID": response.Account.ID,
					"BalanceAmount": response.Account.BalanceAmount,
					"Name": response.Account.Name,
					"CreatedAt": response.Account.CreatedAt,
					"ModifiedAt": response.Account.ModifiedAt,
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Send credit instruction to balance for a given account, then store the transaction into the database
	g.POST("/credit-account", func (ctx *gin.Context) {
		accountID := ctx.PostForm("accountID")
		amount, err := strconv.ParseInt(ctx.PostForm("amount"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter amount"})
			return
		}
		req := &proto.AccountAction{
			Action:               proto.AccountAction_CREDIT,
			AccountID:            accountID,
			Amount:               amount,
		}
		if response, err := client.CreditAccount(context.Background(), req); err == nil {
			if !response.Status {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": response.Message,
				})
			} else {
				CreateTransaction(db, &TransactionModel{
					ID:          GenerateID(20),
					SenderID:    "admin",
					ReceiverID:  accountID,
					CreatedAt:   time.Now(),
					ModifiedAt:  time.Now(),
					Description: "admin credited " + accountID + " an amount of " + string(amount) + "€",
					Amount:      amount,
					Currency:    "€",
					Note:        "Money generated from nowhere",
				})
				ctx.JSON(http.StatusOK, gin.H{
					"Message": response.Message,
					"AccountID": response.Account.ID,
					"BalanceAmount": response.Account.BalanceAmount,
					"Name": response.Account.Name,
					"CreatedAt": response.Account.CreatedAt,
					"ModifiedAt": response.Account.ModifiedAt,
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Send debit instruction to balance for a given account, then store the transaction into the database
	g.POST("/debit-account", func (ctx *gin.Context) {
		accountID := ctx.PostForm("accountID")
		amount, err := strconv.ParseInt(ctx.PostForm("amount"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter amount"})
			return
		}
		req := &proto.AccountAction{
			Action:               proto.AccountAction_DEBIT,
			AccountID:            accountID,
			Amount:               amount,
		}
		if response, err := client.DebitAccount(context.Background(), req); err == nil {
			if !response.Status {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": response.Message,
				})
			} else {
				CreateTransaction(db, &TransactionModel{
					ID:          GenerateID(20),
					SenderID:    "admin",
					ReceiverID:  accountID,
					CreatedAt:   time.Now(),
					ModifiedAt:  time.Now(),
					Description: "admin debited " + accountID + " an amount of: " + string(amount) + "€",
					Amount:      amount,
					Currency:    "€",
					Note:        "Money disappeared by the power of admin.",
				})
				ctx.JSON(http.StatusOK, gin.H{
					"Message": response.Message,
					"AccountID": response.Account.ID,
					"BalanceAmount": response.Account.BalanceAmount,
					"Name": response.Account.Name,
					"CreatedAt": response.Account.CreatedAt,
					"ModifiedAt": response.Account.ModifiedAt,
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Send the transaction to the balance and save it into the database
	g.POST("/make-transaction", func (ctx *gin.Context) {
		fmt.Print("[Client]Making Transaction...\n")
		senderID := ctx.PostForm("senderID")
		receiverID := ctx.PostForm("receiverID")
		description := ctx.PostForm("description")
		note := ctx.PostForm("note")
		amount, err := strconv.ParseInt(ctx.PostForm("amount"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter amount"})
			return
		}
		now := time.Now()
		req := &proto.Transaction{
			ID:                   GenerateID(20),
			SenderID:             senderID,
			ReceiverID:           receiverID,
			CreatedAt:            now.Format("01-02-2006 15:04:05"),
			ModifiedAt:           now.Format("01-02-2006 15:04:05"),
			Description:          description,
			Amount:               amount,
			Currency:             "€",
			Note:                 note,
		}
		if response, err := client.RecordTransaction(context.Background(), req); err == nil {
			if !response.Status {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": response.Message,
				})
			} else {
				CreateTransaction(db, &TransactionModel{
					ID:                   req.ID,
					SenderID:             req.SenderID,
					ReceiverID:           req.ReceiverID,
					CreatedAt:            now,
					ModifiedAt:           now,
					Description:          req.Description,
					Amount:               req.Amount,
					Currency:             req.Currency,
					Note:                 req.Note,
				})
				ctx.JSON(http.StatusOK, gin.H{
					"Message": response.Message,
					"TransactionID": response.Transaction.ID,
					"SenderID": response.Transaction.SenderID,
					"ReceiverID": response.Transaction.ReceiverID,
					"CreatedAt": response.Transaction.CreatedAt,
					"ModifiedAt": response.Transaction.ModifiedAt,
					"Description": response.Transaction.Description,
					"Amount": response.Transaction.Amount,
					"Currency": response.Transaction.Currency,
					"Note": response.Transaction.Note,
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Get information about an account
	g.POST("/get-account", func (ctx *gin.Context) {
		accountID := ctx.PostForm("accountID")
		req := &proto.AccountAction{
			Action:               proto.AccountAction_LOOK,
			AccountID:            accountID,
			Amount:               0,
		}
		if response, err := client.GetAccount(context.Background(), req); err == nil {
			if !response.Status {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": response.Message,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"Message": response.Message,
					"AccountID": response.Account.ID,
					"BalanceAmount": response.Account.BalanceAmount,
					"Name": response.Account.Name,
					"CreatedAt": response.Account.CreatedAt,
					"ModifiedAt": response.Account.ModifiedAt,
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Update a transaction's note or description
	g.POST("/update-transaction", func (ctx *gin.Context) {
		accountID := ctx.PostForm("accountID")
		note := ctx.PostForm("note")
		description := ctx.PostForm("description")

		transaction := TransactionModel{}

		if db.Where("ID = ?", accountID).First(&transaction).Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Account " + accountID + " does not exist.",
			})
		}
		db.Model(&TransactionModel{}).Where("ID = ?", accountID).Updates(map[string]interface{}{
			"note": 			note,
			"description": 		description,
			"modified_at":		time.Now(),
		})
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Transaction updated",
			"TransactionID": transaction.ID,
			"SenderID": transaction.SenderID,
			"ReceiverID": transaction.ReceiverID,
			"CreatedAt": transaction.CreatedAt,
			"ModifiedAt": transaction.ModifiedAt,
			"Description": description,
			"Amount": transaction.Amount,
			"Currency": transaction.Currency,
			"Note": note,
		})
	})

	log.Fatal(http.ListenAndServe(":8080", g))
}