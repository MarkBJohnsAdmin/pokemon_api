package queries

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

//	===============================================================================================

//	URL QUERIES

// ======================================================

type MoveQueries struct {
	Type, Class, Target, Sort string
	// "Sort" orders my power, accuracy, pp, and / or priority
}

type moveFilters struct {
	Where, Order string
	WhereVals    []interface{}
}

// ===============================================================================================

//	DATABASE QUERIES

// ======================================================
func baseMoveQuery(db *gorm.DB) *gorm.DB {
	return db.Table("moves")
}

// -----------------------------------------------------------------------------------------------
func customMoveQuery(db *gorm.DB, queryObj MoveQueries) *gorm.DB {
	query := baseMoveQuery(db)

	filters := getMoveFilters(queryObj)

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
func GetAllMoves(resp Responder, db *gorm.DB) {
	var temp []TempMove

	if err := baseMoveQuery(db).Scan(&temp).Error; err != nil {
		resp.InternalServerErr(MoveQueryErr)
		return
	}

	if err := resp.ReturnAsJSON(ConvertMoves(temp)); err != nil {
		resp.EncodingErr()
		return
	}
}

// -----------------------------------------------------------------------------------------------
func GetMoveById(path string, resp Responder, db *gorm.DB) {
	var temp TempMove

	if err := baseMoveQuery(db).Where("id = ? OR search = ?", path, path).
		Scan(&temp).Error; err != nil {
		resp.InternalServerErr(MoveQueryErr)
		return
	}

	if temp.Id == 0 {
		resp.NotFoundErr(MoveNotFound)
		return
	}

	if err := resp.ReturnAsJSON(temp.Format()); err != nil {
		resp.EncodingErr()
		return
	}
}

// -----------------------------------------------------------------------------------------------
func GetMovesByQuery(resp Responder, db *gorm.DB, queryObj MoveQueries) {
	var temp []TempMove

	if err := customMoveQuery(db, queryObj).Scan(&temp).Error; err != nil {
		resp.InternalServerErr(MoveQueryErr)
		return
	}

	if len(temp) == 0 {
		resp.NotFoundErr(MoveNotFound)
		return
	}

	if err := resp.ReturnAsJSON(ConvertMoves(temp)); err != nil {
		resp.EncodingErr()
		return
	}
}

//	===============================================================================================

//	FILTERING FUNCTIONS

// ======================================================
func getMoveFilters(queryObj MoveQueries) moveFilters {
	var filters moveFilters
	var whereClauses []string
	var whereVals []interface{}
	var orderClause string

	if queryObj.Type != "" {
		typeWhere, typeVal := buildMoveTypeSearch(queryObj.Type)
		whereClauses = append(whereClauses, typeWhere)
		whereVals = append(whereVals, typeVal)
	}

	if queryObj.Class != "" {
		classWhere, classVal := buildMoveClassSearch(queryObj.Class)
		whereClauses = append(whereClauses, classWhere)
		whereVals = append(whereVals, classVal)
	}

	if queryObj.Target != "" {
		targetWhere, targetVal := buildMoveTargetSearch(queryObj.Target)
		whereClauses = append(whereClauses, targetWhere)
		whereVals = append(whereVals, targetVal)
	}

	orderClause = setMoveQueryOrder(queryObj.Sort)

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
func buildMoveTypeSearch(moveType string) (string, interface{}) {
	typeList := strings.Split(moveType, ",")

	if len(typeList) > 1 {
		moveType = typeList[0]
	}

	return "type = ?", moveType
}

// ------------------------------------------------------
func buildMoveClassSearch(dmgClass string) (string, interface{}) {
	classList := strings.Split(dmgClass, ",")

	if len(classList) > 1 {
		dmgClass = classList[0]
	}

	return "class = ?", dmgClass
}

// ------------------------------------------------------
func buildMoveTargetSearch(target string) (string, interface{}) {
	targetList := strings.Split(target, ",")

	if len(targetList) > 1 {
		target = targetList[0]
	}

	return "target = ?", target
}

// ------------------------------------------------------
func setMoveQueryOrder(sort string) string {
	clauseList := strings.Split(sort, ",")

	var orderClauses []string

	sortMap := map[string]bool{
		"power": true, "acc": true, "pp": true, "priority": true,
	}

	for _, clause := range clauseList {
		if sortMap[clause] {
			orderClauses = append(orderClauses, fmt.Sprintf("%s DESC", clause))
		}
	}

	return strings.Join(orderClauses, ", ")
}
