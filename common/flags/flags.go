package flags

import "github.com/urfave/cli"

var (

	// ServerNameFlag ...
	ServerNameFlag = cli.StringFlag{
		Name:   "server_name",
		Usage:  "server name",
		EnvVar: "SERVER_NAME",
		Value:  "Getcare Notication",
	}

	// ServerVersionFlag ...
	ServerVersionFlag = cli.StringFlag{
		Name:   "server_version",
		Usage:  "server sersion",
		EnvVar: "SERVER_VERSION",
		Value:  "Server Version",
	}

	// HttpHostFlag ...
	HttpHostFlag = cli.StringFlag{
		Name:   "http_host",
		Usage:  "Http Host",
		EnvVar: "HTTP_HOST",
		Value:  "",
	}

	// HttpPortFlag ...
	HttpPortFlag = cli.StringFlag{
		Name:   "http_port",
		Usage:  "Http Port",
		EnvVar: "HTTP_PORT",
		Value:  "8080",
	}

	// MongoDatabaseNameFlag ...
	MongoDatabaseNameFlag = cli.StringFlag{
		Name:   "mongodb_name",
		Usage:  "Mongodb name",
		EnvVar: "MONGO_NAME",
		Value:  "notification",
	}

	// MongoURIFlag ...
	MongoURIFlag = cli.StringFlag{
		Name:   "mongodb_uri",
		Usage:  "Mongodb uri",
		EnvVar: "MONGODB_URI",
		Value:  "mongodb://localhost:27017",
	}

	// StorageAccessKeyFlag ...
	StorageAccessKeyFlag = cli.StringFlag{
		Name:   "storage_access_key",
		Usage:  "Storage access key",
		EnvVar: "STORAGE_ACCESS_KEY",
		Value:  "",
	}

	// StorageSecretKeyFlag ...
	StorageSecretKeyFlag = cli.StringFlag{
		Name:   "storage_secret_key",
		Usage:  "Storage secret key",
		EnvVar: "STORAGE_SECRET_KEY",
		Value:  "",
	}

	// StorageRegionFlag ...
	StorageRegionFlag = cli.StringFlag{
		Name:   "storage_region",
		Usage:  "Storage region",
		EnvVar: "STORAGE_REGION",
		Value:  "",
	}

	// StorageName ...
	StorageNameFlag = cli.StringFlag{
		Name:   "storage_name",
		Usage:  "Storage name",
		EnvVar: "STORAGE_NAME",
		Value:  "room-image",
	}

	// JaegerHostFlag ...
	JaegerHostFlag = cli.StringFlag{
		Name:   "Jaeger_Host",
		Usage:  "Jaeger Host",
		EnvVar: "JAEGER_HOST",
		Value:  "tracer",
	}

	// JaegerPortFlag ...
	JaegerPortFlag = cli.StringFlag{
		Name:   "Jaeger_Port",
		Usage:  "Jaeger Port",
		EnvVar: "JAEGER_PORT",
		Value:  "6831",
	}
)

func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			ctx.GlobalSet(name, ctx.String(name))
		}
		return action(ctx)
	}
}
