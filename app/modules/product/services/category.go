package services

import (
	"fmt"
	"gomall/app/libs"
	"gomall/app/modules/product/models"
)

func GetAllCategory() *[]models.ProductCategory {
	var categorys []models.ProductCategory

	// where := map[string]interface{}{
	// 	"parent_id > ": "0",
	// }
	// libs.SqlOptimizedPagedQuery("pdt_product_category", where, 1, 10).Find(&categorys)
	sql := libs.NewSql("pdt_product_category").WithPage(2, 10).WithOrder("ORDER BY id DESC")
	sql.OptimizedPagedQuery().Find(&categorys)
	total := sql.Count()
	fmt.Printf("count: %d \n", total)
	return &categorys
}
