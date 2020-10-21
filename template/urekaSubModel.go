package template

import (
	"fmt"
	"regexp"

	"github.com/EYBlockchain/apigenerator/model"
)

func getSubModelFileName(m model.Model) string {
	return m.Name + ".model.js"
}

func getCustomModelNameFromType(s string) string {

	arrType, err := regexp.Match(`^\[.*\]$`, []byte(s))
	if err != nil {
		panic("error parsing type: " + s)
	}
	if arrType {
		s := s[1:len(s)]
		return s[:len(s)-1]
	}
	return s

}

func getImports(m model.Model) string {

	s := ""
	for _, p := range m.Properties {

		isRefType, isRefArrayType := checkIfRef(p.Type)

		if !OneOfNormalType(p.Type) {
			if !isRefType && !isRefArrayType {
				s += fmt.Sprintf("import { %vSchema } from './%v.model';\n", getCustomModelNameFromType(p.Type), getCustomModelNameFromType(p.Type))
			}
		}

	}

	return s

}

func getRefModelFromType(t string) string {
	return t[4:len(t)]
}

func getRefModelFromArrType(t string) string {
	t = t[5:len(t)]
	return t[:len(t)-1]
}

func checkIfRef(s string) (bool, bool) {

	matched, err := regexp.Match(`^ref-*`, []byte(s))
	if err != nil {
		panic("error parsing type: " + s)
	}
	if matched {
		return true, false
	}
	arrMatched, err := regexp.Match(`^\[ref-*`, []byte(s))
	if err != nil {
		panic("error parsing type: " + s)
	}
	if arrMatched {
		return true, true
	}
	return false, false
}

func GetUrekaSubModelFile(m model.Model) (string, string, string) {
	return getSubModelFileName(m), `/src/models`,
		`
import mongoose from "mongoose";

const { Schema } = mongoose;

` + getImports(m) + `    

export const ` + m.Name + `Schema = new Schema(
  {
` + getModelProperties(m) + `    
  },
  {   
    strict: false
  }
);


export const ` + m.Name + `Model = mongoose.model("` + m.Name + `", ` + m.Name + `Schema);
`
}
