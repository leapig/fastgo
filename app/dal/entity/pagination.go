package entity

type Pagination struct {
	Page    int32  `json:"page"`
	Size    int32  `json:"size"`
	Cursor  int32  `json:"cursor"`
	Total   int32  `json:"total"`
	Keyword string `json:"keyword"`
}
