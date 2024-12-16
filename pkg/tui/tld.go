package tui

//type tldModel struct {
//	textInput textinput.Model
//	callback  func(input string) string
//}
//
////	func initialTLDModel(callback func(input string) string) model {
////		ti := textinput.New()
////		ti.Placeholder = "Phone"
////		ti.Focus()
////		ti.CharLimit = 156
////		ti.Width = 20
////
////		return model{
////			textInput: ti,
////			callback:  callback,
////		}
////	}
////
////	func (m tldModel) Init() tea.Cmd {
////		return textinput.Blink
////	}
////
//// //	func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//// //		var cmd tea.Cmd
//// //
//// //		switch msg := msg.(type) {
//// //		case tea.KeyMsg:
//// //
//// //			switch msg.Type {
//// //			case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
//// //				return m, tea.Quit
//// //			}
//// //
//// //		// We handle errors just like any other message
//// //		case errMsg:
//// //			m.err = msg
//// //			return m, nil
//// //		}
//// //		m.callback(m.textInput.Value())
//// //
//// //		m.textInput, cmd = m.textInput.Update(msg)
//// //		return m, cmd
//// //	}
//func (m tldModel) View() string {
//	input := m.textInput.View()
//	tableString := m.callback(m.textInput.Value())
//	helpText := "Type to filter | q/ctrl+c to quit"
//	label := "Enter country code for country?"
//	return lipgloss.JoinVertical(lipgloss.Left,
//		baseStyle.Render(label),
//		input,
//		color.YellowColorText.Render(helpText),
//		tableString,
//	)
//}

func FilterTLD(callback func(input string) string) {
	callback("foo")
	//p := tea.NewProgram(initial(callback))
	//if _, err := p.Run(); err != nil {
	//	log.Fatal(err)
	//}
}
