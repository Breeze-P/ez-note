package constants

import (
	"net"
	"os"
	"strings"
)

// 这里应用了「环境变量」
func GetIp(key string) string {
	ip := os.Getenv(key)
	if ip == "" {
		ip = "localhost"
	}
	return ip
}

func GetOutBoundIP() (ip string, err error) {
	// 创建一个 UDP 连接到 Google 的 DNS 服务器
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// 获取本地地址信息
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// 从地址字符串中提取 IP 地址部分
	ip = strings.Split(localAddr.String(), ":")[0]

	return ip, nil
}
