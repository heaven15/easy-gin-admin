package initialize

import (
	"flag"
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// InitConfig
// 优先级: 命令行 > 环境变量 > 默认值
func InitConfig() {
	var config string
	// &global.EGVA_DEBUG 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.StringVar(&config, "mode", "dev", "启动模式")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	if config == "" {
		if configEnv := os.Getenv(global.EGVA_ENV); configEnv == "" { // 判断global.CONFIG_ENV 常量存储的环境变量是否为空
			switch gin.Mode() {
			case gin.DebugMode:
				config = "dev"
			case gin.TestMode:
				config = "test"
			case gin.ReleaseMode:
				config = "pro"
			}
			fmt.Printf("您正在使用gin模式的%s环境名称\n", gin.EnvGinMode)
		} else { // global.CONFIG_ENV 常量存储的环境变量不为空 将值赋值于config
			config = configEnv
			fmt.Printf("您正在使用%s环境变量\n", config)
		}
	} else {
		fmt.Printf("您正在使用命令行的-mode参数传值的值%s\n", config)
	}

	dir, file := utils.GetRoot()
	fmt.Println("1111", fmt.Sprintf("%s%sapplication_%s.yaml", dir, file, config))
	v := viper.New()
	v.SetConfigFile(fmt.Sprintf("%s%sapplication_%s.yaml", dir, file, config))
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file：%s\n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(f fsnotify.Event) {
		fmt.Println("config file change:", f.Name)
		//这个对象如何在其他地方如何使用--全局变量
		if err := v.Unmarshal(&global.EGVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.EGVA_CONFIG); err != nil {
		fmt.Println(err)
	}
}
