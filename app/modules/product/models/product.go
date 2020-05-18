package models

type Product struct {
    
	    Id    int       `gorm:"column:id" json:"id"`
    
	    BrandId    int       `gorm:"column:brand_id" json:"brand_id"`
    
	    ProductCategoryId    int       `gorm:"column:product_category_id" json:"product_category_id"`
    
	    FeightTemplateId    int       `gorm:"column:feight_template_id" json:"feight_template_id"`
    
	    ProductAttributeCategoryId    int       `gorm:"column:product_attribute_category_id" json:"product_attribute_category_id"`
    
	    Name    string       `gorm:"column:name" json:"name"`
    
	    Pic    string       `gorm:"column:pic" json:"pic"`
    
	    ProductSn    string       `gorm:"column:product_sn" json:"product_sn"`
    
	    DeleteStatus    int       `gorm:"column:delete_status" json:"delete_status"`
    
	    PublishStatus    int       `gorm:"column:publish_status" json:"publish_status"`
    
	    NewStatus    int       `gorm:"column:new_status" json:"new_status"`
    
	    RecommandStatus    int       `gorm:"column:recommand_status" json:"recommand_status"`
    
	    VerifyStatus    int       `gorm:"column:verify_status" json:"verify_status"`
    
	    Sort    int       `gorm:"column:sort" json:"sort"`
    
	    Sale    int       `gorm:"column:sale" json:"sale"`
    
	    Price    float32       `gorm:"column:price" json:"price"`
    
	    PromotionPrice    float32       `gorm:"column:promotion_price" json:"promotion_price"`
    
	    GiftGrowth    int       `gorm:"column:gift_growth" json:"gift_growth"`
    
	    GiftPoint    int       `gorm:"column:gift_point" json:"gift_point"`
    
	    UsePointLimit    int       `gorm:"column:use_point_limit" json:"use_point_limit"`
    
	    SubTitle    string       `gorm:"column:sub_title" json:"sub_title"`
    
	    Description    string       `gorm:"column:description" json:"description"`
    
	    OriginalPrice    float32       `gorm:"column:original_price" json:"original_price"`
    
	    Stock    int       `gorm:"column:stock" json:"stock"`
    
	    LowStock    int       `gorm:"column:low_stock" json:"low_stock"`
    
	    Unit    string       `gorm:"column:unit" json:"unit"`
    
	    Weight    float32       `gorm:"column:weight" json:"weight"`
    
	    PreviewStatus    int       `gorm:"column:preview_status" json:"preview_status"`
    
	    ServiceIds    string       `gorm:"column:service_ids" json:"service_ids"`
    
	    Keywords    string       `gorm:"column:keywords" json:"keywords"`
    
	    Note    string       `gorm:"column:note" json:"note"`
    
	    AlbumPics    string       `gorm:"column:album_pics" json:"album_pics"`
    
	    DetailTitle    string       `gorm:"column:detail_title" json:"detail_title"`
    
	    DetailDesc    string       `gorm:"column:detail_desc" json:"detail_desc"`
    
	    DetailHtml    string       `gorm:"column:detail_html" json:"detail_html"`
    
	    DetailMobileHtml    string       `gorm:"column:detail_mobile_html" json:"detail_mobile_html"`
    
	    PromotionStartTime    string       `gorm:"column:promotion_start_time" json:"promotion_start_time"`
    
	    PromotionEndTime    string       `gorm:"column:promotion_end_time" json:"promotion_end_time"`
    
	    PromotionPerLimit    int       `gorm:"column:promotion_per_limit" json:"promotion_per_limit"`
    
	    PromotionType    int       `gorm:"column:promotion_type" json:"promotion_type"`
    
	    BrandName    string       `gorm:"column:brand_name" json:"brand_name"`
    
	    ProductCategoryName    string       `gorm:"column:product_category_name" json:"product_category_name"`
    
}