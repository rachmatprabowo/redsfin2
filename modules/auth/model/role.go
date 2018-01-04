package model

import "github.com/rachmatprabowo/redsfin2/core"

// Role struct of role
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	core.BaseModel
}
