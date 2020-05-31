package filter

import (
	"strings"

	"github.com/ynori7/news/bild/model"
)

var (
	categoryBlacklist_ExactMatches = map[string]struct{}{
		//Places
		"RUHRGEBIET":             {},
		"DÜSSELDORF":             {},
		"STUTTGART":              {},
		"HANNOVER":               {},
		"KÖLN":                   {},
		"HAMBURG":                {},
		"SACHSEN-ANHALT":         {},
		"MECKLENBURG-VORPOMMERN": {},
		"BERLIN":                 {},
		"LEIPZIG":                {},

		//Other
		"TV":    {},
		"STARS": {},
		"QUEER": {},
		"MUSIK": {},
	}

	categoryBlackList_FuzzyMatches = []string{
		"BUNDESLIGA",
		"FUSSBALL",
		"SPORT",
		"BASKETBALL",
		"SPIELE",
	}
)

func FilterNewsTickerItems(news []model.NewsTickerItem) []model.NewsTickerItem {
	final := make([]model.NewsTickerItem, 0)

	for _, n := range news {
		//filter junk like videos and clips
		if n.Description == "" {
			continue
		}

		//filter categories
		if isInterestingCategory(n.Category) {
			final = append(final, n)
		}
	}

	return final
}

func isInterestingCategory(category string) bool {
	category = strings.ToUpper(category)

	//filter exact match categories
	if _, ok := categoryBlacklist_ExactMatches[category]; ok {
		return false
	}

	//filter fuzzy match categories
	for _, cat := range categoryBlackList_FuzzyMatches {
		if strings.Contains(category, cat) {
			return false
		}
	}

	return true
}