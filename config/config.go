package config

import (
	"fmt"
	"getcare-notification/utils"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var (
	cmd = &cobra.Command{
		Use:   "start",
		Short: "Start Applications",
	}
)

// Config of application
type Config struct {
	AppVersion string
	Tokens     Tokens
	Gin        Gin
	Http       Http
	Grpc       Grpc
	Logger     Logger
	Jaeger     Jaeger
	MongoDB    MongoDB
	MysqlDB    MysqlDB
	Kafka      Kafka
	Redis      Redis
	Firebase   Firebase
}

// Tokens config
type Tokens struct {
	Server []string
}

// Gin config
type Gin struct {
	Mode string
}

// Grpc config
type Grpc struct {
	Host              string
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

// Http config
type Http struct {
	Host              string
	Port              string
	PprofPort         string
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CookieLifeTime    int
	SessionCookieName string
}

// Logger config
type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Jaeger config
type Jaeger struct {
	Port        string
	Host        string
	ServiceName string
	LogSpans    bool
}

// MysqlDB config
type MysqlDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       string
	Mode     string
	LogLevel int
}

// MongoDB config
type MongoDB struct {
	URI string
	DB  string
	// Username string
	// Password string
}

// Kafka config
type Kafka struct {
	Brokers []string
}

// Redis config
type Redis struct {
	Host           string
	Port           string
	Password       string
	RedisDefaultDB string
	MinIdleConn    int
	PoolSize       int
	PoolTimeout    int
	DB             int
}

// Firebase config
type Firebase struct {
	KeyPath   string
	ProjectID string
}

// Execute executes the root command.
func Execute() error {
	return cmd.Execute()
}

// func init() {
// cobra.OnInitialize(InitConfig)
// cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
// cmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
// cmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
// cmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
// viper.SetDefault("license", "apache")
// rootCmd.AddCommand(addCmd)
// rootCmd.AddCommand(initCmd)
// }

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetConfigName("config.yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// ParseConfig Parse config file
func New() (*Config, error) {
	if err := InitConfig(); err != nil {
		return nil, err
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func GetFlags(itf interface{}, ctx *cli.Context, flagnames ...string) {
	if len(flagnames) == 0 {
		return
	}
	s := reflect.ValueOf(itf)
	for _, flagname := range flagnames {
		v := ctx.GlobalString(flagname)
		fmt.Println(v)
		if v == "" {
			continue
		}
		f := s.FieldByName(flagname)
		if !f.IsValid() && !f.CanSet() {
			continue
		}
		switch f.Kind() {
		case reflect.String:
			f.SetString(v)
		case reflect.Int:
			i64 := int64(utils.StringToInt(v))
			if !f.OverflowInt(i64) {
				f.SetInt(i64)
			}
		}
	}
}

func (m *MysqlDB) Address() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DB)
}

func (r *Redis) Address() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

func (h *Http) Address() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}

func (g *Grpc) Address() string {
	return fmt.Sprintf("%s:%s", g.Host, g.Port)
}

func (j *Jaeger) Address() string {
	return fmt.Sprintf("%s:%s", j.Host, j.Port)
}
