package service

type Base struct{}

// 分页结构
type Pagination struct {
<<<<<<< HEAD
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
=======
	Total    int           `json:"total"`
	List     []interface{} `json:"list"`
	PageNum  int8          `json:"page_num"`
	PageSize int8          `json:"page_size"`
>>>>>>> f5f42371f1c38e0ec7b08a45c1148b681c15329c
}
