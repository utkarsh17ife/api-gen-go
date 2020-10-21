/* eslint-disable func-names */
import mongoose from "mongoose";

const { Schema } = mongoose;

import { addressSchema } from './address.model';
    

export const employeeSchema = new Schema(
  {
    employeeName: { type: String , default: 'somename', required: true } ,
    contactNum: String,
    dateOfBirth: Date,
    permanentAddress:addressSchema,
    createdAt: { type: Date },
    updatedAt: { type: Date }
  },
  {
    collection: "employee",
    strict: false
  }
);

employeeSchema.pre("save", function(next) {
  const currentDate = new Date();
  this.updatedAt = currentDate;

  if (!this.createdAt) {
    this.createdAt = currentDate;
  }

  next();
});

employeeSchema.statics.findOneOrCreate = function findOneOrCreate(
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
      self.findOneAndUpdate(query, _data, { new: true }, (err3, result3) => {
        if (err) {
          callback(err3, result3);
          return;
        }
        callback(err3, result3);
      });
    }
  });
};

export const employeeModel = mongoose.model("employee", employeeSchema);
