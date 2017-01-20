package setting

import (
	"github.com/go-ini/ini"
	log "github.com/Sirupsen/logrus"
)

var (
	configLoaded bool
	Config *ini.File
	App *AppConfig
	Page *PageConfig

	files []interface{}
	mappings []SectionMap
	Initialized bool
)

type SectionMap struct {
	SectionName string
	MapTo       interface{}
}

type AppConfig struct {
	Name        string `ini:"NAME"`
	Version     string `ini:"VERSION"`
	ListenPort  string `ini:"LISTEN_PORT"`
	AllowDomain string `ini:"ALLOW_DOMAIN"`
	AllowFrom   string `ini:"ALLOW_FROM"`
}

type PageConfig struct {
	Size int `ini:"PAGE_SIZE"`
}

func Initialize() error {
	Config, err := ini.LooseLoad("conf/app.ini", files...); if err != nil {
		log.Error("failed to load config files", err)
		return err
	}

	for _, sectionItem := range mappings {
		err = Config.Section(sectionItem.SectionName).MapTo(sectionItem.MapTo); if err != nil {
			log.Error("failed to map section: ", sectionItem.SectionName)
			return err
		}
	}

	Initialized = true
	return nil
}

func AddFile(file string) {
	files = append(files, file)
}

func AddMapping(mapping SectionMap) {
	mappings = append(mappings, mapping)
}

func init() {
	files = []interface{}{}
	Page = &PageConfig{Size: 20}
	App = &AppConfig{}

	mappings = []SectionMap{SectionMap{SectionName: "app", MapTo: App}, SectionMap{SectionName: "page", MapTo: Page}}
}
