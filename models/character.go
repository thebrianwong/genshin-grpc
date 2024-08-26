package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Character struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Height  string `json:"height"`
	Element string `json:"element"`
}

type CharacterModel struct {
	DB *pgx.Conn
}

func (c CharacterModel) All() ([]Character, error) {
	rows, err := c.DB.Query(context.Background(),
		"SELECT character.name, gender, height, element.name FROM genshin.character INNER JOIN genshin.element ON character.element_id=element.id ORDER BY character.id")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var characters []Character

	for rows.Next() {
		var character Character

		err := rows.Scan(&character.Name, &character.Gender, &character.Height, &character.Element)
		if err != nil {
			return nil, err
		}

		characters = append(characters, character)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return characters, nil
}

func (c CharacterModel) ById(id string) ([]Character, error) {
	var character Character

	err := c.DB.QueryRow(
		context.Background(),
		"SELECT character.name, gender, height, element.name FROM genshin.character INNER JOIN genshin.element ON character.element_id=element.id WHERE character.id=$1",
		id,
	).Scan(&character.Name, &character.Gender, &character.Height, &character.Element)
	if err != nil {
		return nil, err
	}

	return []Character{character}, nil
}
