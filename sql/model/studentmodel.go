package model

import (
	"app/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ StudentModel = (*customStudentModel)(nil)

type (
	// StudentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentModel.
	StudentModel interface {
		studentModel

		FindAllBy(ctx context.Context, cond string, data *Student) ([]Student, error)
	}

	customStudentModel struct {
		*defaultStudentModel
	}
)

// NewStudentModel returns a model for the database table.
func NewStudentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StudentModel {
	return &customStudentModel{
		defaultStudentModel: newStudentModel(conn, c, opts...),
	}
}

func (m *customStudentModel) FindAllBy(ctx context.Context, cond string, data *Student) ([]Student, error) {
	var ret []Student
	var whereBuilder []string
	util.NotBlankThen(data.Name, func(name string) {
		whereBuilder = append(whereBuilder, fmt.Sprintf("`name` = '%s'", name))
	})
	//util.IfThen[int64](data.Age.Int64, util.NotZero, func(i int64) {
	//
	//})
	if data.Age.Int64 != 0 {
		whereBuilder = append(whereBuilder, fmt.Sprintf("`age` = %d", data.Age.Int64))
	}
	util.NotBlankThen(data.City.String, func(city string) {
		whereBuilder = append(whereBuilder, fmt.Sprintf("`city` = '%s'", city))
	})

	withCond := fmt.Sprintf(" %s ", strings.ToUpper(cond))
	var whereConditions = strings.Join(whereBuilder, withCond)
	if len(whereConditions) > 0 {
		whereConditions = " WHERE " + whereConditions
	}

	sql := fmt.Sprintf("select %s from %s %s", studentRows, m.table, whereConditions)
	logx.Infof("list all sql: " + sql)
	if err := m.QueryRowsNoCacheCtx(ctx, &ret, sql); err != nil {
		logx.Error(err)
		return nil, err
	}
	return ret, nil
}
