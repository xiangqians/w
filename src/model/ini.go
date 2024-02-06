// @author xiangqian
// @date 22:00 2023/11/02
package model

import (
	"fmt"
	"os"

	// https://pkg.go.dev/gopkg.in/ini.v1
	// https://github.com/go-ini/ini
	pkg_ini "gopkg.in/ini.v1"

	"time"
)

// 配置
type ini struct {
	Sys    sys    // 系统配置
	Data   data   // 数据配置
	Server server // 服务配置
}

// 系统配置
type sys struct {
	Passwd string // 密码
}

// 数据配置
type data struct {
	Dir string // 物理数据文件目录
}

// 服务配置
type server struct {
	Port          uint16        // 监听端口
	ContextPath   string        // 应用的上下文路径，也可以称为项目路径，是构成url地址的一部分
	SessionMaxAge time.Duration // 会话过期时间
}

var Ini ini

func init() {
	var source = os.Getenv("W_INI")
	if source == "" {
		source = "./w.ini"
	}
	file, err := pkg_ini.Load(source)
	if err != nil {
		panic(err)
	}

	// sys
	sys, err := file.GetSection("sys")
	if err != nil {
		panic(err)
	}
	Ini.Sys.Passwd = sys.Key("passwd").String()

	// data
	data, err := file.GetSection("data")
	if err != nil {
		panic(err)
	}
	Ini.Data.Dir = data.Key("dir").String()

	// server
	server, err := file.GetSection("server")
	if err != nil {
		panic(err)
	}
	Ini.Server.Port = uint16(server.Key("port").MustUint64(8080))
	Ini.Server.ContextPath = server.Key("context-path").MustString("")
	Ini.Server.SessionMaxAge = server.Key("session-max-age").MustDuration(12 * time.Hour)
}

// String 返回结构体类型字符串
func (ini ini) String() string {
	sysString := fmt.Sprintf("Sys\t{ Passwd = %s }", ini.Sys.Passwd)
	dataString := fmt.Sprintf("Data\t{ Dir = %s }", ini.Data.Dir)
	serverString := fmt.Sprintf("Server\t{ Port = %d, ContextPath = %s, SessionMaxAge = %s }", ini.Server.Port, ini.Server.ContextPath, ini.Server.SessionMaxAge)
	return fmt.Sprintf("Ini"+
		"\n\t%s"+
		"\n\t%s"+
		"\n\t%s",
		sysString,
		dataString,
		serverString)
}
