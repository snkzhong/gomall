package models

type ProductAttribute struct {
    
	    Id    int       `gorm:"column:id" json:"id"`
    
	    ProductAttributeCategoryId    int       `gorm:"column:product_attribute_category_id" json:"product_attribute_category_id"`
    
	    Name    string       `gorm:"column:name" json:"name"`
    
	    SelectType    int       `gorm:"column:select_type" json:"select_type"`
    
	    InputType    int       `gorm:"column:input_type" json:"input_type"`
    
	    InputList    string       `gorm:"column:input_list" json:"input_list"`
    
	    Sort    int       `gorm:"column:sort" json:"sort"`
    
	    FilterType    int       `gorm:"column:filter_type" json:"filter_type"`
    
	    SearchType    int       `gorm:"column:search_type" json:"search_type"`
    
	    RelatedStatus    int       `gorm:"column:related_status" json:"related_status"`
    
	    HandAddStatus    int       `gorm:"column:hand_add_status" json:"hand_add_status"`
    
	    Type    int       `gorm:"column:type" json:"type"`
    
}