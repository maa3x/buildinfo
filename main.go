package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/pterm/pterm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		os.Exit(1)
	}

	for _, fp := range os.Args[1:] {
		extractInfo(fp)
	}
}

func extractInfo(fp string) {
	info, err := buildinfo.ReadFile(fp)
	if err != nil {
		fmt.Println("Error reading build info:", err)
	}

	box := pterm.DefaultBox.WithTopPadding(1).WithTitle(fmt.Sprintf("%s [%s]", fp, info.Path))
	settings := printSettings(info)
	mods := [][]string{
		{"Kind", "Path", "Version", "Sum"},
	}
	mods = append(mods, module("Main", info.Main)...)
	for _, d := range info.Deps {
		mods = append(mods, module("Dep", *d)...)
	}
	table, _ := pterm.DefaultTable.WithHasHeader().WithData(mods).Srender()
	panels, _ := pterm.DefaultPanel.WithPanels([][]pterm.Panel{
		{{Data: table}},
		{{Data: ""}},
		{{Data: settings}},
	}).Srender()
	box.Println(panels)
}

func module(kind string, m debug.Module) [][]string {
	data := [][]string{
		{kind, m.Path, m.Version, m.Sum},
	}

	if m.Replace != nil {
		data = append(data, module(kind+" [Replace]", *m.Replace)...)
	}

	return data
}

func printSettings(info *debug.BuildInfo) string {
	data := [][]string{
		{"GoVersion", info.GoVersion},
	}
	for _, s := range info.Settings {
		data = append(data, []string{s.Key, s.Value})
	}
	str, _ := pterm.DefaultTable.WithHasHeader().WithData(data).Srender()
	return str
}
