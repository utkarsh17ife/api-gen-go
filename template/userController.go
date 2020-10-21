package template

func GetUserControllerFile() (string, string, string) {
	return `user.controller.js`, `/src/controllers`, `import { createToken } from "../middleware/jwt.middleware";
import { authenticate, getById, save } from "./helpers/user.helper";

export const SignUp = data => {
  return save(data);
};

export const Login = async data => {
  let user;
  try {
    user = await authenticate(data);
    return {
      userName: user.userName,
      auth_token: createToken({
        userId: user._id
      })
    };
  } catch (err) {
    throw new Error(err.message);
  }
};

export const GetUserById = id => {
  return getById(id);
};

`
}
