const mongoose = require("mongoose");
const TransactionDetailsSchema = new mongoose.Schema(
  {
    txnDetailsId: {
      type: String,
      required: true,
      unique: true,
      index: true,
    },

    createdAt: {
      type: Date,
    },

    lastModified: {
      type: Date,
    },

    receiverWalletNumber: {
      type: String,
    },

    receiverName: {
      type: String,
    },

    receiverScheme: {
      type: String,
    },

    receiverAddress: {
      type: String,
    },

    senderWalletNumber: {
      type: String,
    },

    senderName: {
      type: String,
    },

    senderScheme: {
      type: String,
    },

    senderAddress: {
      type: String,
    },

    description: {
      type: String,
    },

    meezaResponseCode: {
      type: String,
    },

    meezaResponseDescription: {
      type: String,
    },

    meezaAdviceResponseCode: {
      type: String,
    },

    meezaAdviceResponseDescription: {
      type: String,
    },

    adviceId: {
      type: String,
    },

    adviceTs: {
      type: Date,
    },

    isAdvised: {
      type: Boolean,
    },

    txnType: {
      type: String,
    },

    transactionReference: {
      type: String,
    },

    txnRequestedAmount: {
      type: Number,
    },

    transactionDomain: {
      type: String,
    },

    transactionStatus: {
      type: String,
    },

    transactionAction: {
      type: String,
    },

    transactionSubStatus: {
      type: String,
    },

    meezaTransactionId: {
      type: String,
    },

    commitTs: {
      type: Date,
    },

    transactionChannel: {
      type: String,
    },

    walletDetailsId: {
      type: String,
    },

    walletBalanceBefore: {
      type: Number,
    },

    walletBalanceAfter: {
      type: Number,
    },

    deviceId: {
      type: String,
    },

    isServiceFeesTxn: {
      type: Boolean,
    },

    hasServiceFees: {
      type: Boolean,
    },

    serviceFees: {
      type: Number,
    },

    serviceFeesTxnDetailsId: {
      type: String,
    },

    interchangeAction: {
      type: String,
    },

    interchangeCurrency: {
      type: String,
    },

    interchangeAmount: {
      type: Number,
    },

    orderingAgencyAddress: {
      type: String,
    },

    orderingAgencyCountry: {
      type: String,
    },

    orderingAgencyName: {
      type: String,
    },

    orderingCustomerName: {
      type: String,
    },

    orderingCustomerAddress: {
      type: String,
    },

    orderingCustomerIdType: {
      type: String,
    },

    orderingCustomerId: {
      type: String,
    },

    orderingCustomerNationality: {
      type: String,
    },

    orderingCustomerBd: {
      type: Date,
    },

    extRefNum: {
      type: String,
    },

    meezaConfirmResponseCode: {
      type: String,
    },

    meezaConfirmResponseDescription: {
      type: String,
    },

    meezaConfirmAdviceResponseCode: {
      type: String,
    },

    meezaConfirmAdviceResponseDescription: {
      type: String,
    },

    confirmAdviceId: {
      type: String,
    },

    confirmAdviceTs: {
      type: Date,
    },

    isConfirmationAdvisedFlag: {
      type: Boolean,
    },

    consumerOnUsConfirmationTs: {
      type: Date,
    },

    referencedTxnId: {
      type: String,
    },

    onUsTransactionId: {
      type: String,
    },

    isReversedFlag: {
      type: Boolean,
    },

    refundedAt: {
      type: Date,
    },

    transactionCurrency: {
      type: String,
    },

    installmentCollectionOrderId: {
      type: String,
    },

    p2mCategory: {
      type: String,
    },

    p2MInitiationType: {
      type: String,
    },

    request2PayId: {
      type: String,
    },

    reference1: {
      type: String,
    },

    reference2: {
      type: String,
    },

    tipsAmount: {
      type: Number,
    },

    qrCode: {
      type: String,
    },

    tipsCurrency: {
      type: String,
    },

    convenienceAmount: {
      type: Number,
    },

    convCurrency: {
      type: String,
    },

    senderCode: {
      type: String,
    },

    txnIconUrl: {
      type: String,
    },

    terminalId: {
      type: String,
    },

    terminalLocation: {
      type: String,
    },

    terminalLongitude: {
      type: Number,
    },

    terminalLatitude: {
      type: Number,
    },
    walletActualBalanceBefore: {
      type: Number,
    },
    walletActualBalanceAfter: {
      type: Number,
    },
  },
  { timestamps: { updatedAt: "lastModifiedTs", createdAt: "createdTs" } },
  { collection: "transactionDetails" }
);

const TransactionModel = mongoose.model(
  "transactionDetails",
  TransactionDetailsSchema,
  "transactionDetails"
);

module.exports = TransactionModel;
