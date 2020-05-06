# Simple Chat Client Go について

## VisualStudioCodeのデバッガ  

VisualStudioCodeでGoを使うときはデバッガのコマンドを忘れずに叩く  
go get -u github.com/derekparker/delve/cmd/dlv  
  
  
## OAuth2パッケージのインストール方法  
go get github.com/stretchr/gomniauth/...  

## objx パッケージのインストール方法
go get github.com/stretchr/objx  

## Chatの認証情報について  

mainFunctionにはClinetIDとkeyを入れる必要がある。  
Google Developer Consoleでこのアプリケーションの認証情報を登録して  
クライアントIDと秘密鍵(key)を取得しないといけない。  

## 今後の方針・やるべきこと
アプリケーションのクライアントIDとkeyを暗号化すること  
docker上で動作できる環境の構築  
クラウド移行  
ポートフォリオに掲載  

#他
## Goでexeファイルを作成する
go build -o [アプリケーション名]

## Windows上で作成する場合は拡張子を付けてbuild
go build -o [アプリケーション名].exe

