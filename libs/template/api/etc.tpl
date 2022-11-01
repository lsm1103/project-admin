Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}

DB:
  DataSource: {{.dataSource}}?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true   //root:pujian123@tcp(172.16.10.183:4306)/im-center
#redis配置
Cache:
  - Host: {{.cacheHost}}    //172.16.10.183:6379
  # Pass: xxx