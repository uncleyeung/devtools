package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/uncleyeung/devtools/mkit/model/sql/template"
)

func genTag(in string) (string, error) {
	if in == "" {
		return in, nil
	}

	text, err := util.LoadTemplate(category, tagTemplateFile, template.Tag)
	if err != nil {
		return "", err
	}

	output, err := util.With("tag").Parse(text).Execute(map[string]interface{}{
		"field": in,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
