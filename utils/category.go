package utils

import "github.com/mairuu/nekopost-api/types"

var Categories = []types.Category{
	{CateName: "Fantasy", CateLink: "fantasy", CateCode: "1"},
	{CateName: "Action", CateLink: "action", CateCode: "2"},
	{CateName: "Drama", CateLink: "drama", CateCode: "3"},
	{CateName: "Sport", CateLink: "sport", CateCode: "5"},
	{CateName: "Sci-fi", CateLink: "sci-fi", CateCode: "7"},
	{CateName: "Comedy", CateLink: "comedy", CateCode: "8"},
	{CateName: "Slice of Life", CateLink: "slice_of_life", CateCode: "9"},
	{CateName: "Romance", CateLink: "romance", CateCode: "10"},
	{CateName: "Adventure", CateLink: "adventure", CateCode: "13"},
	{CateName: "Yaoi", CateLink: "yaoi", CateCode: "23"},
	{CateName: "Seinen", CateLink: "seinen", CateCode: "49"},
	{CateName: "Trap", CateLink: "trap", CateCode: "25"},
	{CateName: "Gender Bender", CateLink: "gender_blender", CateCode: "26"},
	{CateName: "Second Life", CateLink: "second_life", CateCode: "45"},
	{CateName: "Isekai", CateLink: "isekai", CateCode: "44"},
	{CateName: "School Life", CateLink: "school_life", CateCode: "43"},
	{CateName: "Mystery", CateLink: "mystery", CateCode: "32"},
	{CateName: "Horror", CateLink: "horror", CateCode: "47"},
	{CateName: "Shounen", CateLink: "shounen", CateCode: "46"},
	{CateName: "Shoujo", CateLink: "shoujo", CateCode: "42"},
	{CateName: "Yuri", CateLink: "yuri", CateCode: "24"},
	{CateName: "Gourmet", CateLink: "gourmet", CateCode: "41"},
	{CateName: "Harem", CateLink: "harem", CateCode: "50"},
	{CateName: "Reincanate", CateLink: "reincanate", CateCode: "51"},
}

var CategoryByName = make(map[string]types.Category)
var CategoryByLink = make(map[string]types.Category)

func init() {
	for _, category := range Categories {
		CategoryByName[category.CateName] = category
		CategoryByLink[category.CateLink] = category
	}
}

func GetCategoryByName(name string) (types.Category, bool) {
	category, exists := CategoryByName[name]
	return category, exists
}

func GetCategoryByCode(code string) (types.Category, bool) {
	for _, category := range Categories {
		if category.CateCode == code {
			return category, true
		}
	}
	return types.Category{}, false
}

func GetCategoryByLink(link string) (types.Category, bool) {
	category, exists := CategoryByLink[link]
	return category, exists
}

