package config

type GrpcServer struct {
	Host string `yaml:"host" env-required:"true" env-default:""`
	Port int    `yaml:"port" env-required:"true" env-default:"8090"`
}

type HttpServer struct {
	Host string `yaml:"host" env-required:"true" env-default:""`
	Port int    `yaml:"port" env-required:"true" env-default:"8080"`
}

type Logger struct {
	Mode              string `yaml:"mode" env-required:"true" env-default:"local"`
	HttpServerLogFile string `yaml:"httpServerLogFile" env-required:"true"`
	GrpcServerLogFile string `yaml:"grpcServerLogFile" env-required:"true"`
}

type Postgre struct {
	Host     string `yaml:"host" env-required:"true" env-default:"localhost"`
	Port     int    `yaml:"port" env-required:"true" env-default:"5432"`
	User     string `yaml:"user" env-required:"true" env-default:"postgres"`
	Password string `yaml:"password" env-required:"true" env-default:"1234567"`
	Name     string `yaml:"name" env-required:"true" env-default:"postgres"`
	SSLmode  bool   `yaml:"sslmode" env-required:"true" env-default:"false"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Password struct {
	Salt       string `yaml:"salt" env-required:"true" evv-default:"KJSADLHFAKSDJFHh8hwdf980hwa9e8fh89H9HkjH8ouH9H9P8h98HH98h89gqfh98h98H"`
	SecondSalt string `yaml:"secondSalt" env-required:"true" env-default:"JKDGSHF8989uJHB987sdfoljk908098sdfhkj90hjuh98HKjh98l"`
}

type Tokens struct {
	AccessSecret      string `yaml:"accessSecret" env-required:"true"`
	RefreshSize       int    `yaml:"refreshSize" env-required:"true"`
	RefreshSalt       string `yaml:"refreshSalt" env-required:"true"`
	RefreshSecondSalt string `yaml:"refreshSecondSalt" env-required:"true"`
}

type Minio struct {
	Endpoint        string `yaml:"endpoint" env-required:"true" env-default:"localhost:9000"`
	AccessKeyID     string `yaml:"accessKeyID" env-required:"true"`
	SecretAccessKey string `yaml:"secretAccessKey" env-required:"true"`
	UseSSL          bool   `yaml:"useSSL" env-required:"true"`
}

type config struct {
	HttpServer *HttpServer `yaml:"httpServer"`
	Logger     *Logger     `yaml:"logger"`
	GrpcServer *GrpcServer `yaml:"grpcServer"`
	Postgre    *Postgre    `yaml:"postgre"`
	Password   *Password   `yaml:"password"`
	Tokens     *Tokens     `yaml:"tokens"`
	Redis      *Redis      `yaml:"redis"`
	Minio      *Minio      `yaml:"minio"`
}
