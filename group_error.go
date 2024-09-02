package hstorage_common

func ErrGroupIDNotProvided() error {
	return &ErrorMsg{
		Title: "group_id が指定されていません",
		Msg:   "リクエスト内容が正しいか確認してください",
		Err:   "group_id not provided",
	}
}

func ErrGroupUIDNotProvided() error {
	return &ErrorMsg{
		Title: "group_uid が指定されていません",
		Msg:   "リクエスト内容が正しいか確認してください",
		Err:   "group_uid not provided",
	}
}

func ErrGroupTypeNotProvided() error {
	return &ErrorMsg{
		Title: "type が指定されていません",
		Msg:   "リクエスト内容が正しいか確認してください",
		Err:   "?type is not provided",
	}
}

func ErrGroupNotFound() error {
	return &ErrorMsg{
		Title: "グループが見つかりません",
		Msg:   "グループが見つかりませんでした。",
		Err:   "group not provided",
	}
}

func ErrGroupNotBelongToUser() error {
	return &ErrorMsg{
		Title: "グループが見つかりません",
		Msg:   "グループが見つかりませんでした。",
		Err:   "group not belong to user",
	}
}

func ErrDuplicatedGroupName() error {
	return &ErrorMsg{
		Title: "グループ名が重複しています",
		Msg:   "別のグループ名を指定してください。",
		Err:   "duplicated_group_name",
	}
}

func ErrGroupNotPublicView() error {
	return &ErrorMsg{
		Title: "グループが公開されていません",
		Msg:   "グループの作成者へグループの設定を変更してもらってください。",
		Err:   "group_not_public_view",
	}
}

func ErrGroupNotPublicUpload() error {
	return &ErrorMsg{
		Title: "グループへのアップロードが許可されていません",
		Msg:   "グループの作成者へグループの設定を変更してもらってください。",
		Err:   "group_not_public_upload",
	}
}
