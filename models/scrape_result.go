package models

// Movie struct corresponds to IMovies interface in TypeScript
type Movie struct {
	ID                string   `json:"_id"`
	Title             string   `json:"title"`
	Type              string   `json:"type"`
	PosterImg         string   `json:"posterImg"`
	Rating            string   `json:"rating"`
	URL               string   `json:"url"`
	QualityResolution string   `json:"qualityResolution"`
	Genres            []string `json:"genres"`
}

// MovieDetails struct corresponds to IMovieDetails interface in TypeScript
type MovieDetails struct {
	ID          string   `json:"_id"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	PosterImg   string   `json:"posterImg"`
	Rating      string   `json:"rating"`
	Quality     string   `json:"quality"`
	ReleaseDate string   `json:"releaseDate"`
	Synopsis    string   `json:"synopsis"`
	Duration    string   `json:"duration"`
	TrailerURL  string   `json:"trailerUrl"`
	Directors   []string `json:"directors"`
	Countries   []string `json:"countries"`
	Casts       []string `json:"casts"`
	Genres      []string `json:"genres"`
}

// Series struct corresponds to ISeries interface in TypeScript
type Series struct {
	ID        string   `json:"_id"`
	Title     string   `json:"title"`
	Type      string   `json:"type"`
	PosterImg string   `json:"posterImg"`
	Rating    string   `json:"rating"`
	URL       string   `json:"url"`
	Episode   int      `json:"episode"`
	Genres    []string `json:"genres"`
}

// SeasonsList struct corresponds to ISeasonsList interface in TypeScript
type SeasonsList struct {
	Season        int `json:"season"`
	TotalEpisodes int `json:"totalEpisodes"`
}

// SeriesDetails struct corresponds to ISeriesDetails interface in TypeScript
type SeriesDetails struct {
	ID          string        `json:"_id"`
	Title       string        `json:"title"`
	Type        string        `json:"type"`
	PosterImg   string        `json:"posterImg"`
	Rating      string        `json:"rating"`
	Status      string        `json:"status"`
	ReleaseDate string        `json:"releaseDate"`
	Synopsis    string        `json:"synopsis"`
	Duration    string        `json:"duration"`
	TrailerURL  string        `json:"trailerUrl"`
	Directors   []string      `json:"directors"`
	Countries   []string      `json:"countries"`
	Casts       []string      `json:"casts"`
	Genres      []string      `json:"genres"`
	Seasons     []SeasonsList `json:"seasons"`
}

type StreamSources struct {
	Provider    string   `json:"provider"`
	URL         string   `json:"url"`
	Resolutions []string `json:"resolutions"`
}

type SearchedMoviesOrSeries struct {
	ID        string   `json:"_id"`
	Title     string   `json:"title"`
	Type      string   `json:"type"`
	PosterImg string   `json:"posterImg"`
	URL       string   `json:"url"`
	Genres    []string `json:"genres"`
}
