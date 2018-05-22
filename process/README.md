## About Source Architecture

```
.
├── README.md
├── q2_1
│   ├── Makefile
│   ├── capthread.c
│   └── capthread.h
├── q2_2
│   ├── Makefile
│   ├── capfork.c
│   └── capfork.h
└── q2_3
    ├── Makefile
    ├── capfork2.c
    └── capfork2.h

3 directories, 10 files
```

## Q. 2_1

### 実行方法
```
cd Q2_1 && make
```

```
$ type something... >>> hello
HELLO
```

### 概要
入力した値を別スレッドに渡してcapitalize後に表示

## Q. 2_2

### 実行方法
```
cd Q2_2 && make
```

```
$ type your message >>> hello
write message : HELLO
read message from pipe: HELLO
```

### 概要
forkした子プロセス内で入力を受け取り、capitalize後にパイプに書き込む。
親プロセスはパイプを読み込み文字を出力する

## Q. 2_3
### 実行方法
```
cd Q2_3 && make
```

```
$ type your message >>> hello
[parent] write message : hello

[child] read message from pipe: hello

[child] write message : HELLO

[parent] read message from pipe: HELLO
```

### 概要

1. 親プロセスが入力を受け取り、パイプに書き込む。
2. 子プロセスがパイプを読み取り、値をcapitalizeした値を別のパイプに書き込む
3. 親プロセスは別のパイプから読み取り、値を表示する

 