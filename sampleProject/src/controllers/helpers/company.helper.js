import {companyModel } from "../../models/company.model";

export const save = (data, passedQuery) => {
  return new Promise((resolve, reject) => {
    const query = passedQuery || { _id: data.Id };
    companyModel.findOneOrCreate(query, data, (err, company) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(company);
    });
  });
};

export const getById = id => {
 
    const query = { _id: id };
    return companyModel.find(query).populate('employees')
 
};

export const getList = () => {
  return new Promise((resolve, reject) => {
    companyModel.find({}, (err, result) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(result);
    });
  });
};

export const remove = id => {
  return new Promise((resolve, reject) => {
    companyModel.deleteOne({ _id: id }, (err, result) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(result);
    });
  });
};

