package config

import (
	"github.com/sqshq/sampler/console"
)

const (
	defaultRefreshRateMs = 1000
	defaultScale         = 1
	defaultTheme         = console.ThemeDark
)

func (c *Config) setDefaults() {
	c.setDefaultValues()
	c.setDefaultColors()
	c.setDefaultLayout()
}

func (c *Config) setDefaultValues() {

	if c.Theme == nil {
		t := defaultTheme
		c.Theme = &t
	}

	for i, chart := range c.RunCharts {

		setDefaultTriggersValues(chart.Triggers)

		if chart.RefreshRateMs == nil {
			r := defaultRefreshRateMs
			chart.RefreshRateMs = &r
		}
		if chart.Scale == nil {
			p := defaultScale
			chart.Scale = &p
		}
		if chart.Legend == nil {
			chart.Legend = &LegendConfig{true, true}
			c.RunCharts[i] = chart
		}
		c.RunCharts[i] = chart
	}

	for i, chart := range c.BarCharts {

		setDefaultTriggersValues(chart.Triggers)

		if chart.RefreshRateMs == nil {
			r := defaultRefreshRateMs
			chart.RefreshRateMs = &r
		}
		if chart.Scale == nil {
			p := defaultScale
			chart.Scale = &p
		}
		c.BarCharts[i] = chart
	}

	for i, g := range c.Gauges {

		setDefaultTriggersValues(g.Triggers)

		if g.RefreshRateMs == nil {
			r := defaultRefreshRateMs
			g.RefreshRateMs = &r
		}
		if g.Scale == nil {
			p := defaultScale
			g.Scale = &p
		}
		var items []Item
		for label, script := range g.Values {
			l := label
			items = append(items, Item{Label: &l, Script: script})
		}
		g.Items = items
		c.Gauges[i] = g
	}

	for i, box := range c.AsciiBoxes {

		setDefaultTriggersValues(box.Triggers)

		if box.RefreshRateMs == nil {
			r := defaultRefreshRateMs
			box.RefreshRateMs = &r
		}
		if box.Label == nil {
			label := string(i)
			box.Label = &label
		}
		if box.Font == nil {
			font := console.AsciiFontFlat
			box.Font = &font
		}
		if box.Color == nil {
			color := console.ColorWhite
			box.Color = &color
		}
		c.AsciiBoxes[i] = box
	}
}

func setDefaultTriggersValues(triggers []TriggerConfig) {

	defaultTerminalBell := true
	defaultSound := false
	defaultVisual := true

	for i, trigger := range triggers {

		if trigger.Actions == nil {
			trigger.Actions = &ActionsConfig{TerminalBell: &defaultTerminalBell, Sound: &defaultSound, Visual: &defaultVisual, Script: nil}
		} else {
			if trigger.Actions.TerminalBell == nil {
				trigger.Actions.TerminalBell = &defaultTerminalBell
			}
			if trigger.Actions.Sound == nil {
				trigger.Actions.Sound = &defaultSound
			}
			if trigger.Actions.Visual == nil {
				trigger.Actions.Visual = &defaultVisual
			}
		}

		triggers[i] = trigger
	}
}

func (c *Config) setDefaultLayout() {
	// TODO auto-arrange components
}

func (c *Config) setDefaultColors() {

	palette := console.GetPalette(*c.Theme)
	colorsCount := len(palette.Colors)

	for _, chart := range c.RunCharts {
		for j, item := range chart.Items {
			if item.Color == nil {
				item.Color = &palette.Colors[j%colorsCount]
				chart.Items[j] = item
			}
		}
	}

	for _, chart := range c.BarCharts {
		for j, item := range chart.Items {
			if item.Color == nil {
				item.Color = &palette.Colors[j%colorsCount]
				chart.Items[j] = item
			}
		}
	}

	for i, g := range c.Gauges {
		if g.Color == nil {
			g.Color = &palette.Colors[i%colorsCount]
			c.Gauges[i] = g
		}
	}
}
