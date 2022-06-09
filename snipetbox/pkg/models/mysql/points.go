package mysql

import (
	"database/sql"
	"math/rand"

	"github.com/rmcs87/cc5m/pkg/models"
)

type PointsModel struct{
  DB *sql.DB
}

func (m *PointsModel)Insert(name string)(error){
  stmt := `INSERT INTO Pontuacao (name, pontos) VALUES(?, ?)`

  points := rand.Intn(150)
  
  _, err := m.DB.Exec(stmt, name, points)
  if err != nil{
    return err
  }
  return nil
}

func(m *PointsModel) Latest()([]*models.Pontuacao, error){
  stmt := `SELECT * FROM Pontuacao`
  rows, err := m.DB.Query(stmt)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  ranking := []*models.Pontuacao{}
  for rows.Next(){
    s := &models.Pontuacao{}
    err = rows.Scan(&s.ID, &s.Name, &s.Pontos)

    if err != nil{
      return nil, err
    }
    ranking = append(ranking, s)
  }
  err = rows.Err()
  if err != nil{
    return nil, err
  }
  return ranking, nil
}

// func (m *SnippetModel)Insert(title, content, expires string)(int, error){
//   stmt := `INSERT INTO snippets(title, content, created, expires)                                    VALUES(?,?,UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

//   result, err := m.DB.Exec(stmt, title, content, expires)
//   if err != nil{
//     return 0,err
//   }

//   id, err := result.LastInsertId()
//   if err != nil{
//     return 0, err
//   }
//   return int(id), nil
// }

// func(m *SnippetModel) Latest()([]*models.Snippet, error){
//   stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

//   rows, err := m.DB.Query(stmt)
//   if err != nil{
//     return nil, err
//   }
//   defer rows.Close()

//   snippets := []*models.Snippet{}
//   for rows.Next(){
//     s := &models.Snippet{}
//     err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

//     if err != nil{
//       return nil, err
//     }
//     snippets = append(snippets, s)
//   }
//   err = rows.Err()
//   if err != nil{
//     return nil, err
//   }
//   return snippets, nil
// }