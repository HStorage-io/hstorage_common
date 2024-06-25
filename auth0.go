package hstorage_common

type Auth0Roles struct {
	Roles []string `json:"roles"`
}

type Auth0FreeTrialCount struct {
	Premium  int `json:"premium"`
	Business int `json:"business"`
	API      int `json:"api"`
}

type Auth0AppMetadata struct {
	Key            string              `json:"key" mapstructure:"key"`
	MD5            string              `json:"md5" mapstructure:"md5"`
	Roles          []string            `json:"roles" mapstructure:"roles"`
	FreeTrialCount Auth0FreeTrialCount `json:"free_trial_count" mapstructure:"free_trial_count"`
}

type Auth0FreeTrial struct {
	FreeTrialCount Auth0FreeTrialCount `json:"free_trial_count"`
}

type Auth0UserInfo struct {
	Sub               string `json:"sub"`
	Name              string `json:"name"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	MiddleName        string `json:"middle_name"`
	Nickname          string `json:"nickname"`
	PreferredUsername string `json:"preferred_username"`
	Profile           string `json:"profile"`
	Picture           string `json:"picture"`
	Website           string `json:"website"`
	Email             string `json:"email"`
	// EmailVerified       bool   `json:"email_verified"` // apple provider returns as string
	Gender              string `json:"gender"`
	Birthdate           string `json:"birthdate"`
	Zoneinfo            string `json:"zoneinfo"`
	Locale              string `json:"locale"`
	PhoneNumber         string `json:"phone_number"`
	PhoneNumberVerified bool   `json:"phone_number_verified"`
	Address             struct {
		Country string `json:"country"`
	} `json:"address"`
	UpdatedAt string `json:"updated_at"`
}
