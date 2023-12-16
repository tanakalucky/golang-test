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
	newComment.CommentID = int(articleID)
	newComment.ArticleID = int(articleID)
	newComment.Message = comment.Message

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select
			*
		from
			comments
		where
			article_id = ?
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentList = make([]models.Comment, 0)

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime

		err := rows.Scan(&comment.ArticleID, &comment.ArticleID, &comment.Message, &createdTime)
		if err != nil {
			return nil, err
		}

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentList = append(commentList, comment)
	}

	return commentList, nil
}
