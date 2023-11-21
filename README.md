# README
- QRオーダーシステムのインフラソース\
*qr-order-apiのみでは動きません*

## 環境構築
### 階層構造
- 下記ツリーのように同一階層に３つのソースを配置してください。
```
任意のディレクトリ
    ├── qr-order-infra
    ├── qr-order-ui
    └── qr-order-api
```
### 初回dokcerコンテナ起動
```
docker-compose up -d
```