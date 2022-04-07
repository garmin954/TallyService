package service

type LedgerRecord struct{}

// 获取记录
func (ledgerRecord *LedgerRecord) Pagination() (pg Pagination, err error) {
	pg.List = make([]interface{}, 5)
	pg.Total = 100
	pg.PageNum = 1
	pg.PageSize = 10

	return pg, nil
}
