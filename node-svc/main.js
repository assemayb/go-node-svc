// start express
const express = require("express");
const app = express();
const server = require("http").createServer(app);

const dotenv = require("dotenv");
dotenv.config();
const { initDB, getDb } = require("./database/config.js");

app.use(express.json({ limit: "50mb" }));
app.use(express.urlencoded({ limit: "50mb" }));

const TransactionModel = require("./models/Transaction.js");

app.post("/txns", async function (req, res) {
  console.log("list all txns");
  const rps = req.query.rps || 15000;
  console.log("rps", rps);
  const filter = {};

  if (req.body.senderWalletNumber) {
    filter.senderWalletNumber = req.body.senderWalletNumber;
  }
  const transactions = await TransactionModel.find(filter).limit(rps);
  console.log(transactions && transactions.length);
  res.status(200).json(transactions);
});

app.get("/txns/:id", async function (req, res) {
  console.log("get txn by id");
  const transaction = await TransactionModel.findOne({
    txnDetailsId: req.params.id,
  });
  res.status(200).json(transaction);
});

server.listen(3000, async () => {
  console.log("Server listening at port %d", 3000);
  await initDB();
});
