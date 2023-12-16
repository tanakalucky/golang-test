package services

import (
	"github.com/tanakalucky/golang-test/models"
	"github.com/tanakalucky/golang-test/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, nil
	}

	return newComment, nil
}
