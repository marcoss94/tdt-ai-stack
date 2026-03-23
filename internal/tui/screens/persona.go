package screens

import (
	"strings"

	"github.com/gentleman-programming/gentle-ai/internal/model"
	"github.com/gentleman-programming/gentle-ai/internal/tui/styles"
)

func PersonaOptions() []model.PersonaID {
	return []model.PersonaID{model.PersonaGentleman, model.PersonaNeutral, model.PersonaCustom}
}

func RenderPersona(selected model.PersonaID, cursor int) string {
	var b strings.Builder

	b.WriteString(styles.TitleStyle.Render("Choose your Persona"))
	b.WriteString("\n\n")
	b.WriteString(styles.SubtextStyle.Render("TDT is direct, professional, and explains the why before the how."))
	b.WriteString("\n\n")

	for idx, persona := range PersonaOptions() {
		isSelected := persona == selected
		focused := idx == cursor
		b.WriteString(renderRadio(personaDisplayName(persona), isSelected, focused))
	}

	b.WriteString("\n")
	b.WriteString(renderOptions([]string{"Back"}, cursor-len(PersonaOptions())))
	b.WriteString("\n")
	b.WriteString(styles.HelpStyle.Render("j/k: navigate • enter: select • esc: back"))

	return b.String()
}

func personaDisplayName(persona model.PersonaID) string {
	if persona == model.PersonaGentleman {
		return "TDT"
	}

	return string(persona)
}
