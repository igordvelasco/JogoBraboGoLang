package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record Found")

type Image struct{
  ID int
  Name string
  Image_path string
  Categoria string
  Palavra_chave string
  Dificuldade string
}

type Pontuacao struct{
  ID int
  Name string
  Pontos int
}