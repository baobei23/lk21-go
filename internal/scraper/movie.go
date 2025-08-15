package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/baobei23/lk21-go/models"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeMovies(req *http.Request, res *http.Response) ([]models.Movie, error) {
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat dokumen goquery: %w", err)
	}

	selections := doc.Find("#grid-wrapper .infscroll-item")

	movies := make([]models.Movie, 0, selections.Length())

	selections.Each(func(i int, s *goquery.Selection) {
		article := s.Find("article.mega-item")
		linkTag := article.Find("figure > a")

		href, _ := linkTag.Attr("href")

		movieIDParts := strings.Split(href, "/")
		movieID := movieIDParts[len(movieIDParts)-2]

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

		movie := models.Movie{
			ID:                movieID,
			Title:             imgTag.AttrOr("alt", ""),
			Type:              "movie",
			PosterImg:         fmt.Sprintf("https:%s", posterImg),
			Rating:            article.Find(".rating").Text(),
			URL:               fmt.Sprintf("%s://%s/movies/%s", scheme, req.Host, movieID),
			QualityResolution: article.Find(".quality").Text(),
			Genres:            genres,
		}

		movies = append(movies, movie)
	})

	return movies, nil
}

// ScrapeMovieDetails scrapes movie details with a focus on direct efficiency.
func ScrapeMovieDetails(req *http.Request, res *http.Response) (models.MovieDetails, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return models.MovieDetails{}, fmt.Errorf("gagal membuat dokumen goquery: %w", err)
	}

	var details models.MovieDetails

	contentWrapper := doc.Find("div.content-wrapper")

	// --- Ekstrak ID, Poster, dan Judul ---
	pathParts := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	details.ID = pathParts[len(pathParts)-1]

	imgTag := contentWrapper.Find(".content-poster img")
	details.PosterImg, _ = imgTag.Attr("src")
	if details.PosterImg != "" {
		details.PosterImg = fmt.Sprintf("https:%s", details.PosterImg)
	}
	details.Title = imgTag.AttrOr("alt", details.ID) // Fallback ke ID jika alt kosong
	details.Type = "movie"

	// --- Ekstrak detail dari setiap baris info ---
	contentWrapper.Find(".content > div").Each(func(i int, s *goquery.Selection) {
		// Konversi teks heading ke huruf kecil
		heading := strings.ToLower(s.Find("h2").Text())

		switch heading {
		case "durasi":
			details.Duration = s.Find("h3").Text()
		case "imdb":
			// BUG FIX: Ambil elemen h3 pertama untuk mendapatkan rating yang benar
			details.Rating = s.Find("h3").First().Text()
		case "diterbitkan":
			details.ReleaseDate = s.Find("h3").Text()
		case "kualitas":
			details.Quality = s.Find("h3 > a").Text()
		case "sutradara":
			s.Find("h3 > a").Each(func(_ int, director *goquery.Selection) {
				details.Directors = append(details.Directors, director.Text())
			})
		case "negara":
			s.Find("h3 > a").Each(func(_ int, country *goquery.Selection) {
				details.Countries = append(details.Countries, country.Text())
			})
		case "genre":
			s.Find("h3 > a").Each(func(_ int, genre *goquery.Selection) {
				details.Genres = append(details.Genres, genre.Text())
			})
		case "bintang film":
			s.Find("h3 > a").Each(func(_ int, cast *goquery.Selection) {
				details.Casts = append(details.Casts, cast.Text())
			})
		}
	})

	// --- Ekstrak Sinopsis & Trailer ---
	synopsisBlock := contentWrapper.Find("blockquote")
	synopsisBlock.Find("strong, a, br").Remove()
	details.Synopsis = strings.TrimSpace(synopsisBlock.Text())

	details.TrailerURL, _ = doc.Find(".action-player a.fancybox").Attr("href")

	return details, nil
}
