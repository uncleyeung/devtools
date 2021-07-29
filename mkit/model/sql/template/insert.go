package template

// Insert defines a template for insert code in model
var Insert = `
func (m *default{{.upperStartCamelObject}}Dao) Insert(data *{{.upperStartCamelObject}}) (sql.Result, error) {
	appzaplog.Debug("{{.upperStartCamelObject}}Dao Insert start", zap.Any("data", data))
	create := m.mizaDB.Create(data)
	err := create.Error
	if err != nil {
		appzaplog.Error("{{.upperStartCamelObject}}Dao Insert error", zap.Error(err))
		return driver.RowsAffected(0), err
	}
	// todo 如需增加缓存自己增加
	appzaplog.Debug("{{.upperStartCamelObject}}Dao Insert", zap.Any("Rows", create.RowsAffected), zap.Any("data", data))
	return driver.RowsAffected(create.RowsAffected), nil
}
`

// InsertMethod defines a interface method template for insert code in model
var InsertMethod = `Insert(data *{{.upperStartCamelObject}}) (sql.Result, error)`
