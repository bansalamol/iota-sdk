package modules

import (
	"encoding/json"
	"github.com/iota-agency/iota-erp/internal/configuration"
	"github.com/iota-agency/iota-erp/internal/modules/elxolding"
	"github.com/iota-agency/iota-erp/internal/modules/iota"
	"github.com/iota-agency/iota-erp/internal/modules/shared"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"slices"
)

var (
	AllModules = []shared.Module{
		iota.NewModule(),
		elxolding.NewModule(),
	}
	LoadedModules = Load()
)

func Load() []shared.Module {
	jsonConf := configuration.UseJsonConfig()
	modules := make([]shared.Module, 0, len(AllModules))
	for _, module := range AllModules {
		if slices.Contains(jsonConf.Modules, module.Name()) {
			modules = append(modules, module)
		}
	}
	return modules
}

func LoadBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("pkg/locales/en.json")
	bundle.MustLoadMessageFile("pkg/locales/ru.json")
	for _, module := range LoadedModules {
		for _, localeFile := range module.LocaleFiles() {
			bundle.MustLoadMessageFile(localeFile)
		}
	}
	return bundle
}
