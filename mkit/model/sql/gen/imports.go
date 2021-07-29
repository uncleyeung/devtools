package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/uncleyeung/devtools/mkit/model/sql/template"
)

func genImports(table Table, withCache, timeImport bool) (string, error) {
	if withCache {
		text, err := util.LoadTemplate(category, importsTemplateFile, template.Imports)
		if err != nil {
			return "", err
		}

		buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
			"time":   timeImport,
			"gitlab": table.Table.Gitlab,
			"repo":   table.Table.Repo,
		})
		if err != nil {
			return "", err
		}

		return buffer.String(), nil
	}

	text, err := util.LoadTemplate(category, importsWithNoCacheTemplateFile, template.ImportsNoCache)
	if err != nil {
		return "", err
	}

	buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
		"time":   timeImport,
		"gitlab": table.Table.Gitlab,
		"repo":   table.Table.Repo,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
