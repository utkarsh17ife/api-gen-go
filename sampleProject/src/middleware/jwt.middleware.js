const jwt = require("jsonwebtoken");

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
