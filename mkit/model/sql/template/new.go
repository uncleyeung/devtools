package template

// New defines an template for creating model instance
var New = `
func (obj {{.upperStartCamelObject}}) TableName() string {
	return {{.table}}
}

func new{{.upperStartCamelObject}}Dao() {{.upperStartCamelObject}}Dao {
	return &default{{.upperStartCamelObject}}Dao{
		{{if .withCache}}mizaDB:      mainDB{{else}}mizaDB:      mainDB{{end}},
		redisClient: RedisClient,
		table:       {{.table}},
	}
}
`
