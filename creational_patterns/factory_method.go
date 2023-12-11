package creational_patterns

type Dialog struct{
	render()
	createButton()(*models.Button)
}
type WindowsDialog struct{
	Dialog
}
type WebDialog struct{
	Dialog
}

func (WindowsDialog) createButton()(*models.WindowsButton){
	return &models.WindowsButton{}
}
func (HTMLButton) createButton()(*models.HTMLButton){
	return &models.HTMLButton{}
}

