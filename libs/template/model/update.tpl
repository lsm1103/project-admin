func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}err:=m.FindOne(ctx, session, data.{{.upperStartCamelPrimaryKey}}, &{{.upperStartCamelObject}}{})
	if err!=nil{
		return err
	}
	{{end}}{{.keys}}

    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join(sqlUtils.BuildFields(data, sqlUtils.IsEmptyValue),", "))
        if session != nil {
            return session.Exec(query, data.Id)
        }
        return conn.Exec(query, data.Id)
	}, {{.keyValues}})
	{{else}}
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}
