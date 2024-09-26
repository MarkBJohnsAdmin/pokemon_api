package tables

import (
	frmt "server/database/format"
	"sync"
)

//	===============================================================================================

func PrepAllTables(frmtData frmt.FormattedData) DataTables {
	var dataTables DataTables
	var wg sync.WaitGroup

	tmMap := CreateTmMap(frmtData.Machines)

	tableFuncs := []func(){
		func() { dataTables.Pokemon = prepPokemonTable(frmtData.Pokemon) },
		func() { dataTables.LearnedMoves = prepLearnedMovesTable(frmtData.LrndMoves, &tmMap) },
		func() { dataTables.Pokedex = prepPokedexTable(frmtData.Pokedex) },
		func() { dataTables.Abilities = prepAbilityTable(frmtData.Abilities) },
		func() { dataTables.AbltyDescs = prepAbltyDescsTable(frmtData.AbltyDescs) },
		func() { dataTables.Moves = prepMovesTable(frmtData.Moves) },
		func() { dataTables.MoveDescs = prepMoveDescsTable(frmtData.MoveDescs) },
	}

	wg.Add(len(tableFuncs))

	for _, fn := range tableFuncs {
		go func(f func()) {
			defer wg.Done()
			f()
		}(fn)
	}
	wg.Wait()

	return dataTables
}

//	===============================================================================================
//	TABLE PREP

func prepPokemonTable(pkmnData []frmt.Pokemon) []Pokemon {
	return addAllRows(pkmnData, MakePokemonRow)
}

//	-----------------------------------------

func prepLearnedMovesTable(lrndData []frmt.LearnedMove, tmMap *map[string][]tmName) []LearnedMove {
	convFunc := func(lrnd *frmt.LearnedMove) LearnedMove {
		return MakeLearnedMoveRow(lrnd, tmMap)
	}
	return addAllRows(lrndData, convFunc)
}

//	-----------------------------------------

func prepPokedexTable(pdexData []frmt.Pokedex) []Pokedex {
	return addAllRows(pdexData, MakePokedexRow)
}

//	-----------------------------------------

func prepAbilityTable(abltyData []frmt.Ability) []Ability {
	return addAllRows(abltyData, MakeAbilityRow)
}

//	-----------------------------------------

func prepAbltyDescsTable(descData []frmt.AbltyDesc) []AbltyDesc {
	return addAllRows(descData, MakeAbltyDescRow)
}

//	-----------------------------------------

func prepMovesTable(moveData []frmt.Move) []Move {
	return addAllRows(moveData, MakeMoveRow)
}

//	-----------------------------------------

func prepMoveDescsTable(descData []frmt.MoveDesc) []MoveDesc {
	return addAllRows(descData, MakeMoveDescRow)
}

//	===============================================================================================
//	GENERIC TABLE PREP

type structConv[A any, B any] func(*A) B

func addAllRows[A any, B any](data []A, convFunc structConv[A, B]) []B {
	result := make([]B, len(data))

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(data))

	for i, item := range data {
		go func(i int, a *A) {
			defer wg.Done()

			b := convFunc(a)

			mu.Lock()
			result[i] = b
			mu.Unlock()

		}(i, &item)
	}
	wg.Wait()

	return result
}
