package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Test struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Pass  string `json:"pass"`
	//Enable bool   `json:"enable"`
}

func main() {
	f := &Test{}

	a := app.NewWithID("dev.suah.fyne_struct")
	bg := canvas.NewRectangle(color.Gray{Y: 0x16})
	w := a.NewWindow("Fyne Struct")
	w.SetContent(container.NewMax(bg))
	w.Resize(fyne.NewSize(300, 400))

	items, err := makeForm(f)
	if err != nil {
		log.Fatal(err)
	}

	dialog.ShowForm("Fyne Struct Test", "Ok", "Cancel", items,
		func(ok bool) {
			if ok {
				fmt.Printf("%#v\n", f)
			}
			os.Exit(0)
		}, w)
	w.ShowAndRun()
}

func makeForm(strct interface{}) ([]*widget.FormItem, error) {
	var items []*widget.FormItem

	strVal := reflect.ValueOf(strct)
	strType := reflect.Indirect(strVal).Type()

	for i := 0; i < strType.NumField(); i++ {
		v := strType.Field(i)
		field := strVal.Elem().FieldByName(v.Name)

		if !field.IsValid() || !field.CanSet() {
			return nil, fmt.Errorf("can't set field")
		}
		if field.Kind() != reflect.String || field.String() != "" {
			return nil, fmt.Errorf("not string")
		}

		//w := widget.NewEntry()
		//items = append(items, widget.NewFormItem(v.Tag.Get("json"), w))

		if s, ok := field.Addr().Interface().(*string); ok {
			boundString := binding.BindString(s)
			w := widget.NewEntryWithData(boundString)
			items = append(items, widget.NewFormItem(v.Tag.Get("json"), w))
		} else {
			fmt.Println("not ok")
		}

	}
	return items, nil
}
