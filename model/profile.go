package model


type Profile struct {
	ID int `gorm:"-;primary_key;AUTO_INCREMENT"` // 自增主键
	ProjectBuilding 	string  // 项目楼栋情况
	Block	string				// 座号
	BlockType	string			// 户型
	Contract	string			// 合同号
	FilingPrice	float32	`gorm:"type:decimal(10,2);not null;default '';index:idx_filing_price;"`			// 备案价格	 元/平方米(按建筑面积计)
	Floor	string				// 楼层
	Room	int64				// 房号
	HouseType	string			// 用途
	PreSalesArea float32	`gorm:"type:decimal(10,2);not null;default '';index:idx_pre_sales_area;"` 			// 建筑面积 (预售查丈)
    PreUnitCstArea float32	`gorm:"type:decimal(10,2);not null;default '';"`		// 户内面积 (预售查丈)
	PreSharedPublicArea float32	`gorm:"type:decimal(10,2);not null;default '';"`	// 分摊面积 (预售查丈)
	CSalesArea float32	`gorm:"type:decimal(10,2);not null;default '';"`			// 户内面积 (竣工查丈)
	CUnitCstArea float32	`gorm:"type:decimal(10,2);not null;default '';"`			// 分摊面积 (竣工查丈)
	CSharedPublicArea float32	`gorm:"type:decimal(10,2);not null;default '';index:idx_c_shared_public_area;"`	// 建筑面积 (竣工查丈)
}
