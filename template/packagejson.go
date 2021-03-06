package template

func GetPackageJson() (string, string, string) {
	return "package.json", `/`, `{
  "name": "project_name",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "start": "node build/index.js",
    "dev": "nodemon build/index.js",
    "build": "babel src -d build",
    "watch": "babel src -d build --watch",
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "@babel/cli": "^7.4.3",
    "@babel/core": "^7.4.3",
    "@babel/preset-env": "^7.4.3"
  },
  "dependencies": {
    "bcrypt": "^3.0.5",
    "body-parser": "^1.18.3",
    "cors": "^2.8.5",
    "express": "^4.16.4",
    "jsonwebtoken": "^8.5.1",
    "lodash": "^4.17.11",
    "mongodb": "^3.2.3",
    "mongoose": "^5.5.1",
    "mongoose-bcrypt": "^1.6.0",
    "mongoose-unique-validator": "^2.0.2",
    "nodemon": "^1.18.11"
  }
}`

}
