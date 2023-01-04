package types

type (
	ById struct {
		Id uint `form:"id"  binding:"required" json:"id"`
	}

	ByIdStr struct {
		Id string `form:"id"  binding:"required" json:"id"`
	}

	Pages struct {
		Page     int `form:"page"      json:"page"`
		PageSize int `form:"page_size" json:"page_size"`
	}
)
