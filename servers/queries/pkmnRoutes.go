package queries

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

//	===============================================================================================

//	URL QUERIES

// ======================================================
type PkmnQueries struct {
	Types, Color, Abilities, Size, Stats string
}

type pkmnFilters struct {
	Where, Order string
	WhereVals    []interface{}
}

// ===============================================================================================

//	DATABASE QUERIES

// ======================================================
func pkmnRowQuery(db *gorm.DB) *gorm.DB {
	return db.Table("pkmn").Select("id, search, name, sprite, type1, type2")
}

// -----------------------------------------------------------------------------------------------
func basePkmnQuery(db *gorm.DB) *gorm.DB {
	return db.Table("pkmn p").
		Select("p.*, a1.name AS a1Name, a1.summary AS a1Sum, a2.name AS a2Name, " +
			"a2.summary AS a2Sum, ah.name AS ahName, ah.summary AS ahSum").
		Joins("LEFT JOIN abltys a1 ON p.ablty1 = a1.search").
		Joins("LEFT JOIN abltys a2 ON p.ablty2 = a2.search").
		Joins("LEFT JOIN abltys ah ON p.ablty_h = ah.search")
}

// -----------------------------------------------------------------------------------------------
func customPkmnQuery(db *gorm.DB, queryObj PkmnQueries) *gorm.DB {
	query := basePkmnQuery(db)

	filters := getPkmnFilters(queryObj)

	if filters.Where != "" {
		query = query.Where(filters.Where, filters.WhereVals...)
	}

	if filters.Order != "" {
		query = query.Order(filters.Order)
	}

	return query
}

//	===============================================================================================

//	ROUTE FUNCTIONS

// ======================================================
func GetAllPokemon(resp Responder, db *gorm.DB) {
	var tempPkmn []TempPkmnRow

	if err := pkmnRowQuery(db).Scan(&tempPkmn).Error; err != nil {
		resp.InternalServerErr(PkmnQueryErr)
		return
	}

	if err := resp.ReturnAsJSON(ConvertPkmnRow(tempPkmn)); err != nil {
		resp.EncodingErr()
		return
	}
}

// -----------------------------------------------------------------------------------------------
func GetPokemonById(path string, resp Responder, db *gorm.DB) {
	var tempPkmn TempPokemon

	if err := basePkmnQuery(db).Where("p.id = ? OR p.search = ?", path, path).
		Scan(&tempPkmn).Error; err != nil {
		resp.InternalServerErr(PkmnQueryErr)
		return
	}

	if tempPkmn.Id == 0 {
		resp.NotFoundErr(PkmnNotFound)
		return
	}

	if err := resp.ReturnAsJSON(tempPkmn.Format()); err != nil {
		resp.EncodingErr()
		return
	}
}

// -----------------------------------------------------------------------------------------------
func GetPokemonByQueries(resp Responder, db *gorm.DB, queryObj PkmnQueries) {
	var tempPkmn []TempPokemon

	if err := customPkmnQuery(db, queryObj).Scan(&tempPkmn).Error; err != nil {
		resp.InternalServerErr(PkmnQueryErr)
		return
	}

	if len(tempPkmn) == 0 {
		resp.NotFoundErr(PkmnNotFound)
		return
	}

	if err := resp.ReturnAsJSON(ConvertPokemon(tempPkmn)); err != nil {
		resp.EncodingErr()
		return
	}
}

//	===============================================================================================

//	FILTERING FUNCTIONS

// ======================================================
func getPkmnFilters(queryObj PkmnQueries) pkmnFilters {
	var filters pkmnFilters
	var whereClauses []string
	var whereVals []interface{}
	var orderClause string

	if queryObj.Types != "" {
		typeWhere, typeVals := buildPkmnTypeSearch(queryObj.Types)
		whereClauses = append(whereClauses, typeWhere)
		whereVals = append(whereVals, typeVals...)
	}

	if queryObj.Color != "" {
		colorWhere, colorVal := buildColorSearch(queryObj.Color)
		if colorWhere != "" {
			whereClauses = append(whereClauses, colorWhere)
			whereVals = append(whereVals, colorVal)
		}
	}

	if queryObj.Abilities != "" {
		abltyWhere, albltyVals := buildAbltySearch(queryObj.Abilities)
		whereClauses = append(whereClauses, abltyWhere)
		whereVals = append(whereVals, albltyVals...)
	}

	orderClause = setPkmnQueryOrder(queryObj.Stats, queryObj.Size)

	if len(whereClauses) > 0 {
		filters.Where = strings.Join(whereClauses, " AND ")
		filters.WhereVals = whereVals
	}

	if orderClause != "" {
		filters.Order = orderClause
	}

	return filters
}

// ------------------------------------------------------
func buildPkmnTypeSearch(types string) (string, []interface{}) {
	typesList := strings.Split(types, ",")

	if len(typesList) > 2 {
		typesList = typesList[:2]
	}

	var whereClauses []string
	var whereVals []interface{}

	switch len(typesList) {
	case 1:
		whereClauses = append(whereClauses, "(type1 = ? OR type2 = ?)")
		whereVals = append(whereVals, typesList[0], typesList[0])
	case 2:
		whereClauses = append(whereClauses, ("((type1 = ? AND type2 = ?) OR (type1 = ? AND type2 = ?))"))
		whereVals = append(whereVals, typesList[0], typesList[1], typesList[1], typesList[0])
	}
	return strings.Join(whereClauses, " "), whereVals
}

// ------------------------------------------------------
func buildColorSearch(color string) (string, interface{}) {
	colorList := strings.Split(color, ",")

	if len(colorList) > 1 {
		color = colorList[0]
	}

	colorMap := map[string]bool{
		"red": true, "yellow": true, "green": true, "blue": true, "purple": true,
		"white": true, "black": true, "gray": true, "pink": true, "brown": true,
	}

	if !colorMap[color] {
		return "", nil
	}

	return "color = ?", color
}

// ------------------------------------------------------
func buildAbltySearch(abilities string) (string, []interface{}) {
	abilityList := strings.Split(abilities, ",")

	if len(abilityList) > 3 {
		abilityList = abilityList[:3]
	}

	var whereClauses []string
	var whereVals []interface{}

	switch len(abilityList) {
	case 1:
		whereClauses = append(whereClauses, "(ablty1 = ? OR ablty2 = ? OR ablty_h = ?)")
		whereVals = append(whereVals, abilityList[0], abilityList[0], abilityList[0])
	case 2:
		whereClauses = append(whereClauses,
			"((ablty1 = ? OR ablty2 = ? OR  ablty_h = ?) AND (ablty1 = ? OR ablty2 = ? OR ablty_h = ?))")
		whereVals = append(whereVals,
			abilityList[0], abilityList[0], abilityList[0], abilityList[1], abilityList[1], abilityList[1])
	case 3:
		whereClauses = append(whereClauses,
			"((ablty1 = ? OR ablty2 = ? OR ablty_h = ?) AND (ablty1 = ? OR ablty2 = ? OR ablty_h = ?) AND "+
				"(ablty1 = ? OR ablty2 = ? OR ablty_h = ?))")
		whereVals = append(whereVals,
			abilityList[0], abilityList[0], abilityList[0], abilityList[1], abilityList[1], abilityList[1],
			abilityList[2], abilityList[2], abilityList[2])
	}
	return strings.Join(whereClauses, " "), whereVals
}

// ------------------------------------------------------
func setPkmnQueryOrder(stats, size string) string {
	var orderClauses []string

	if stats != "" {
		statList := strings.Split(stats, ",")
		statsMap := map[string]bool{
			"hp": true, "atk": true, "def": true, "spa": true, "spd": true, "spe": true,
		}

		for _, stat := range statList {
			if statsMap[stat] {
				orderClauses = append(orderClauses, fmt.Sprintf("%s DESC", stat))
			}
		}
	}

	if size != "" {
		sizeList := strings.Split(size, ",")
		sizeMap := map[string]bool{
			"height": true, "weight": true,
		}

		for _, s := range sizeList {
			if sizeMap[s] {
				orderClauses = append(orderClauses, fmt.Sprintf("%s DESC", s))
			}
		}
	}
	return strings.Join(orderClauses, ", ")
}
