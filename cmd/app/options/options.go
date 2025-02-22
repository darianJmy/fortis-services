package options

import (
	"context"
	"fmt"
	"github.com/darianJmy/fortis-services/cmd/app/config"
	"github.com/darianJmy/fortis-services/pkg/controller"
	"github.com/darianJmy/fortis-services/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	defaultConfigFile = "/etc/fortis/config.yaml"
)

type ServerRunOptions struct {
	Config config.Config

	ConfigFile string

	HttpEngine *gin.Engine

	Control controller.FortisInterface
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		ConfigFile: defaultConfigFile,
		HttpEngine: gin.Default(),
	}
}

func (s *ServerRunOptions) Complete() error {
	var componentConfig config.Config

	data, err := ioutil.ReadFile(s.ConfigFile)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(data, &componentConfig); err != nil {
		return err
	}

	s.Config = componentConfig

	return nil
}

func (s *ServerRunOptions) Registry() error {

	mongoCN, err := s.registerMongo()
	if err != nil {
		return err
	}

	factory, err := db.NewDaoFactory(mongoCN, s.Config.Default.AutoMigrate)
	if err != nil {
		return err
	}

	s.Control = controller.New(s.Config, factory)

	return nil
}

func (s *ServerRunOptions) registerMongo() (*mongo.Database, error) {
	mongoConfig := s.Config.Mongo
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d/?replicaSet=rs0&authSource=cmdb",
		mongoConfig.User,
		mongoConfig.Password,
		mongoConfig.Host,
		mongoConfig.Port)

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return client.Database(mongoConfig.Name), nil
}

func (s *ServerRunOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&s.ConfigFile, "config", s.ConfigFile, "The location of the fortis configuration file")
}
