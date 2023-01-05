/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    http
	@Date    2022/5/15 18:24
	@Desc
*/

package logic

import (
	applicationChannel "github.com/heshaofeng1991/ddd-sample/application/channel"
)

type HTTPServer struct {
	application applicationChannel.Application
}

func NewHTTPServer(application applicationChannel.Application) HTTPServer {
	return HTTPServer{
		application: application,
	}
}
