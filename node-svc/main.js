// start express
const express = require("express");
const app = express();
const server = require("http").createServer(app);
const promClient = require("prom-client");

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

const cpuUsageGauge = new promClient.Gauge({
  name: "nodejs_cpu_usage",
  help: "CPU usage for Node.js service",
});

const memoryUsageGauge = new promClient.Gauge({
  name: "nodejs_memory_usage",
  help: "Memory usage for Node.js service",
});

app.get("/metrics", async function (req, res) {
  console.log("---- get metrics ---");
  const cpuUsage = process.cpuUsage().user / 1000000;
  cpuUsageGauge.set(cpuUsage);

  const memoryUsage = process.memoryUsage().rss / 1024 / 1024;
  memoryUsageGauge.set(memoryUsage);

  res.set("Content-Type", promClient.register.contentType);
  res.send(await promClient.register.metrics());
  res.end();
});

const PORT = 3001;
server.listen(PORT, async () => {
  console.log("Server listening at port %d", PORT);
  await initDB();
});
