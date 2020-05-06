# Simple Chat Client Go について

## VisualStudioCodeのデバッガ  

VisualStudioCodeでGoを使うときはデバッガのコマンドを忘れずに叩く  
go get -u github.com/derekparker/delve/cmd/dlv  
  
  
## Chatの認証情報について  

mainFunctionにはClinetIDとkeyを入れる必要がある。  
Google Developer Consoleでこのアプリケーションの認証情報を登録して  
クライアントIDと秘密鍵(key)を取得しないといけない。  

## 今後の方針・やるべきこと
アプリケーションのクライアントIDとkeyを暗号化すること  
docker上で動作できる環境の構築  
クラウド移行  
ポートフォリオに掲載  
