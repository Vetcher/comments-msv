package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Text     string `json:"text"`
	AuthorID uint   `gorm:"index" json:"author"`
}

func (db *Database) PostComment(c *Comment) (*Comment, error) {
	if err := db.db.Create(c); err.Error != nil {
		return nil, DBError(err.Error)
	} else {
		return c, nil
	}
}

func (db *Database) DeleteComment(id uint) (bool, error) {
	var temp Comment
	if err := db.db.Where("id = ?", id).First(&temp); err.Error != nil {
		return false, DBError(err.Error)
	} else {
		if err := db.db.Delete(temp); err.Error != nil {
			return false, DBError(err.Error)
		}
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
