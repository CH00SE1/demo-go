package main

import (
	"database/sql"
	"demoGo/model/topen"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 结构体
var op topen.TOpen

/**
模型
*/
type (
	// 数据表字段
	t_opens struct {
		gorm.Model
		ID           int
		Name         string
		Email        string
		Age          int
		MemberNumber int
		Birthday     *time.Time
		ActivatedAt  sql.NullTime
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	// 展示数据
	info struct {
		ID   int
		Name string
		Age  int
	}
)

// 数据库连接访问
func db() (*gorm.DB, error) {
	dsn := "root:11098319@tcp(192.168.10.87:3306)/djwk_test?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印执行日志
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func getTime() {
	now := time.Now()                  //获取当前时间
	timestamp := now.Unix()            //时间戳
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

/**
结果集展示
*/
func list(c *gin.Context) {
	db, err := db()
	if err != nil {
		fmt.Println("连接错误")
	}
	// 定义存储查询结果的变量
	var result []*info
	// 指定查询model(对应的数据结构体名称)，FIND(存储查询结果的变量地址)
	db.Model(&t_opens{}).Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   result,
	})
}

/**
接受数据 @RequestParam
*/
func save_request_param(c *gin.Context) {
	db, err := db()
	if err != nil {
		fmt.Println("success")
	}
	name_web := c.Query("name")
	email_web := c.Query("email")
	age_web, err := strconv.Atoi(c.Query("age"))
	if err != nil {
		age_web = 0
	}
	// 转为结构体
	opens := t_opens{Name: name_web, Age: age_web, Email: email_web}
	// 保存数据库
	db.Create(&opens)
	c.JSON(http.StatusOK, gin.H{
		"data": opens,
	})
}

// @PathVariable 查询
func select_id_path_variable(c *gin.Context) {
	id, err2 := strconv.Atoi(c.Param("id"))
	if err2 != nil {
		id = 1
	}
	log.Printf("id = %d \n", id)
	db, err := db()
	if err != nil {
		fmt.Println("success")
	}
	var result []*info
	db.Model(&t_opens{}).Find(&result, id)
	c.JSON(http.StatusOK, gin.H{
		"id": result,
	})
}

//@RequestBody
func json_request_body(c *gin.Context) {

	var opens []t_opens
	err := c.ShouldBindJSON(&opens)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := db()
	if err != nil {
		fmt.Println("success")
	}
	for i, open := range opens {
		fmt.Printf("i --> %d insert time:", i)
		db.Create(&open)
		getTime()
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "操作成功",
	})
}

/**
判断string是否在数组
*/
func IsContainInArray(item string, items []*topen.Seltid) bool {
	for _, eachItem := range items {
		if eachItem.SetlID == item {
			return true
		}
	}
	return false
}

/**
保存医保数据
存在如何获取json对象某个标签
*/
func open_json(c *gin.Context) {
	var opens []topen.TOpen
	err := c.ShouldBindJSON(&opens)
	db, err := db()
	if err != nil {
		return
	}
	var result []*topen.Seltid
	db.Debug().Select("SetlID").Find(&result)
	num := 0
	for _, open := range opens {
		if IsContainInArray(open.SetlID, result) == false {
			db.Debug().Create(&open)
			num++
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "操作成功",
		"insert": num,
		"size":   len(opens),
		"exist":  len(opens) - num,
	},
	)
}
