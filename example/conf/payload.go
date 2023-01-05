package conf

import (
	_ "embed"
)

//go:embed payload/top-movies.json
var TopMoviesPayload string

//go:embed payload/top-actors.json
var TopActorsPayload string

//go:embed payload/movies/399566.json
var Movie399566Payload string

//go:embed payload/movies/399566/actors/15556.json
var Movie399566Actor15556Payload string

//go:embed payload/movies/suggest.json
var MovieSuggestPayload string
