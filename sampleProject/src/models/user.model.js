/* eslint-disable func-names */
import bcrypt from "bcrypt";
import mongoose from "mongoose";

const { Schema } = mongoose;

const userSchema = new Schema(
  {
    userName: { type: String, required: true, unique: true },
    password: { type: String, required: true, bcrypt: true },
    email: { type: String },
    createdAt: { type: Date },
    updatedAt: { type: Date }
  },
  {
    collection: "user",
    strict: false
  }
);

userSchema.plugin(require("mongoose-bcrypt"));
userSchema.plugin(require("mongoose-unique-validator"));

userSchema.pre("save", function(next) {
  const currentDate = new Date();
  this.updatedAt = currentDate;

  if (!this.createdAt) {
    this.createdAt = currentDate;
  }

  next();
});

userSchema.pre("findOneAndUpdate", function(next) {
  this.updatedAt = new Date();
  next();
});

userSchema.statics.verifyPassword = function(query, password, callback) {
  const self = this;
  self.find(query, (err, result) => {
    if (err) {
      callback(err, false);
      return;
    }
    if (result.length === 0) {
      callback(new Error("User not found in the system"), false);
      return;
    }

    bcrypt.compare(password, result[0].password, (err2, isMatch) => {
      if (err2) {
        callback(err, null);
        return;
      }
      if (isMatch) {
        callback(null, result[0]);
        return;
      }
      callback(new Error("Invalid username/password"), null);
    });
  });
};

userSchema.statics.findOneOrCreate = function findOneOrCreate(
  query,
  data,
  callback
) {
  const _data = data;
  const self = this;
  self.find(query, (err, result) => {
    if (err) {
      callback(err, result);
      return;
    }
    if (result.length === 0) {
      self.create(_data, (err2, result2) => {
        if (err2) {
          callback(err2, result2);
          return;
        }
        callback(err2, result2);
      });
    } else {
      _data.updatedAt = new Date();
      self.findOneAndUpdate(query, _data, { upsert: true }, (err3, result3) => {
        if (err3) {
          callback(err3, result3);
          return;
        }
        callback(err3, result3);
      });
    }
  });
};

const userModel = mongoose.model("user", userSchema);
export default userModel;

