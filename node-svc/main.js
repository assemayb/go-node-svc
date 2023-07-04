// start express
const os = require("os");
const express = require("express");
const pidusage = require("pidusage");

const app = express();
const server = require("http").createServer(app);
const promClient = require("prom-client");

const dotenv = require("dotenv");
dotenv.config();
const { initDB } = require("./database/config.js");

app.use(express.json({ limit: "50mb" }));
app.use(express.urlencoded({ limit: "50mb" }));

const TransactionModel = require("./models/Transaction.js");

app.get("/txns", async function (req, res) {
  // console.log("----- list all txns -----");
  const rps = req.query.rps || 15000;
  const filter = {};
  const transactions = await TransactionModel.find(filter).limit(rps);
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

setInterval(() => {
  pidusage(process.pid, (err, stats) => {
    if (err) {
      console.error(err);
      return;
    }
    const memoryUsg = stats.memory / 1024 / 1024;
    const cpuUsg = stats.cpu;
    cpuUsageGauge.set(cpuUsg);
    memoryUsageGauge.set(memoryUsg);
    console.log(`CPU: ${cpuUsg}%, Memory: ${memoryUsg} MB`);
  });
}, 1000);

app.get("/metrics", async function (req, res) {
  console.log("---- get metrics ---");
  res.set("Content-Type", promClient.register.contentType);
  res.end(await promClient.register.metrics());
});

const PORT = 3001;
server.listen(PORT, async () => {
  console.log("Server listening at port %d", PORT);
  await initDB();
});
