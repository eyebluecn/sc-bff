package session

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/eyebluecn/sc-bff/src/common/constant"
	"github.com/eyebluecn/sc-bff/src/common/enums"
	"github.com/eyebluecn/sc-bff/src/common/errs"
	"github.com/eyebluecn/sc-bff/src/model/vo_model"
	"sync"
)

type Session struct {
	loginReaders map[string]*vo_model.ReaderVO
	loginEditors map[string]*vo_model.EditorVO
}

var instance *Session
var instanceOnce sync.Once

// 使用session应从这里入口
func DefaultSession() *Session {
	if instance == nil {
		instanceOnce.Do(func() {
			instance = &Session{
				loginReaders: make(map[string]*vo_model.ReaderVO),
				loginEditors: make(map[string]*vo_model.EditorVO),
			}
		})
	}
	return instance
}

// 存储登录的读者
func (receiver Session) StoreLoginReader(key string, reader *vo_model.ReaderVO) {
	receiver.loginReaders[key] = reader
}

// 存储登录的编辑
func (receiver Session) StoreLoginEditor(key string, editor *vo_model.EditorVO) {
	receiver.loginEditors[key] = editor
}

// 删除登录的读者
func (receiver Session) DelLoginReader(key string) {
	delete(receiver.loginReaders, key)
}

// 删除登录的编辑
func (receiver Session) DelLoginEditor(key string) {
	delete(receiver.loginEditors, key)
}

// 查出当前登录的读者 找不到返回nil
func (receiver Session) FindLoginReader(appCtx *app.RequestContext) *vo_model.ReaderVO {
	bytes := appCtx.Cookie(constant.READER_COOKIE_KEY)
	str := string(bytes)
	if readerVo, ok := receiver.loginReaders[str]; ok {
		return readerVo
	}
	return nil
}

// 查出当前登录的编辑 找不到返回nil
func (receiver Session) FindLoginEditor(appCtx *app.RequestContext) *vo_model.EditorVO {
	bytes := appCtx.Cookie(constant.EDITOR_COOKIE_KEY)
	str := string(bytes)
	if editorVo, ok := receiver.loginEditors[str]; ok {
		return editorVo
	}
	return nil
}

// 查出当前登录的读者 找不到返回错误
func (receiver Session) CheckLoginReader(appCtx *app.RequestContext) (*vo_model.ReaderVO, error) {

	reader := receiver.FindLoginReader(appCtx)
	if reader == nil {
		return nil, errs.CodeErrorf(enums.StatusCodeLogin, "未登录，禁止访问！")
	}
	return reader, nil
}

// 查出当前登录的编辑 找不到返回错误
func (receiver Session) CheckLoginEditor(appCtx *app.RequestContext) (*vo_model.EditorVO, error) {

	editor := receiver.FindLoginEditor(appCtx)
	if editor == nil {
		return nil, errs.CodeErrorf(enums.StatusCodeLogin, "未登录，禁止访问！")
	}
	return editor, nil
}
