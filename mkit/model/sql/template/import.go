package template

var (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	appzaplog "{{if .gitlab}}{{.gitlab}}{{else}}your.gitlab.com{{end}}/gomod/utils/ulog"
	utildao "{{if .gitlab}}{{.gitlab}}{{else}}your.gitlab.com{{end}}/{{if .repo}}{{.repo}}{{else}}your-repository{{end}}/server/src/util/dao"
	"github.com/jinzhu/gorm"
	"{{if .gitlab}}{{.gitlab}}{{else}}your.gitlab.com{{end}}/{{if .repo}}{{.repo}}{{else}}your-repository{{end}}/server/src/util/uredis"
	"go.uber.org/zap"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	appzaplog "{{if .gitlab}}{{.gitlab}}{{else}}your.gitlab.com{{end}}/gomod/utils/ulog"
	utildao "{{if .gitlab}}{{.gitlab}}{{else}}your.gitlab.com{{end}}/{{if .repo}}{{.repo}}{{else}}your-repository{{end}}/server/src/util/dao"
	"github.com/jinzhu/gorm"
	"{{if .gitlab}}{{.gitlab}}{{else}}your.gitlab.com{{end}}/{{if .repo}}{{.repo}}{{else}}your-repository{{end}}/server/src/util/uredis"
	"go.uber.org/zap"
)
`
)
