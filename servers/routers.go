package servers

import (
	"net/http"
	"reflect"
	"server/servers/queries"
	"strings"

	"gorm.io/gorm"
)

// ===============================================================================================
const AccessControlAllowOrigin = "Access-Control-Allow-Origin"

const AccessControlAllowMethods = "Access-Control-Allow-Methods"

const GET_OPTIONS = "GET, OPTIONS"

const AccessControlAllowHeaders = "Access-Control-Allow-Headers"

const ContentType = "Content-Type"

const MethodNotAllowed = "Method not allowed"

// ===============================================================================================
func SetUpPokemonRouter(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/pokemon/", func(w http.ResponseWriter, r *http.Request) {
		resp := queries.Responder{Writer: w, Status: http.StatusOK}

		resp.SetCORSHeaders()

		if r.Method == http.MethodOptions {
			resp.SetStatus(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			resp.MethodNotAllowedErr()
			return
		}

		queryList := queries.PkmnQueries{
			Types:     r.URL.Query().Get("types"),
			Color:     r.URL.Query().Get("color"),
			Abilities: r.URL.Query().Get("abilities"),
			Size:      r.URL.Query().Get("size"),
			Stats:     r.URL.Query().Get("stats"),
		}

		if hasQueries(queryList) {
			queries.GetPokemonByQueries(resp, db, queryList)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/pokemon/")
		if path != "" {
			queries.GetPokemonById(path, resp, db)
			return
		}

		queries.GetAllPokemon(resp, db)
	})
	return mux
}

// -----------------------------------------------------------------------------------------------
func SetUpPokedexRouter(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/pokedex/", func(w http.ResponseWriter, r *http.Request) {
		resp := queries.Responder{Writer: w, Status: http.StatusOK}

		resp.SetCORSHeaders()

		if r.Method == http.MethodOptions {
			resp.SetStatus(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			resp.MethodNotAllowedErr()
			return
		}

		queryList := queries.PdexQueries{
			Version: r.URL.Query().Get("version"),
		}
		path := strings.TrimPrefix(r.URL.Path, "/pokedex/")

		if hasQueries(queryList) && path != "" {
			queries.GetPokedexByPkmnAndVersion(path, resp, db, queryList)
			return
		}

		if hasQueries(queryList) && path == "" {
			queries.GetPokedexByVersion(resp, db, queryList)
			return
		}

		if !hasQueries(queryList) && path != "" {
			queries.GetPokedexByPokemon(path, resp, db)
			return
		}
	})
	return mux
}

// -----------------------------------------------------------------------------------------------
func SetUpAbilityRouter(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ability/", func(w http.ResponseWriter, r *http.Request) {
		resp := queries.Responder{Writer: w, Status: http.StatusOK}

		resp.SetCORSHeaders()

		if r.Method == http.MethodOptions {
			resp.SetStatus(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			resp.MethodNotAllowedErr()
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/ability/")
		if path != "" {
			queries.GetAbilityById(path, resp, db)
			return
		}

		queries.GetAllAbilities(resp, db)
	})
	return mux
}

// -----------------------------------------------------------------------------------------------
func SetUpMoveRouter(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/move/", func(w http.ResponseWriter, r *http.Request) {
		resp := queries.Responder{Writer: w, Status: http.StatusOK}

		resp.SetCORSHeaders()

		if r.Method == http.MethodOptions {
			resp.SetStatus(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			resp.MethodNotAllowedErr()
			return
		}

		queryList := queries.MoveQueries{
			Type:   r.URL.Query().Get("type"),
			Class:  r.URL.Query().Get("class"),
			Target: r.URL.Query().Get("target"),
			Sort:   r.URL.Query().Get("sort"),
		}

		if hasQueries(queryList) {
			queries.GetMovesByQuery(resp, db, queryList)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/move/")
		if path != "" {
			queries.GetMoveById(path, resp, db)
			return
		}

		queries.GetAllMoves(resp, db)
	})
	return mux
}

//	===============================================================================================

func hasQueries(v interface{}) bool {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.String() != "" {
			return true
		}
		if field.Kind() == reflect.Bool && field.Bool() {
			return true
		}

	}
	return false
}

func SetUpTestRouter(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		resp := queries.Responder{Writer: w, Status: http.StatusOK}

		w.Header().Set(AccessControlAllowOrigin, "*")
		w.Header().Set(AccessControlAllowMethods, GET_OPTIONS)
		w.Header().Set(AccessControlAllowHeaders, ContentType)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		queryList := queries.PdexQueries{
			Version: r.URL.Query().Get("version"),
		}

		if hasQueries(queryList) {
			queries.GetPokedexByVersion(resp, db, queryList)
			return
		}

		queries.GetPokedexByPokemon("absol", resp, db)
	})
	return mux
}
