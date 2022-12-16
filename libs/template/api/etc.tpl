Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
#日志配置
Log:
  Mode: file
  Level: info
  Compress: true
  KeepDays: 3

DB:
  DataSource: {{.dataSource}}?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true
#redis配置
Cache:
  - Host: {{.cacheHost}}
  # Pass: xxx