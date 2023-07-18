package abstractFactory

import "fmt"

type Button interface {
	Paint()
}

type Window interface {
	Display()
}

type WindowsButton struct{}

func (wb *WindowsButton) Paint() {
	fmt.Println("Windows button is painted.")
}

type WindowsWindow struct{}

func (ww *WindowsWindow) Display() {
	fmt.Println("Windows window is displayed.")
}

type MacOSButton struct{}

func (mb *MacOSButton) Paint() {
	fmt.Println("macOS button is painted.")
}

type MacOSWindow struct{}

func (mw *MacOSWindow) Display() {
	fmt.Println("macOS window is displayed.")
}

type GUIFactory interface {
	CreateButton() Button
	CreateWindow() Window
}

type WindowsFactory struct{}

func (wf *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (wf *WindowsFactory) CreateWindow() Window {
	return &WindowsWindow{}
}

type MacOSFactory struct{}

func (mf *MacOSFactory) CreateButton() Button {
	return &MacOSButton{}
}

func (mf *MacOSFactory) CreateWindow() Window {
	return &MacOSWindow{}
}

func AbstractFactory() {

	windowsFactory := &WindowsFactory{}
	windowsButton := windowsFactory.CreateButton()
	windowsWindow := windowsFactory.CreateWindow()

	windowsButton.Paint()
	windowsWindow.Display()

	macosFactory := &MacOSFactory{}
	macosButton := macosFactory.CreateButton()
	macosWindow := macosFactory.CreateWindow()

	macosButton.Paint()
	macosWindow.Display()
}
