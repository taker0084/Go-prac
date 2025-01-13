package config

import (
	"html/template"
	"log"
)

//さまざまなファイルでデータの受け渡しを行いたい(データベースのコネクトなど)場合、
//下のようにAppConfigの中に書くのが良い
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
}