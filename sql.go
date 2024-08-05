package option

import (
	"database/sql"
	"database/sql/driver"
)

var (
	_ sql.Scanner   = &Option[any]{}
	_ driver.Valuer = Option[any]{}
)

func (opt *Option[T]) Scan(src any) error {
	if src == nil {
		*opt = None[T]()
		return nil
	}
	var dest T
	if err := sqlConvertAssign(&dest, src); err != nil {
		return err
	}
	*opt = Some(dest)
	return nil
}

func (opt Option[T]) Value() (driver.Value, error) {
	if opt.IsNone() {
		return nil, nil //nolint:nilnil
	}
	return driver.DefaultParameterConverter.ConvertValue(opt.Unwrap())
}
