package upload

import (
	"fmt"
	"server/database/fetch"
	frmt "server/database/format"
	"server/database/tables"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//	===============================================================================================

var batchSize = 100

//	===============================================================================================

func UploadAllTables(db *gorm.DB, tableData tables.DataTables) error {
	var uploadErrs []error
	var wg sync.WaitGroup
	var mu sync.Mutex

	uploadFuncs := []func() error{
		func() error { return insertPokemonBatch(db, tableData.Pokemon) },
		func() error { return insertLearnedMoveBatch(db, tableData.LearnedMoves) },
		func() error { return insertPokedexBatch(db, tableData.Pokedex) },
		func() error { return insertAbilityBatch(db, tableData.Abilities) },
		func() error { return insertAbltyDescBatch(db, tableData.AbltyDescs) },
		func() error { return insertMoveBatch(db, tableData.Moves) },
		func() error { return insertMoveDescBatch(db, tableData.MoveDescs) },
	}

	wg.Add(len(uploadFuncs))

	for _, fn := range uploadFuncs {
		go func(f func() error) {
			defer wg.Done()
			if err := f(); err != nil {
				mu.Lock()
				uploadErrs = append(uploadErrs, err)
				mu.Unlock()
			}
		}(fn)
	}
	wg.Wait()

	if len(uploadErrs) > 0 {
		return uploadErrs[0]
	}

	return nil
}

//	===============================================================================================

func insertPokemonBatch(db *gorm.DB, pkmnData []tables.Pokemon) error {
	return insertBatch(db, "pokemon", pkmnData)
}

func insertLearnedMoveBatch(db *gorm.DB, lrndData []tables.LearnedMove) error {
	return insertBatch(db, "learned moves", lrndData)
}

func insertPokedexBatch(db *gorm.DB, pdexData []tables.Pokedex) error {
	return insertBatch(db, "pokedex", pdexData)
}

func insertAbilityBatch(db *gorm.DB, abltyData []tables.Ability) error {
	return insertBatch(db, "ability", abltyData)
}

func insertAbltyDescBatch(db *gorm.DB, descData []tables.AbltyDesc) error {
	return insertBatch(db, "ability description", descData, "desc", "version")
}

func insertMoveBatch(db *gorm.DB, moveData []tables.Move) error {
	return insertBatch(db, "move", moveData)
}

func insertMoveDescBatch(db *gorm.DB, descData []tables.MoveDesc) error {
	return insertBatch(db, "move description", descData)
}

//	===============================================================================================

func GetTableData() tables.DataTables {
	var apiData fetch.ApiData

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		apiData = fetch.FetchAllPokeApiData()
	}()

	wg.Wait()

	formatted := frmt.FormatAllData(apiData)

	return tables.PrepAllTables(formatted)
}

//	===============================================================================================

func insertBatch[T any](db *gorm.DB, table string, data []T, conflicts ...string) error {
	if err := db.AutoMigrate(new(T)); err != nil {
		return fmt.Errorf("failed to migrate %s table: %w", table, err)
	}

	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}

		batchData := data[i:end]

		err := db.Clauses(clause.OnConflict{
			Columns:   toClauseColumns(conflicts),
			DoUpdates: clause.AssignmentColumns(conflicts),
		}).CreateInBatches(batchData, batchSize).Error

		if err != nil {
			if gorm.ErrDuplicatedKey == err {
				for _, entry := range batchData {
					fmt.Printf("conflicting entry: %+v\n", entry)
				}
			}
			return fmt.Errorf("conflicting entries: %w", err)
		}
	}
	return nil
}

//	-----------------------------------------------------------------------------------------------

func toClauseColumns(columns []string) []clause.Column {
	var clauseCols []clause.Column

	for _, col := range columns {
		clauseCols = append(clauseCols, clause.Column{Name: col})
	}

	return clauseCols
}
