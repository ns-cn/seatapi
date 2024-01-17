package seatapi

import "time"

// RowRecord 行记录
type RowRecord struct {
	Id    string    `json:"_id"`
	Ctime time.Time `json:"_ctime"`
	Mtime time.Time `json:"_mtime"`
}
