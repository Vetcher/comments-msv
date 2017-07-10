package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Text     string `json:"text"`
	AuthorID uint   `gorm:"index" json:"author_id"`
}

func (db *Database) PostComment(c *Comment) (*Comment, error) {
	if err := db.db.Create(c); err.Error != nil {
		return nil, DBError(err.Error)
	} else {
		return c, nil
	}
}

func (db *Database) DeleteComment(id uint) (bool, error) {
	var c Comment
	err := db.db.Where("id = ?", id).First(&c)
	if err != nil {
		log.Println(err.Error)
	}
	log.Println(c)
	if err := db.db.Delete(&Comment{}, "id = ?", id); err.Error != nil {
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
	if err := db.db.Where("author_id = ?", userId); err.Error != nil {
		return nil, DBError(err.Error)
	}
	return comments, nil
}
