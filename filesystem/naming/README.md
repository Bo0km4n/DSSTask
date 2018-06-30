# v6ファイルシステム

## How to compile and run

- requirements: go >= 1.8

```
go run main.go
```


## Architecture

```
.
├── README.md
├── byte
│   └── entity.go
├── disk
│   └── entity.go
├── entry
│   └── entity.go
├── filesys
│   └── entity.go
├── inode
│   └── entity.go
├── main.go
├── repl
│   └── repl.go
└── v6root
```

### byte
byte配列のラッパー

### disk
v6rootの持つdisk領域全体のラッパー

### entry
ファイル、ディレクトリのもつデータのリストをinodeから名前とサイズをマッピングした構造体
主にlsコマンドで使用

### filesys
filesys構造体のラッパー

### inode
inode構造体のラッパー

### repl
対話型コンソールのエントリーポイント

## Screen shot

pngフォルダ以下を参照
