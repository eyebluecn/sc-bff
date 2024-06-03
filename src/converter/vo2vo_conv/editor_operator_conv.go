package vo2vo_conv

import (
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
)

func ConvertEditorOperator(readerVO *vo_model.EditorVO) *sc_misc_base.Operator {
	if readerVO == nil {
		return nil

	}
	return &sc_misc_base.Operator{
		OperatorId: readerVO.ID,
		Type:       sc_misc_base.OperatorType_EDITOR,
		Username:   readerVO.Username,
	}
}
