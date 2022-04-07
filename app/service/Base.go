package service

type Base struct{}

// 分页结构
type Pagination struct {
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
}
