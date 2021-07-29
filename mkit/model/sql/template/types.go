package template

// Types defines a template for types in model
var Types = `
type (
	{{.upperStartCamelObject}}Dao interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Dao struct {
		{{if .withCache}}mizaDB      *gorm.DB{{else}}mizaDB      *gorm.DB{{end}}
		redisClient *uredis.Client
		table       string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
`
