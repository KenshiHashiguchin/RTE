# RTE(TrustReserve)

# 目次
- [プロダクト概要](#プロダクト概要)
- [事前準備](#事前準備)
- [サービスログイン](#サービスログイン)
- [tips](#tips)

# プロダクト概要
TrustReserveは「予約通り飲食店に来店すること」にWeb3の技術を用いて経済的インセンティブをつけることで、飲食店の無断キャンセル減少を目指す決済サービスである。
## サービスの特徴
*①来店予約/決済に対するトークン付与*
- 予約者は、ウォレットから事業者が指定する額のデポジットを支払うことで飲食店の来店予約ができる。
- 事前の予約通りに来店してサービス内で決済を行うことで予約履行とみなされ、予約者のウォレットにトークン（TRES）が付与される。
  - TRESは市場で売買できる他、TrustReserveでの支払い手段として利用ができる。
- 来店して決済しなかった場合、無断キャンセル（=予約不履行）とみなされてデポジットされたトークンが事業者側へ支払われる。

*②SBTを使った予約者の信頼度の蓄積*
- 予約者は、サービス初回利用時にSBTを獲得し、来店実績に応じて予約者の信頼度が蓄積される。
  - 予約通りの来店をするほど信頼度が上がり、無断キャンセルを行うと信頼度が下がる。
- 事業者が信頼度に応じて予約者を優遇したり、特別なプランを用意するなどの利用を想定。
  - 将来的には信頼度に応じて1予約で付与されるTRESを増やすことも可能。

# 事前準備
## Mac
*1. docker for mac のインストール*  
https://docs.docker.com/desktop/install/mac-install/

*2. 環境変数の設定*
```bash
cp .env.example .env
cp web/api/.env.example web/api/.env
cp web/frontend/.env.example web/frontend/.env
```  
*3. コンテナの作成、起動*
- 下記いずれかを実行
```bash
# foreground
docker-compose up
```

```bash
# background
docker-compose up -d 
```

*4. コンテナ起動確認*
- 下記を実行してdocker-compose.yml定義のコンテナが起動していればOK
```bash
docker ps
```

*5. /etc/hostsの変更*

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

*6. api起動*
```bash
docker exec -it rte-golang bash
go mod tidy
go run main/*
```
起動確認はブラウザで[http://localhost/api/token/0xF272453bf221D5e09374aDa6607Fc8d834E2d25b](http://localhost/api/token/0xF272453bf221D5e09374aDa6607Fc8d834E2d25b)にアクセスできればOK
<br>（証明書エラーが起こる場合は[http://localhost/api/token/0xF272453bf221D5e09374aDa6607Fc8d834E2d25b](http://localhost/api/token/0xF272453bf221D5e09374aDa6607Fc8d834E2d25b)にアクセスして確認）

*7. フロントエンド build&nuxt server起動*
 ※nuxtはユーザ側、事業者側でそれぞれ起動する構成としているので以下必要に応じて実施  
**ユーザ側の起動**
```
# nuxtコンテナに接続
docker exec -it rte-user-nuxt ash
yarn install
yarn dev-user
```
起動確認はブラウザで[https://rte-local.com](https://rte-local.com)にアクセスできればOK
<br>(証明書エラーが起こる場合は[http://rte-local.com](http://rte-local.com)にアクセスして確認)

**事業者側の起動**
```
# nuxtコンテナに接続
docker exec -it rte-merchant-nuxt ash
yarn install
yarn dev-merchant
```
起動確認はブラウザで[https://merchant.rte-local.com](https://merchant.rte-local.com)にアクセスできればOK
<br>(証明書エラーが起こる場合は[http://merchant.rte-local.com](http://merchant.rte-local.com)にアクセスして確認)

*8. DBスキーマ更新方法*
```shell
 docker-compose exec ruby bundle exec ridgepole -a -c /db/config/database.yml -f /db/SchemaFile.rb -E default
```
## Windows
# サービスログイン

## ユーザー側
 **ログイン**
 <br>[https://rte-local.com/admin](https://rte-local.com/admin)
 <br>(証明書エラーが起こる場合は[http://rte-local.com/admin](http://rte-local.com/admin))

 **サイトトップ**
 <br>[https://rte-local.com/](https://rte-local.com/)
 <br>(証明書エラーが起こる場合は[http://rte-local.com/](http://rte-local.com/))
## 事業者側
 [https://merchant.rte-local.com/](https://merchant.rte-local.com/)
 <br>(証明書エラーが起こる場合は[http://merchant.rte-local.com/](http://merchant.rte-local.com/))

# tips
### 証明書エラーが出る

---
ブラウザで証明書関連の警告が出た場合以下参照し、自PCに証明書を信頼させる
<br>https://www.curict.com/item/2a/2ace791.html
