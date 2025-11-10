package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	// Faz o unwappring de um modelo do Gorm pra struct
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
