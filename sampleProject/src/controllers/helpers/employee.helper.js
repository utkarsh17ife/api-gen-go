import {employeeModel } from "../../models/employee.model";

export const save = (data, passedQuery) => {
  return new Promise((resolve, reject) => {
    const query = passedQuery || { _id: data.Id };
    employeeModel.findOneOrCreate(query, data, (err, employee) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(employee);
    });
  });
};

export const getById = id => {
 
    const query = { _id: id };
    return employeeModel.find(query)
 
};

export const getList = () => {
  return new Promise((resolve, reject) => {
    employeeModel.find({}, (err, result) => {
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
    employeeModel.deleteOne({ _id: id }, (err, result) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(result);
    });
  });
};

