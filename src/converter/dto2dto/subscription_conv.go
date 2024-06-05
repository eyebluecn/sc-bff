package dto2dto

import (
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertSubscriptionStatus(status *int32) *sc_subscription_api.SubscriptionStatus {
	if status == nil {
		return nil
	}
	res := sc_subscription_api.SubscriptionStatus(*status)

	return &res
}
