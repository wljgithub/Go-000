package repository

import "gorm.io/gorm"
import "week4/pkg/database/sql"

func NewMysql()(db *gorm.DB,cf func(),err error)  {
	db = sql.NewMysql()
	return
}
