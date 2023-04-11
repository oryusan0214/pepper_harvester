# API リファレンス
基本的にはテスト環境のためhttp通信のローカルで限られた機械しか接続しないものと考える。
インターネットに公開する場合は別途認証処理が必要だと思います。

## 農作者の情報

### GET

#### 農作者の追加
ユーザーの追加は外部サービスを利用している事もあり、ユーザーの名前とEmailをつかって識別を考えています。
- http://localhost:8080/GET/:EMAIL/:USER_NAME
  >※:NAME＝ユーザー固有ID
  　http://localhost:8000/GET/email_email@sample.jp/yamadaryunosuke
  ```json
  json[
    {
        "ID":"1"
    },
    ]
  ```

#### 農作者が所有している農耕地と農作物と機械の情報

- http://localhost:8080/GET/agrist/:ID
  >※:NAME＝ユーザー固有ID
  　http://localhost:8000/agrist/yamadaryunosuke
  ```json
  json[
    {
        "UserID":"1",
        "Name": "農耕地A",
        "CropsName": "ピーマン",
        "MachineNum": "123455678"
    },
    {
        "UserID":"1",
        "Name": "農耕地A",
        "CropsName": "ピーマン",
        "MachineNum": "000000000"
    }
        ・
        ・
        ・
    ]
  ```
  * return:情報がない場合はNULLで返る
  * 注意:データがない場合は詰められる

#### 農作者が所有している農耕地の温度と湿度情報(取得した日付のみ)

- http://localhost:8080/GET/temphumi/:NAME
  >※:NAME＝農耕地の名前
  　http://localhost:8080/temphumi/農耕地A
  ```json
  json[
    {
        "temp":"20",
        "humi":"40"
    }
  ]
  ```
  * return:情報がない場合はNULLで返る

#### 農作者が所有している農耕地の温度・湿度・収穫写真データ(日ごとに指定)

- http://localhost:8080/GET/farmland/:NAME/DATE/:DAY/HARVEST
  >※:NAME＝農耕地の名前
  ※:DAY=日付データ
  　http://localhost:8080/farmland/農耕地A/DATE/2023-01-17/HARVEST
  ```json
  json[
    {
        "url":"https://???????/??????.jpg",
    }
        ・
        ・
        ・
  ]
  ```

  * return:情報がない場合はNULLで返る
  * 注意:データがない場合は詰められる


#### 農作者が所有している農耕地の温度・湿度・収穫写真データ(週ごと)

- http://localhost:8080/GET/farmland/:NAME/WEEK/:DAY/HARVEST
  >※:NAME＝農耕地の名前
  ※:DAY=日付データ
  　http://localhost:8080/farmland/農耕地A/WEEK/2023-01-17/HARVEST
  ```json
  json[
    {
        "day":"2023-01-17",
        "url":"https://???????/??????.jpg",
        "temp":"20",
        "humi":"40"
    }
    {
        "day":"2023-01-18",
        "url":"https://???????/??????.jpg",
        "temp":"21",
        "humi":"41"
    }
        ・
        ・
        ・
  ]
  ```
  * return:情報がない場合はNULLで返る
  * 注意:データがない場合は詰められる

#### 農作者が所有している農耕地の温度・湿度・収獲写真データ(月ごと)

- http://localhost:8080/GET/farmland/:NAME/MONTH/:DAY/HARVEST
  >※:NAME＝農耕地の名前
  ※:DAY=日付データ
  　http://localhost:8080/farmland/農耕地A/MONTH/2023-01-17/HARVEST
  ```json
  json[
    {
        "day":"2023-01-17",
        "url":"https://???????/??????.jpg",
        "temp":"20",
        "humi":"40"
    }
    {
        "day":"2023-01-18",
        "url":"https://???????/??????.jpg",
        "temp":"21",
        "humi":"41"
    }
        ・
        ・
        ・
  ]
  ```
  * return:情報がない場合はNULLで返る
  * 注意:データがない場合は詰められる
#### 農作者が所有している農耕地の温度・湿度・未収穫写真データ(日ごとに指定)

- http://localhost:8080/GET/farmland/:NAME/DATE/:DAY/UNHARVEST
  >※:NAME＝農耕地の名前
  ※:DAY=日付データ
  　http://localhost:8080/farmland/農耕地A/DATE/2023-01-17/UNHARVEST
  ```json
  json[
    {
        "url":"https://???????/??????.jpg",
    }
        ・
        ・
        ・
  ]
  ```

  * return:情報がない場合はNULLで返る
  * 注意:データがない場合は詰められる


#### 農作者が所有している農耕地の温度・湿度・未収穫写真データ(週ごと)

- http://localhost:8080/GET/farmland/:NAME/WEEK/:DAY/UNHARVEST
  >※:NAME＝農耕地の名前
  ※:DAY=日付データ
  　http://localhost:8080/farmland/農耕地A/WEEK/2023-01-17/UNHARVEST
  ```json
  json[
    {
        "day":"2023-01-17",
        "url":"https://???????/??????.jpg",
        "temp":"20",
        "humi":"40"
    }
    {
        "day":"2023-01-18",
        "url":"https://???????/??????.jpg",
        "temp":"21",
        "humi":"41"
    }
        ・
        ・
        ・
  ]
  ```
  * return:情報がない場合はNULLで返る
  * 注意:データがない場合は詰められる

#### 農作者が所有している農耕地の温度・湿度・未収獲写真データ(月ごと)

- http://localhost:8080/GET/farmland/:NAME/MONTH/:DAY/UNHARVEST
  >※:NAME＝農耕地の名前
  ※:DAY=日付データ
  　http://localhost:8080/GET/farmland/農耕地A/MONTH/2023-01-17/UNHARVEST
  ```json
  json[
    {
        "day":"2023-01-17",
        "url":"https://???????/??????.jpg",
        "temp":"20",
        "humi":"40"
    }
    {
        "day":"2023-01-18",
        "url":"https://???????/??????.jpg",
        "temp":"21",
        "humi":"41"
    }
        ・
        ・
        ・
  ]
  ```
  * return:情報がない場合はNULLで返る
  * 注意:データがない場合は詰められる

  

#### 農作者の情報

### PUT(更新)
#### 農作者が所有している農耕地と農作物と機械の情報
ユーザーの登録されている機械、農耕地、収穫物の情報を更新する。

- http://localhost:8080/PUT/agrist/:ID
  >※:NAME＝ユーザーの固有ID  
    　http://localhost:8000/PUT/agrist/1
  ```json
  json[
    {
        "Name": "農耕地A",
        "CropsName": "ピーマン",
        "MachineNum": "123455678"
    },
    {
        "Name": "農耕地A",
        "CropsName": "ピーマン",
        "MachineNum": "000000000"
    }
        ・
        ・
        ・
    ]
  ```
### POST(追加)
#### 農作者の追加
ユーザーの追加は外部サービスを利用している事もあり、ユーザーの名前とEmailをつかって識別を考えています。
- http://localhost:8080/POST/user
  >※:NAME＝ユーザー固有ID  
    　http://localhost:8000/POST/user
  ```json
  json[
    {
        "UserName":"yamadaryunosuke",
        "Email": "Email@sample.jp",
    },
    ]
  ```

#### 農作者が所有している農耕地と農作物と機械の情報
ユーザーの登録されている機械、農耕地、収穫物の情報を更新する。

- http://localhost:8080/POST/agrist/:ID
  >※:NAME＝ユーザーの固有ID  
    　http://localhost:8000/POST/agrist/1
  ```json
  json[
    {
        "Name": "農耕地A",
        "CropsName": "ピーマン",
        "MachineNum": "123455678"
    },
    {
        "Name": "農耕地A",
        "CropsName": "ピーマン",
        "MachineNum": "000000000"
    }
        ・
        ・
        ・
    ]
  ```

#### 農作者が所有している農耕地の温度と湿度情報

- http://localhost:8080/POST/temphumi/:ID
  >※:ID=機械number  
    　http://localhost:8080/POST/temphumi/0000a0000-b
  ```json
  json[
    {
        "temp":"20",
        "humi":"40"
    }
  ]
  ```
  * return:情報がない場合はNULLで返る

#### 農作者が所有している農耕地の温度・湿度・収穫写真データ

- http://localhost:8080/POST/HARVEST/:ID
  >※:NAME＝農耕地の名前
  ※:DAY=日付データ
  　http://localhost:8080/POST/HARVEST/0000a0000-b
  ```json
  json[
    {   
        "amount":10,
        "url":"https://???????/??????.jpg",
        "temp":"20",
        "humi":"40"
    }
        ・
        ・
        ・
  ]
  ```

#### 農作者が所有している農耕地の温度・湿度・未収穫写真データ

- http://localhost:8080/POST/UNHARVEST/:ID
  >※:NAME＝農耕地の名前
  ※:DAY=日付データ
  　http://localhost:8080/POST/UNHARVEST/0000a0000-b
  ```json
  json[
    {   
        "amount":10,
        "url":"https://???????/??????.jpg",
        "temp":"20",
        "humi":"40"
    }
        ・
        ・
        ・
  ]
  ```

### Delete

***開発中***
