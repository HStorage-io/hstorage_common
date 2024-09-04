package hstorage_common

func ErrSubscriptionAlready() *ErrorMsg {
	return &ErrorMsg{
		Title: "既にプレミアム・ビジネスプランに加入しています",
		Msg:   "既にプレミアム・ビジネスプランに加入しています。再ログインを行い、最新の情報を反映させてください。",
		Err:   "subscription_already",
	}
}

func ErrSubscriptionTypeNotProvidedInQuery() *ErrorMsg {
	return &ErrorMsg{
		Title: "プランが指定されていません",
		Msg:   "?type= としてプランを指定してください。",
		Err:   "subscription_type_not_provided_in_query",
	}
}

func ErrSubscriptionEmailNotProvidedInQuery() *ErrorMsg {
	return &ErrorMsg{
		Title: "メールアドレスが指定されていません",
		Msg:   "?email= としてメールアドレスを指定してください。",
		Err:   "subscription_email_not_provided_in_query",
	}
}

func ErrSubscriptionNotFound() *ErrorMsg {
	return &ErrorMsg{
		Title: "サブスクリプションが見つかりません",
		Msg:   "サブスクリプションが見つかりませんでした。今も引き落としがありましたらお手数ですがお問い合わせください。",
		Err:   "subscription_not_found",
	}
}

func ErrUserDuplicated() *ErrorMsg {
	return &ErrorMsg{
		Title: "複数のアカウントを検知しました",
		Msg:   "フリープランでは複数アカウントで利用することを禁じております。プレミアム・ビジネスプランへのアップグレードをご検討いただけますと幸いです。",
		Err:   "user_duplicated",
	}
}

func ErrLimitExceededFiles() *ErrorMsg {
	return &ErrorMsg{
		Title: "月間アップロード数の制限に達しました",
		Msg:   "フリープランでは月間にアップロードできるファイル数に制限があります。プレミアム・ビジネスプランへのアップグレードをご検討いただけますと幸いです。",
		Err:   "limit_exceeded_files",
	}
}

func ErrLimitExceededCapacity() *ErrorMsg {
	return &ErrorMsg{
		Title: "容量の制限に達しました",
		Msg:   "フリープランでは保存できる容量に制限があります。プレミアム・ビジネスプランへのアップグレードをご検討いただけますと幸いです。",
		Err:   "limit_exceeded_capacity",
	}
}

func ErrDownloadNotificationNotEnabled() *ErrorMsg {
	return &ErrorMsg{
		Title: "ダウンロード通知が有効になっていません",
		Msg:   "ダウンロード通知が有効になっていません。",
		Err:   "download_notification_not_enabled",
	}
}
