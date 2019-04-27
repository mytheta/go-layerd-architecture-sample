# API Server



## Version
* Go -> 1.12
* Gin -> 1.3.0

### Run api server

```
make run
```

### Go Modules の設定
標準ライブラリを除く全てのパッケージをモジュールとして管理するため，以下を設定してください．  

```
 export GO111MODULE=on
```

## architecture
This project architect by layerd architecture

 
### ディレクトリ構成
```
go-layerd-architecture-sample
├── Makefile
├── README.md
├── application  // アプリケーション層
│   └── usecase
│       └── user_usecase.go
│
├── domain      // ドメイン層
│   ├── model
│   │   └── user.go
│   ├── repository
│   │   └── user_repository.go
│   └── service
│       └── user_service.go
├── go.mod  
├── go.sum
├── infrastructure　// インフラストラクチャ層
│   └── datastore
│       └── user_repository.go
└── interface　　　// UI(Presentation)層
    ├── api
    │   └── server
    │       ├── handler
    │       │   └── user_handler.go
    │       ├── main.go  // main
    │       └── router
    │           ├── api.go
    │           └── router.go
    └── env  //configの設定
        ├── conf
        │   ├── dev.toml
        │   ├── local.toml
        │   └── prd.toml
        └── env.go

```

#### UI(Presentation)層
外部からの入力を受け取り，情報を出力する責務を担います．

#### アプリケーション層
domain層が提供しているビジネスロジックを呼び出し，use caseに合わせてタスクを調整する責務を担います．

#### ドメイン層
ビジネスロジックとそれに伴う概念を全てここに配置します．

#### インフラストラクチャ層
永続化，メッセージ送信などの技術的機能を提供します． 
repository等の実装詳細をここで定義します．