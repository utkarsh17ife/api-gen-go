package model

import "strings"

type Property struct {
	Name         string `json:"name"`         // model name
	Type         string `json:"type"`         // string, number, [string], [number], object, [object], refId, [refId], date, bool
	Default      string `json:"default"`      // some value
	Required     bool   `json:"required"`     // boolean
	Ref          string `json:"ref"`          // reference model name
	ObjectString string `json:"objectString"` // complex object string
}

type Model struct {
	Name       string     `json:"modelname"`
	Api        bool       `json:"api"`
	Properties []Property `json:"properties"`
}

type Project struct {
	Name   string  `json:"projectName"`
	Models []Model `json:"models"`
}

func (m Model) FirstLetterCaps() string {
	return strings.Title(m.Name)
}
