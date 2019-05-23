package category_words

import "testing"

func TestGetCategoryWords(t *testing.T) {
	_, err := GetCategoryWords("www.gametea.com", []string{"麻将游戏", "游戏大厅", "棋牌游戏", "斗牛", "双扣", "拼十", "斗地主"}, "", "")
	if err != nil {
		panic(err)
	}
}
