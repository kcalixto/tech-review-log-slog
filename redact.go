package main

import "log/slog"

type Person struct {
	Name          string `json:"name"`
	FavoriteColor string `json:"favorite_color"`
}

func (p Person) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", p.Name),
		slog.String("favorite_color", "***"),
	)
}

func PrintWithRedactedFields() {
	person := Person{
		Name:          "Ca Calixto lixto",
		FavoriteColor: "green (don't tell anyone)",
	}

	slog.Info("redact log information", "person", person)
}
