package template

import "github.com/EYBlockchain/apigenerator/model"

func getControllerFileName(m model.Model) string {
	return m.Name + ".controller.js"
}

func GetUrekaControllerFile(m model.Model) (string, string, string) {
	return getControllerFileName(m), `/src/controllers`, `import { getById, getList, save } from "./helpers/` + m.Name + `.helper";

export const Create` + m.FirstLetterCaps() + ` = data => {
  return save(data);
};

export const Get` + m.FirstLetterCaps() + `ById = id => {
  return getById(id);
};

export const Get` + m.FirstLetterCaps() + `List = () => {
  return getList();
};

export const Update` + m.FirstLetterCaps() + ` = (id, data) => {
  return save(data, { _id: id });
};

`
}
