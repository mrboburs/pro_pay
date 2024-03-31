package function

import (
	"pro_pay/model"
	"pro_pay/tools/logger"
	"database/sql"
	"errors"
	"math"

	"github.com/jmoiron/sqlx"
)

func GetListCount(db *sqlx.DB, loggers *logger.Logger, pagination *model.Pagination, countFilterQuery string, filtersArray []interface{}) (err error) {
	var count int64
	err = db.Get(&count, countFilterQuery, filtersArray...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	if count < pagination.Offset+pagination.Limit {
		pageCount := int64(math.Ceil(float64(count) / float64(pagination.PageSize)))
		var offset int64
		if pageCount > 1 {
			offset = (pageCount - 1) * pagination.PageSize
		}
		pagination.Limit = count - offset
		pagination.Offset = offset
	}
	pagination.ItemTotal = count
	pagination.PageTotal = int64(math.Ceil(float64(count) / float64(pagination.PageSize)))
	if pagination.Page > pagination.PageTotal {
		if pagination.PageTotal == 0 {
			pagination.Page = 1
		}
		if pagination.PageTotal != 0 {
			pagination.Page = pagination.PageTotal
		}
	}
	return nil
}
