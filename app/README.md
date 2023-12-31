# **Golang Clean Architecture**

## **Introduction**
本篇使用golang & gin框架建構後端，並以符合clean architecture的方式規劃專案code的layout
以此為後續專案code的example，此layout適合較大型專案，若小型專案相比之下程式碼會顯得較多餘。

---
## **clean architecture**

此架構符合clean architecture
* 可測試
* 獨立於框架
* 獨立於通訊協議
* 獨立於 Database

此架構以domain (module)來區分資料夾，並主要將程式碼分成**controller**、**service**、**repository**三部分，並使用依賴注入進行建構，能夠符合**關注點分離原則**，
也能夠更好的進行解耦，保證核心邏輯(service)程式碼的健壯性。

**controller**
```
* 職責 : 處理http協議、解出參數、響應response

需處理service中返回之result與error，
由 /pkg/except 中定義之error種類轉換對應之http status code
``` 

**service**
```
* 職責 : 實現業務邏輯

dto實現與controller層的解耦，讓此服務不用關注協議部分、或使用何種框架，若
後續專案架構變動(換框架 or 切分微服務)，都不會影響此業務邏輯。

interface實現與repository的解耦，使用interface規範repository的func，讓service依賴於抽象而不是實作，實現控制反轉。

參數檢查由dto func實現，避免service內寫入過多檢查的程式碼。

service丟出之error盡量避免使用error.New()，因若使用error.New()，controller將會轉換出之
status code 400，錯誤碼較不明確。

error種類由 pkg/base/service中 import pkg/except中的error種類來定義，
此方式能夠在多人開發時統一管理error種類並能夠更好的維護error轉換之status code。
```

**repository**
```
* 職責 : 處理資料庫、快取部分

撰寫repository時，需同步維護 pkg/repos/interfaces 中的interface。
```

***hint**
```
* 定義錯誤類型，並定義特定錯誤轉換之status code會放在pkg/except/exception.go中，讓middleware、controller、
  service的base class實作，再藉由繼承調用這些error type。

* 若有使用第三方API or IO相關操作(ex:mail、第三方金流..)，開新的domainModule資料夾並寫成service，再使用依賴
  注入注進欲使用之service。

* 寫成service跟單純的寫成helper，分別在於helper不會撰寫過多的邏輯且不會有與第三方交互的動作
  (第三方API、redis、mail等等...)，主要是順手可以直接拿來用的小程式，service可以撰寫
  較多業務邏輯且能夠與第三方交互，回歸可測試性，若 helpers/mailer 沒寫成 mailModule/mailService 
  那在撰寫測試程式碼時，就較難對mailer進行mock的動作。
```

<br>

---
## **migration**
詳細請參考 : ./pkg/migrate/readme.md
進行資料庫同步
```
> go run ./pkg/migrate/migrate.go
``` 

<br>

---
## **go mod**
管理套件
初始化套件管理 , project name會影響到import路徑
```
> go mod init  `project name`
```
自動裝載引入之套件、自動卸載未使用之套件
```
> go mod tidy
```

<br>

---
## dependency injection
依賴注入
在 XXXModule 資料夾底下新增 wire.go
寫入各個layer的construct func，下command後套件會自動完成所有注入
* wire不會自動找對應之interface與實作，因此interface需使用wire中較高級語法。
```
wire.go
-------------------
//+build wireinject

package userModule

import (
	"github.com/google/wire"
	"golang/pkg/repos"
	"golang/pkg/helpers"
)

func InitializeUserController() *UserController{
	wire.Build(NewUserController , NewUserService , repos.NewUserRepo , helpers.NewSqlSession)
	return &UserController{}
}
```

```
> wire
```

會自動建立出 wire_gen.go 供router調用
```
wire_gen.go
-------------------
// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package userModule

import (
	"golang/pkg/helpers"
	"golang/pkg/repos"
)

// Injectors from wire.go:

func InitializeUserController() *UserController {
	db := helpers.NewSqlSession()
	userRepo := repos.NewUserRepo(db)
	userService := NewUserService(userRepo)
	userController := NewUserController(userService)
	return userController
}
```
