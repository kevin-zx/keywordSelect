package category_words

import (
	"fmt"
	"testing"
)

func TestGetCategoryWordsBySiteUrl(t *testing.T) {
	f, err := GetCategoryWordsBySiteUrl("http://chengzhitianhong.mvp.baixing.com", "", "")
	println(f)
	fmt.Println(err)
}
