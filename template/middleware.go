package template

func GetErrorMiddleWare() (string, string, string) {
	return `errorHandling.js`, `/src/middleware`, `
export function defaultErrorHandler(err, req, res, next) {
  if (res.headersSent) {
    return next(err);
  }

  console.error(err);
  return res
    .status((res.statusCode > 200 && res.statusCode) || 500)
    .send({ message: err.message });
}

export function validateContentType(req, res, next) {
  if (
    req.method !== "GET" &&
    req.headers["content-type"] !== "application/json"
  ) {
    res.status(400).json({ message: "Body should be a JSON object" });
  } else {
    next();
  }
}

export function validateJSONSyntax(err, req, res, next) {
  if (err instanceof SyntaxError && err.status === 400) {
    res.status(400).json({ message: "Problems parsing JSON" });
  }

  return next();
}
`
}

func GetJWTMiddleWare() (string, string, string) {
	return `jwt.middleware.js`, `/src/middleware`, `const jwt = require("jsonwebtoken");

export function verifyJWT(req, res, next) {
  const token = req.body.token || req.query.token || req.headers.authorization;
  if (token) {
    jwt.verify(token, process.env.JWT_SECRET, function cb(err, decoded) {
      if (err) {
        res.json({
          message: "Failed to authenticate token"
        });
        return;
      }
      req.decoded = decoded;
      next();
    });
  } else {
    res.status(403).send({
      message: "No token provided."
    });
  }
}

export function createToken(details) {
  return jwt.sign(details, process.env.JWT_SECRET, { expiresIn: "100 days" });
}
`
}
