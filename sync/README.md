## 14 - Sync

Lamportを用いた全順序マルチキャストの実装

## how to run

```
go run main.go
```

## tree

.
├── README.md
├── main.go
├── process
│   └── process.go
└── request
    └── request.go

### main.go

エントリーポイント

プロセスの生成、リクエストの生成、送信を行う

### request.go

requestとackの構造体を以下のように定義した

```go

type Request struct {
	ID       int
	ClientID int
	Tick     int
}

type Ack struct {
	Request   *Request
	ProcessID int
	Tick      int
}

```

#### type request

- ID リクエスト固有のID
- Client ID リクエストを発行したクライアントのID
- Tick クライアントが付与したタイムスタンプ

#### type ack

- Request 対応するrequest構造体へのポインタ
- ProcessID ackを返したプロセスのID
- Tick プロセスが付与したタイムスタンプ

### process.go

リクエストを受付処理し、他プロセスとの通信も行う構造体を以下のように定義した

```go
type Process struct {
	ID          int
	ProcessList []*Process
	Tick        int
	reqQueue    []*request.Request
	ackQueue    []*request.Ack
	reqChan     chan *request.Request
	ackChan     chan *request.Ack
	indent      string
}
```

- ID プロセス固有のID
- ProcessList 自身を含めたマルチキャストするためのポインタ配列
- Tick プロセスが起動中に増加していく時刻
- reqQueue リクエストをソートしてackと比較するための配列
- ackQueue ackQueueと同様
- reqChan 非同期処理のために別プロセスからリクエストを受け取るためのバッファ
- ackChan reqChanと同様
- indent 表示のためのタブフォーマット

それぞれのプロセス間ではchannelを通じてデータをやり取りする。
