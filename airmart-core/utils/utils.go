package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net"
	"reflect"

	"github.com/spf13/viper"
)

// GetEnvInfo 获取环境变量
func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func MD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}

func ReflectRestore(a interface{}, b interface{}) (interface{}, error) {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := reflect.TypeOf(b)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		iv := v.Field(i)
		if iv.Type() == t {
			if iv.Kind() == reflect.Ptr {
				iv = iv.Elem()
			}
			return iv.Interface(), nil
		}
	}
	return nil, errors.New("类型不存在")
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func GetIpAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("获取Ip失败", err)
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
