package userdao

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	table = "user"
)

type Model struct {
	ID          int       `gorm:"column:id"`
	UserAccount string    `gorm:"column:user_account"`
	Gender      string    `gorm:"column:gender"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

//QueryModel lists all queryable columns
type QueryModel struct {
	ID          int
	UserAccount string
	Gender      string
}

func New(tx *gorm.DB, model *Model) {
	err := tx.Table(table).
		Create(&model).Error

	if err != nil {
		panic(err)
	}
}

func Count(tx *gorm.DB, query *QueryModel) int {
	var count int
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Count(&count).Error
	if err != nil {
		panic(err)
	}
	return count
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

func Modify(tx *gorm.DB, user *Model) {
	attrs := map[string]interface{}{
		"gender": user.Gender,
	}
	err := tx.Table(table).
		Model(Model{}).
		Where("id = ?", user.ID).
		Updates(attrs).Error

	// update("gender",gender)

	if err != nil {
		panic(err)
	}
}

func Delete(tx *gorm.DB, query *QueryModel){

	result := &Model{}
	err := tx.Table(table).Scopes(queryChain(query)).Delete(&result).Error

	if err != nil {
		panic(err)
	}

}


func userAccountEqualScope(userAccount string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userAccount != "" {
			return db.Where(table+".user_account = ?", userAccount)
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

func genderEqualScope(gender string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if gender != "" {
			return db.Where(table+".gender = ?", gender)
		}
		return db
	}
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(userAccountEqualScope(query.UserAccount)).
			Scopes(idEqualScope(query.ID)).
			Scopes(genderEqualScope(query.Gender))
	}
}
