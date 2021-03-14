# はじめに

## VisualStudioCode のデバッガ

VisualStudioCode で Go を使うときはデバッガのコマンドを忘れずに叩く  
go get -u github.com/derekparker/delve/cmd/dlv

## OAuth2 パッケージのインストール方法

go get github.com/stretchr/gomniauth/...

## objx パッケージのインストール方法

go get github.com/stretchr/objx

## gin をインストール

go get github.com/gin-gonic/gin

## 他

## Go で exe ファイルを作成する

go build -o [アプリケーション名]

## Windows 上で作成する場合は拡張子を付けて build

go build -o [アプリケーション名].exe
