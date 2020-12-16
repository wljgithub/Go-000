package repository

import (
	"context"
	"gorm.io/gorm"
	"week4/internal/model"
)

var _ Repository = &repo{}


type Repository interface {
	Ping(ctx context.Context) error
	Close()
	GetArticle(ctx context.Context, id int) (model.Article, error)
}

type repo struct {
	sql   *gorm.DB
}

var Repo Repository

func init() {
	mysql,_,err:=NewMysql()
	if err != nil{
		panic(err)
	}
	Repo,_,_= NewRepo(mysql)
}
func NewRepo(sql *gorm.DB) (r Repository, cf func(), err error) {
	r = &repo{sql: sql}
	cf = r.Close
	return
}

func (this *repo) GetArticle(ctx context.Context, id int) (article model.Article, err error) {
	err = this.sql.First(&article).Error
	return
}

func (this *repo) Ping(ctx context.Context) error {
	return nil
}

func (this *repo) Close() {

}
