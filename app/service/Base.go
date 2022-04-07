package service

type Base struct{}

// 分页结构
type Pagination struct {
	Total    int           `json:"total"`
	List     []interface{} `json:"list"`
	PageNum  int8          `json:"page_num"`
	PageSize int8          `json:"page_size"`
}
