package widgets

import (
    "log"
    // "fmt"
    // "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
)

func WindowNew(title string, width, height int) (*gtk.Window){
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Unable to create window: ", err)
    }

    win.SetTitle(title)
    win.SetDefaultSize(width, height)
    win.SetPosition(gtk.WIN_POS_CENTER)
    win.Connect("destroy", func(){
        gtk.MainQuit()
    })  
    return win 
}

func LabelNew(title string, wrapText bool) (*gtk.Label){
    label, err := gtk.LabelNew(title)
    if err != nil {
        log.Fatal("Unable to add label: ", err)
    }
    label.SetLineWrap(wrapText)
    return label
}


func EntryNew() (*gtk.Entry){
    entry, err := gtk.EntryNew()
    if err != nil {
        log.Fatal("Unable to create entry: ", err)
    }
    return entry
}

func ButtonNew(label string, onClick func(), args ...interface{}) (*gtk.Button){
    button, err := gtk.ButtonNewWithLabel(label)
    if err != nil {
        log.Fatal("Unable to create button: ", err)
    }
    button.Connect("clicked", onClick, args)
    return button
}

func BoxNew(orient gtk.Orientation, spacing int) (*gtk.Box){
    box, err := gtk.BoxNew(orient, spacing)
    if err != nil{
        log.Fatal("Unable to create a Box")
    }
    return box
}

func StackSwitcherNew() (*gtk.StackSwitcher){
    stackSwitcher,err := gtk.StackSwitcherNew()
    if err != nil {
        log.Fatal("Unable to add StackSwitcher: ", err)
    }
    return stackSwitcher
}

func StackNew() (*gtk.Stack){
    stack, err := gtk.StackNew()
    if err != nil {
        log.Fatal("Unable to add Stack: ", err)
    }
    return stack
}

func GridNew(columnHomogeneous, rowHomogeneous bool, colSpacing, rowSpacing uint) (*gtk.Grid) {
    grid, err := gtk.GridNew()
    if err != nil {
        log.Fatal("Unable to add grid: ", err)
    }
    grid.SetColumnHomogeneous(columnHomogeneous)
    grid.SetRowHomogeneous(rowHomogeneous)
    grid.SetColumnSpacing(colSpacing)
    grid.SetRowSpacing(rowSpacing)
    return grid
}

func DialogNew(title string, width, height int) (*gtk.Dialog){
    dialog, err := gtk.DialogNew()
    if err != nil {
        log.Fatal("Unable to add Dialog: ", err)
    }
    dialog.SetDefaultSize(width, height)
    dialog.SetTitle(title)
    return dialog
}

func DialogNewWithButtons(title string, parent gtk.IWindow, flag gtk.DialogFlags, buttons []interface{}) (*gtk.Dialog){
    dialog, err := gtk.DialogNewWithButtons(title, parent, flag, buttons)
    if err != nil {
        log.Fatal("Unable to create Dialog with Buttons: ", err)
    }
    return dialog
}

func RadioButtonNew(grpMember *gtk.RadioButton, label string, callback func()) (*gtk.RadioButton){
    radio, err := gtk.RadioButtonNewWithLabelFromWidget(grpMember, label)
    if err != nil {
        log.Fatal("Unable to generate radio button: ", err)
    }
    radio.Connect("toggled", callback)
    return radio
}
