package frmt

import (
	"server/database/fetch"
	"sync"
)

//	===============================================================================================

func FormatAllData(apiData fetch.ApiData) FormattedData {
	var frmtData FormattedData
	var wg sync.WaitGroup

	frmtFuncs := []func(){
		func() { frmtData.Pokemon = convertPokemonData(apiData.Pokemon, apiData.Species) },
		func() { frmtData.LrndMoves = convertLrndMovesData(apiData.Pokemon) },
		func() { frmtData.Pokedex = convertPokedexData(apiData.Species) },
		func() { frmtData.Abilities = convertAbilitiesData(apiData.Abilities) },
		func() { frmtData.AbltyDescs = convertAbltyDescData(apiData.Abilities) },
		func() { frmtData.Moves = convertMovesData(apiData.Moves) },
		func() { frmtData.MoveDescs = convertMoveDescData(apiData.Moves) },
		func() { frmtData.Machines = convertMachinesData(apiData.Machines) },
	}
	wg.Add(len(frmtFuncs))

	for _, fn := range frmtFuncs {
		go func(f func()) {
			defer wg.Done()
			f()
		}(fn)
	}
	wg.Wait()

	return frmtData
}

//	=============================================================================================
//	CONVERT POKEMON DATA

func convertPokemonData(pkmnData []fetch.Pokemon, spcsData []fetch.PokemonSpecies) []Pokemon {
	speciesMap := make(map[int]fetch.PokemonSpecies, len(spcsData))
	for _, s := range spcsData {
		speciesMap[s.Id] = s
	}

	pokemonSlice := make([]Pokemon, len(pkmnData))

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(pkmnData))

	for i, pkmn := range pkmnData {
		go func(i int, fPkmn fetch.Pokemon) {
			defer wg.Done()

			p := MakeTempPkmn(fPkmn, i+1)
			s := MakeTempSpcs(speciesMap[p.BaseId])

			pokemon := MakePokemonStruct(p, s)

			mu.Lock()
			pokemonSlice[i] = pokemon
			mu.Unlock()

		}(i, pkmn)
	}
	wg.Wait()

	return pokemonSlice
}

//	=============================================================================================
//	CONVERT LEARNED_MOVE DATA

func convertLrndMovesData(pkmnData []fetch.Pokemon) []LearnedMove {
	pool := &sync.Pool{
		New: func() interface{} {
			return &LearnedMove{}
		},
	}
	return convertWithSyncPool(pkmnData, MakeLearnedMoveSlice, pool)
}

//	=============================================================================================
//	CONVERT POKEDEX DATA

func convertPokedexData(spcsData []fetch.PokemonSpecies) []Pokedex {
	pool := &sync.Pool{
		New: func() interface{} {
			return &Pokedex{}
		},
	}
	return convertWithSyncPool(spcsData, MakePokedexSlice, pool)
}

//	=============================================================================================
//	CONVERT ABILITIES DATA

func convertAbilitiesData(abltyData []fetch.Ability) []Ability {
	return convertSlice(abltyData, MakeAbilityStruct)
}

//	=============================================================================================
//	CONVERT ABLTY_DESC DATA

func convertAbltyDescData(abltyData []fetch.Ability) []AbltyDesc {
	pool := &sync.Pool{
		New: func() interface{} {
			return &AbltyDesc{}
		},
	}
	return convertWithSyncPool(abltyData, MakeAbltyDescSlice, pool)
}

//	=============================================================================================
//	CONVERT MOVES DATA

func convertMovesData(moveData []fetch.Move) []Move {
	return convertSlice(moveData, MakeMoveStruct)
}

//	=============================================================================================
//	CONVERT MOVE_DESC DATA

func convertMoveDescData(moveData []fetch.Move) []MoveDesc {
	pool := &sync.Pool{
		New: func() interface{} {
			return &MoveDesc{}
		},
	}
	return convertWithSyncPool(moveData, MakeMoveDescSlice, pool)
}

//	=============================================================================================
//	CONVERT MACHINES DATA

func convertMachinesData(mchnData []fetch.Machine) []Machine {
	return convertSlice(mchnData, MakeMachineStruct)
}

//	=============================================================================================
//	GENERIC CONVERSIONS

type sliceConv[A any, B any] func(A) []B

func convertWithSyncPool[A any, B any](data []A, convFunc sliceConv[A, B], pool *sync.Pool) []B {
	var result []B

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(data))

	for _, item := range data {
		go func(a A) {
			defer wg.Done()

			convSlice := convFunc(a)

			for _, conv := range convSlice {
				b := pool.Get().(*B)
				*b = conv

				mu.Lock()
				result = append(result, *b)
				mu.Unlock()

				pool.Put(b)
			}
		}(item)
	}
	wg.Wait()

	return result
}

//	---------------------------------------------------------------------------------------------

type structConv[A any, B any] func(A) B

func convertSlice[A any, B any](data []A, convFunc structConv[A, B]) []B {
	result := make([]B, len(data))

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(data))

	for i, item := range data {
		go func(i int, a A) {
			defer wg.Done()

			b := convFunc(a)

			mu.Lock()
			result[i] = b
			mu.Unlock()

		}(i, item)
	}
	wg.Wait()

	return result
}
