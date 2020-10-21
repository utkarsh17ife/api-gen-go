package template

import "github.com/EYBlockchain/apigenerator/model"

func getRouteImportStrings(m []model.Model) string {
	s := ""
	for _, m := range m {
		if m.Api {
			s += `import ` + m.Name + `V1Routes from "./routes/v1/` + m.Name + `.route";` + "\n"
		}		
	}
	return s
}

func getRouteInitStrings(m []model.Model) string {
	s := ""
	for _, m := range m {
		if m.Api {
			s += `app.use("/v1/` + m.Name + `", ` + m.Name + `V1Routes)` + "\n"
		}
	}
	return s
}

func GetAppFile(m []model.Model) (string, string, string) {

	return `app.js`, `/src`, `import bodyParser from "body-parser";
import cors from "cors";
import express from "express";
import {
  defaultErrorHandler,
  validateContentType,
  validateJSONSyntax
} from "./middleware/errorHandling";

import userV1Routes from "./routes/v1/user.route";

` + getRouteImportStrings(m) + `



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

` + getRouteInitStrings(m) + `


app.use(validateContentType);
app.use(defaultErrorHandler);

export default app;`

}
