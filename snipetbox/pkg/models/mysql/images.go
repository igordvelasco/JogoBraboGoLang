package mysql

import (
	"database/sql"

	"github.com/rmcs87/cc5m/pkg/models"
)

type ImageModel struct{
  DB *sql.DB
}

func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func(m *ImageModel) AllCat()([]string, error){
  stmt := `SELECT * FROM Imagens`
  rows, err := m.DB.Query(stmt)
  
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  images := []*models.Image{}
  for rows.Next(){
    s := &models.Image{}
    err = rows.Scan(&s.ID, &s.Name, &s.Image_path, &s.Palavra_chave, &s.Dificuldade, &s.Categoria)
    if err != nil{
      return nil, err
    }
    images = append(images, s)
  }
  err = rows.Err()
  if err != nil {
    return  nil, err
  }

  var allCat[] string
  for i := range(images){
    allCat = append(allCat, images[i].Categoria)
  }

  allCat = removeDuplicateStr(allCat)
  return allCat, nil
}

func(m *ImageModel) LatestCatDif(categoria string, dificuldade string)([]*models.Image, error){
  stmt := `SELECT * FROM Imagens WHERE categoria = ? AND dificuldade = ?`
  rows, err := m.DB.Query(stmt, categoria, dificuldade)
  
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  images := []*models.Image{}
  for rows.Next(){
    s := &models.Image{}
    err = rows.Scan(&s.ID, &s.Name, &s.Image_path, &s.Palavra_chave, &s.Dificuldade, &s.Categoria)
    if err != nil{
      return nil, err
    }
    images = append(images, s)
  }
  err = rows.Err()
  if err != nil {
    return  nil, err
  }
  
  return images, nil
}