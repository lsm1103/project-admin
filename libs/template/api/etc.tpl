Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}

DB:
  DataSource: {{.dataSource}}?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true
#redis配置
Cache:
  - Host: {{.cacheHost}}
  # Pass: xxx