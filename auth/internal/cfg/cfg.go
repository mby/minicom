package cfg

type Config struct {
	JWTSecret       string
	MongoURI        string
	MongoDB         string
	UsersCollection string
}

var cfgs = map[string]Config{
	"qa": {
		JWTSecret:       "secret",
		MongoURI:        "mongodb://localhost:27017", // TODO: mongodb
		MongoDB:         "auth-qa",
		UsersCollection: "users",
	},
	"prod": {
		JWTSecret:       "secret",
		MongoURI:        "mongodb://localhost:27017", // TODO: mongodb
		MongoDB:         "auth",
		UsersCollection: "users",
	},
}

func GetConfig(env string) Config {
	cfg, ok := cfgs[env]

	if !ok {
		panic("ENV is invalid")
	}

	return cfg
}
