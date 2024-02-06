// @author xiangqian
// @date 11:08 2023/02/04
package os

import (
	"fmt"
	"runtime"
	"strings"
)

// IsWindows 是否是windows系统
var IsWindows = false

// IsLinux 是否是linux系统
var IsLinux = false

func init() {
	switch runtime.GOOS {
	// windows
	case "windows":
		IsWindows = true

	// linux
	case "linux":
		fallthrough // 执行穿透
	case "android":
		IsLinux = true

	// 不支持当前操作系统
	default:
		panic(fmt.Sprintf("不支持当前操作系统：%s", runtime.GOOS))
	}
}

// Path 拼接路径
func Path(more ...string) string {
	if IsWindows {
		return strings.Join(more, "\\")
	}

	if IsLinux {
		return strings.Join(more, "/")
	}

	panic(fmt.Sprintf("不支持当前操作系统：%s", runtime.GOOS))
}

// HumanizFileSize 人性化文件大小
// size: 文件大小，单位：byte
func HumanizFileSize(size int64) string {

	// B, Byte
	// 1B  = 8b
	// 1KB = 1024B
	// 1MB = 1024KB
	// 1GB = 1024MB
	// 1TB = 1024GB

	if size <= 0 {
		return "0 B"
	}

	// GB
	gb := float64(size) / (1024 * 1024 * 1024)
	if gb > 1 {
		return fmt.Sprintf("%.2f GB", gb)
	}

	// MB
	mb := float64(size) / (1024 * 1024)
	if mb > 1 {
		return fmt.Sprintf("%.2f MB", mb)
	}

	// KB
	kb := float64(size) / 1024
	if kb > 1 {
		return fmt.Sprintf("%.2f KB", kb)
	}

	// B
	return fmt.Sprintf("%d B", size)
}
