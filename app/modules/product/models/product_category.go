package models

type ProductCategory struct {
    
	    Id    int       `gorm:"column:id" json:"id"`
    
	    ParentId    int       `gorm:"column:parent_id" json:"parent_id"`
    
	    Name    string       `gorm:"column:name" json:"name"`
    
	    Level    int       `gorm:"column:level" json:"level"`
    
	    ProductCount    int       `gorm:"column:product_count" json:"product_count"`
    
	    ProductUnit    string       `gorm:"column:product_unit" json:"product_unit"`
    
	    NavStatus    int       `gorm:"column:nav_status" json:"nav_status"`
    
	    ShowStatus    int       `gorm:"column:show_status" json:"show_status"`
    
	    Sort    int       `gorm:"column:sort" json:"sort"`
    
	    Icon    string       `gorm:"column:icon" json:"icon"`
    
	    Keywords    string       `gorm:"column:keywords" json:"keywords"`
    
	    Description    string       `gorm:"column:description" json:"description"`
    
}