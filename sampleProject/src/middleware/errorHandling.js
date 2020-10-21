
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
