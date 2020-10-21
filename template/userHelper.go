package template

func GetUserHelper() (string, string, string) {
	return `user.helper.js`, `/src/controllers/helpers`, `import userModel from "../../models/user.model";

export const save = (data, passedQuery) => {
  return new Promise((resolve, reject) => {
    const query = passedQuery || { _id: data.Id };
    userModel.findOneOrCreate(query, data, (err, result) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(result);
    });
  });
};

export const authenticate = data => {
  return new Promise((resolve, reject) => {
    const query = { userName: data.userName };
    userModel.verifyPassword(query, data.password, (err, result) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(result);
    });
  });
};

export const getById = id => {
  return new Promise((resolve, reject) => {
    const query = { _id: id };
    userModel.find(query, (err, data) => {
      if (err) {
        reject(err);
        return;
      }
      if (data && data.length > 0) {
        resolve(data[0]);
        return;
      }
      resolve(false);
    });
  });
};
`
}
