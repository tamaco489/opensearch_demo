# opensearch_demo

1. 環境立ち上げ（APIサーバ、OpenSearchを起動）
```bash
$ cd api/shop

$ make up
docker compose up -d shop-api opensearch opensearch-dashboards

$ docker ps -a
CONTAINER ID   IMAGE                                            COMMAND                  CREATED              STATUS          PORTS                                                                                                      NAMES
b9c9b7d98c47   opensearchproject/opensearch-dashboards:2.19.0   "./opensearch-dashbo…"   About a minute ago   Up 59 seconds   0.0.0.0:5601->5601/tcp, :::5601->5601/tcp                                                                  opensearch-dashboards
795fc7dfa735   opensearchproject/opensearch:2.19.0              "./opensearch-docker…"   About a minute ago   Up 59 seconds   0.0.0.0:9200->9200/tcp, :::9200->9200/tcp, 9300/tcp, 0.0.0.0:9600->9600/tcp, :::9600->9600/tcp, 9650/tcp   opensearch
e169fb4f9f3f   cosmtrek/air:v1.61.0                             "/go/bin/air"            About a minute ago   Up 59 seconds   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp                                                                  shop-api

$ make logs
docker compose logs -f shop-api
```

2. OpenSearch インデックス作成、テストデータ投入、商品ID検索
```bash
$ cd api

# インデックス作成
$ make create-product-comments-index

# テストデータ投入
$ make bulk-insert-comments

# 商品ID検索
$ make search-comment-by-product-id
curl -sX GET "http://localhost:9200/product_comments/_search?pretty" \
-H 'Content-Type: application/json' \
-d '{"query": {"bool": {"must": [{ "match": { "product_id": 20010001 } }]}}}' | jq .
{
  "took": 12,
  "timed_out": false,
  "_shards": {
    "total": 1,
    "successful": 1,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 68,
      "relation": "eq"
    },
    "max_score": 1,
    "hits": [
      {
        "_index": "product_comments",
        "_id": "70235594",
        "_score": 1,
        "_source": {
          "product_id": 20010001,
          "user_id": 10010004,
          "title": "良いですね",
          "content": "価格と性能のバランスが良い",
          "rate": 4,
          "created_at": "2025-02-16 15:00:00"
        }
      },

      ... 省略

      {
        "_index": "product_comments",
        "_id": "70235603",
        "_score": 1,
        "_source": {
          "product_id": 20010001,
          "user_id": 10010013,
          "title": "素晴らしい",
          "content": "とても満足しています",
          "rate": 5,
          "created_at": "2025-02-16 15:45:00"
        }
      }
    ]
  }
}
```

3. コメント登録、登録したコメントの検索
```bash
# コメント登録
$ make create-product-comment
curl -i -sX 'POST' \
        'http://localhost:8080/shop/v1/products/20010001/comments' \
        -H 'accept: application/json' \
        -H 'Content-Type: application/json' \
        -d '{"title": "思っていた以上に中々良い商品でした。", "content": "この商品は非常に良いです。特にデザインが素晴らしい。", "rate": 4}'
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sat, 22 Feb 2025 19:09:38 GMT
Content-Length: 16

{"id":70235692}

# 登録したコメントの検索（product_idは適宜変更）
$ make search-comment-by-product-id-and-comment-id
curl -sX GET "http://localhost:9200/product_comments/_search?pretty" \
-H 'Content-Type: application/json' \
-d '{"query": {"bool": {"must": [{ "match": { "product_id": 20010001 } }, { "term": { "_id": "70235692" } }]}}}' | jq .
{
  "took": 7,
  "timed_out": false,
  "_shards": {
    "total": 1,
    "successful": 1,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 1,
      "relation": "eq"
    },
    "max_score": 2,
    "hits": [
      {
        "_index": "product_comments",
        "_id": "70235692",
        "_score": 2,
        "_source": {
          "id": 70235692,
          "product_id": 20010001,
          "user_id": 25540992,
          "title": "思っていた以上に中々良い商品でした。",
          "content": "この商品は非常に良いです。特にデザインが素晴らしい。",
          "rate": 4,
          "created_at": "2025-02-23 04:09:38"
        }
      }
    ]
  }
}
```
