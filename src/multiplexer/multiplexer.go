package multiplexer

import (
	"net/http"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/router"
)

var Mux *http.ServeMux

func Initialize() {
	if Mux == nil {
		Mux = http.NewServeMux()
	}
	router.Init(Mux)
}
