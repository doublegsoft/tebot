package database

import (
  "context"
  "fmt"

  "github.com/jackc/pgx/v5"
)

// QueryPostgres 执行查询：连接参数由 dsn 传入，SQL 与查询参数由 query/args 传入。
// 返回每一行的列名 -> 值映射；value 中 []byte 会被转换成 string。
func QueryPostgres(
  dsn string,
  query string,
  expected string,
) ([]map[string]any, error) {
  ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
  if err != nil {
    return nil, fmt.Errorf("new pgx pool: %w", err)
  }
  defer conn.Close(ctx)

  rows, err := conn.Query(ctx, query)
  if err != nil {
    return nil, fmt.Errorf("query postgres: %w", err)
  }
  defer rows.Close()

  fds := rows.FieldDescriptions()
  cols := make([]string, len(fds))
  for i, fd := range fds {
    cols[i] = fd.Name
  }

  out := make([]map[string]any, 0)
  for rows.Next() {
    values, err := rows.Values()
    if err != nil {
      return nil, fmt.Errorf("read row values: %w", err)
    }

    rowMap := make(map[string]any, len(cols))
    for i, col := range cols {
      v := values[i]
      if b, ok := v.([]byte); ok {
        rowMap[col] = string(b)
      } else {
        rowMap[col] = v
      }
    }
    out = append(out, rowMap)
  }

  if err := rows.Err(); err != nil {
    return nil, fmt.Errorf("rows error: %w", err)
  }

  return out, nil
}