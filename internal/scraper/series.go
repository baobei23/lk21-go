package scraper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/baobei23/lk21-go/models"
)

// ScrapeSeries scrapes a list of series with a focus on direct efficiency.
func ScrapeSeries(req *http.Request, res *http.Response) ([]models.Series, error) {
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat dokumen goquery: %w", err)
	}

	selections := doc.Find("#grid-wrapper .infscroll-item")

	seriesList := make([]models.Series, 0, selections.Length())

	selections.Each(func(i int, s *goquery.Selection) {
		article := s.Find("article.mega-item")
		linkTag := article.Find("figure > a")

		href, _ := linkTag.Attr("href")

		seriesIDParts := strings.Split(href, "/")
		seriesID := seriesIDParts[len(seriesIDParts)-2]

		imgTag := linkTag.Find("img")
		posterImg, _ := imgTag.Attr("src")

		var genres []string
		article.Find(".grid-categories a").Each(func(_ int, genreLink *goquery.Selection) {
			genreHref, _ := genreLink.Attr("href")
			parts := strings.Split(genreHref, "/")
			if len(parts) > 2 && parts[1] == "genre" {
				genres = append(genres, parts[2])
			}
		})

		episodeStr := article.Find(".last-episode span").Text()
		episode, _ := strconv.Atoi(episodeStr)

		series := models.Series{
			ID:        seriesID,
			Title:     imgTag.AttrOr("alt", seriesID),
			Type:      "series",
			PosterImg: fmt.Sprintf("https:%s", posterImg),
			Rating:    article.Find(".rating").Text(),
			URL:       fmt.Sprintf("%s://%s/series/%s", scheme, req.Host, seriesID),
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
