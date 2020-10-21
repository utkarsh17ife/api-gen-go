import { getById, getList, save } from "./helpers/company.helper";

export const CreateCompany = data => {
  return save(data);
};

export const GetCompanyById = id => {
  return getById(id);
};

export const GetCompanyList = () => {
  return getList();
};

export const UpdateCompany = (id, data) => {
  return save(data, { _id: id });
};

