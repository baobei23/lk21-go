package scraper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/baobei23/lk21-go/models"
)

// ScrapeSeries scrapes a list of series from a given URL
func ScrapeSeries(req *http.Request, res *http.Response) ([]models.Series, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var seriesList []models.Series

	doc.Find("main > div.container > section.archive div.grid-archive > div#grid-wrapper > div.infscroll-item").Each(func(i int, s *goquery.Selection) {
		parent := s.Find("article.mega-item")
		var genres []string

		parent.Find("footer div.grid-categories > a").Each(func(i int, s2 *goquery.Selection) {
			href, _ := s2.Attr("href")
			parts := strings.Split(href, "/")
			if len(parts) > 2 && parts[1] == "genre" {
				genres = append(genres, parts[2])
			}
		})

		href, _ := parent.Find("figure > a").Attr("href")
		seriesID := strings.Split(href, "/")

		posterImg, _ := parent.Find("figure > a > picture > img").Attr("src")
		title, _ := parent.Find("figure > a > picture > img").Attr("alt")
		episodeStr := parent.Find("figure > div.grid-meta > div.last-episode > span").Text()
		episode, _ := strconv.Atoi(episodeStr)

		series := models.Series{
			ID:        seriesID[len(seriesID)-2],
			Title:     title,
			Type:      "series",
			PosterImg: fmt.Sprintf("https:%s", posterImg),
			Rating:    parent.Find("figure div.rating").Text(),
			URL:       fmt.Sprintf("%s://%s/series/%s", req.URL.Scheme, req.Host, seriesID[len(seriesID)-2]),
			Episode:   episode,
			Genres:    genres,
		}

		seriesList = append(seriesList, series)
	})

	return seriesList, nil
}

// ScrapeSeriesDetails scrapes the details of a single series
func ScrapeSeriesDetails(req *http.Request, res *http.Response) (models.SeriesDetails, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return models.SeriesDetails{}, err
	}

	var details models.SeriesDetails
	var genres, directors, countries, casts []string

	doc.Find("div.content blockquote strong").Remove()

	originalURL := req.URL.Path
	parts := strings.Split(originalURL, "/")
	details.ID = parts[len(parts)-1]

	posterImg, _ := doc.Find("div.content-poster figure > picture > img").Attr("src")
	title, _ := doc.Find("div.content-poster figure > picture > img").Attr("alt")

	details.Title = title
	details.Type = "series"
	details.PosterImg = fmt.Sprintf("https:%s", posterImg)

	doc.Find("div.content > div").Each(func(i int, s *goquery.Selection) {
		switch strings.ToLower(s.Find("h2").Text()) {
		case "durasi":
			details.Duration = strings.TrimSpace(s.Find("h3").Text())
		case "imdb":
			details.Rating = strings.TrimSpace(s.Find("h3:nth-child(2)").Text())
		case "diterbitkan":
			details.ReleaseDate = strings.TrimSpace(s.Find("h3").Text())
		case "status":
			details.Status = strings.TrimSpace(strings.ToLower(s.Find("h3 > span").Text()))
		case "sutradara":
			s.Find("h3 > a").Each(func(i int, s2 *goquery.Selection) {
				directors = append(directors, strings.TrimSpace(s2.Text()))
			})
		case "negara":
			s.Find("h3 > a").Each(func(i int, s2 *goquery.Selection) {
				countries = append(countries, s2.Text())
			})
		case "genre":
			s.Find("h3 > a").Each(func(i int, s2 *goquery.Selection) {
				genres = append(genres, s2.Text())
			})
		case "bintang film":
			s.Find("h3").Each(func(i int, s2 *goquery.Selection) {
				casts = append(casts, s2.Find("a").Text())
			})
		}
	})

	details.Synopsis = doc.Find("div.content blockquote").Text()
	trailerURL, _ := doc.Find("div.player-content > iframe").Attr("src")
	details.TrailerURL = trailerURL
	details.Genres = genres
	details.Directors = directors
	details.Countries = countries
	details.Casts = casts

	var seasons []models.SeasonsList
	epsElem := doc.Find("div.serial-wrapper > div.episode-list")
	for i := len(epsElem.Nodes); i >= 1; i-- {
		season := models.SeasonsList{
			Season:        i,
			TotalEpisodes: epsElem.Eq(len(epsElem.Nodes) - i).Find("a.btn-primary").Length(),
		}
		seasons = append(seasons, season)
	}
	details.Seasons = seasons

	return details, nil
}
