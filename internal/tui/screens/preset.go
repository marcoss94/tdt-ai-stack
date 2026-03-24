package screens

import (
	"strings"

	"github.com/gentleman-programming/gentle-ai/internal/model"
	"github.com/gentleman-programming/gentle-ai/internal/tui/styles"
)

func PresetOptions() []model.PresetID {
	return []model.PresetID{
		model.PresetFullGentleman,
		model.PresetEcosystemOnly,
		model.PresetMinimal,
		model.PresetCustom,
	}
}

var presetDescriptions = map[model.PresetID]string{
	model.PresetFullGentleman: "Everything: memory, SDD, skills, docs, persona & security",
	model.PresetEcosystemOnly: "Core tools only: memory, SDD, skills & docs (no persona/security)",
	model.PresetMinimal:       "Just Engram persistent memory",
	model.PresetCustom:        "Pick individual components yourself",
}

func visiblePresetName(preset model.PresetID) string {
	switch preset {
	case model.PresetFullGentleman:
		return "TDT Standard"
	case model.PresetEcosystemOnly:
		return "Ecosystem Only"
	case model.PresetMinimal:
		return "Minimal"
	case model.PresetCustom:
		return "Custom"
	default:
		return string(preset)
	}
}

func RenderPreset(selected model.PresetID, cursor int) string {
	var b strings.Builder

	b.WriteString(styles.TitleStyle.Render("Select Ecosystem Preset"))
	b.WriteString("\n\n")

	for idx, preset := range PresetOptions() {
		isSelected := preset == selected
		focused := idx == cursor
		b.WriteString(renderRadio(visiblePresetName(preset), isSelected, focused))
		b.WriteString(styles.SubtextStyle.Render("    "+presetDescriptions[preset]) + "\n")
	}

	b.WriteString("\n")
	b.WriteString(renderOptions([]string{"Back"}, cursor-len(PresetOptions())))
	b.WriteString("\n")
	b.WriteString(styles.HelpStyle.Render("j/k: navigate • enter: select • esc: back"))

	return b.String()
}
