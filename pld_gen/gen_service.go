package pld_gen

import (
	"bytes"
	"fmt"
	"github.com/michaelzx/pld/pld_fs"
	"github.com/michaelzx/pld/pld_print"
	"github.com/pkg/errors"
	"go/format"
	"io/ioutil"
	"path/filepath"
)

func (g *Generator) genService() {
	entDir := filepath.Join(g.OutDir, g.PackageService)
	pld_fs.CreateIfNotExist(entDir)
	tmpl := g.loadTpl(g.TplService)
	for tableName, meta := range g.tableMetaMap {

		var doc bytes.Buffer
		err := tmpl.Execute(&doc, &entTplData{
			StructName:  meta.StructName,
			PackageName: g.PackageService,
		})
		if err != nil {
			panic(errors.Wrap(err, "模板执行失败"+tableName))
		}

		reFormatDoc, err := format.Source(doc.Bytes())
		if err != nil {
			panic(errors.Wrap(err, "格式化go文件失败"+tableName))
		}
		fileName := tableName + "_service.go"
		if g.FileNameProcessor != nil {
			fileName = g.FileNameProcessor(fileName)
		}
		outPath := filepath.Join(g.OutDir, g.PackageService, fileName)
		ioutil.WriteFile(filepath.Join(outPath), reFormatDoc, 0777)
		pld_print.Green(fmt.Sprintf("成功生成 : %s", filepath.Join(g.rootPath, outPath)))
	}
}
