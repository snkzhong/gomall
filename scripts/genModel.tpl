package models

type {{ .TableName }} struct {
    {{ range .Fields }}
	    {{ .FieldName }}    {{ .FieldType }}       `gorm:"column:{{ .FieldOriginal }}" json:"{{ .FieldOriginal }}"`
    {{ end }}
}