package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/baobei23/lk21-go/models"
)

func ScrapeSearchedMoviesOrSeries(r *http.Request, res *http.Response) ([]models.SearchedMoviesOrSeries, error) {
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal memuat dokumen HTML: %w", err)
	}

	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}
	host := r.Host

	var payload []models.SearchedMoviesOrSeries

	doc.Find("div.search-wrapper > div.search-item").Each(func(i int, s *goquery.Selection) {
		content := s.Find("div.search-content")
		titleLink := content.Find("h3 > a")

		title := titleLink.Text()
		itemType := "movies"
		if strings.Contains(strings.ToLower(title), "series") {
			itemType = "series"
		}

		href, _ := titleLink.Attr("href")
		movieID := strings.Trim(href, "/")

		posterImg, _ := s.Find("figure > a > img").Last().Attr("src")

		fullURL := fmt.Sprintf("%s://%s/%s/%s", protocol, host, itemType, movieID)

		var genres []string
		content.Find("p").Each(func(j int, p *goquery.Selection) {

			strongTag := p.Find("strong")
			if strings.ToLower(strings.TrimSpace(strongTag.Text())) == "genres:" {
				strongTag.Remove()
				genresText := strings.TrimSpace(p.Text())
				if genresText != "" {
					genres = strings.Split(genresText, ", ")
				}
			}
		})

		payload = append(payload, models.SearchedMoviesOrSeries{
			ID:        movieID,
			Title:     title,
			Type:      itemType,
			PosterImg: posterImg,
			URL:       fullURL,
			Genres:    genres,
		})
	})

	return payload, nil
}
