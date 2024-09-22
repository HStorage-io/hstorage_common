package hstorage_common

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type ErrorMsg struct {
	Title string `json:"title"` // api
	Msg   string `json:"msg"`   // api
	Err   string `json:"error"` // web で利用
}

func ErrCommon() *ErrorMsg {
	return &ErrorMsg{
		Title: "エラーが発生しました",
		Msg:   "しばらく時間をおいてから再度お試しください。",
		Err:   "unknown_error",
	}
}

func ErrHealth() *ErrorMsg {
	return &ErrorMsg{
		Title: "エラーが発生しました",
		Msg:   "サーバーが正常に動作していません。しばらく時間をおいてから再度お試しください。",
		Err:   "health_check_failed",
	}
}

func ErrInvalidRequest() *ErrorMsg {
	return &ErrorMsg{
		Title: "リクエストが無効です",
		Msg:   "リクエスト形式が正しいか確認してください",
		Err:   "invalid_request",
	}
}

func ErrInvalidToken() *ErrorMsg {
	return &ErrorMsg{
		Title: "トークンが無効です",
		Msg:   "トークンが無効です。",
		Err:   "invalid_token",
	}
}

func ErrInvalidWorkerKey() *ErrorMsg {
	return &ErrorMsg{
		Title: "キーが無効です",
		Msg:   "Worker に指定しているキーが無効です。",
		Err:   "invalid_worker_key",
	}
}

func ErrNotAllowed() *ErrorMsg {
	return &ErrorMsg{
		Title: "その操作は許可されておりません",
		Msg:   "ログインしているユーザーが正しいか確認してください。",
		Err:   "not_allowed",
	}
}

func ErrEmailFileInvalidEmailFormat() *ErrorMsg {
	return &ErrorMsg{
		Title: "メールアドレスの形式が正しくありません",
		Msg:   "メールアドレスの形式が正しくありません。",
		Err:   "invalid_email_format",
	}
}

func ErrStripeSignatureNotMatched() *ErrorMsg {
	return &ErrorMsg{
		Title: "Stripe のシグネチャーが一致しません",
		Msg:   "Stripe のシグネチャーが一致しません。",
		Err:   "stripe_signature_not_matched",
	}
}

func (e *ErrorMsg) Error() string {
	jsonMsg, err := json.Marshal(e)
	if err != nil {
		log.Printf("Error marshaling ErrorMsg: %v", err)
		return fmt.Sprintf("Internal error: %s", e.Title)
	}
	return string(jsonMsg)
}

func (e *ErrorMsg) Is(tgt error) bool {
	var target *ErrorMsg
	if !errors.As(tgt, &target) {
		return false
	}
	return e.Title == target.Title &&
		e.Msg == target.Msg &&
		e.Err == target.Err
}
