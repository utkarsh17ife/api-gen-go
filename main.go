package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/EYBlockchain/apigenerator/model"
	"github.com/EYBlockchain/apigenerator/template"
)

var projectName string

func main() {

	//read project.json
	projectModel := ReadProjectModel()

	projectName = projectModel.Name
	//loop in the model and create the files
	for _, m := range projectModel.Models {
		log.Println("creating stuff for " + m.Name + " model")

		if m.Api {
			//create route file
			PlaceCode(template.GetUrekaRouteFile(m))
			//create controller file
			PlaceCode(template.GetUrekaControllerFile(m))
			//create helper file
			PlaceCode(template.GetUrekaHelper(m))
			//create model file
			PlaceCode(template.GetUrekaModelFile(m))
		} else {
			//create model file
			PlaceCode(template.GetUrekaSubModelFile(m))
		}

	}

	PlaceCode(template.GetAppFile(projectModel.Models))
	PlaceCode(template.GetDBConnectionFile())
	PlaceCode(template.GetPackageJson())
	PlaceCode(template.GetBabelConfig())
	PlaceCode(template.GetGitIgnore())
	PlaceCode(template.GetIndexFile())
	PlaceCode(template.GetErrorMiddleWare())
	PlaceCode(template.GetJWTMiddleWare())
	PlaceCode(template.GetUserControllerFile())
	PlaceCode(template.GetUserHelper())
	PlaceCode(template.GetUserModelFile())
	PlaceCode(template.GetUserRouteFile())
	PlaceCode(template.GetDockerCompose())
	PlaceCode(template.GetDockerFile())

}

func ReadProjectModel() *model.Project {

	projectConf, err := ioutil.ReadFile("project.json")
	if err != nil {
		panic("error reading the project configuration")
	}

	//parse into project struct
	p := &model.Project{}
	err = json.Unmarshal(projectConf, p)
	if err != nil {
		panic("error parsing the project configuration")
	}

	return p

}

func PlaceCode(fileName, filePath, code string) {

	var err error
	newpath := filepath.Join("./"+projectName+"/", filePath)
	os.MkdirAll(newpath, os.ModePerm)

	f, err := os.Create(newpath + "/" + fileName)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(code)
	if err != nil {
		panic(err)
	}

}
