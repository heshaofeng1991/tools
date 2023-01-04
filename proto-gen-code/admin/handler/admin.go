package handler

import (
	"encoding/json"

	"proto/admin"
)

type AdminServer struct {
	admin.UnimplementedAdminServer
}

// StructCopyTo 把一个结构体赋值给另一个结构体 to 传指针.
func StructCopyTo(source interface{}, to interface{}) error {
	buff, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(buff, to)
}
