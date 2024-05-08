package datasources

import (
	"errors"
	"github.com/adrianrudnik/ddev-configure-ide/internal/config"
	"github.com/beevik/etree"
	"github.com/rs/zerolog/log"
	"os"
	"path"
	"path/filepath"
)

const configPath = ".idea/dataSources.xml"

func MustHave(conf *config.RuntimeConfig) {
	configFilePath := getAbsLocalConfigPath(conf)

	// Check if dataSources.local.xml exists
	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		log.Info().Str("path", configFilePath).Msg("Configuration file could not be found")
		rebuild(configFilePath, conf)
	} else {
		log.Info().Str("path", configFilePath).Msg("Configuration file found")
		reconfigure(configFilePath, conf)
	}
}

func getAbsLocalConfigPath(conf *config.RuntimeConfig) string {
	return path.Join(conf.WorkingDirectory, configPath)
}

func reconfigure(path string, conf *config.RuntimeConfig) {
	// Load the existing document into a simplified lib to avoid complicated struct juggling
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		log.Fatal().Err(err).Msg("Unable to read local configuration XML")
	}

	if len(doc.FindElements("/project/component/data-source[@name='DDEV']")) > 0 {
		// Update existing elements
		for _, ds := range doc.FindElements("/project/component/data-source[@name='DDEV']") {
			log.Info().Msg("Found existing DDEV data source, updating the values")

			configureDataSourceNodeValues(conf, ds)
		}

		saveDocument(conf, doc, path)
	} else {
		// Create a new DDEV element
		root := doc.FindElement("/project/component")
		if root == nil {
			log.Fatal().Msg("Weird component configuration")
		}

		configureDataSourceNode(conf, root)
		saveDocument(conf, doc, path)
	}
}

func rebuild(path string, conf *config.RuntimeConfig) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	configureDataSourceRoot(conf, doc)
	saveDocument(conf, doc, path)
}

func saveDocument(conf *config.RuntimeConfig, doc *etree.Document, path string) {
	doc.Indent(2)

	if conf.DryRun {
		log.Info().Bool("dry-run", true).Msg("Writing updated config console")

		if _, err := doc.WriteTo(os.Stdout); err != nil {
			log.Fatal().Err(err).Msg("Unable to write updated config to os.Stdout")
		}
	} else {
		// Ensure the .idea folder exists
		folder := filepath.Dir(path)
		if _, err := os.Stat(folder); errors.Is(err, os.ErrNotExist) {
			err := os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				log.Fatal().Err(err).Str("folder", folder).Msg("Unable to create folder")
			}

			log.Info().Str("folder", folder).Msg("Created IDE base folder")
		}

		err := doc.WriteToFile(path)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to write updated config to file")
		}

		log.Info().Str("path", path).Msg("Wrote updated config to file")
	}
}
