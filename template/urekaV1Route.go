package template

import (
	"github.com/EYBlockchain/apigenerator/model"
)

func getRouteFileName(m model.Model) string {
	return m.Name + ".route.js"
}

func GetUrekaRouteFile(m model.Model) (string, string, string) {
	return getRouteFileName(m), `/src/routes/v1`, `import express from "express";
import {
  Create` + m.FirstLetterCaps() + `,
  Get` + m.FirstLetterCaps() + `ById,
  Get` + m.FirstLetterCaps() + `List,
  Update` + m.FirstLetterCaps() + `
} from "../../controllers/` + m.Name + `.controller";

const router = express.Router();

router.post("/", async (req, res, next) => {
  let result;

  try {
    result = await Create` + m.FirstLetterCaps() + `(req.body.data);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(201).send({ data: result });
});

router.get("/list", async (req, res, next) => {
  let result;

  try {
    result = await Get` + m.FirstLetterCaps() + `List();
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

router.get("/:` + m.Name + `Id", async (req, res, next) => {
  let result;

  try {
    result = await Get` + m.FirstLetterCaps() + `ById(req.params.` + m.Name + `Id);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

router.put("/:` + m.Name + `Id", async (req, res, next) => {
  let result;

  try {
    result = await Update` + m.FirstLetterCaps() + `(req.params.` + m.Name + `Id, req.body.data);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

export default router;`

}
