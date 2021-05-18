package layout

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
	tele "gopkg.in/tucnak/telebot.v3"
)

type Settings struct {
	URL     string
	Token   string
	Updates int

	LocalesDir string `json:"locales_dir"`
	TokenEnv   string `json:"token_env"`
	ParseMode  string `json:"parse_mode"`

	Webhook    *tele.Webhook    `json:"webhook"`
	LongPoller *tele.LongPoller `json:"long_poller"`
}

func (lt *Layout) UnmarshalYAML(data []byte) error {
	var aux struct {
		Settings *Settings
		Config   map[string]interface{}
		Buttons  yaml.MapSlice
		Markups  yaml.MapSlice
		Locales  map[string]map[string]string
	}
	if err := yaml.Unmarshal(data, &aux); err != nil {
		return err
	}

	lt.Config = &Config{
		v: aux.Config,
	}

	if pref := aux.Settings; pref != nil {
		lt.pref = &tele.Settings{
			URL:       pref.URL,
			Token:     pref.Token,
			Updates:   pref.Updates,
			ParseMode: pref.ParseMode,
		}

		if pref.TokenEnv != "" {
			lt.pref.Token = os.Getenv(pref.TokenEnv)
		}

		if pref.Webhook != nil {
			lt.pref.Poller = pref.Webhook
		} else if pref.LongPoller != nil {
			lt.pref.Poller = pref.LongPoller
		}
	}

	lt.buttons = make(map[string]Button, len(aux.Buttons))
	for _, item := range aux.Buttons {
		k, v := item.Key.(string), item.Value

		// 1. Shortened reply button

		if v, ok := v.(string); ok {
			lt.buttons[k] = Button{Text: v}
			continue
		}

		// 2. Extended reply or inline button

		data, err := yaml.Marshal(v)
		if err != nil {
			return err
		}

		var btn Button
		if err := yaml.Unmarshal(data, &btn); err != nil {
			return err
		}

		lt.buttons[k] = btn
	}

	lt.markups = make(map[string]Markup, len(aux.Markups))
	for _, item := range aux.Markups {
		k, v := item.Key.(string), item.Value

		data, err := yaml.Marshal(v)
		if err != nil {
			return err
		}

		var shortenedMarkup [][]string
		if yaml.Unmarshal(data, &shortenedMarkup) == nil {
			// 1. Shortened reply or inline markup

			kb := make([][]Button, len(shortenedMarkup))
			for i, btns := range shortenedMarkup {
				row := make([]Button, len(btns))
				for j, btn := range btns {
					b, ok := lt.buttons[btn]
					if !ok {
						return fmt.Errorf("telebot/layout: no %s button for %s markup", btn, k)
					}
					row[j] = b
				}
				kb[i] = row
			}

			data, err := yaml.Marshal(kb)
			if err != nil {
				return err
			}

			tmpl, err := template.New(k).Funcs(lt.funcs).Parse(string(data))
			if err != nil {
				return err
			}

			markup := Markup{keyboard: tmpl}
			for _, row := range kb {
				for _, btn := range row {
					inline := btn.URL != "" ||
						btn.Unique != "" ||
						btn.InlineQuery != "" ||
						btn.InlineQueryChat != ""

					if markup.inline == nil {
						markup.inline = &inline
					} else if *markup.inline != inline {
						return fmt.Errorf("telebot/layout: mixed reply and inline buttons in %s markup", k)
					}
				}
			}

			lt.markups[k] = markup
		} else {
			// 2. Extended reply markup

			var markup struct {
				Markup   `yaml:",inline"`
				Keyboard [][]string `json:"keyboard"`
			}
			if err := yaml.Unmarshal(data, &markup); err != nil {
				return err
			}

			kb := make([][]tele.ReplyButton, len(markup.Keyboard))
			for i, btns := range markup.Keyboard {
				row := make([]tele.ReplyButton, len(btns))
				for j, btn := range btns {
					row[j] = *lt.buttons[btn].Reply()
				}
				kb[i] = row
			}

			data, err := yaml.Marshal(kb)
			if err != nil {
				return err
			}

			tmpl, err := template.New(k).Funcs(lt.funcs).Parse(string(data))
			if err != nil {
				return err
			}

			markup.inline = new(bool)
			markup.Markup.keyboard = tmpl
			lt.markups[k] = markup.Markup
		}
	}

	if aux.Locales == nil {
		if aux.Settings.LocalesDir == "" {
			aux.Settings.LocalesDir = "locales"
		}
		return lt.parseLocales(aux.Settings.LocalesDir)
	}

	return nil
}

func (lt *Layout) parseLocales(dir string) error {
	lt.locales = make(map[string]*template.Template)

	return filepath.Walk(dir, func(path string, fi os.FileInfo, _ error) error {
		if fi == nil || fi.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var texts map[string]string
		if err := yaml.Unmarshal(data, &texts); err != nil {
			return err
		}

		name := fi.Name()
		name = strings.TrimSuffix(name, filepath.Ext(name))

		tmpl := template.New(name)
		for key, text := range texts {
			text = strings.Trim(text, "\r\n")
			tmpl, err = tmpl.New(key).Funcs(lt.funcs).Parse(text)
			if err != nil {
				return err
			}
		}

		lt.locales[name] = tmpl
		return nil
	})
}
