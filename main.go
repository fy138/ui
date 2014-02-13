// 11 february 2014
package main

func main() {
	w := NewWindow("Main Window", 320, 240)
	w.Closing = make(chan struct{})
	b := NewButton("Click Me")
	err := w.Open(b)
	if err != nil {
		panic(err)
	}

	w2 := NewWindow("Checkbox Window", 200, 100)
	c := NewCheckbox("Check Me")
	err = w2.Open(c)
	if err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case <-w.Closing:
			break mainloop
		case <-b.Clicked:
			err := w.SetTitle("Button Clicked")
			if err != nil {
				panic(err)
			}
		}
	}
	w.Hide()
}

