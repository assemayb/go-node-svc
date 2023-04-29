package Axis.Transactions.demo.Model;


import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import org.bson.types.ObjectId;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.util.Date;

@Document(collection = "transactionDetails")
@Getter
@AllArgsConstructor
@NoArgsConstructor
public class TransactionModel {

    @Id
    private ObjectId id;

    @Field("txnDetailsId")
    private String txnDetailsId;

    @Field("createdAt")
    private Date createdAt;

    @Field("lastModified")
    private Date lastModified;

    @Field("receiverWalletNumber")
    private String receiverWalletNumber;

    @Field("receiverName")
    private String receiverName;

    @Field("receiverScheme")
    private String receiverScheme;

    @Field("receiverAddress")
    private String receiverAddress;

    @Field("senderWalletNumber")
    private String senderWalletNumber;

    @Field("senderName")
    private String senderName;

    @Field("senderScheme")
    private String senderScheme;

    @Field("senderAddress")
    private String senderAddress;

    @Field("description")
    private String description;

    @Field("isAdvised")
    private boolean isAdvised;

    @Field("txnType")
    private String txnType;

    @Field("transferId")
    private String transferId;

    @Field("transferTransactionId")
    private String transferTransactionId;

    @Field("transactionReference")
    private String transactionReference;

    @Field("txnRequestedAmount")
    private double txnRequestedAmount;

    @Field("transactionDomain")
    private String transactionDomain;

    @Field("transactionStatus")
    private String transactionStatus;

    @Field("transactionAction")
    private String transactionAction;

    @Field("commitTs")
    private Date commitTs;

    @Field("transactionChannel")
    private String transactionChannel;

    @Field("walletDetailsId")
    private String walletDetailsId;

    @Field("walletBalanceBefore")
    private double walletBalanceBefore;

    @Field("walletBalanceAfter")
    private double walletBalanceAfter;

    @Field("isServiceFeesTxn")
    private boolean isServiceFeesTxn;

    @Field("hasServiceFees")
    private boolean hasServiceFees;

    @Field("serviceFees")
    private double serviceFees;

    @Field("serviceFeesTxnDetailsId")
    private String serviceFeesTxnDetailsId;

    @Field("isConfirmationAdvisedFlag")
    private boolean isConfirmationAdvisedFlag;

    @Field("onUsTransactionId")
    private String onUsTransactionId;

    @Field("isReversedFlag")
    private boolean isReversedFlag;

    @Field("transactionCurrency")
    private String transactionCurrency;

    // getters and setters omitted for brevity
}