package scraper

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/baobei23/lk21-go/models"
)

func ScrapeStreamSources(res *http.Response) ([]models.StreamSources, error) {

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {

		return nil, err
	}

	payload := []models.StreamSources{}

	doc.Find("div#load-sources ul > li").Each(func(i int, s *goquery.Selection) {
		var resolutions []string

		s.Find("div > span").Each(func(j int, span *goquery.Selection) {
			resolutions = append(resolutions, span.Text())
		})

		provider := s.Find("a").Text()

		url, _ := s.Find("a").Attr("href")

		payload = append(payload, models.StreamSources{
			Provider:    provider,
			URL:         url,
			Resolutions: resolutions,
		})
	})

	return payload, nil
}
