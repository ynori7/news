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
		"THÜRINGEN":              {},
		"FRANKFURT":              {},
		"NÜRNBERG":               {},
		"DRESDEN":                {},
		"SAARLAND":               {},

		//Other
		"TV":     {},
		"STARS":  {},
		"QUEER":  {},
		"MUSIK":  {},
		"TENNIS": {},
		"KINO":   {},
	}

	categoryBlackList_FuzzyMatches = []string{
		"BUNDESLIGA",
		"FUSSBALL",
		"SPORT",
		"BASKETBALL",
		"SPIELE",
	}
)

func FilterCoronaNewsItems(news []model.NewsTickerItem) []model.NewsTickerItem {
	final := make([]model.NewsTickerItem, 0)

	for _, n := range news {
		//filter junk like videos and clips
		if n.Description == "" {
			continue
		}

		if strings.HasPrefix(n.Title, "Grafik") ||
			strings.HasPrefix(n.Title, "BILD-Grafik") ||
			strings.HasPrefix(n.Title, "VIDEO") {
			continue
		}

		final = append(final, n)
	}

	return final
}

func FilterNewsTickerItems(news []model.NewsTickerItem) []model.NewsTickerItem {
	final := make([]model.NewsTickerItem, 0)

	for _, n := range news {
		//filter junk like videos and clips
		if n.Description == "" {
			continue
		}

		//filter out junk like "bild plus"
		if !isNormalNewsArticle(n.Link) {
			continue
		}

		//filter categories
		if isInterestingCategory(n.Category) {
			final = append(final, n)
		}
	}

	return final
}

func isNormalNewsArticle(link string) bool {
	if strings.Contains(link, "/bild-plus") {
		return false
	}
	return true
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
