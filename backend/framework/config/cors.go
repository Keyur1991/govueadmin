package config

type Cors struct {
	AllowedOrigins   []string
	AllowCredentials bool
	Debug            bool
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowedMethods   []string
}
