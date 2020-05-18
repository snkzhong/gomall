package libs

import (
	"fmt"
	"gomall/app/global"
	"strings"

	"github.com/jinzhu/gorm"
)

type sqlOption struct {
	Table string
	Where string
	Values []interface{}
	Order string
	Page int
	Pagesize int
}

func NewSql(table string) *sqlOption {
	return &sqlOption{
		Table: table,
		Page: 1,
		Pagesize: 10,
	}
}

func (s *sqlOption) WithPage(page, pagesize int) *sqlOption {
	s.Page = page
	s.Pagesize = pagesize
	return s
}

func (s *sqlOption) WithOrder(order string) *sqlOption {
	s.Order = order
	return s
}

func (s *sqlOption) OptimizedPagedQuery() *gorm.DB {
	offset := (s.Page - 1) * s.Pagesize
	sql := fmt.Sprintf(`SELECT t.* FROM (
		SELECT id FROM %s WHERE 1 LIMIT %d, %d
		) a, %s t WHERE a.id=t.id
	`, s.Table, offset, s.Pagesize, s.Table)
	return global.GlobalDb.Raw(sql)
}

func (s *sqlOption) Count() int {
	total := 0
	sql := fmt.Sprintf("SELECT count(*) AS total FROM %s WHERE 1", s.Table)
	global.GlobalDb.Raw(sql).Count(&total)
	return total
}

func SqlOptimizedPagedQuery(table string, condition map[string]interface{}, page int, pagesize int) *gorm.DB {
	whereStr, values := SqlBuildWhere(condition)
	offset := (page - 1) * pagesize

	sql := fmt.Sprintf(`SELECT t.* FROM (
		SELECT id FROM %s WHERE %s LIMIT %d, %d
		) a, %s t WHERE a.id=t.id
	`, table, whereStr, offset, pagesize, table)
	return global.GlobalDb.Raw(sql, values...)
}

func SqlBuildWhere(condition map[string]interface{}) (whereStr string, values []interface{}) {
	where := []string{}
	for k, v := range condition {
		if strings.Contains(k, " ") {
			where = append(where, k+" ?")
		} else {
			where = append(where, k+" = ?")
		}

		values = append(values, v)
	}

	whereStr = strings.Join(where, " AND ")
	return
}