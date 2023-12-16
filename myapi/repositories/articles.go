package repositories

import (
	"database/sql"
	"fmt"

	"github.com/tanakalucky/golang-test/models"
)

const (
	articleNumPerPage = 5
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

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select
			article_id,
			title,
			contents,
			username,
			nice
		from
			articles
		limit ?
		offset ?
	`

	offset := articleNumPerPage * (page - 1)

	rows, err := db.Query(sqlStr, articleNumPerPage, offset)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	articles := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article

		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName)
		articles = append(articles, article)
	}

	return articles, nil
}
