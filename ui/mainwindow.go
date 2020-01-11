package ui

// MainWindow struct
// type MainWindow struct {
// 	Close func() bool
// 	PushButtonExit *widgets.QPushButton
// 	widget *widgets.QWidget
// }
//
// func (c *MainWindow) init(formWidget *widgets.QWidget) {
// 	c.widget = formWidget
// 	c.Close = formWidget.Close
// 	c.PushButtonExit = widgets.NewQPushButtonFromPointer(formWidget.FindChild("pushButtonExit", core.Qt__FindChildrenRecursively).Pointer())
//
// 	c.PushButtonExit.ConnectPressed( func() {
// 		c.Close()
// 	})
// }
//
// // NewMainWindow xaxa
// func NewMainWindow() *widgets.QWidget {
// 	file := core.NewQFile2("./ui/mainwindow.ui")
// 	file.Open(core.QIODevice__ReadOnly)
// 	formWidget := uitools.NewQUiLoader(nil).Load(file, nil)
// 	file.Close()
//
// 	mw := MainWindow{}
//
// 	mw.init(formWidget)
//
// 	return formWidget
// }
