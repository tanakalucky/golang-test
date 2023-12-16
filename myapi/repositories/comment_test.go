package repositories_test

import (
	"testing"

	"github.com/tanakalucky/golang-test/models"
	"github.com/tanakalucky/golang-test/repositories"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	commentList, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	expectedNum := 2
	if expectedNum != len(commentList) {
		t.Errorf("want %d but get %d comments\n", expectedNum, len(commentList))
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "testtest",
	}

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	expectedNum := 3
	if newComment.CommentID != expectedNum {
		t.Errorf("new comment id is expected %d but get %d\n", expectedNum, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from
				comments
			where
				article_id = ? and
				message = ?;
		`

		testDB.Exec(sqlStr, newComment.ArticleID, newComment.Message)
	})
}
