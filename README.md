# Logrus Config

提供`Logrus`的相关配置

<br/>

## 快速上手

### 全局配置

``` go
package main

import (
  "github.com/sirupsen/logrus"
  
  logrus_conf "github.com/surh-go/logrus_conf"
)

func main() {
  logrus_conf.Init("./logs/logrus.log")
  logrus.Info("INFO")
}
```

<br/>

### Gin

``` go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/surh-go/logrus_conf/middleware"

  _ "github.com/surh-go/logrus_conf"
)

func main() {
  r := gin.New()
  r.Use(middleware.GinMiddleware())
  r.Run(":8080")
}
```
