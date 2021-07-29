package template

// Update defines a template for generating update codes
var Update = `
func (m *default{{.upperStartCamelObject}}Dao) Update(data {{.upperStartCamelObject}}) error {
	appzaplog.Debug("{{.upperStartCamelObject}}Dao Update start", zap.Any("data", data))
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    err := m.mizaDB.Exec(query, {{.expressionValues}}).Error
	// todo 如需清理缓存自己增加
	return err
}
`

// UpdateMethod defines an interface method template for generating update codes
var UpdateMethod = `Update(data {{.upperStartCamelObject}}) error`
