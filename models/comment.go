package models

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Text      string     `json:"text"`
	AuthorID  uint       `gorm:"index" json:"author_id"`
}

func (db *Database) PostComment(c *Comment) (*Comment, error) {
	if err := db.db.Create(c); err.Error != nil {
		return nil, DBError(err.Error)
	} else {
		return c, nil
	}
}

func (db *Database) DeleteComment(id uint) (bool, error) {
	if err := db.db.Delete(Comment{}, "id = ?", id); err.Error != nil {
		return false, DBError(err.Error)
	} else {
		return true, nil
	}
}

func (db *Database) SelectCommentByID(id uint) (*Comment, error) {
	var c Comment
	if err := db.db.Where("id = ?", id).First(&c); err.Error != nil {
		return nil, DBError(err.Error)
	}
	return &c, nil
}

func (db *Database) LoadCommentsForUser(userId uint) ([]*Comment, error) {
	var comments []*Comment
	if err := db.db.Where("author_id = ?", userId).Find(&comments); err.Error != nil {
		return nil, DBError(err.Error)
	}
	return comments, nil
}
