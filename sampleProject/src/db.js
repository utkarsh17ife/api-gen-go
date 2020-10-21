import mongoose from "mongoose";

let db = null;

export default {
  connect: () => {
    const mongoDBURI = process.env.OFFCHAIN_DB_URL;

    mongoose.connect(mongoDBURI, {
      dbName: process.env.OFFCHAIN_DB_NAME,
      useNewUrlParser: true,
      useFindAndModify: false,
      useCreateIndex: true
    });

    mongoose.Promise = global.Promise;
    db = mongoose.connection;
    db.on("error", console.error.bind(console, "MongoDB connection error:"));
    db.on("connected", () => {
      console.info("Mongoose default connection open to "+mongoDBURI);
    });
  },

  get() {
    return db;
  }
};
