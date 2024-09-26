package queries

import (
	"gorm.io/gorm"
)

//	===============================================================================================

//	URL QUERIES

// ======================================================

// ===============================================================================================

//	DATABASE QUERIES

// ======================================================
func baseAbltyQuery(db *gorm.DB) *gorm.DB {
	return db.Table("abltys")
}

// -----------------------------------------------------------------------------------------------

// -----------------------------------------------------------------------------------------------

//	===============================================================================================

//	ROUTE FUNCTIONS

// ======================================================
func GetAllAbilities(resp Responder, db *gorm.DB) {
	var tempAbltys []TempAbility

	if err := baseAbltyQuery(db).Scan(&tempAbltys).Error; err != nil {
		resp.InternalServerErr(AbltyQueryErr)
		return
	}

	if err := resp.ReturnAsJSON(ConvertAbilities(tempAbltys)); err != nil {
		resp.EncodingErr()
		return
	}
}

// -----------------------------------------------------------------------------------------------
func GetAbilityById(path string, resp Responder, db *gorm.DB) {
	var tempAblty TempAbility

	if err := baseAbltyQuery(db).Where("id = ? OR search = ?", path, path).
		Scan(&tempAblty).Error; err != nil {
		resp.InternalServerErr(AbltyQueryErr)
		return
	}

	if tempAblty.Id == 0 {
		resp.NotFoundErr(AbltyNotFound)
		return
	}

	if err := resp.ReturnAsJSON(tempAblty.Format()); err != nil {
		resp.EncodingErr()
		return
	}
}
