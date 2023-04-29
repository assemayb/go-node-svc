package models

import "time"

type Transaction struct {
	TxnDetailsId                   string    `json:"txnDetailsId" bson:"txnDetailsId"`
	CreatedAt                      time.Time `json:"createdAt" bson:"createdAt"`
	LastModifiedAt                 string    `json:"updatedAt" bson:"updatedAt"`
	ReceiverWalletNumber           string    `json:"receiverWalletNumber" bson:"receiverWalletNumber"`
	SenderWalletNumber             string    `json:"senderWalletNumber" bson:"senderWalletNumber"`
	ReceiverName                   string    `json:"receiverName" bson:"receiverName"`
	ReceiverScheme                 string    `json:"receiverScheme" bson:"receiverScheme"`
	ReceiverAddress                string    `json:"receiverAddress" bson:"receiverAddress"`
	SenderName                     string    `json:"senderName" bson:"senderName"`
	SenderScheme                   string    `json:"senderScheme" bson:"senderScheme"`
	SenderAddress                  string    `json:"senderAddress" bson:"senderAddress"`
	Description                    string    `json:"description" bson:"description"`
	MeezaResponseCode              string    `json:"meezaResponseCode" bson:"meezaResponseCode"`
	MeezaResponseDescription       string    `json:"meezaResponseDescription" bson:"meezaResponseDescription"`
	MeezaAdviceResponseCode        string    `json:"meezaAdviceResponseCode" bson:"meezaAdviceResponseCode"`
	MeezaAdviceResponseDescription string    `json:"meezaAdviceResponseDescription" bson:"meezaAdviceResponseDescription"`
	AdviceId                       string    `json:"adviceId" bson:"adviceId"`
	AdviceTs                       time.Time `json:"adviceTs" bson:"adviceTs"`
	IsAdvised                      bool      `json:"isAdvised" bson:"isAdvised"`
	TxnType                        string    `json:"txnType" bson:"txnType"`
	TransactionReference           string    `json:"transactionReference" bson:"transactionReference"`
	TxnRequestedAmount             float64   `json:"txnRequestedAmount" bson:"txnRequestedAmount"`
	TransactionDomain              string    `json:"transactionDomain" bson:"transactionDomain"`
	TransactionStatus              string    `json:"transactionStatus" bson:"transactionStatus"`
	TransactionAction              string    `json:"transactionAction" bson:"transactionAction"`
	TransactionSubStatus           string    `json:"transactionSubStatus" bson:"transactionSubStatus"`
	MeezaTransactionId             string    `json:"meezaTransactionId" bson:"meezaTransactionId"`
	CommitTs                       time.Time `json:"commitTs" bson:"commitTs"`
	TransactionChannel             string    `json:"transactionChannel" bson:"transactionChannel"`
	WalletDetailsId                string    `json:"walletDetailsId" bson:"walletDetailsId"`
	WalletBalanceBefore            float64   `json:"walletBalanceBefore" bson:"walletBalanceBefore"`
	WalletBalanceAfter             float64   `json:"walletBalanceAfter" bson:"walletBalanceAfter"`
}
