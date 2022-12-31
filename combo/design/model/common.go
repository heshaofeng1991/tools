package model

import (
	. "goa.design/goa/v3/dsl"
)

var Pagination = Type("Pagination", func() {
	Field(1, "page", Int, "page 页", func() {
		Minimum(1)
		Example(1)
	})
	Field(2, "page_size", Int, "page_size 页大小", func() {
		Minimum(1)
		Example(1)
	})
	Extend(AuthToken)
	Required("page", "page_size")
})

var AuthToken = Type("AuthToken", func() {
	Field(1, "Authorization", String, "Authorization")
	TokenField(2, "token", String, func() {
		Description("JWT used for authentication")
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
})

var Base = Type("Base", func() {
	Field(1, "id", Int, "id 自增ID", func() {
		Minimum(1)
		Example(1)
	})
	Field(2, "create_time", String, "create_time 创建时间")
	Field(3, "update_time", String, "update_time 更新时间")
})

var Author = Type("Author", func() {
	Field(1, "create_by", String, "create_by 创建人")
	Field(2, "update_by", String, "update_by 更新人")
})

var BaseAndRevision = Type("BaseAndRevision", func() {
	Field(1, "revision", Int, "revision 乐观锁")
	Extend(Base)
})

var BaseAndRevisionAndAuthor = Type("BaseAndRevisionAndAuthor", func() {
	Extend(BaseAndRevision)
	Extend(Author)
})

var BaseResponse = Type("BaseResponse", func() {
	Field(1, "code", Int, "code", func() {
		Example(0)
	})
	Field(2, "message", String, "message", func() {
		Example("description error information")
	})
	Required("code", "message")
})
