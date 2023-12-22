package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/haykh/tuigo"
)

var food_menu = map[string]float32{
	"pizza":  9.99,
	"burger": 5.99,
	"salad":  6.99,
	"pie":    4.95,
	"cake":   3.95,
}
var drinks_menu = map[string]float32{
	"water":  1.99,
	"juice":  2.99,
	"coffee": 3.95,
	"tea":    2.95,
}

func main() {
	type pickupOrDeliveryMsg struct{}

	backend := tuigo.Backend{
		States: []tuigo.AppState{"food", "order"},
		Constructors: map[tuigo.AppState]tuigo.Constructor{
			"food": func(tuigo.Window) tuigo.Window {
				title := tuigo.Text("shell foods : pick food & beverage options", tuigo.NormalText)
				food_options := []string{}
				for item, price := range food_menu {
					food_options = append(food_options, fmt.Sprintf("%s ($%.2f)", item, price))
				}
				food_selector := tuigo.SelectorWithID(1, food_options, true, nil)
				food_container := tuigo.Container(
					true, tuigo.VerticalContainer, tuigo.Text("Food", tuigo.NormalText), food_selector,
				)

				drinks_options := []string{}
				for item, price := range drinks_menu {
					drinks_options = append(drinks_options, fmt.Sprintf("%s ($%.2f)", item, price))
				}
				drinks_selector := tuigo.SelectorWithID(2, drinks_options, true, nil)
				drinks_container := tuigo.Container(
					true, tuigo.VerticalContainer, tuigo.Text("Drinks", tuigo.NormalText), drinks_selector,
				)
				return tuigo.Container(
					true, tuigo.VerticalContainer, title, food_container, drinks_container,
				)
			},
			"order": func(prev tuigo.Window) tuigo.Window {
				_, selected_food_options_acc := prev.GetElementByID(1)
				selected_food_options := selected_food_options_acc.Data().([]string)
				_, selected_drinks_options_acc := prev.GetElementByID(2)
				selected_drinks_options := selected_drinks_options_acc.Data().([]string)
				item_str := "items:\n\n"
				var price float32 = 0.0
				for _, option := range append(selected_food_options, selected_drinks_options...) {
					item := option[:len(option)-8]
					item_str += fmt.Sprintf("+ %s\n", item)
					if _, ok := food_menu[item]; ok {
						price += food_menu[item]
					} else {
						price += drinks_menu[item]
					}
				}
				subtotal := tuigo.Text(fmt.Sprintf("subtotal: $%.2f", price), tuigo.DimmedText)
				utencils := tuigo.RadioWithID(1, "include utencils", nil)
				pickup_or_delivery := tuigo.SelectorWithID(2, []string{"pickup", "delivery"}, false, pickupOrDeliveryMsg{})
				address_entry := tuigo.InputWithID(4, "delivery address", "<address>", "<address>", tuigo.TextInput, nil).Hide()
				container := tuigo.Container(
					true,
					tuigo.HorizontalContainer,
					tuigo.TextWithID(3, item_str, tuigo.NormalText),
					tuigo.Container(true, tuigo.VerticalContainer, utencils, pickup_or_delivery, address_entry),
				)
				return tuigo.Container(true, tuigo.VerticalContainer, container, subtotal)
			},
		},
		Updaters: map[tuigo.AppState]tuigo.Updater{
			"order": func(window tuigo.Window, msg tea.Msg) (tuigo.Window, tea.Cmd) {
				hide_address := tuigo.TgtCmd(
					4,
					func(cont tuigo.Wrapper, inp tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
						return cont.Hide().(tuigo.Wrapper), inp
					},
				)
				unhide_address := tuigo.TgtCmd(
					4,
					func(cont tuigo.Wrapper, inp tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
						return cont.Unhide().(tuigo.Wrapper), inp
					},
				)
				switch msg.(type) {
				case pickupOrDeliveryMsg:
					_, pickup_or_delivery_acc := window.GetElementByID(2)
					pickup_or_delivery := pickup_or_delivery_acc.Data()
					if pickup_or_delivery != nil {
						if pickup_or_delivery.(string) == "delivery" {
							return window, unhide_address
						} else {
							return window, hide_address
						}
					}
				}
				return window, nil
			},
		},
		Finalizer: func(containers map[tuigo.AppState]tuigo.Window) tuigo.Window {
			prev := containers["order"]
			_, utencils_acc := prev.GetElementByID(1)
			utencils := utencils_acc.Data().(bool)
			_, pickup_or_delivery_acc := prev.GetElementByID(2)
			pickup_or_delivery := pickup_or_delivery_acc.Data().(string)
			_, order_acc := prev.GetElementByID(3)
			order := order_acc.Data().(string)
			items := []string{}
			for _, line := range strings.Split(order, "\n") {
				if len(line) > 0 && line[0] == '+' {
					items = append(items, line[2:])
				}
			}
			if utencils {
				items = append(items, "utencils")
			}
			var text string
			items_txt := strings.Join(items, ", ")
			ind := strings.LastIndex(items_txt, ",")
			items_txt = items_txt + strings.Replace(items_txt[ind:], ",", " and", 1)
			if pickup_or_delivery == "pickup" {
				text = fmt.Sprintf("Your order of %s will soon be ready for pickup", items_txt)
			} else {
				_, address_acc := prev.GetElementByID(4)
				address := address_acc.Data().(string)
				text = fmt.Sprintf("Your order of %s\nis on its way to %s", items_txt, address)
			}
			return tuigo.Container(
				false,
				tuigo.VerticalContainer,
				tuigo.Text(text, tuigo.LabelText),
			)
		},
	}
	p := tea.NewProgram(tuigo.App(backend, false))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
