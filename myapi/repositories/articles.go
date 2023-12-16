package repositories

import (
	"database/sql"
	"fmt"

	"github.com/tanakalucky/golang-test/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles
		(title, contents, username, nice, created_at)
		values (?, ?, ?, 0, now());
	`

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	var newArticle models.Article

	newArticle.ID = int(id)
	newArticle.Title = article.Title
	newArticle.Contents = article.Contents
	newArticle.UserName = article.UserName

	return newArticle, nil
}
