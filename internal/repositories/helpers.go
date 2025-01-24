package repositories

import "eskept/internal/utils"

func GenerateSearchValue(keyword string) string {
	tokens := utils.Tokenize(keyword)
	searchValue := ""
	for _, token := range tokens {
		searchValue += "%" + token
	}
	searchValue += "%"
	return searchValue
}
