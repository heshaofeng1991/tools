package api

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type coreMiddlewares struct {
	conf interface{}
}

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func New(conf interface{}) *coreMiddlewares {
	return &coreMiddlewares{
		conf: conf,
	}
}

func (middle *coreMiddlewares) getConfig(conf interface{}) (interface{}, error) {
	v := reflect.ValueOf(middle.conf)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := reflect.TypeOf(conf)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Type() == t {
			return v.Field(i).Interface(), nil
		}
	}
	return nil, fmt.Errorf("参数错误 %+v : %s=%s", middle.conf, v.String(), t.String())
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
