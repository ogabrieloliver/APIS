package configs

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	webServerPort string
	JWTSecret     string
	JwtExperesIN  int
}

func LoadConfig(path string) (*conf, error) {
	//...
}
