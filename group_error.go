package hstorage_common

func ErrGroupIDNotProvided() *ErrorMsg {
	return &ErrorMsg{
		Title: "group_id が指定されていません",
		Msg:   "リクエスト内容が正しいか確認してください",
		Err:   "group_id not provided",
	}
}

func ErrGroupUIDNotProvided() *ErrorMsg {
	return &ErrorMsg{
		Title: "group_uid が指定されていません",
		Msg:   "リクエスト内容が正しいか確認してください",
		Err:   "group_uid not provided",
	}
}

func ErrGroupTypeNotProvided() *ErrorMsg {
	return &ErrorMsg{
		Title: "type が指定されていません",
		Msg:   "リクエスト内容が正しいか確認してください",
		Err:   "?type is not provided",
	}
}

func ErrGroupNotFound() *ErrorMsg {
	return &ErrorMsg{
		Title: "グループが見つかりません",
		Msg:   "グループが見つかりませんでした。",
		Err:   "group not found",
	}
}

func ErrGroupNotBelongToUser() *ErrorMsg {
	return &ErrorMsg{
		Title: "グループが不適切です",
		Msg:   "あなたのグループではありません。",
		Err:   "group not belong to user",
	}
}

func ErrDuplicatedGroupName() *ErrorMsg {
	return &ErrorMsg{
		Title: "グループ名が重複しています",
		Msg:   "別のグループ名を指定してください。",
		Err:   "duplicated_group_name",
	}
}

func ErrGroupNotPublicView() *ErrorMsg {
	return &ErrorMsg{
		Title: "グループが公開されていません",
		Msg:   "グループの作成者へグループの設定を変更してもらってください。",
		Err:   "group_not_public_view",
	}
}

func ErrGroupNotPublicUpload() *ErrorMsg {
	return &ErrorMsg{
		Title: "グループへのアップロードが許可されていません",
		Msg:   "グループの作成者へグループの設定を変更してもらってください。",
		Err:   "group_not_public_upload",
	}
}
