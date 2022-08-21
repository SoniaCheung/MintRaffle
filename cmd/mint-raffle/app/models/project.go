package models

import "time"

type Project struct {
	Id          int           `json:"id" xorm:"not null pk autoincr INT"`
	Name        string        `json:"name" xorm:"not null VARCHAR(255)"`
	Description string        `json:"description" xorm:"TEXT"`
	OfficalLink string        `json:"offical_link" xorm:"VARCHAR(255)"`
	MaxWinner   int           `json:"max_winner" xorm:"not null INT"`
	DueTime     time.Time     `json:"due_time" xorm:"not null DATETIME"`
	Status      ProjectStatus `json:"status" xorm:"default 'opening' not null VARCHAR(255)"`
	CreatedAt   time.Time     `json:"created_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
	UpdatedAt   time.Time     `json:"updated_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

type ProjectStatus string

const (
	// project is opening
	ProjectStatusOpening ProjectStatus = "opening"
	// project is pending for raffle
	ProjectStatusPending ProjectStatus = "pending"
	// project is closed
	ProjectStatusClosed ProjectStatus = "closed"
)

func (*Project) TableName() string {
	return "projects"
}
