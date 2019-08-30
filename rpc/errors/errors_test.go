package errors_test

import (
	"github.com/jinmukeji/plat-pkg/rpc/errors"

	goerr "errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ErrorsTestSuite struct {
	suite.Suite
}

func (suite *ErrorsTestSuite) TestExtractErrCode() {
	// 用例被测目标函数输入参数
	type args struct {
		err error
	}

	// 用例清单
	tests := []struct {
		name     string // 用例名称
		args     args   // 输入参数
		wantCode int
		wantOk   bool
	}{
		{
			name:     "correct code [12345]",
			args:     args{err: goerr.New("[errcode:12345] this is 12345")},
			wantCode: 12345,
			wantOk:   true,
		},
		{
			name:     "correct code [0]",
			args:     args{err: goerr.New("[errcode:0] this is 0")},
			wantCode: 0,
			wantOk:   true,
		},
		{
			name:     "no code",
			args:     args{err: goerr.New("[errcode:] this is no code")},
			wantCode: 0,
			wantOk:   false,
		},
		{
			name:     "wrong format",
			args:     args{err: goerr.New("[errcode:123abc] this is wrong format")},
			wantCode: 0,
			wantOk:   false,
		},
		{
			name:     "wrong format",
			args:     args{err: goerr.New("[errcode:abc123] this is wrong format")},
			wantCode: 0,
			wantOk:   false,
		},
	}

	for _, tt := range tests {

		suite.Run(tt.name, func() {
			code, ok := errors.ExtractErrCode((tt.args.err))
			suite.Assert().Equal(tt.wantCode, code)
			suite.Assert().Equal(tt.wantOk, ok)
		})
	}
}

func TestErrorsTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorsTestSuite))
}
