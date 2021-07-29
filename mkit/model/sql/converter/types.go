package converter

import (
	"fmt"
	"strings"

	"github.com/zeromicro/ddl-parser/parser"
)

var commonMysqlDataTypeMap = map[int]string{
	// For consistency, all integer types are converted to int64
	// number
	parser.Bool:      "uint8",
	parser.Boolean:   "uint8",
	parser.TinyInt:   "int8",
	parser.SmallInt:  "int16",
	parser.MediumInt: "int32",
	parser.Int:       "int32",
	parser.MiddleInt: "int64",
	parser.Int1:      "int64",
	parser.Int2:      "int64",
	parser.Int3:      "int64",
	parser.Int4:      "int64",
	parser.Int8:      "int8",
	parser.Integer:   "int32",
	parser.BigInt:    "int64",
	parser.Float:     "float64",
	parser.Float4:    "float64",
	parser.Float8:    "float64",
	parser.Double:    "float64",
	parser.Decimal:   "float64",
	// date&time
	parser.Date:      "time.Time",
	parser.DateTime:  "time.Time",
	parser.Timestamp: "time.Time",
	parser.Time:      "string",
	parser.Year:      "int64",
	// string
	parser.Char:       "string",
	parser.VarChar:    "string",
	parser.Binary:     "string",
	parser.VarBinary:  "string",
	parser.TinyText:   "string",
	parser.Text:       "string",
	parser.MediumText: "string",
	parser.LongText:   "string",
	parser.Enum:       "string",
	parser.Set:        "string",
	parser.Json:       "string",
	parser.Blob:       "string",
	parser.LongBlob:   "string",
	parser.MediumBlob: "string",
	parser.TinyBlob:   "string",
}

var commonMysqlDataTypeMap2 = map[string]string{
	// For consistency, all integer types are converted to int64
	// number
	"bool":      "uint8",
	"boolean":   "uint8",
	"tinyint":   "int8",
	"smallint":  "int16",
	"mediumint": "int32",
	"int":       "int32",
	"integer":   "int32",
	"bigint":    "int64",
	"float":     "float64",
	"double":    "float64",
	"decimal":   "float64",
	// date&time
	"date":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",
	"time":      "string",
	"year":      "int64",
	// string
	"char":       "string",
	"varchar":    "string",
	"binary":     "string",
	"varbinary":  "string",
	"tinytext":   "string",
	"text":       "string",
	"mediumtext": "string",
	"longtext":   "string",
	"enum":       "string",
	"set":        "string",
	"json":       "string",
	"blob":       "string",
	"longblob":   "string",
	"mediumblob": "string",
	"tinyblob":   "string",
}

// ConvertDataType converts mysql column type into golang type
func ConvertDataType(dataBaseType int, isDefaultNull bool) (string, error) {
	tp, ok := commonMysqlDataTypeMap[dataBaseType]
	if !ok {
		return "", fmt.Errorf("unsupported database type: %v", dataBaseType)
	}

	return mayConvertNullType(tp, isDefaultNull), nil
}

// ConvertStringDataType converts mysql column type into golang type
func ConvertStringDataType(dataBaseType string, isDefaultNull bool) (string, error) {
	tp, ok := commonMysqlDataTypeMap2[strings.ToLower(dataBaseType)]
	if !ok {
		return "", fmt.Errorf("unsupported database type: %s", dataBaseType)
	}

	return mayConvertNullType(tp, isDefaultNull), nil
}

func mayConvertNullType(goDataType string, isDefaultNull bool) string {
	if !isDefaultNull {
		return goDataType
	}

	switch goDataType {
	case "int64":
		return "sql.NullInt64"
	case "int32":
		return "sql.NullInt32"
	case "float64":
		return "sql.NullFloat64"
	case "bool":
		return "sql.NullBool"
	case "string":
		return "sql.NullString"
	case "time.Time":
		return "sql.NullTime"
	default:
		return goDataType
	}
}
