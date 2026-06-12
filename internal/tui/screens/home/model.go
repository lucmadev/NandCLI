package home

type HomeScreen struct {
	cursor int
	items  []MenuItem
}

type MenuItem struct {
	Title string
	Route string
}