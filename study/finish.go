package main

import (
	"demoGo/model/topen"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"net/http"
)

// 入参
type YiBaoInfo struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Data    struct {
		PageNum      int         `json:"pageNum"`
		PageSize     int         `json:"pageSize"`
		Size         int         `json:"size"`
		StartRow     int         `json:"startRow"`
		EndRow       int         `json:"endRow"`
		Pages        int         `json:"pages"`
		RecordCounts int         `json:"recordCounts"`
		PrePage      int         `json:"prePage"`
		NextPage     int         `json:"nextPage"`
		OrderField   interface{} `json:"orderField"`
		OrderType    interface{} `json:"orderType"`
		Data         []struct {
			SetlId            string      `json:"setlId"`
			MdtrtId           string      `json:"mdtrtId"`
			PsnName           string      `json:"psnName"`
			PsnCertType       string      `json:"psnCertType"`
			PsnCertTypeName   string      `json:"psnCertTypeName"`
			Certno            string      `json:"certno"`
			PsnNo             string      `json:"psnNo"`
			Gend              string      `json:"gend"`
			GendName          string      `json:"gendName"`
			Naty              string      `json:"naty"`
			NatyName          string      `json:"natyName"`
			Insutype          string      `json:"insutype"`
			InsutypeName      string      `json:"insutypeName"`
			PsnType           string      `json:"psnType"`
			PsnTypeName       string      `json:"psnTypeName"`
			Begndate          string      `json:"begndate"`
			Enddate           string      `json:"enddate"`
			SetlTime          string      `json:"setlTime"`
			MedfeeSumamt      float64     `json:"medfeeSumamt"`
			HiAgreSumfee      float64     `json:"hiAgreSumfee"`
			FundPaySumamt     float64     `json:"fundPaySumamt"`
			PsnPay            float64     `json:"psnPay"`
			AcctPay           float64     `json:"acctPay"`
			CashPayamt        float64     `json:"cashPayamt"`
			FulamtOwnpayAmt   float64     `json:"fulamtOwnpayAmt"`
			PreselfpayAmt     float64     `json:"preselfpayAmt"`
			InscpAmt          float64     `json:"inscpAmt"`
			DedcStd           float64     `json:"dedcStd"`
			CrtDedc           float64     `json:"crtDedc"`
			ActPayDedc        float64     `json:"actPayDedc"`
			HifpPay           float64     `json:"hifpPay"`
			PoolPropSelfpay   float64     `json:"poolPropSelfpay"`
			CvlservPay        float64     `json:"cvlservPay"`
			HifesPay          float64     `json:"hifesPay"`
			HifmiPay          float64     `json:"hifmiPay"`
			HifobPay          float64     `json:"hifobPay"`
			HifdmPay          float64     `json:"hifdmPay"`
			MafPay            float64     `json:"mafPay"`
			OthfundPay        float64     `json:"othfundPay"`
			CvlservFlag       string      `json:"cvlservFlag"`
			CvlservFlagName   string      `json:"cvlservFlagName"`
			CvlservLv         interface{} `json:"cvlservLv"`
			CvlservLvName     interface{} `json:"cvlservLvName"`
			SpPsnType         *string     `json:"spPsnType"`
			SpPsnTypeName     *string     `json:"spPsnTypeName"`
			SpPsnTypeLv       *string     `json:"spPsnTypeLv"`
			SpPsnTypeLvName   *string     `json:"spPsnTypeLvName"`
			ClctGrde          interface{} `json:"clctGrde"`
			ClctGrdeName      interface{} `json:"clctGrdeName"`
			FlxempeFlag       string      `json:"flxempeFlag"`
			FlxempeFlagName   string      `json:"flxempeFlagName"`
			NwbFlag           interface{} `json:"nwbFlag"`
			NwbFlagName       interface{} `json:"nwbFlagName"`
			InsuAdmdvs        string      `json:"insuAdmdvs"`
			InsuAdmdvsName    string      `json:"insuAdmdvsName"`
			EmpNo             string      `json:"empNo"`
			EmpName           string      `json:"empName"`
			EmpType           string      `json:"empType"`
			EmpTypeName       string      `json:"empTypeName"`
			EmpMgtType        interface{} `json:"empMgtType"`
			EmpMgtTypeName    interface{} `json:"empMgtTypeName"`
			PayLoc            string      `json:"payLoc"`
			PayLocName        string      `json:"payLocName"`
			FixmedinsCode     string      `json:"fixmedinsCode"`
			FixmedinsName     string      `json:"fixmedinsName"`
			HospLv            string      `json:"hospLv"`
			HospLvName        string      `json:"hospLvName"`
			MdtrtCertType     string      `json:"mdtrtCertType"`
			MdtrtCertTypeName string      `json:"mdtrtCertTypeName"`
			MedType           string      `json:"medType"`
			MedTypeName       string      `json:"medTypeName"`
			SetlType          string      `json:"setlType"`
			SetlTypeName      string      `json:"setlTypeName"`
			ClrType           string      `json:"clrType"`
			ClrTypeName       string      `json:"clrTypeName"`
			ClrWay            string      `json:"clrWay"`
			ClrWayName        string      `json:"clrWayName"`
			ClrOptins         string      `json:"clrOptins"`
			ClrOptinsName     string      `json:"clrOptinsName"`
			RefdSetlFlag      string      `json:"refdSetlFlag"`
			RefdSetlFlagName  string      `json:"refdSetlFlagName"`
			MdtrtCertNo       string      `json:"mdtrtCertNo"`
			DiseNo            string      `json:"diseNo"`
			DiseName          string      `json:"diseName"`
		} `json:"data"`
		FirstPage bool `json:"firstPage"`
		LastPage  bool `json:"lastPage"`
	} `json:"data"`
}

func swap(a *int, b *int) {
	*a, *b = *b, *a+10
}

// 指针基本数据
func zf_num() {
	a, b := 1515, 100
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}

func zf_array() {
	a := []int{10, 100, 200}
	MAX := 3
	var par [3]*int
	for i := 0; i < MAX; i++ {
		par[i] = &a[i]
	}
	fmt.Println(*par[1])
}

func gin_request() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func gin2() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func gin3() {
	router := gin.Default()
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}

func gin4() {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func gin5() {
	r := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

// mysql数据库连接访问
func db() (*gorm.DB, error) {
	dsn := "root:11098319@tcp(192.168.10.87:3306)/djwk_test?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印执行日志
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func yibaoinfosave(c *gin.Context) {
	r := gin.Default()
	var yiBaoInfo *YiBaoInfo
	c.ShouldBindJSON(&yiBaoInfo)
	opens := yiBaoInfo.Data.Data
	r.POST("/yibaoinfojson", func(c *gin.Context) {
		if yiBaoInfo.Type == "success" {
			db, _ := db()
			for _, open := range opens {
				db.Table("t_open").Create(&open)
			}
			c.JSON(http.StatusOK, gin.H{
				"msg":        "操作成功",
				"array_size": len(opens),
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"msg":  "入参不对",
				"info": "入参状态 - " + yiBaoInfo.Message,
			})
		}
	})
	r.Run(":8512")
}

func getSales(c *gin.Context) {
	url := "http://192.168.191.253:8500/djwk/sales/shop/" + c.Param("time") + "/32"
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"url": url,
			"msg": "url请求不对",
		})
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var jsonObeject topen.Sale
	json.Unmarshal(data, &jsonObeject)
	var doc topen.TSales
	var detail topen.SalesInfoDetails
	doc.ShopId = jsonObeject.ShopId
	doc.Type = jsonObeject.Type
	doc.Title = jsonObeject.Title
	mysql, _ := db()
	mysql.Debug().Create(&doc)
	for _, info := range jsonObeject.SalesInfoDetails {
		detail.SumFlMoney = info.SumFlMoney
		detail.SumMoney = info.SumMoney
		detail.Name = info.Name
		detail.TSalesId = doc.TSalesId
		mysql.Debug().Create(detail)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "数据保存成功",
		"data": jsonObeject,
	})
}

func auditMac(c *gin.Context) {
	mac := c.Query("mac")
	url := "http://192.168.191.253:8500/djwk/inck/auditNpMac?mac=" + mac
	macInfo, _ := http.Get(url)
	defer macInfo.Body.Close()
	res, _ := ioutil.ReadAll(macInfo.Body)
	var macResult topen.MacResult
	json.Unmarshal(res, &macResult)
	c.JSON(http.StatusOK, macResult)
}

func main() {
	r := gin.Default()
	r.GET("/:time/wuyi", getSales)
	r.GET("/inckMacAudit", auditMac)
	r.Run(":8500")
}
