package model

import (
	"github.com/jinzhu/gorm"
)

// Query represents a query that is being stored in the database.
type Query struct {
	gorm.Model        // Embed GORM's standard model fields (ID, CreatedAt, UpdatedAt, DeletedAt)
	ChatID     int64  `gorm:"index:chatid"` // Chat ID associated with the query, indexed for faster queries
	LastAds    []Ad   // Slice of Ad records associated with this query (one-to-many relationship)
	Term       string `gorm:"type:varchar(100)"` // Search term for the query, stored as a string
	Radius     int    // Radius associated with the query
	City       int    // City code associated with the query
	CityName   string `gorm:"type:varchar(100)"` // Name of the city, stored as a string
	MaxPrice   *int   // Maximum price filter for the query
	MinPrice   *int   // Minimum price filter for the query
}

// AfterDelete is a GORM hook method that is executed after a Query record is deleted.
// It deletes all associated Ad records belonging to the deleted query.
func (q *Query) AfterDelete(tx *gorm.DB) (err error) {
	tx.Where("query_id = ?", q.ID).Delete(&Ad{})
	return
}
