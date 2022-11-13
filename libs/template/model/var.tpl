var (
	{{.lowerStartCamelObject}}FieldNames          = builder.RawFieldNames(&{{.upperStartCamelObject}}{}{{if .postgreSql}}, true{{end}})
	{{.lowerStartCamelObject}}Rows                = strings.Join({{.lowerStartCamelObject}}FieldNames, ",")
	{{.lowerStartCamelObject}}RowsExpectAutoSet   = {{if .postgreSql}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, {{if .autoIncrement}}"{{.originalPrimaryKey}}",{{end}} "create_time", "update_time"), ","){{else}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, {{if .autoIncrement}}"{{.originalPrimaryKey}}",{{end}} {{.ignoreColumns}}, "`state`"), ","){{end}}
	{{.lowerStartCamelObject}}RowsWithPlaceHolder = {{if .postgreSql}}builder.PostgreSqlJoin(stringx.Remove({{.lowerStartCamelObject}}FieldNames, "{{.originalPrimaryKey}}", "create_time", "update_time")){{else}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, "{{.originalPrimaryKey}}", {{.ignoreColumns}}), "=?,") + "=?"{{end}}
	{{.lowerStartCamelObject}}ListRows 			  = {{if .postgreSql}} builder.PostgreSqlJoin( builder.RawFieldNames( &{{.upperStartCamelObject}}{} ), ","){{else}} strings.Join( builder.RawFieldNames( &{{.upperStartCamelObject}}{} ), ",") {{end}}

	{{if .withCache}}{{.cacheKeys}}{{end}}
)
