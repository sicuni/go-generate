package go_generate

import (
	"go-generate/config"
	"go-generate/database"
	"go-generate/internal/struct"
)

type Generator struct {
	db       database.Database
	executor *Executor
	Err      error
}

func New(config *config.Config) *Generator {
	return &Generator{
		executor: &Executor{
			Data:   make(map[string]*_struct.BaseStruct),
			Config: config,
		},
	}
}

func (gen *Generator) UseDB(db database.Database, err error) {
	if err != nil {
		gen.Err = err
		return
	}
	gen.db = db
}

func (gen *Generator) BindModel(base *_struct.BaseStruct) {
	if gen.Err != nil {
		return
	}

	gen.executor.Data[base.StructName] = base
}

func (gen *Generator) GenModel(tableName string) *_struct.BaseStruct {
	if gen.Err != nil {
		return nil
	}
	return gen.GenModelAs(tableName, "")
}

func (gen *Generator) GenModelAs(tableName, modelName string) *_struct.BaseStruct {
	if gen.Err != nil {
		return nil
	}
	baseStruct, err := _struct.GenBaseStructs(gen.db, tableName, modelName)
	if err != nil {
		gen.Err = err
		return nil
	}

	return baseStruct
}

func (gen *Generator) Execute() error {
	if gen.Err != nil {
		return gen.Err
	}

	return gen.executor.Execute()
}
