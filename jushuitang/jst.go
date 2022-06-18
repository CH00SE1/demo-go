package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

// 订单请求对象
type JstorderReq struct {
	PageIndex     int    `json:"page_index"`
	ModifiedBegin string `json:"modified_begin"`
	ModifiedEnd   string `json:"modified_end"`
	PageSize      int    `json:"page_size"`
	Status        string `json:"status"`
}

// 订单返回对象
type JstorderRes struct {
	PageSize  int  `json:"page_size"`
	PageIndex int  `json:"page_index"`
	DataCount int  `json:"data_count"`
	PageCount int  `json:"page_count"`
	HasNext   bool `json:"has_next"`
	Datas     []struct {
		CoId             int         `json:"co_id"`
		ShopId           int         `json:"shop_id"`
		ShopName         string      `json:"shop_name"`
		IoId             int         `json:"io_id"`
		OId              int         `json:"o_id"`
		SoId             string      `json:"so_id"`
		Created          string      `json:"created"`
		Modified         string      `json:"modified"`
		Status           string      `json:"status"`
		OrderType        string      `json:"order_type"`
		InvoiceTitle     *string     `json:"invoice_title"`
		ShopBuyerId      *string     `json:"shop_buyer_id"`
		BuyerId          int         `json:"buyer_id"`
		OpenId           string      `json:"open_id"`
		ReceiverCountry  *string     `json:"receiver_country"`
		ReceiverState    *string     `json:"receiver_state"`
		ReceiverCity     *string     `json:"receiver_city"`
		ReceiverDistrict *string     `json:"receiver_district"`
		ReceiverTown     *string     `json:"receiver_town"`
		ReceiverAddress  *string     `json:"receiver_address"`
		ReceiverName     *string     `json:"receiver_name"`
		ReceiverPhone    *string     `json:"receiver_phone"`
		ReceiverMobile   *string     `json:"receiver_mobile"`
		BuyerMessage     interface{} `json:"buyer_message"`
		Remark           interface{} `json:"remark"`
		IsCod            bool        `json:"is_cod"`
		PayAmount        float64     `json:"pay_amount"`
		LId              string      `json:"l_id"`
		LogisticsCompany string      `json:"logistics_company"`
		IoDate           string      `json:"io_date"`
		LcId             string      `json:"lc_id"`
		StockEnabled     string      `json:"stock_enabled"`
		DrpCoIdFrom      string      `json:"drp_co_id_from"`
		Labels           *string     `json:"labels"`
		PaidAmount       float64     `json:"paid_amount"`
		FreeAmount       float64     `json:"free_amount"`
		Freight          float64     `json:"freight"`
		FFreight         interface{} `json:"f_freight"`
		Weight           float64     `json:"weight"`
		FWeight          interface{} `json:"f_weight"`
		MergeSoId        *string     `json:"merge_so_id"`
		WmsCoId          int         `json:"wms_co_id"`
		BusinessStaff    string      `json:"business_staff"`
		Currency         interface{} `json:"currency"`
		Node             interface{} `json:"node"`
		PayDate          string      `json:"pay_date"`
		SellerFlag       interface{} `json:"seller_flag"`
		WaveId           interface{} `json:"wave_id"`
		OrderStaffId     int         `json:"order_staff_id"`
		OrderStaffName   string      `json:"order_staff_name"`
		IsPrintExpress   bool        `json:"is_print_express"`
		IsPrint          bool        `json:"is_print"`
		Items            []struct {
			OId             int
			IoiId           int         `json:"ioi_id"`
			Pic             *string     `json:"pic"`
			SkuId           string      `json:"sku_id"`
			Qty             int         `json:"qty"`
			Name            string      `json:"name"`
			PropertiesValue string      `json:"properties_value"`
			SalePrice       float64     `json:"sale_price"`
			SaleAmount      float64     `json:"sale_amount"`
			IId             string      `json:"i_id"`
			FreeAmount      interface{} `json:"free_amount"`
			SaleBasePrice   float64     `json:"sale_base_price"`
			OrderBasePrice  interface{} `json:"order_base_price"`
			IsGift          bool        `json:"is_gift"`
			CombineSkuId    *string     `json:"combine_sku_id"`
			RawSoId         *string     `json:"raw_so_id"`
		} `json:"items"`
	} `json:"datas"`
	Code      int    `json:"code"`
	Issuccess bool   `json:"issuccess"`
	Msg       string `json:"msg"`
}

// 合作请求参数
const (
	url         = "https://open.erp321.com/api/open/query.aspx"
	partner_id  = "aba2cb02b01af36a4958a45c937e6945"
	partner_key = "a0364d3e5f87264ac29195f6407a0cfd"
	token       = "e7b4241d69dfe335901f6ab42a301dae"
)

// 接口对象(三个对接模板)
const (
	goodInfoInterface  = "jushuitan.itemsku.upload"
	goodsSqtyInterface = "jushuitan.inventory.upload"
	orderInterface     = "orders.out.simple.query"
)

// string 转 []byte
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// md5 加密string 返回 string
func md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

// 请求url拼接
func JstRequest(InterfaceName string) string {
	unix := time.Now().Local().Unix()
	timestamp := strconv.FormatInt(unix, 10)
	key := InterfaceName + partner_id + "token" + token + "ts" + timestamp + partner_key
	sign := md5V3(key)
	URL := url + "?method=" + InterfaceName + "&partnerid=" + partner_id + "&token=" + token + "&ts=" + timestamp + "&sign=" + sign
	return URL
}

// 获取订单返回
func orderRequest(c *gin.Context) {
	client := &http.Client{}
	URL := JstRequest(orderInterface)
	// 请求参数
	orderReq1 := JstorderReq{
		PageIndex:     1,
		ModifiedBegin: "2022-06-17 00:00:00",
		ModifiedEnd:   "2022-06-17 23:59:59",
		PageSize:      51,
		Status:        "Confirmed",
	}
	req1, _ := json.Marshal(&orderReq1)
	reader := bytes.NewReader(req1)
	request, _ := http.NewRequest("POST", URL, reader)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, _ := client.Do(request)
	fmt.Println(*response)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var orderRes1 JstorderRes
	json.Unmarshal(body, &orderRes1)
	msg := saveMysql(orderRes1)
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": len(orderRes1.Datas),
	})
}

// mysql数据库连接访问
func db() (*gorm.DB, error) {
	dsn := "root:11098319@tcp(192.168.10.87:3306)/djwk_test?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印执行日志
		Logger: logger.Default.LogMode(logger.Info),
	})
}

// 保存数据库
func saveMysql(orderRes1 JstorderRes) string {
	mysql, _ := db()
	for _, orderdoc := range orderRes1.Datas {
		mysql.Table("t_order_jst").Create(&orderdoc)
		for _, orderdetail := range orderdoc.Items {
			orderdetail.OId = orderdoc.OId
			mysql.Table("t_order_detail_jst").Create(&orderdetail)
		}
	}
	return "保存成功"
}

// 启动类
func main() {
	r := gin.Default()
	// 销售订单获取
	r.GET("/getOrder", orderRequest)
	r.Run(":8015")
}
