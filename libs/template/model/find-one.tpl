func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}, resp interface{}) (err error) {
	{{if .withCache}}{{.cacheKey}}
	err = m.QueryRowCtx(ctx, resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query :=  fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		if session != nil {
			return session.QueryRowCtx(ctx, v, query, {{.lowerStartCamelPrimaryKey}})
		}
		return conn.QueryRowCtx(ctx, v, query, {{.lowerStartCamelPrimaryKey}})
	}){{else}}
	query := fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
	//var resp {{.upperStartCamelObject}}
	err = m.conn.QueryRowCtx(ctx, resp, query, {{.lowerStartCamelPrimaryKey}}){{end}}
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return sqlUtils.ErrNotFound
	default:
		return err
	}
}
