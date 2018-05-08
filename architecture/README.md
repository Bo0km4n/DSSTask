## Environment

- go 1.10

## Example

### run server
`make example-server`

### run example client
- `make example-tcp-cli`
- `make example-udp-cli`

### run example with log file
- `make example-tcp-with-output`
- `make example-udp-with-output`
- `make example-server-with-outpt`

## About source code
```
./
├── Makefile
├── README.md
├── client
│   └── main.go
├── common
│   └── util.go
├── layer1
│   └── dip.go
├── layer2
│   ├── dtcp.go
│   └── dudp.go
├── layer3
│   └── data.go
├── lib
│   └── message.go
└── server
    └── main.go
```

### client
クライアント側実行ファイル

dependencies (
    layer1
    layer2
    layer3
    lib
    common
)

### server
サーバー側実行ファイル

dependencies (
    layer1
    layer2
    layer3
    lib
    common
)

### common
共通処理の切り出し

### lib
送信するバイト配列のモデルとそれに関する処理

### layer1
ip層モデル

### layer2
- tcp層モデル
- udp層モデル

### layer3
ペイロード部分モデル