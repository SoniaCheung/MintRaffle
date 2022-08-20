package models

import "time"

type Project struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	Name        string    `xorm:"not null VARCHAR(255)"`
	Description string    `xorm:"TEXT"`
	OfficalLink string    `xorm:"VARCHAR(255)"`
	MaxWinner   int       `xorm:"not null INT"`
	DueTime     time.Time `xorm:"not null DATETIME"`
	Status      string    `xorm:"default 'opening' not null VARCHAR(255)"`
	CreatedAt   time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	UpdatedAt   time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (*Project) TableName() string {
	return "projects"
}
