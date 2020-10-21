package template

import (
	"fmt"

	"github.com/EYBlockchain/apigenerator/model"
)

func getModelFileName(m model.Model) string {
	return m.Name + ".model.js"
}

func getModelProperties(m model.Model) string {
	s := ""
	for _, p := range m.Properties {

		if OneOfNormalType(p.Type) {
			if p.Default != "" || p.Required == true {
				s += fmt.Sprintf("    %v: { type: %v , default: '%v', required: %v } ,\n", p.Name, p.Type, p.Default, p.Required)
			} else {
				s += fmt.Sprintf("    %v: %v,\n", p.Name, p.Type)
			}
		} else {
			isRefType, isRefArrayType := checkIfRef(p.Type)
			if isRefType && !isRefArrayType {
				s += fmt.Sprintf("    %v: { type:  Schema.Types.ObjectId , ref: '%v' } ,\n", p.Name, getRefModelFromType(p.Type))
			} else if isRefType && isRefArrayType {
				s += fmt.Sprintf("    %v: [{ type:  Schema.Types.ObjectId , ref: '%v' }] ,\n", p.Name, getRefModelFromArrType(p.Type))
			} else {
				s += fmt.Sprintf("    %v:%vSchema,\n", p.Name, getCustomModelNameFromType(p.Type))
			}
		}

	}
	return s[:len(s)-1]
}

// type: String, Number, [String], [Number], object, [object], refId, [refId], date, bool (in this case the objectString should be provided)
func OneOfNormalType(str string) bool {
	if str == "String" || str == "Number" || str == "[String]" || str == "[Number]" || str == "Boolean" || str == "Date" {
		return true
	}
	return false
}

func GetUrekaModelFile(m model.Model) (string, string, string) {
	return getModelFileName(m), `/src/models`, `/* eslint-disable func-names */
import mongoose from "mongoose";

const { Schema } = mongoose;

` + getImports(m) + `    

export const ` + m.Name + `Schema = new Schema(
  {
` + getModelProperties(m) + `
    createdAt: { type: Date },
    updatedAt: { type: Date }
  },
  {
    collection: "` + m.Name + `",
    strict: false
  }
);

` + m.Name + `Schema.pre("save", function(next) {
  const currentDate = new Date();
  this.updatedAt = currentDate;

  if (!this.createdAt) {
    this.createdAt = currentDate;
  }

  next();
});

` + m.Name + `Schema.statics.findOneOrCreate = function findOneOrCreate(
  query,
  data,
  callback
) {
  const _data = data;
  const self = this;
  self.find(query, (err, result) => {
    if (err) {
      callback(err, result);
      return;
    }
    if (result.length === 0) {
      self.create(_data, (err2, result2) => {
        if (err2) {
          callback(err2, result2);
          return;
        }
        callback(err2, result2);
      });
    } else {
      _data.updatedAt = new Date();
      self.findOneAndUpdate(query, _data, { new: true }, (err3, result3) => {
        if (err) {
          callback(err3, result3);
          return;
        }
        callback(err3, result3);
      });
    }
  });
};

export const ` + m.Name + `Model = mongoose.model("` + m.Name + `", ` + m.Name + `Schema);
`
}
