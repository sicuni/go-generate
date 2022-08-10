package go_generate

import (
	"bytes"
	"fmt"
	"github.com/sicuni/go-generate/config"
	"github.com/sicuni/go-generate/internal/struct"
	"github.com/sicuni/go-generate/internal/tmpl"
	"golang.org/x/tools/imports"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Executor struct {
	*config.Config
	Data map[string]*_struct.BaseStruct
}

// Execute 执行自动生成
func (e *Executor) Execute() error {
	var err error
	for _, baseStruct := range e.Data {
		if baseStruct == nil {
			continue
		}
		// 生成strcut
		if err = generateFile(e.ModelPath, e.ModelPkg, tmpl.StructTemplate, baseStruct.TableName, baseStruct); err != nil {
			return err
		}
		// 生成dao
		if err = generateFile(e.DaoPath, e.DaoPkg, tmpl.DaoTemplate, fmt.Sprint(baseStruct.TableName, "_", "dao"), baseStruct); err != nil {
			return err
		}
	}
	return err
}

// generateFile 根据模板生成代码
func generateFile(path, pkgName, tmp, fileName string, baseStruct *_struct.BaseStruct) error {
	outPath, err := getOutPath(path, pkgName)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	baseStruct.PackageName = pkgName
	if err := render(tmp, &buf, baseStruct); err != nil {
		return err
	}

	modelFile := fmt.Sprint(outPath, fileName, ".go")
	fmt.Printf("generate %s file(table <%s> -> {%s.%s}): %s \n", pkgName, baseStruct.TableName, baseStruct.PackageName, baseStruct.StructName, modelFile)

	return output(modelFile, buf.Bytes())
}

func output(fileName string, content []byte) error {
	result, err := imports.Process(fileName, content, nil)
	if err != nil {
		fmt.Println(string(result))
		return err
	}
	return outputFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, result)
}

func outputFile(filename string, flag int, data []byte) error {
	out, err := os.OpenFile(filename, flag, 0640)
	if err != nil {
		return fmt.Errorf("open out file fail: %w", err)
	}
	return write(out, data)
}

func write(wr io.WriteCloser, data []byte) (err error) {
	defer func() {
		if e := wr.Close(); e != nil {
			err = fmt.Errorf("close file fail: %w", e)
		}
	}()

	if _, err = wr.Write(data); err != nil {
		return fmt.Errorf("write file fail: %w", err)
	}
	return nil
}

// render 将data信息写入
func render(tmpl string, wr io.Writer, data interface{}) error {
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

// getOutPath 得到输出路径
func getOutPath(path, pkg string) (string, error) {
	outPath, err := genAbsPath(path)
	if err != nil {
		return outPath, err
	}
	outPath = fmt.Sprint(outPath, "/", filepath.Base(pkg), "/")
	if err := mkdirAll(os.ModePerm, outPath); err != nil {
		return outPath, err
	}
	return outPath, nil
}

// genAbsPath 生成绝对路径 路径不存在创建路径文件夹
func genAbsPath(path string) (string, error) {
	outPath, err := filepath.Abs(path)
	if err != nil {
		return outPath, err
	}
	if err := mkdirAll(os.ModePerm, outPath); err != nil {
		return outPath, err
	}

	return outPath, nil
}

// mkdirAll 创建绝对路径文件夹（支持传递多路径）
func mkdirAll(perm fs.FileMode, path ...string) error {
	for _, p := range path {
		if _, err := os.Stat(p); err != nil {
			if err := os.MkdirAll(p, perm); err != nil {
				return err
			}
		}
	}
	return nil
}
