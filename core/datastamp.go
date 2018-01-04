package core

// BaseModel _
type BaseModel struct {
	Status      int    `json:"status"`
	CreatedDate string `json:"created_date"`
	CreatedBy   string `json:"created_by"`
	UpdatedDate string `json:"updated_date"`
	UpdatedBy   string `json:"updated_by"`
	DeletedDate string `json:"deleted_date"`
	DeletedBy   string `json:"deleted_by"`
}
