package config

type Config struct {
	DBURL string `json:db_url`
	CurrentUserName `json:current_user_name`
}
