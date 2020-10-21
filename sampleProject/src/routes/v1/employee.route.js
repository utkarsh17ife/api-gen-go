import express from "express";
import {
  CreateEmployee,
  GetEmployeeById,
  GetEmployeeList,
  UpdateEmployee
} from "../../controllers/employee.controller";

const router = express.Router();

router.post("/", async (req, res, next) => {
  let result;

  try {
    result = await CreateEmployee(req.body.data);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(201).send({ data: result });
});

router.get("/list", async (req, res, next) => {
  let result;

  try {
    result = await GetEmployeeList();
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

router.get("/:employeeId", async (req, res, next) => {
  let result;

  try {
    result = await GetEmployeeById(req.params.employeeId);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

router.put("/:employeeId", async (req, res, next) => {
  let result;

  try {
    result = await UpdateEmployee(req.params.employeeId, req.body.data);
  } catch (err) {
    res.status(500);
    return next(err);
  }

  return res.status(200).send({ data: result });
});

export default router;