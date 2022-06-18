package main

import "github.com/gin-gonic/gin"

// 定义端口号
var port = ":8511"

func main() {

	// 接受变脸
	r := gin.Default()
	// 列表
	r.GET("/list", list)
	// @RequestParam 数据保存
	r.GET("/save", save_request_param)
	// @PathVariable 数据保存
	r.GET("/get/:id", select_id_path_variable)
	// @RequestBody
	r.POST("/json", json_request_body)

	r.POST("/open_json", open_json)

	r.Run(port)
}
