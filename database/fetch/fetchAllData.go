package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

//	===============================================================================================

var client = &http.Client{}

type reqs struct {
	id  string
	idx int
}

//	===============================================================================================

func FetchAllPokeApiData() ApiData {
	var apiData ApiData

	var wg sync.WaitGroup

	fetchFuncs := []func(){
		func() { apiData.Pokemon = fetchPokemon() },
		func() { apiData.Species = fetchSpecies() },
		func() { apiData.Abilities = fetchAbilties() },
		func() { apiData.Moves = fetchMoves() },
		func() { apiData.Machines = fetchMachines() },
	}

	wg.Add(len(fetchFuncs))

	for _, fn := range fetchFuncs {
		go func(f func()) {
			defer wg.Done()
			f()
		}(fn)
	}

	wg.Wait()

	return apiData
}

//	===============================================================================================

//	Generic fetching functions

func fetchResp(route, id string) ([]byte, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/%s/%s", route, id)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("no response from %s: \n%v", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: \n%v", url, err)
	}
	return body, nil
}

func fetchAndConvertAll[T any](route string, ids []string) []T {
	responses := make([]T, len(ids))

	maxRetry := 5

	var wg sync.WaitGroup
	wg.Add(len(ids))

	for i, id := range ids {
		go func(i int, id string) {
			defer wg.Done()

			queue := []reqs{{id: id, idx: i}}

			for len(queue) > 0 {
				crntReq := queue[0]
				queue = queue[1:]

				var resp T

				data, err := fetchResp(route, crntReq.id)
				if err != nil {
					if len(queue) >= maxRetry {
						log.Printf("no resp from %s/%s: \n%v", route, crntReq.id, err)
					} else {
						backoff := time.Duration(500) * time.Millisecond
						time.Sleep(backoff)
						queue = append(queue, crntReq)
					}
					continue
				}

				if err := json.Unmarshal(data, &resp); err != nil {
					if len(queue) >= maxRetry {
						log.Printf("could not format %s/%s: \n%v", route, crntReq.id, err)
					} else {
						backoff := time.Duration(500*(len(queue)+1)) * time.Millisecond
						time.Sleep(backoff)
						queue = append(queue, crntReq)
					}
					continue
				}
				responses[crntReq.idx] = resp

			}
		}(i, id)
	}
	wg.Wait()

	return responses
}

//	===============================================================================================

//	Route Specific Functions

func fetchPokemon() []Pokemon {
	return fetchAndConvertAll[Pokemon]("pokemon", PokemonIds())
}

func fetchSpecies() []PokemonSpecies {
	return fetchAndConvertAll[PokemonSpecies]("pokemon-species", SpeciesIds())
}

func fetchAbilties() []Ability {
	return fetchAndConvertAll[Ability]("ability", AbilityIds())
}

func fetchMoves() []Move {
	return fetchAndConvertAll[Move]("move", MoveIds())
}

func fetchMachines() []Machine {
	return fetchAndConvertAll[Machine]("machine", MachineIds())
}
