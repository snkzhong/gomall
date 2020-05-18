package libs

import (
	"fmt"
	"gomall/app/global"
	"log"
	"strconv"
	"strings"
	"time"
)

type Logger struct{}

func (logger *Logger) Print(values ...interface{}) {

	if values[0] == "sql" {
		source := values[1]
		ms := values[2]
		sql := values[3].(string)
		vars := values[4].([]interface{})
		sql = buildSql(sql, vars)

		// t := time.Now()
		// now := fmt.Sprintf("%v-%v-%v %v:%v:%v:%v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
		now := time.Now().Format("2006-01-02 15:04:05")

		message := fmt.Sprintf("\n[%v] %v sql=\"%s\" file=\"%s\" \n", now, ms, sql, source)

		log.Println(message)
		if global.GlobalSqlLogFile != nil {
			global.GlobalSqlLogFile.WriteString(message)
		}
	}
}

func buildSql(sql string, vars []interface{}) string {
	for _, v := range vars {
		t := fmt.Sprintf("%T", v)
		if t == "string" {
			sql = strings.Replace(sql, "?", fmt.Sprintf("'%s'", v), 1)
		} else if t == "int" || t == "float64" {
			sql = strings.Replace(sql, "?", v.(string), 1)
		} else if t == "[]string" {
			sv := v.([]string)
			s := fmt.Sprintf("'%s'", strings.Join(sv, "','"))
			sql = strings.Replace(sql, "?", s, 1)
		} else if t == "[]int" {
			iv := v.([]int)
			sl := []string{}
			for _, i := range iv {
				sl = append(sl, strconv.Itoa(i))
			}
			is := fmt.Sprintf("'%s'", strings.Join(sl, "','"))
			sql = strings.Replace(sql, "?", is, 1)
		} else {
			sql = strings.Replace(sql, "?", fmt.Sprintf("%v", v), 1)
		}
	}

	return sql
}
