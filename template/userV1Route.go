package template

func GetUserRouteFile() (string, string, string) {
	return `user.route.js`, `/src/routes/v1`, `import express from "express";
import { GetUserById, Login, SignUp } from "../../controllers/user.controller";
import { verifyJWT } from "../../middleware/jwt.middleware";

const router = express.Router();

router.post("/signup", async (req, res, next) => {
  let result;

  try {
    result = await SignUp(req.body.data);
    result.password = undefined;
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(201).send({ data: result });
});

router.post("/login", async (req, res, next) => {
  let result;

  try {
    result = await Login(req.body.data);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

router.get("/", verifyJWT, async (req, res, next) => {
  let result;

  try {
    result = await GetUserById(req.decoded.userId);
    result.password = undefined;
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

export default router;`
}
