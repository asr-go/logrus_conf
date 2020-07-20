# Logrus Config

提供`Logrus`的相关配置

<br/>

## 快速上手

### 全局配置

``` go
package main

import (
  "github.com/sirupsen/logrus"
  
  "github.com/asr-go/logrusconf"
)

func main() {
  logrusconf.Init("./logs/logrus.log")
  logrus.Info("INFO")
}
```

<br/>

### Gin

``` go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/asr-go/logrusconf/middleware"

  _ "github.com/asr-go/logrusconf"
)

func main() {
  r := gin.New()
  r.Use(middleware.GinMiddleware(true))
  r.Run(":8080")
}
```
