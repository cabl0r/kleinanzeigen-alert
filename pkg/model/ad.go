package model

import "time"

// Ad represents an advertisement that is being stored in the database.
type Ad struct {
	ID        uint      `gorm:"primary_key"`       // Primary key of the Ad in the database
	EbayID    string    `gorm:"type:varchar(255)"` // Unique eBay ID for the ad, stored as a string
	QueryID   uint      `gorm:"index:ad_queryid"`  // Foreign key referencing the associated query
	Location  string    `gorm:"type:varchar(510)"` // Location information for the ad, stored as a string
	CreatedAt time.Time // Timestamp representing when the ad was created
}
