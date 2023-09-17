package main

import (
	"epc/cmd"
	"epc/pkg"
)

func main() {
	cmd.SetStaticTime(&pkg.RealTime{})
	cmd.Execute()
	//components := &view.Components{
	//	TzTable:       ui.NewTimezoneTable(),
	//	TimeInfoTable: ui.NewTimeInfoTable(),
	//}
	//
	//v := view.New(components)
	//view.Init(v)
	//
	//v.Layout.Application.SetFocus(v.Components.TzTable.Table.Primitive())
	//v.Layout.Application.Run()
}
