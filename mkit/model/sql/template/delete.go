package template

// Delete defines a delete template
var Delete = `
func (m *default{{.upperStartCamelObject}}Dao) Delete({{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	appzaplog.Debug("{{.upperStartCamelObject}}Dao Delete", zap.Any("{{.lowerStartCamelPrimaryKey}}", {{.lowerStartCamelPrimaryKey}}))
	query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
	err := m.mizaDB.Exec(query, {{.lowerStartCamelPrimaryKey}}).Error
	// todo 如需清理缓存自己增加
	return err
}
`

// DeleteMethod defines a delete template for interface method
var DeleteMethod = `Delete({{.lowerStartCamelPrimaryKey}} {{.dataType}}) error`
