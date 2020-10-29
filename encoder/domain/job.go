package domain

import "time"

type Job struct {
	ID              string
	OutputBuckePath string
	Status          string
	Video           *Video
	Error           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
