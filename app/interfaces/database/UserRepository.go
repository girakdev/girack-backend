package database

import "app/domain"

type SqlHandler interface {
  Execute(string, ...interface{}) (Result, error)
  Query(string, ...interface{}) (Row, error)
}

type Result interface {
  LastInsertId() (int64, error)
  RowsAffected() (int64, error)
}

type Row interface {
  Scan(...interface{}) error
  Next() bool
  Close() error
}

type UserRepository struct {
  SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error){
  result, err := repo.Execute(
    "INSERT INTO users (id_name, name) VALUES(?,?)", u.ID, u.Name,
  )
  if err != nil {
    return
  }
  id64, err := result.LastInsertId()
  if err != nil {
    return
  }
  id = int(id64)
  return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
  row, err := repo.Query("SELECT id, real_name, ideal_name FROM users WEHERE id = ?", identifier)
  if err != nil {
    return
  }
  var id int
  var idealName string
  var realName string

  row.Next()
  if err = row.Scan(&id, &idealName, &realName); err != nil {
    return
  }
  user.ID = id
  user.IdealName = idealName
  user.RealName = realName
  return
}


func (repo *UserRepository) FindAll() (users domain.Users, err error) {
  rows, err := repo.query("SELECT id, ideal_name, real_name FROM users")
  defer rows.Close()
  if err != nil {
    return
  }
  for rows.Next() {
    var id int
    var idealName string
    var realName string
    if err := rows.Scan(&id, &idealName, &realName); err != nil {
      continue
    }
    user := domain.User{
      ID: id,
      RealName: realName,
      IdealName: idealName,
    }
    users = append(users, user)
  }
  return
}
