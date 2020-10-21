package template

func GetIndexFile() (string,string, string) {
	return `index.js`, `/src`,`import app from "./app";
import db from "./db";

/**
 * Bootstrap the application. Start express.
 */
const main = async () => {
  try {
    await db.connect();

    app.listen(process.env.PORT);
  } catch (err) {
    console.error(err);
    process.exit(1);
  }
};

main();`

}
