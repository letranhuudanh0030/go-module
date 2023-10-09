package util

import (
	"reflect"
	"time"

	"gorm.io/gorm"
)

func ConvertTimeFieldsToTimeZone(s interface{}, timezone *time.Location) {
	val := reflect.ValueOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := field.Type()

		if fieldType == reflect.TypeOf(time.Time{}) || fieldType == reflect.TypeOf(&time.Time{}) {
			// Convert time.Time and *time.Time fields
			if field.CanSet() {
				if field.Kind() == reflect.Ptr && !field.IsNil() {
					convertedTime := field.Interface().(*time.Time).In(timezone)
					field.Set(reflect.ValueOf(&convertedTime))
				} else {
					convertedTime := field.Interface().(time.Time).In(timezone)
					field.Set(reflect.ValueOf(convertedTime))
				}
			}
		} else if fieldType == reflect.TypeOf(gorm.DeletedAt{}) {
			// Convert gorm.DeletedAt fields
			if field.CanSet() {
				deletedAt := field.Interface().(gorm.DeletedAt)
				if deletedAt.Valid {
					convertedTime := deletedAt.Time.In(timezone)
					newDeletedAt := gorm.DeletedAt{Time: convertedTime, Valid: deletedAt.Valid}
					field.Set(reflect.ValueOf(newDeletedAt))
				}
			}
		}
	}
}

func ConvertToDate(t int64) string {
	return time.UnixMilli(t).Format("2006-01-02")
}
