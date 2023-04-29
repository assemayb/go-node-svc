const mongoose = require("mongoose");

let _db;

const initDB = async () => {
  const mongoBinding = { url: process.env["MONGODB_CONNECT_STRING"] };
  try {
    const connection = await mongoose.connect(mongoBinding.url, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
    });
    _db = connection;
    console.log("Connected to MongoDB");
  } catch (error) {
    console.log("Error connecting to MongoDB", error);
    throw error;
  }
};

const getDb = () => _db;

module.exports = {
  initDB,
  getDb,
};
