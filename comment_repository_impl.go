package repository

import (
	"belajar-golang-mysql-raul/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepositoryImpl(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlExec := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, sqlExec, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlQuery := "select id, email, comment from comments where id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, sqlQuery, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
	return comment, nil
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlQueryAll := "select id, email, comment from comments"
	rows, err := repo.DB.QueryContext(ctx, sqlQueryAll)
	comment := entity.Comment{}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment = entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
