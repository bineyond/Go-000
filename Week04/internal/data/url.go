// solmyr

package data

import (
	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	Url    string
	Shorty string
}
