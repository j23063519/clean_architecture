package domain

import (
	"time"

	"github.com/jackc/pgtype"
)

// base id
type BSModel struct {
	ID uint64 `json:"id" gorm:"column:id;primaryKey;uniqueIndex;autoIncrement;comment:id;"`
}

// uuid
type UuidModel struct {
	ID string `json:"id" gorm:"column:id;type:uuid;primaryKey;uniqueIndex;comment:uuid;"`
}

// ip
type IPModel struct {
	IP pgtype.Inet `json:"ip" gorm:"type:inet;not null;index;comment:ip address" swaggertype:"string"`
}

// set ip
func (i *IPModel) Set(ip string) (inet pgtype.Inet) {
	inet.Set(ip)
	i.IP = inet
	return
}

type CommonTimestmp struct {
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP;comment:create time;"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;index;default:null;comment:update time;"`
}
