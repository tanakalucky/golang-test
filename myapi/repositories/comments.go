package repositories

import (
	"database/sql"

	"github.com/tanakalucky/golang-test/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into
			comments
		(article_id, message, created_at)
		values (?, ?, now());
	`

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	articleID, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	var newComment models.Comment
	newComment.ArticleID = int(articleID)
	newComment.Message = comment.Message

	return newComment, nil
}
