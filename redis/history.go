package redis

func AddSearchHistory(userId, searchText string) {
	DB.LPush(userId, searchText)
}

func GetSearchHistory(userId string) (v []string, ok bool) {
	v, err := DB.LRange(userId, 0, -1).Result()
	return v, err == nil
}
