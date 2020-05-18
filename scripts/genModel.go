package scripts

import (
	"fmt"
	"gomall/app/global"
	"os"
	"strings"
	"text/template"
)

type TmplData struct {
	TableName string
	Fields    []FieldData
}

type FieldData struct {
	FieldName     string
	FieldType     string
	FieldOriginal string
}

func GenModel(module, table string) {
	tableArr := strings.SplitN(table, "_", 2)
	tableNameUpper := convertHumpformat(tableArr[1])
	rows, _ := global.GlobalDb.Raw(fmt.Sprintf("desc %s;", table)).Rows()
	defer rows.Close()

	var fields []FieldData
	for rows.Next() {
		var (
			Field   string
			Type    string
			Null    string
			Key     string
			Default string
			Extra   string
		)
		rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)
		// fmt.Println(Field, Type, Null, Key, Default, Extra)
		field := FieldData{convertHumpformat(Field), getGoType(Type), Field}
		fields = append(fields, field)
	}

	data := &TmplData{TableName: tableNameUpper, Fields: fields}
	tmpl, _ := template.ParseFiles("./scripts/genModel.tpl")
	f, _ := os.OpenFile(fmt.Sprintf("./app/modules/%s/models/%s.go", module, tableArr[1]), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	tmpl.Execute(f, data)
}

func firstCharUpper(str string) string {
	charSlice := []rune(str)
	return strings.ToUpper(string(charSlice[0:1])) + string(charSlice[1:])
}

func convertHumpformat(str string) string {
	strSlice := strings.Split(str, "_")
	for k, v := range strSlice {
		strSlice[k] = firstCharUpper(v)
	}

	return strings.Join(strSlice, "")
}

func getGoType(fieldType string) string {
	if isStringType(fieldType) {
		return "string"
	}
	if isIntType(fieldType) {
		return "int"
	}
	if isFloatType(fieldType) {
		return "float32"
	}

	return "string"
}

func isStringType(fieldType string) bool {
	return strings.Contains(fieldType, "varchar") || strings.Contains(fieldType, "char") || strings.Contains(fieldType, "tinytext") || strings.Contains(fieldType, "text") || strings.Contains(fieldType, "mediumtext") || strings.Contains(fieldType, "longtext") || strings.Contains(fieldType, "datetime")
}

func isIntType(fieldType string) bool {
	return strings.Contains(fieldType, "int") || strings.Contains(fieldType, "tinyint") || strings.Contains(fieldType, "bigint")
}

func isFloatType(fieldType string) bool {
	return strings.Contains(fieldType, "decimal")
}
