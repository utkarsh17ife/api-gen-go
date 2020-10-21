import express from "express";
import {
  CreateCompany,
  GetCompanyById,
  GetCompanyList,
  UpdateCompany
} from "../../controllers/company.controller";

const router = express.Router();

router.post("/", async (req, res, next) => {
  let result;

  try {
    result = await CreateCompany(req.body.data);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(201).send({ data: result });
});

router.get("/list", async (req, res, next) => {
  let result;

  try {
    result = await GetCompanyList();
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

router.get("/:companyId", async (req, res, next) => {
  let result;

  try {
    result = await GetCompanyById(req.params.companyId);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

router.put("/:companyId", async (req, res, next) => {
  let result;

  try {
    result = await UpdateCompany(req.params.companyId, req.body.data);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

export default router;