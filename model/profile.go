package model


type Profile struct {
	ID	uint	`gorm:"AUTO_INCREMENT"` // 自增
	ProjectBuilding 	string  // 项目楼栋情况
	Block	string				// 座号
	BlockType	string			// 户型
	Contract	string			// 合同号
	FilingPrice	float32	`gorm:"type:decimal(10,2);"`			// 备案价格	 元/平方米(按建筑面积计)
	Floor	string				// 楼层
	Room	int64				// 房号
	HouseType	string			// 用途
	PreSalesArea float32	`gorm:"type:decimal(10,2);"` 			// 建筑面积 (预售查丈)
    PreUnitCstArea float32	`gorm:"type:decimal(10,2);"`		// 户内面积 (预售查丈)
	PreSharedPublicArea float32	`gorm:"type:decimal(10,2);"`	// 分摊面积 (预售查丈)
	CSalesArea float32	`gorm:"type:decimal(10,2);"`			// 户内面积 (竣工查丈)
	CUnitCstArea float32	`gorm:"type:decimal(10,2);"`			// 分摊面积 (竣工查丈)
	CSharedPublicArea float32	`gorm:"type:decimal(10,2);"`	// 建筑面积 (竣工查丈)
}
