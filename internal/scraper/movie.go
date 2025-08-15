package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/baobei23/lk21-go/models"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeMovies scrapes a list of movies from a given URL
func ScrapeMovies(req *http.Request, res *http.Response) ([]models.Movie, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var movies []models.Movie

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
		movieID := strings.Split(href, "/")

		posterImg, _ := parent.Find("figure > a > picture > img").Attr("src")
		title, _ := parent.Find("figure > a > picture > img").Attr("alt")

		movie := models.Movie{
			ID:                movieID[len(movieID)-2],
			Title:             title,
			Type:              "movie",
			PosterImg:         fmt.Sprintf("https:%s", posterImg),
			Rating:            parent.Find("figure div.rating").Text(),
			URL:               fmt.Sprintf("%s://%s/movies/%s", req.URL.Scheme, req.Host, movieID[len(movieID)-2]),
			QualityResolution: parent.Find("figure div.quality").Text(),
			Genres:            genres,
		}

		movies = append(movies, movie)
	})

	return movies, nil
}

// ScrapeMovieDetails scrapes the details of a single movie
func ScrapeMovieDetails(req *http.Request, res *http.Response) (models.MovieDetails, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return models.MovieDetails{}, err
	}

	var details models.MovieDetails
	var genres, directors, countries, casts []string

	doc.Find("div.content blockquote strong").Remove()

	originalURL := req.URL.Path
	parts := strings.Split(originalURL, "/")
	details.ID = parts[len(parts)-1]

	posterImg, _ := doc.Find("div.content-poster figure > picture > img").Attr("src")
	title, _ := doc.Find("div.content-poster figure > picture > img").Attr("alt")

	details.Title = title
	details.Type = "movie"
	details.PosterImg = fmt.Sprintf("https:%s", posterImg)

	doc.Find("div.content > div").Each(func(i int, s *goquery.Selection) {
		switch strings.ToLower(s.Find("h2").Text()) {
		case "durasi":
			details.Duration = strings.TrimSpace(s.Find("h3").Text())
		case "imdb":
			details.Rating = strings.TrimSpace(s.Find("h3:nth-child(2)").Text())
		case "diterbitkan":
			details.ReleaseDate = strings.TrimSpace(s.Find("h3").Text())
		case "kualitas":
			details.Quality = strings.TrimSpace(s.Find("h3 > a").Text())
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
	trailerURL, _ := doc.Find("div.action-player a.fancybox").Attr("href")
	details.TrailerURL = trailerURL
	details.Genres = genres
	details.Directors = directors
	details.Countries = countries
	details.Casts = casts

	return details, nil
}
