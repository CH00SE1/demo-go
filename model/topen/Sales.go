package topen

// 销售返回类型
type Sale struct {
	Title            string `json:"title"`
	ShopId           int    `json:"shop_id"`
	Type             string `json:"type"`
	SalesInfoDetails []struct {
		Name       string  `json:"name"`
		SumMoney   float64 `json:"sum_money"`
		SumFlMoney float64 `json:"sum_fl_money"`
		Fl         struct {
			A string `json:"A"`
			B string `json:"B"`
			C string `json:"C"`
			D string `json:"D"`
			E string `json:"E"`
		} `json:"fl"`
	} `json:"salesInfoDetails"`
	SalesName      []string `json:"sales_name"`
	Classification []string `json:"classification"`
}

type MacResult struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		Approvedate     int64  `json:"approvedate"`
		Credate         int64  `json:"credate"`
		Licensetype     string `json:"licensetype"`
		Mac             string `json:"mac"`
		Memo            string `json:"memo"`
		Reqdate         int64  `json:"reqdate"`
		Reqemployeeid   int    `json:"reqemployeeid"`
		Reqemployeename string `json:"reqemployeename"`
		Reqip           string `json:"reqip"`
		Seqid           int    `json:"seqid"`
	} `json:"data"`
}

// t_sales
type TSales struct {
	TSalesId int64  `gorm:"column:t_sales_id;type:bigint(100);primary_key;AUTO_INCREMENT;comment:t_sales_id" json:"t_sales_id"`
	Title    string `gorm:"column:title;type:varchar(100);comment:今天销售数据" json:"title"`
	ShopId   int    `gorm:"column:shop_id;type:bigint(11);comment:32" json:"shop_id"`
	Type     string `gorm:"column:type;type:varchar(100);comment:day" json:"type"`
}

// sales_info_details
type SalesInfoDetails struct {
	Name       string  `gorm:"column:name;type:varchar(100);comment:许晓晴" json:"name"`
	TSalesId   int64   `gorm:"column:t_sales_id;type:bigint(100);comment:t_sales_id" json:"t_sales_id"`
	SumMoney   float64 `gorm:"column:sum_money;type:double;comment:1529.8" json:"sum_money"`
	SumFlMoney float64 `gorm:"column:sum_fl_money;type:double;comment:62.4449" json:"sum_fl_money"`
}

func (TSales) TableName() string {
	return "t_sales"
}

func (SalesInfoDetails) TableName() string {
	return "sales_info_details"
}
