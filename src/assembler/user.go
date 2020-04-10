package assembler

import (
	bookmarkdao "github.com/robin019/bookmark/persistance/gorm/bookmarkDao"
	userdao "github.com/robin019/bookmark/persistance/gorm/userDao"
)

func User(users []*userdao.Model) (result []map[string]interface{}) {
	if len(users) == 0 {
		return []map[string]interface{}{}
	}

	for _, user := range users {
		m := map[string]interface{}{
			"userId":      user.ID,
			"userAccount": user.UserAccount,
			"gender":      user.Gender,
		}

		result = append(result, m)
	}
	return result
}

func Bookmark(bookmarks []*bookmarkdao.Model)(result []map[string]interface{}){

	if len(bookmarks) == 0 {
		return []map[string]interface{}{}
	}

	for _, bookmark := range bookmarks {
		m := map[string]interface{}{
			"user_id":  bookmark.UserID,
			"name": 	bookmark.Name,
			"url":      bookmark.URL,
		}

		result = append(result, m)
	}
	return result
}