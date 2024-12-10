package types

import (
	"time"
)

type DBImage struct {
	ID          uint `gorm:"primaryKey;unique;autoIncrement;index"`
	ImageID     string
	ImageName   string
	ImageSize   uint
	ContentType string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
