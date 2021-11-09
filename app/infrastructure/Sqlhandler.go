package infrastructure

import (
  "database/sql"
  _ "github.com/lib/pq"
  "girack/interfaces/database"
)

const (
  conf = "host=postgres port=5555 user=girak password=password dbname=girack sslmode=disable"
)

type SqlHandler struct {
  Conn *sql.DB
}

func NewSqlHandler() database.SqlHandler {
  conn, err := sql.Open("postgres", conf)
  if err != nil {
    panic(err.Error)
  }
  defer conn.Close()

  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn

  return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
  res := SqlResult{}
  result, err := handler.Conn.Exec(statement, args...)
  if err != nil {
    return res, err
  }
  res.Result = result
  return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
  rows, err := handler.Conn.Query(statement, args...)
  if err != nil {
    return new(SqlRow), err
  }
  row := new(SqlRow)
  row.Rows = rows
  return row, nil
}

type SqlResult struct {
  Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
    return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
    return r.Result.RowsAffected()
}

type SqlRow struct {
    Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
    return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
    return r.Rows.Next()
}

func (r SqlRow) Close() error {
    return r.Rows.Close()
}