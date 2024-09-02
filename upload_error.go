package hstorage_common

func ErrFileNameNotProvided() error {
	return &ErrorMsg{
		Title: "file_name が指定されていません",
		Msg:   "リクエスト内容が正しいか確認してください",
		Err:   "file_name not provided",
	}
}

func ErrFileNotUploaded() error {
	return &ErrorMsg{
		Title: "ファイルがアップロードされていません",
		Msg:   "ファイルがアップロードされていません",
		Err:   "file_not_uploaded",
	}
}

func ErrFileInfoPasswordNeed() error {
	return &ErrorMsg{
		Title: "パスワードが必要です",
		Msg:   "パスワードが必要です",
		Err:   "need_password",
	}
}

func ErrFileInfoPasswordInvalid() error {
	return &ErrorMsg{
		Title: "パスワードが無効です",
		Msg:   "パスワードが無効です",
		Err:   "invalid_password",
	}
}

func ErrFileInfoHashedKeyNeed() error {
	return &ErrorMsg{
		Title: "ハッシュ化されたキーが必要です",
		Msg:   "ハッシュ化されたキーが必要です",
		Err:   "need_hashed_key",
	}
}

func ErrFileInfoHashedKeyInvalid() error {
	return &ErrorMsg{
		Title: "ハッシュ化されたキーが無効です",
		Msg:   "ハッシュ化されたキーが無効です",
		Err:   "invalid_hashed_key",
	}
}

func ErrFileNotFound() error {
	return &ErrorMsg{
		Title: "ファイルが見つかりません",
		Msg:   "ファイルが見つかりませんでした。",
		Err:   "not_found",
	}
}
