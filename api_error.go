package hstorage_common

func ErrCantAuthenticate() *ErrorMsg {
	return &ErrorMsg{
		Title: "認証に失敗しました",
		Msg:   "認証に失敗しました",
		Err:   "cant_authenticate",
	}
}

func ErrInvalidAPIKey() *ErrorMsg {
	return &ErrorMsg{
		Title: "APIキーが無効です",
		Msg:   "APIキーが無効です",
		Err:   "invalid_api_key",
	}
}
