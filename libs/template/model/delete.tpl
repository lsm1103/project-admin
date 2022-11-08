func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}{{if .containsIndexCache}}{{if ne .upperStartCamelObject "data" }}data := &{{.upperStartCamelObject}}{}{{end}}
	err:=m.FindOne(ctx, session, {{.lowerStartCamelPrimaryKey}}, data)
	if err!=nil{
		return err
	}c

	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
		}
		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}})
	{{else}}query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		_,err:=m.conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	{{end}}return err
}
