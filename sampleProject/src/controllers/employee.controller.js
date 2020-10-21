import { getById, getList, save } from "./helpers/employee.helper";

export const CreateEmployee = data => {
  return save(data);
};

export const GetEmployeeById = id => {
  return getById(id);
};

export const GetEmployeeList = () => {
  return getList();
};

export const UpdateEmployee = (id, data) => {
  return save(data, { _id: id });
};

