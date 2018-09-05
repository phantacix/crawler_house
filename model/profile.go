package model


type Profile struct {
	ProjectBuilding 	string  // 项目楼栋情况
	Block	string				// 座号
	BlockType	string			// 户型
	Contract	string			// 合同号
	FilingPrice	string			// 备案价格	 元/平方米(按建筑面积计)
	Floor	string				// 楼层
	Room	string				// 房号
	HouseType	string			// 用途
	PreSalesArea string			// 建筑面积 (预售查丈)
    PreUnitCstArea string		// 户内面积 (预售查丈)
	PreSharedPublicArea string	// 分摊面积 (预售查丈)
	CSalesArea string			// 户内面积 (竣工查丈)
	CUnitCstArea string			// 分摊面积 (竣工查丈)
	CSharedPublicArea string	// 建筑面积 (竣工查丈)
}