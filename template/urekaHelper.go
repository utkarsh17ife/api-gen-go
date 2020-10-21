package template

import (
	"github.com/EYBlockchain/apigenerator/model"
)

func getHelperFileName(m model.Model) string {
	return m.Name + ".helper.js"
}

func GetPopulateString(m model.Model) string {
	s := ""
	for _, p := range m.Properties {
		isRefType, _ := checkIfRef(p.Type)
		if isRefType {
			// 	s += `.populate('` + getRefModelFromType(p.Type) + `')`
			// } else if isRefArrayType {
			// 	s += `.populate('` + getRefModelFromArrType(p.Type) + `')`
			s += `.populate('` + p.Name + `')`
		}
	}
	return s
}

func GetUrekaHelper(m model.Model) (string, string, string) {
	return getHelperFileName(m), `/src/controllers/helpers`, `import {` + m.Name + `Model } from "../../models/` + m.Name + `.model";

export const save = (data, passedQuery) => {
  return new Promise((resolve, reject) => {
    const query = passedQuery || { _id: data.Id };
    ` + m.Name + `Model.findOneOrCreate(query, data, (err, ` + m.Name + `) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(` + m.Name + `);
    });
  });
};

export const getById = id => {
 
    const query = { _id: id };
    return ` + m.Name + `Model.find(query)` + GetPopulateString(m) + `
 
};

export const getList = () => {
  return new Promise((resolve, reject) => {
    ` + m.Name + `Model.find({}, (err, result) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(result);
    });
  });
};

export const remove = id => {
  return new Promise((resolve, reject) => {
    ` + m.Name + `Model.deleteOne({ _id: id }, (err, result) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(result);
    });
  });
};

`
}
