package models

import "time"

type Submission struct {
	Id            int       `xorm:"not null pk autoincr INT"`
	ProjectId     int       `xorm:"INT"`
	WalletAddress string    `xorm:"VARCHAR(255)"`
	Winner        bool      `xorm:"TINYINT(1)"`
	CreatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	UpdatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (*Submission) TableName() string {
	return "submissions"
}
