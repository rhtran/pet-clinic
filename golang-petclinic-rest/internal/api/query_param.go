package api

type NameParam struct {
	Name string `form:"name" binding:"required"`
}

type LastNameParam struct {
	LastName string `form:"last-name" binding:"required"`
}
