package gen

import (
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/uncleyeung/devtools/mkit/model/sql/template"
)

func genNew(table Table, withCache, postgreSql bool) (string, error) {
	text, err := util.LoadTemplate(category, modelNewTemplateFile, template.New)
	if err != nil {
		return "", err
	}

	t := fmt.Sprintf(`"%s"`, wrapWithRawString(table.Name.Source(), postgreSql))
	if postgreSql {
		t = "`" + fmt.Sprintf(`"%s"."%s"`, table.Db.Source(), table.Name.Source()) + "`"
	}

	output, err := util.With("new").
		Parse(text).
		Execute(map[string]interface{}{
			"table":                 strings.ReplaceAll(t, "`", ""),
			"withCache":             withCache,
			"upperStartCamelObject": table.Name.ToCamel(),
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
