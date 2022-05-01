package codegen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/pflag"

	"github.com/robertsong9972/gorm-gen/internal/gormfilter"
	"github.com/robertsong9972/gorm-gen/util"
)

type config struct {
	filepath   string
	genType    string
	modulePath string
}

const (
	TypeFilter = "filter"
)

func Gen() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic!", err)
		}
	}()

	// parse flags
	cfg := &config{}
	parseFlag(cfg)

	util.Assert(strings.Trim(cfg.filepath, "") != "", "--filepath cannot be empty")
	util.Assert(strings.Trim(cfg.genType, "") != "", "--gen_type cannot be empty")
	util.Assert(strings.Trim(cfg.modulePath, "") != "", "--module_path cannot be empty")

	switch cfg.genType {
	case TypeFilter:
		genFilter(cfg)
	}
}

func genFilter(cfg *config) {
	models := gormfilter.Parse(gormfilter.ParseOptions{
		FilePath: cfg.filepath,
	})
	parentDir := filepath.Dir(cfg.filepath)

	genDir := parentDir + "/gen"
	if err := os.Mkdir(genDir, os.ModeDir|os.ModePerm); !os.IsExist(err) && err != nil {
		log.Panicf("failed to create directory|dir=%s|err=%v", genDir, err)
	}

	for _, m := range models {
		if IgnoreMessageName(m.Name) {
			continue
		}
		m.DbPackage = cfg.modulePath + "/" + m.DbPackage
		gormfilter.Generate(m, gormfilter.GenOptions{
			FilePath: genDir + "/" + strcase.ToSnake(m.Name) + "_filter.gen.go",
		})
	}
}

func parseFlag(cfg *config) {
	pflag.StringVar(&cfg.genType, "gen_type", "", "generation type")
	pflag.StringVar(&cfg.filepath, "filepath", "", "the file path to read the proto file")
	pflag.StringVar(&cfg.modulePath, "module_path", "", "the path to pb gen file")
	pflag.Parse()
}
