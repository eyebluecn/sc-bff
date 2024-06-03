include "column/column_api.thrift"
include "reader/reader_api.thrift"
include "editor/editor_api.thrift"
include "payment/payment_api.thrift"
include "subscription/subscription_api.thrift"
include "rich/rich_api.thrift"
namespace go sc_bff_api

//定义bff应用中的所有HTTP服务
service BffService {


    //**************** 专栏START ****************//
    //演示接口 创建专栏等相关所有实体
    column_api.ColumnOmnibusResponse ColumnOmnibus(1: column_api.ColumnOmnibusRequest request)
    (api.get="/api/column/omnibus")

    //查看专栏列表
    rich_api.ColumnRichPageResponse ColumnPage(1: rich_api.ColumnRichPageRequest request)
    (api.get="/api/column/rich/page")

    //查看专栏详情
    column_api.ColumnDetailResponse ColumnDetail(1: column_api.ColumnDetailRequest request)
    (api.get="/api/column/detail")


    //**************** 专栏END ****************//



    //**************** 小编START ****************//

    //小编获取自己信息
    editor_api.EditorTinyLoginResponse EditorTinyLogin(1: editor_api.EditorTinyLoginRequest request)
    (api.get="/api/editor/tiny/login")

    //小编登录
    editor_api.EditorLoginResponse EditorLogin(1: editor_api.EditorLoginRequest request)
    (api.get="/api/editor/login")

    //小编退出登录
    editor_api.EditorLogoutResponse EditorLogout(1: editor_api.EditorLogoutRequest request)
    (api.get="/api/editor/logout")


    //**************** 小编END ****************//



    //**************** 读者START ****************//

    //读者获取自己信息
    reader_api.ReaderTinyLoginResponse ReaderTinyLogin(1: reader_api.ReaderTinyLoginRequest request)
    (api.get="/api/reader/tiny/login")

    //读者登录
    reader_api.ReaderLoginResponse ReaderLogin(1: reader_api.ReaderLoginRequest request)
    (api.get="/api/reader/login")

    //读者退出登录
    reader_api.ReaderLogoutResponse ReaderLogout(1: reader_api.ReaderLogoutRequest request)
    (api.get="/api/reader/logout")


    //**************** 读者END ****************//


    //**************** 支付单START ****************//

    //第三方支付平台，支付成功后的回调接口。
    payment_api.PaymentPaidCallbackResponse PaymentPaidCallback(1: payment_api.PaymentPaidCallbackRequest request)
    (api.get="/api/payment/paid/callback")


    //**************** 支付单END ****************//


    //**************** 订阅START ****************//

    //下单
    subscription_api.SubscriptionPrepareResponse SubscriptionPrepare(1: subscription_api.SubscriptionPrepareRequest request)
    (api.get="/api/subscription/prepare")


    //查看订阅列表
    subscription_api.SubscriptionPageResponse SubscriptionPage(1: subscription_api.SubscriptionPageRequest request)
    (api.get="/api/subscription/page")


    //查看订阅富列表
    rich_api.SubscriptionRichPageResponse SubscriptionRichPage(1: rich_api.SubscriptionRichPageRequest request)
    (api.get="/api/subscription/rich/page")


    //**************** 订阅END ****************//


}
