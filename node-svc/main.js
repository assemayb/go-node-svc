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

app.get("/txns", async function (req, res) {
  console.log("----- list all txns -----");
  const rps = req.query.rps || 15000;
  const filter = {};
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
  name: "nodejs_svc_cpu_usage",
  help: "CPU usage for Node.js service",
});

const memoryUsageGauge = new promClient.Gauge({
  name: "nodejs_svc_memory_usage",
  help: "Memory usage for Node.js service",
});

const primaryCpuUsage = process.cpuUsage();

app.get("/metrics", async function (req, res) {
  console.log("---- get metrics ---");

  const cpuUsage = process.cpuUsage(primaryCpuUsage);
  const userUsageInMicroSeconds = cpuUsage.user;
  const userUsageInSeconds = userUsageInMicroSeconds / 1000000;
  cpuUsageGauge.set(userUsageInSeconds);
  const memoryUsage = process.memoryUsage();
  memoryUsageGauge.set(memoryUsage.heapUsed);
  res.set("Content-Type", promClient.register.contentType);
  res.end(await promClient.register.metrics());
});

const PORT = 3001;
server.listen(PORT, async () => {
  console.log("Server listening at port %d", PORT);
  await initDB();
});
