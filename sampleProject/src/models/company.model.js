/* eslint-disable func-names */
import mongoose from "mongoose";

const { Schema } = mongoose;

import { addressSchema } from './address.model';
    

export const companySchema = new Schema(
  {
    companyname: { type: String , default: 'ey', required: true } ,
    dateOfRegistration: Date,
    officialAddress:addressSchema,
    employees: [{ type:  Schema.Types.ObjectId , ref: 'employee' }] ,
    revenue: String,
    ceo: String,
    createdAt: { type: Date },
    updatedAt: { type: Date }
  },
  {
    collection: "company",
    strict: false
  }
);

companySchema.pre("save", function(next) {
  const currentDate = new Date();
  this.updatedAt = currentDate;

  if (!this.createdAt) {
    this.createdAt = currentDate;
  }

  next();
});

companySchema.statics.findOneOrCreate = function findOneOrCreate(
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

export const companyModel = mongoose.model("company", companySchema);
