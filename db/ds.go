package db

import "time"

// database operation error
type OPError struct {
    Code int
    Err error
}

type BasicModel struct {
    ID        uint `gorm:"primary_key" json:"id"`
    CreatedAt time.Time     `json:"-"`
    UpdatedAt time.Time     `json:"-"`
    DeletedAt *time.Time    `json:"-"`
}
