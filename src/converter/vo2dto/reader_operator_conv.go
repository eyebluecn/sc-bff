package vo2dto

import (
	"github.com/eyebluecn/sc-bff/src/model/vo"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
)

func ConvertReaderOperator(readerVO *vo.ReaderVO) *sc_misc_base.Operator {
	if readerVO == nil {
		return nil

	}
	return &sc_misc_base.Operator{
		OperatorId: readerVO.ID,
		Type:       sc_misc_base.OperatorType_READER,
		Username:   readerVO.Username,
	}
}
