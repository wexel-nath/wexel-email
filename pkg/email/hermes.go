package email

// This file adds json keys to the structs in github.com/matcornic/hermes
// Some elements have not been included for simplicity

type Product struct {
	Name      string `json:"name"`
	Logo      string `json:"logo"`      // e.g. https://matcornic.github.io/img/logo.png
	Copyright string `json:"copyright"` // Copyright © 2017 Hermes. All rights reserved.
}

// Entry is a simple entry of a map
// Allows using a slice of entries instead of a map
// Because Golang maps are not ordered
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Button defines an action to launch
type Button struct {
	Color     string `json:"color"`
	TextColor string `json:"text_color"`
	Text      string `json:"text"`
	Link      string `json:"link"`
}

// Action is an action the user can do on the email (click on a button)
type Action struct {
	Instructions string `json:"instructions"`
	Button       Button `json:"button"`
}

// Columns contains meta-data for the different columns
type Columns struct {
	CustomWidth     map[string]string `json:"custom_width"`
	CustomAlignment map[string]string `json:"custom_alignment"`
}

// Table is an table where you can put data (pricing grid, a bill, and so on)
type Table struct {
	Data    [][]Entry `json:"data"`    // Contains data
	Columns Columns   `json:"columns"` // Contains meta-data for display purpose (width, alignement)
}

// Body is the body of the email, containing all interesting data
type Body struct {
	Name         string   `json:"name"`       // The name of the contacted person
	Intros       []string `json:"intros"`     // Intro sentences, first displayed in the email
	Dictionary   []Entry  `json:"dictionary"` // A list of key+value (useful for displaying parameters/settings/personal info)
	Table        Table    `json:"table"`      // Table is an table where you can put data (pricing grid, a bill, and so on)
	Actions      []Action `json:"actions"`    // Actions are a list of actions that the user will be able to execute via a button click
	Outros       []string `json:"outros"`     // Outro sentences, last displayed in the email
	Greeting     string   `json:"greeting"`   // Greeting for the contacted person (default to 'Hi')
	Signature    string   `json:"signature"`  // Signature for the contacted person (default to 'Yours truly')
	Title        string   `json:"title"`      // Title replaces the greeting+name when set
}
