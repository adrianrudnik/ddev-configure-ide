package datasources

import (
	"fmt"
	"github.com/adrianrudnik/ddev-configure-ide/internal/config"
	"github.com/adrianrudnik/ddev-configure-ide/internal/xmlknife"
	"github.com/beevik/etree"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/url"
)

func configureDataSourceRoot(conf *config.RuntimeConfig, d *etree.Document) *etree.Element {
	p := d.CreateElement("project")
	p.CreateAttr("version", "4")

	c := p.CreateElement("component")
	c.CreateAttr("name", "DataSourceManagerImpl")
	c.CreateAttr("format", "xml")
	c.CreateAttr("multifile-model", "true")

	return configureDataSourceNode(conf, c)
}

func configureDataSourceNode(conf *config.RuntimeConfig, e *etree.Element) *etree.Element {
	dsn := e.CreateElement("data-source")
	dsn.CreateAttr("source", "LOCAL")
	dsn.CreateAttr("name", "DDEV")
	dsn.CreateAttr("uuid", uuid.New().String())

	configureDataSourceNodeValues(conf, dsn)

	return dsn
}

func configureDataSourceNodeValues(conf *config.RuntimeConfig, e *etree.Element) {
	xmlknife.OverwriteXmlElement(e, "synchronize", "true")
	xmlknife.OverwriteXmlElement(e, "configured-by-url", "true")
	xmlknife.OverwriteXmlElement(e, "remarks", "DDEV generated data source by ddev-configure-ide")
	xmlknife.OverwriteXmlElement(e, "schema-control", "AUTOMATIC")
	xmlknife.OverwriteXmlElement(e, "working-dir", "$ProjectFileDir$")

	switch conf.DDEVConfig.Raw.Database.Type {
	case "postgres":
		jdbcUrl := url.URL{
			Scheme: "jdbc:postgresql",
			Host:   fmt.Sprintf("%s:%d", "127.0.0.1", conf.DDEVConfig.Raw.Database.MappedPort),
			Path:   conf.DDEVConfig.Raw.Database.Name,
		}

		jdbcQuery := jdbcUrl.Query()
		jdbcQuery.Set("user", conf.DDEVConfig.Raw.Database.Username)
		jdbcQuery.Set("password", conf.DDEVConfig.Raw.Database.Password)
		jdbcUrl.RawQuery = jdbcQuery.Encode()

		xmlknife.OverwriteXmlElement(e, "driver-ref", "postgresql")
		xmlknife.OverwriteXmlElement(e, "jdbc-driver", "org.postgresql.Driver")
		xmlknife.OverwriteXmlElement(e, "jdbc-url", jdbcUrl.String())

	case "mariadb":
		jdbcUrl := url.URL{
			Scheme: "jdbc:mariadb",
			Host:   fmt.Sprintf("%s:%d", "127.0.0.1", conf.DDEVConfig.Raw.Database.MappedPort),
			Path:   conf.DDEVConfig.Raw.Database.Name,
		}

		jdbcQuery := jdbcUrl.Query()
		jdbcQuery.Set("user", conf.DDEVConfig.Raw.Database.Username)
		jdbcQuery.Set("password", conf.DDEVConfig.Raw.Database.Password)
		jdbcUrl.RawQuery = jdbcQuery.Encode()

		xmlknife.OverwriteXmlElement(e, "driver-ref", "mariadb")
		xmlknife.OverwriteXmlElement(e, "jdbc-driver", "org.mariadb.jdbc.Driver")
		xmlknife.OverwriteXmlElement(e, "jdbc-url", jdbcUrl.String())

	case "mysql":
		jdbcUrl := url.URL{
			Scheme: "jdbc:mysql",
			Host:   fmt.Sprintf("%s:%d", "127.0.0.1", conf.DDEVConfig.Raw.Database.MappedPort),
			Path:   conf.DDEVConfig.Raw.Database.Name,
		}

		jdbcQuery := jdbcUrl.Query()
		jdbcQuery.Set("user", conf.DDEVConfig.Raw.Database.Username)
		jdbcQuery.Set("password", conf.DDEVConfig.Raw.Database.Password)
		jdbcUrl.RawQuery = jdbcQuery.Encode()

		xmlknife.OverwriteXmlElement(e, "driver-ref", "mysql.8")
		xmlknife.OverwriteXmlElement(e, "jdbc-driver", "com.mysql.cj.jdbc.Driver")
		xmlknife.OverwriteXmlElement(e, "jdbc-url", jdbcUrl.String())

	default:
		log.Fatal().Msg("Unsupported database type: " + conf.DDEVConfig.Raw.Database.Type)
	}
}
