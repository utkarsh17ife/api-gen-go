import bodyParser from "body-parser";
import cors from "cors";
import express from "express";
import {
  defaultErrorHandler,
  validateContentType,
  validateJSONSyntax
} from "./middleware/errorHandling";

import userV1Routes from "./routes/v1/user.route";

import companyV1Routes from "./routes/v1/company.route";
import employeeV1Routes from "./routes/v1/employee.route";




const app = express();

// cors & body parser middleware should come before any routes are handled
app.use(cors({ exposedHeaders: ["Total-Count", "Report-Total"] }));
app.use(bodyParser.json({ limit: "2mb" }));
app.use(bodyParser.urlencoded({ limit: "2mb", extended: false }));

(function defaultEnv() {
  process.env.OFFCHAIN_DB_URL =
    process.env.OFFCHAIN_DB_URL || "mongodb://127.0.0.1:27017";
  process.env.PORT = process.env.PORT || 3000;
  process.env.JWT_SECRET = process.env.JWT_SECRET || "verysecret";
  process.env.OFFCHAIN_DB_NAME = process.env.OFFCHAIN_DB_NAME || "templatebase";
})();

// catch body parser json errors. must come directly after body parser
app.use(validateJSONSyntax);

app.use("/v1/user", userV1Routes);

app.use("/v1/company", companyV1Routes)
app.use("/v1/employee", employeeV1Routes)



app.use(validateContentType);
app.use(defaultErrorHandler);

export default app;