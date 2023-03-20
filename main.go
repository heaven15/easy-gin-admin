/*
 * @Author: null 1060236395@qq.com
 * @Date: 2023-03-20 16:32:53
 * @LastEditors: null 1060236395@qq.com
 * @LastEditTime: 2023-03-20 17:02:35
 * @FilePath: \easy-gin-vue-admin\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/develop-kevin/easy-gin-vue-admin/initialize"
)

func main() {
	//1.初始化config
	initialize.InitConfig()
	//2.初始化zap
	initialize.InitZap()
	//3.初始化routers
	Router := initialize.Routers()
	//4.初始化验证器翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(fmt.Sprintf("初始化验证器翻译：%v", err.Error()))
	}
	initialize.InitDB()
	//5.初始化数据库
	if global.EGVA_DB != nil {
		//TODO  读取文件安装文件是否存在，存在，初始化表不执行
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.EGVA_DB.DB()
		defer db.Close()
	}
	//6.初始化Redis
	initialize.InitRedis()
	port := global.EGVA_CONFIG.System.Port
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        Router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		//启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("服务启动失败：%v", err.Error()))
		}
	}()
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("|               easy-gin-vue-admin              |")
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("|    Go Gin Activity Server Start Successful    |")
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("|                  Port：" + strconv.Itoa(port) + "                   |")
	fmt.Println("|-----------------------------------------------|")
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.EGVA_LOG.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Service Shutdown Failed：%v", err.Error()))
	}
	global.EGVA_LOG.Info("Server exit...")
}
