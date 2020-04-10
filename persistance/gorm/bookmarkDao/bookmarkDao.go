package bookmarkdao

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	table = "bookmark"
)

type Model struct {
	ID          int       `gorm:"column:id"`
	UserID 		int    	  `gorm:"column:user_id"`
	Name      	string    `gorm:"column:name"`
	URL 		string	  `gorm:"column:url"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

//QueryModel lists all queryable columns
type QueryModel struct {
	ID          int
	UserID 		int
	Name      	string
	URL 		string
}


func New(tx *gorm.DB, model *Model) {
	err := tx.Table(table).
		Create(&model).Error
	if err != nil {
		panic(err)
	}
}

func Delete(tx *gorm.DB, query *QueryModel) {

	result := &Model{}
	err := tx.Table(table).Scopes(queryChain(query)).
		Delete(&result).Error
	if err != nil {
		panic(err)
	}
}


func Get(tx *gorm.DB, query *QueryModel) *Model {
	result := &Model{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(result).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}
	return result
}

func Gets(tx *gorm.DB, query *QueryModel) []*Model {
	var rows []*Model
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&rows).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}

	return rows
}


func userIDEqualScope(userId int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userId != 0 {
			return db.Where(table+".user_id = ?", userId)
		}
		return db
	}
}

func idEqualScope(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id != 0 {
			return db.Where(table+".id = ?", id)
		}
		return db
	}
}

func nameEqualScope(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {
			return db.Where(table+".name = ?", name)
		}
		return db
	}
}

func urlEqualScope(url string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if url != ""{
			return db.Where(table+".url = ?", url)
		}
		return db
	}
}


func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(userIDEqualScope(query.UserID)).
			Scopes(idEqualScope(query.ID)).
			Scopes(nameEqualScope(query.Name)).
			Scopes(urlEqualScope(query.URL))
	}
}