#ピーマン収獲機 ベルペット
### 企画趣旨
本プロジェクトは農家の負担をへらすことを目的として作られた。

### フォルダー構成
<pre>
.
├── .vscode
│   └── settings.json
├── API資料
│   └── API.md
├── DB資料
│   ├── ~$AWS構成案.pptm
│   ├── ER図-Ver0.32.drawio.png
│   ├── ER図.drawio
│   ├── テーブル設計 (version 1).xlsx
│   └── テーブル設計.xlsx
├── docker_image
│   ├── db
│   ├── golang
│   │   ├── controller
│   │   ├── drivers
│   │   ├── entities
│   │   │   └── entities.go
│   │   ├── gateways
│   │   │   ├── gateways_interface.go
│   │   │   └── gateways.go
│   │   ├── get
│   │   │   ├── getFarmland.go
│   │   │   └── getUser.go
│   │   ├── dockerfil
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── main.go
│   ├── mysql
│   │   ├── conf.d
│   │   │   └── my.cnf
│   │   ├── sql
│   │   │   └── init.sql
│   │   └── dockerfile
│   ├── .env
│   └── docker-compose.yml
├── flutter
│   └── pepper_harvester
├── システム資料
│   ├── AWS.drawio
│   ├── AWS構成案.pptm
│   ├── システム構成図Ver0.2-Ver0.2　オンプレ.drawio.png
│   ├── システム構成図Ver0.2.drawio
│   ├── ユースケース図-Ver0.2.drawio.png
│   └── ユースケース図.drawio
├── README.md
├── 企画.md
└── 要件定義.md
</pre>

#### フォルダー説明
「～資料」フォルダーには設計の時に作成した資料が入っています。
- [API資料](API資料)  
  APIの仕様について書いた資料があります。
- [DB資料](DB資料)  
  ER図やテーブル導出した時のExcel資料があります。
  ![](DB資料\ER図-Ver0.34.drawio.png)
- [システム資料](システム資料)  
  システムの構成図やAWSに移行した場合の構成図などを考えたときの資料があります。
  ![](システム資料\システム構成図Ver0.2-Ver0.2　オンプレ.drawio.png)
- [docker_image](docker_image)  
  dockerの設定ファイルがあります。
    * golang  
    APIサーバー用のGOのソースコードがあります。
    * mysql  
      mysqlの設定ファイルと初期化ファイルがあります。
- [flutter](flutter)
  * pepper_harvester  
    flutterのアプリケーションのコードがあります。
### チーム

  メカ：1人  
  ハード：2人  
  ソフトウェア：1人(設計フェーズ　+1)  
  画像処理：1人  
  システム＋アプリケーション：1人  (担当)
  
    計5+1名


### 要件定義
ピーマン収穫機の要件定義は他のGoogleDriveの書類を参考にしてください。
要件定義については「要件定義.md」に記載しています。
そちらをご覧ください。

### 開発環境
#### 開発資料資料
vscode:
バージョン: 1.74.3 (user setup)
OS: Windows_NT x64 10.0.19044
#### RDBMS
MYSQL:8.0.31

#### API
go:1.19.4

#### アプリケーション
fluter:flutter 3.3.9





### 参考にしたサイト  
#### docker(test環境)  
<br>
dockerについての概要  
- https://matsuand.github.io/docs.docker.jp.onthefly/reference/  
<br>
docker compose についての参考サイト  
- https://docs.docker.jp/compose/compose-file/index.html?  
<br>
.cnfファイルのエラー（読み取り専用にすることで読み込み可）
- https://qiita.com/rabbitbeef/items/14433a2c0a6f85c3b476
<br>
mysql(test環境)(参考)  
 - https://jeannie-blog.online/docker-mysql-windows/  
docker mysql 初期ログインの手引き  
 - https://prograshi.com/platform/docker/mysql-on-docker/  
docker-composed(参考)  
 - https://zenn.dev/re24_1986/articles/153cdc5db96dc0  
 composeに置ける環境変数の設定  
 - https://matsuand.github.io/docs.docker.jp.onthefly/compose/environment-variables/  
逆引きSQL（参考）  
 - https://style.potepan.com/articles/25967.html  
 - https://www.sql-reference.com/  
docker-compoe volumesのマウントのされ方について  
 - https://qiita.com/daemokra/items/322270091cf41a853226  
go  
GORMについて  
 - https://gorm.io/ja_JP/docs/connecting_to_the_database.html  
http package  
 - https://pkg.go.dev/net/http  
Ginについて  
公式  
 - https://pkg.go.dev/github.com/gin-gonic/gin  
参考記事  
 - https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a  
ginのcontextについての参考記事  
 - https://ipeblog.com/golang-gin-json/304/  
 構造体の非公開フィールドについて  
- https://stop-the-world.hatenablog.com/entry/2019/12/31/214058  
設計思想  
- https://zenn.dev/nakaryo/articles/87b15866d67efe  