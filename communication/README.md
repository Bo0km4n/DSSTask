# How to run

## client

```
$ cd src
$ mkdir bin
$ make run.client
```

## server

```
$ cd src
$ make run.server
```

## Protocol

HTTP

### server endpoint

- `GET /api/v1/struct` Task構造体情報をserializeしてclientに返信
- `POST /api/v1/post_struct` Person構造体情報を受け取ってdeserializeしてTask.hello()を呼び出す

### client process

1. `GET /api/v1/struct` を叩いて構造体情報を取得
2. デシリアライズしてC言語の構造体に流し込む
3. 文字列を連結
4. シリアライズしてrequest bodyに構造体情報を埋め込み `POST /api/v1/post_struct` を叩く

## code architecture

```
.
├── README.md
├── sample.tar.gz
└── src
    ├── Makefile
    ├── Middleware.class
    ├── Middleware.java
    ├── Person.class
    ├── Person.java
    ├── Server.class
    ├── Server.java
    ├── StructFetchHandler.class
    ├── StructFetchHandler.java
    ├── StructPostHandler.class
    ├── StructPostHandler.java
    ├── Task.class
    ├── Task.java
    └── client
        ├── a.out
        ├── blist.c
        ├── blist.h
        ├── client.c
        ├── descriptor.c
        ├── descriptor.h
        ├── net.c
        ├── net.h
        ├── parser.c
        ├── parser.h
        ├── person.c
        ├── person.h
        ├── serializer.c
        ├── serializer.h
        ├── task.c
        ├── task.h
        ├── util.c
        └── util.h
```

## Result example


### server
```
javac Server.java
java Server
start http listening :18888
GET / HTTP/1.1
Accept: [*/*]
Host: [localhost:18888]

POST / HTTP/1.1
Accept: [*/*]
Host: [localhost:18888]
Content-type: [application/x-www-form-urlencoded]
Content-length: [69]

HelloWorld!
```

### client

```
fetched data:
0xac, 0xed, 0x00, 0x05, 0x73, 0x72, 0x00, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x02, 0x00, 0x04, 0x49, 0x00, 0x01, 0x76, 0x42, 0x00, 0x01, 0x78, 0x4c, 0x00, 0x04, 0x73, 0x74, 0x72, 0x31, 0x74, 0x00, 0x12, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x4c, 0x00, 0x04, 0x73, 0x74, 0x72, 0x32, 0x71, 0x00, 0x7e, 0x00, 0x01, 0x78, 0x70, 0x12, 0x34, 0x56, 0x78, 0x01, 0x74, 0x00, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x74, 0x00, 0x06, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21,
newHandle 0
newHandle 1
handle 7e0001
newHandle 2
newHandle 3
newHandle 4
len: 95, read: 95
task.v: 0x12345678
task.x: 1
task.str1: Hello
task.str2: World!
person.name: HelloWorld!
serialize_result:
  0000  ac ed 00 05 73 72 00 06 50 65 72 73 6f 6e 11 22  ....sr..Person."
  0010  33 44 55 66 77 88 02 00 01 4c 00 04 6e 61 6d 65  3DUfw....L..name
  0020  74 00 12 4c 6a 61 76 61 2f 6c 61 6e 67 2f 53 74  t..Ljava/lang/St
  0030  72 69 6e 67 3b 78 70 74 00 0b 48 65 6c 6c 6f 57  ring;xpt..HelloW
  0040  6f 72 6c 64 21                                   orld!
ok
```