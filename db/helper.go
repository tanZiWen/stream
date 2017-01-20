package db

import "strconv"

func InOrEqual(column string, ids []int64) string {
	if len(ids) == 0 {
		return ""
	}

	if len(ids) == 1 {
		return column + " = " + strconv.FormatInt(ids[0], 10)
	}

	var sql string = column + " in ("

	for index, id := range ids {
		if index > 0 {
			sql = sql + "," + strconv.FormatInt(id, 10)
		} else {
			sql = sql + strconv.FormatInt(id, 10)
		}
	}

	sql = sql + ")"

	return sql
}