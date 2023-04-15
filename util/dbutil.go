package util

import "database/sql"

func ToNullString(str string) sql.NullString {
	return sql.NullString{
		String: str, Valid: true,
	}
}

func ToNullInt(numInt int) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64(numInt), Valid: true,
	}
}
