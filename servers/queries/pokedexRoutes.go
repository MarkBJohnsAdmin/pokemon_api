package queries

import (
	"gorm.io/gorm"
)

//	===============================================================================================

//	URL QUERIES

// ======================================================
type PdexQueries struct {
	Version string
}

// ===============================================================================================

//	DATABASE QUERIES

// ======================================================
func basePdexPkmnQuery(db *gorm.DB) *gorm.DB {
	return db.Table("pkmn").
		Select("id, search, name, sprite, offcl, shiny, home, hshiny")
}

// -----------------------------------------------------------------------------------------------
func pdexVersionIds(db *gorm.DB, queryObj PdexQueries) *gorm.DB {
	return db.Table("pdex").Select("pkmn_id").
		Where("version = ?", queryObj.Version)
}

// -----------------------------------------------------------------------------------------------
func basePdexVersionQuery(db *gorm.DB) *gorm.DB {
	return db.Table("pkmn p").
		Select("p.id, p.search, p.name, p.sprite, d.entry").
		Joins("LEFT JOIN pdex d ON p.base_id = d.pkmn_id")
}

//	===============================================================================================

//	ROUTE FUNCTIONS

// ======================================================
func GetPokedexByPokemon(path string, resp Responder, db *gorm.DB) {
	var temp TempPdexPkmn
	var dex []pdexVersion

	if err := basePdexPkmnQuery(db).Where("base_id = ? OR search = ?", path, path).
		Scan(&temp).Error; err != nil {
		resp.InternalServerErr(PkmnQueryErr)
		return
	}

	if temp.Id == 0 {
		resp.NotFoundErr(PkmnNotFound)
		return
	}

	if err := db.Table("pdex").Where("pkmn_id = ? OR search = ?", path, path).
		Scan(&dex).Error; err != nil {
		resp.InternalServerErr(PdexQueryErr)
		return
	}

	if len(dex) == 0 {
		resp.NotFoundErr(PdexNotFound)
		return
	}

	if err := resp.ReturnAsJSON(temp.Join(dex)); err != nil {
		resp.EncodingErr()
		return
	}
}

// -----------------------------------------------------------------------------------------------
func GetPokedexByVersion(resp Responder, db *gorm.DB, queryObj PdexQueries) {
	var ids []int
	var temp []pdexPkmn

	if err := pdexVersionIds(db, queryObj).Scan(&ids).Error; err != nil {
		resp.InternalServerErr(PdexQueryErr)
		return
	}

	if len(ids) == 0 {
		resp.NotFoundErr(PdexNotFound)
		return
	}

	if err := basePdexVersionQuery(db).
		Where("p.base_id IN ? AND d.version = ?", ids, queryObj.Version).
		Scan(&temp).Error; err != nil {
		resp.InternalServerErr(PkmnQueryErr)
		return
	}

	if len(temp) == 0 {
		resp.NotFoundErr(PkmnNotFound)
		return
	}

	if err := resp.ReturnAsJSON(ConvertPdexVersion(queryObj.Version, temp)); err != nil {
		resp.EncodingErr()
		return
	}
}

// -----------------------------------------------------------------------------------------------
func GetPokedexByPkmnAndVersion(path string, resp Responder, db *gorm.DB, queryObj PdexQueries) {
	var temp TempPdexPkmn
	var dex []pdexVersion

	if err := basePdexPkmnQuery(db).Where("id = ? OR search = ?", path, path).
		Scan(&temp).Error; err != nil {
		resp.InternalServerErr(PkmnQueryErr)
		return
	}

	if temp.Id == 0 {
		resp.NotFoundErr(PkmnNotFound)
		return
	}

	if err := db.Table("pdex").
		Where("(pkmn_id = ? OR search = ?) AND version = ?", path, path, queryObj.Version).
		Scan(&dex).Error; err != nil {
		resp.InternalServerErr(PdexQueryErr)
		return
	}

	if len(dex) == 0 {
		resp.NotFoundErr(PdexNotFound)
		return
	}

	if err := resp.ReturnAsJSON(temp.Join(dex)); err != nil {
		resp.EncodingErr()
		return
	}
}
