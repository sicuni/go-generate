package _struct

import (
	"fmt"
	"github.com/sicuni/go-generate/database"
	"github.com/sicuni/go-generate/internal/model"
	"regexp"
	"strings"
)

type BaseStruct struct {
	S               string
	PackageName     string
	StructName      string
	TableName       string
	StructFieldName string
	Fields          []*model.Field
}

func GenBaseStructs(db database.Database, tableName string, modelName string) (*BaseStruct, error) {
	if err := checkModelName(modelName); err != nil {
		return nil, fmt.Errorf("model name %q is invalid: %w", modelName, err)
	}

	fields, err := db.GetStructFields(tableName)
	if err != nil {
		return nil, err
	}

	s := strings.ToLower(modelName)[:1]
	base := BaseStruct{
		Fields:          fields,
		TableName:       tableName,
		StructName:      modelName,
		StructFieldName: s + modelName[1:],
		S:               s,
	}

	return &base, err
}

var modelNameReg = regexp.MustCompile(`^\w+$`)

func checkModelName(name string) error {
	if name == "" {
		return nil
	}
	if !modelNameReg.MatchString(name) {
		return fmt.Errorf("model name cannot contains invalid character")
	}
	if name[0] < 'A' || name[0] > 'Z' {
		return fmt.Errorf("model name must be initial capital")
	}
	return nil
}
