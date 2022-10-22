# RTE

# 事前準備
## Mac
docker for mac のインストール  
https://docs.docker.com/desktop/install/mac-install/

# setup
- 環境変数の設定
```bash
cp .env.example .env
cp web/api/.env.example web/api/.env
```  
- コンテナの作成、起動（下記いずれか）
```bash
# foreground
docker-compose up
# background
docker-compose up -d 
```

- 起動確認
```bash
docker ps
```
でdocker-compose.yml定義のコンテナが起動していればOK

- /etc/hostsの変更
```bash
sudo vi /etc/hosts

>>>>>
# tokyoweb3hackathon
127.0.0.1 rte-local.com 
127.0.0.1 merchant.rte-local.com
>>>>>
```
※rte-local.comがユーザドメイン <br>
※merchant.rte-local.comが事業者ドメインとなる

## api起動
```bash
docker exec -it rte-golang bash
go run main/*
```
起動確認はブラウザで http://localhost/api/token/0xF272453bf221D5e09374aDa6607Fc8d834E2d25b にアクセスできればOK

## フロントエンド build&nuxt server起動  
※nuxtはユーザ側、事業者側でそれぞれ起動する構成としているので以下必要に応じて実施  
**ユーザ側起動する場合**
```
# nuxtコンテナに接続
docker exec -it rte-user-nuxt ash
yarn dev-user
```
起動確認はブラウザで https://rte-local.com にアクセスできればOK

**事業者側起動する場合**
```
# nuxtコンテナに接続
docker exec -it rte-merchant-nuxt ash
yarn dev-merchant
```
起動確認はブラウザで https://merchant.rte-local.com にアクセスできればOK

## DBスキーマ更新方法
```shell
 docker-compose exec ruby bundle exec ridgepole -a -c /db/config/database.yml -f /db/SchemaFile.rb -E default
```


## tips
### 証明書エラーが出る

---
ブラウザで証明書関連の警告が出た場合以下参照し、自PCに証明書を信頼させる
https://www.curict.com/item/2a/2ace791.html

### DBスキーマを更新する

---
```shell
 docker-compose exec ruby bundle exec ridgepole -a -c /db/config/database.yml -f /db/SchemaFile.rb -E default
```

