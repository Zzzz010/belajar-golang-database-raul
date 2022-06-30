package repository

import (
	golang_mysql_raul "belajar-golang-mysql-raul"
	"belajar-golang-mysql-raul/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(golang_mysql_raul.GetDatabase())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository2@example.com",
		Comment: "Ini Test Komen Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(golang_mysql_raul.GetDatabase())

	comment, err := commentRepository.FindById(context.Background(), 12)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)

}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(golang_mysql_raul.GetDatabase())

	comment, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comment {
		fmt.Println(comment)
	}
}
