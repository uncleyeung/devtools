package template

// FindOne defines find row by id.
var FindOne = `
func (m *default{{.upperStartCamelObject}}Dao) FindOne({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	appzaplog.Debug("{{.upperStartCamelObject}}Dao FindOne", zap.Any("{{.lowerStartCamelPrimaryKey}}", {{.lowerStartCamelPrimaryKey}}))
	// todo 如需查询缓存自己增加
	resp := new({{.upperStartCamelObject}})
	err := m.mizaDB.Where("{{.lowerStartCamelPrimaryKey}} = ?", {{.lowerStartCamelPrimaryKey}}).First(resp).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	appzaplog.Debug("{{.upperStartCamelObject}}Dao FindOne", zap.Any("resp", resp))
	return resp, err
}
`

// FindOneByField defines find row by field.
var FindOneByField = `
func (m *default{{.upperStartCamelObject}}Model) FindOneBy{{.upperField}}({{.in}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowIndex(&resp, {{.cacheKeyVariable}}, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where {{.originalField}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		if err := conn.QueryRow(&resp, query, {{.lowerStartCamelField}}); err != nil {
			return nil, err
		}
		return resp.{{.upperStartCamelPrimaryKey}}, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}{{else}}var resp {{.upperStartCamelObject}}
	query := fmt.Sprintf("select %s from %s where {{.originalField}} limit 1", {{.lowerStartCamelObject}}Rows, m.table )
	err := m.conn.QueryRow(&resp, query, {{.lowerStartCamelField}})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}{{end}}
`

// FindOneByFieldExtraMethod defines find row by field with extras.
var FindOneByFieldExtraMethod = `
func (m *default{{.upperStartCamelObject}}Model) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", {{.primaryKeyLeft}}, primary)
}

func (m *default{{.upperStartCamelObject}}Model) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where {{.originalPrimaryField}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table )
	return conn.QueryRow(v, query, primary)
}
`

// FindOneMethod defines find row method.
var FindOneMethod = `FindOne({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)`

// FindOneByFieldMethod defines find row by field method.
var FindOneByFieldMethod = `FindOneBy{{.upperField}}({{.in}}) (*{{.upperStartCamelObject}}, error) `
