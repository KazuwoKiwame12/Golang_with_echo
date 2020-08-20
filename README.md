# Golang_with_echo
## 本番環境
heroku: https://test-echo-with-postgres.herokuapp.com/test
## setup_プロジェクト作成からherokuでのリリースまで
1. Go Moduleで依存ライブラリの管理を行う
```bash
$ go mod init
```
2. ライブラリをinstallする
```bash
$ go get -u github.com/labstack/echo/...
```

3. Procfileの作成
```Procfile
web: bin/<project-name>
```
※go.modのmoduleの名前 = "project-name"が、モジュールのパスとして扱われる。  
結果、binの下に作成されるのは"project-name"である
  
 4. herokuでアプリを作成
 5. Githubのリポジトリ連携
 6. deploy

## 開発
※exportする関数は、関数名の最初の文字を大文字にしなければならない