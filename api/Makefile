# =================================================================
# api request (local)
# =================================================================
# 管理用API
.PHONY: api-all-request
sleep:
	@sleep 0.5

api-all-request: healthcheck sleep \
	create-user sleep \
	get-me sleep \
	create-profile sleep \
	get-profile-me sleep \
	get-products sleep \
	get-product-by-id sleep \
	get-credit-cards sleep \
	create-credit-card sleep \
	delete-credit-card sleep \
	create-reservations sleep \
	create-charge sleep \
	get-charge-histories sleep \
	get-customer sleep \
	create-customer sleep \
	delete-customer sleep \
	get-product-comments sleep \
	get-product-comment-by-id sleep \
	create-product-comment sleep \
	get-product-my-comment sleep \
	update-product-my-comment sleep \
	delete-product-my-comment sleep \
	create-like-product-comment sleep \
	delete-like-product-comment sleep \
	get-violations-comment sleep \
	delete-violations-comment

# 生存確認API
.PHONY: healthcheck
healthcheck:
	curl -i -X 'GET' \
		'http://localhost:8080/shop/v1/healthcheck' \
		-H 'accept: application/json'

# ユーザ系API
.PHONY: create-user get-me create-profile get-profile-me
create-user:
	curl -i -X 'POST' \
		'http://localhost:8080/shop/v1/users' \
		-H 'accept: application/json' \
		-d ''

get-me:
	curl -i -X 'GET' \
		'http://localhost:8080/shop/v1/users/me' \
		-H 'accept: application/json'

create-profile:
	curl -sX 'POST' \
		'http://localhost:8080/shop/v1/users/profiles' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '{"name": {"first_name": "世一","last_name": "潔","first_name_roman": "Yoichi","last_name_roman": "Isagi"}, "address": { "region": "関東", "zip_code": "150-8377", "prefecture": "東京都", "city": "渋谷区", "street": "宇田川町", "other": "15番1号"}, "birthdate": "2000-01-01", "image_url": "https://example"}' | jq .

get-profile-me:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/users/profiles/me' \
		-H 'accept: application/json' | jq .

# 商品系API
.PHONY: get-products get-product-by-id
get-products: # 商品の一覧情報取得API
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products?cursor=MjAwMDEwMDE%3D&limit=10' \
		-H 'accept: application/json' | jq .

get-product-by-id: # 商品の詳細情報取得API
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products/20001001' \
		-H 'accept: application/json' | jq .

# コメント系API
get-product-comments: # 商品に対してのコメント一覧取得API
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products/20010001/comments?cursor=NTQwMDkwMTI%3D&limit=10' \
		-H 'accept: application/json' | jq .

create-product-comment: # 商品に対してのコメント投稿API
	curl -i -sX 'POST' \
		'http://localhost:8080/shop/v1/products/20010001/comments' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '{"title": "思っていた以上に中々良い商品でした。", "content": "この商品は非常に良いです。特にデザインが素晴らしい。", "rate": 4}'

get-product-comment-by-id: # 商品に対しての詳細なコメント取得API
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products/20010001/comments/70235692' \
		-H 'accept: application/json' | jq .

get-product-my-comment: # 自身のコメント取得API
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products/20010001/users/comments/54009221' \
		-H 'accept: application/json' | jq .

update-product-my-comment: # 自身のコメント更新API
	curl -i -X 'PUT' \
		'http://localhost:8080/shop/v1/products/20010001/users/comments/54009221' \
		-H 'accept: */*' \
		-H 'Content-Type: application/json' \
		-d '{"title": "まずまず商品でした", "content": "価格が少し高いです。", "rate": 3}'

delete-product-my-comment: # 自身のコメント削除API
	curl -i -X 'DELETE' \
		'http://localhost:8080/shop/v1/products/20010001/users/comments/54009221' \
		-H 'accept: */*'

# いいね系API
create-like-product-comment: # コメントに対しての「いいね」API
	curl -i -X 'POST' \
		'http://localhost:8080/shop/v1/products/20010001/comments/54009221/like' \
		-H 'accept: */*' \
		-d ''

delete-like-product-comment: # コメントに対しての「いいね」を取り消すAPI
	curl -i -X 'DELETE' \
		'http://localhost:8080/shop/v1/products/20010001/comments/54009221/like' \
		-H 'accept: */*' \
		-d ''

# 外部決済サービス向け顧客情報系API
.PHONY: get-customer create-customer delete-customer
get-customer:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/payment/customers/10001001' \
		-H 'accept: application/json' | jq .

create-customer:
	curl -i -sX 'POST' \
		'http://localhost:8080/shop/v1/payment/customers' \
		-H 'accept: application/json' \
		-d ''

delete-customer:
	curl -i -sX 'DELETE' \
		'http://localhost:8080/shop/v1/payment/customers' \
		-H 'accept: */*'

# カード系API
.PHONY: get-credit-cards create-credit-card delete-credit-card
get-credit-cards:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/payments/cards' \
		-H 'accept: application/json' | jq .

create-credit-card:
	curl -i -sX 'POST' \
		'http://localhost:8080/shop/v1/payments/cards' \
		-H 'accept: */*' \
		-H 'Content-Type: application/json' \
		-d '{"token": "9f9213b2-9687-6910-d2ec-3a247582be2d", "is_default": true}'

delete-credit-card:
	curl -i -sX 'DELETE' \
		'http://localhost:8080/shop/v1/payments/cards' \
		-H 'accept: */*'

# 予約系API
.PHONY: create-reservations
create-reservations:
	curl -sX 'POST' \
		'http://localhost:8080/shop/v1/payments/reservations' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '[{"product_id": 10001001, "quantity": 2}, {"product_id": 10001002, "quantity": 3}]' | jq .

# 注文系API
.PHONY: create-charge get-charge-histories
create-charge:
	curl -i -sX 'POST' \
		'http://localhost:8080/shop/v1/payments/charges' \
		-H 'accept: */*' \
		-H 'Content-Type: application/json' \
		-d '{"reservation_id": "9f9213b2-9687-6910-d2ec-3a247582be2d"}'

get-charge-histories:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/payments/charges/histories?offset=0&limit=10' \
		-H 'accept: application/json' | jq .

# 管理者系API
.PHONY: get-violations-comment delete-violations-comment
get-violations-comment:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/admin/products/comments/violations?cursor=NTQwMDkwMTI%3D&limit=10' \
		-H 'accept: application/json' | jq .

delete-violations-comment:
	curl -i -X 'DELETE' \
		'http://localhost:8080/shop/v1/admin/products/comments/violations/70235692' \
		-H 'accept: */*'


# =================================================================
# open search api request
# =================================================================
# インデックス作成・削除
.PHONY: create-product-comments-index delete-product-comments-index
create-product-comments-index: # インデックス作成
	curl -i -X PUT "http://localhost:9200/product_comments" -H 'Content-Type: application/json' -d @./local/opensearch/product_comments_index.json

delete-product-comments-index: # インデックス削除
	curl -i -X DELETE "http://localhost:9200/product_comments"

# データ投入、全件カウント
.PHONY: bulk-insert-comments count-data
bulk-insert-comments: # データ投入
	curl -sX POST "http://localhost:9200/_bulk" -H "Content-Type: application/json" --data-binary @./local/opensearch/bulk_data.json

count-data: # 全件カウント
	curl -sX GET "http://localhost:9200/product_comments/_count?pretty"

# データ検索
.PHONY: search-comment-all search-comment-by-product-id search-comment-by-product-id-and-created_at-desc doc-comment-by-product-id-and-comment-id
search-comment-all:
	curl -sX GET "http://localhost:9200/product_comments/_search?size=50&pretty" | jq .

search-comment-by-product-id:
	curl -sX GET "http://localhost:9200/product_comments/_search?pretty" \
	-H 'Content-Type: application/json' \
	-d '{"query": {"bool": {"must": [{ "match": { "product_id": 20010001 } }]}}}' | jq .

search-comment-by-product-id-and-created_at-desc:
	curl -sX GET "localhost:9200/product_comments/_search?pretty" \
		-H 'Content-Type: application/json' \
		-d '{"query": {"match": {"product_id": 20010001}}, "sort": [{"created_at": {"order": "desc"}}], "size": 1}' | jq .

doc-comment-by-product-id-and-comment-id: # 商品IDとコメントIDを指定してコメントを取得（こちらの方が効率的）
	curl -sX GET "http://localhost:9200/product_comments/_doc/70235709?pretty" \
    -H 'Content-Type: application/json' | jq .

# search-comment-by-product-id-and-comment-id: # 商品IDとコメントIDを指定してコメントを取得（こちらは非効率のためコメントアウト）
# 	curl -sX GET "http://localhost:9200/product_comments/_search?pretty" \
# 	-H 'Content-Type: application/json' \
# 	-d '{"query": {"bool": {"must": [{ "match": { "product_id": 20010001 } }, { "term": { "_id": "70235692" } }]}}}' | jq .

search-ngword-comment:
	curl -sX GET "localhost:9200/product_comments/_search?pretty" \
			-H 'Content-Type: application/json' \
			-d '{"query": {"bool": {"should": [{ "match_phrase": { "content": "クズ" }}, { "match_phrase": { "content": "ゴミ" }}, { "match_phrase": { "content": "ぶっ潰す" }}, { "match_phrase": { "content": "http://" }}, { "match_phrase": { "content": "https://" }}, { "match_phrase": { "content": "最悪" }}, { "match_phrase": { "title": "クズ" }}, { "match_phrase": { "title": "ゴミ" }}, { "match_phrase": { "title": "ぶっ潰す" }}, { "match_phrase": { "title": "http://" }}, { "match_phrase": { "title": "https://" }}, { "match_phrase": { "title": "最悪" }}], "minimum_should_match": 1}}, "size": 10}' | jq .

# データ削除
.PHONY: delete-all-ops
delete-all-ops:
	curl -i -sX POST "http://localhost:9200/product_comments/_delete_by_query" \
		-H 'Content-Type: application/json' \
		-d '{"query": {"match_all": {}}}'


# =================================================================
# api request (aws)
# =================================================================
AWS_API_v1_DOMAIN := apiv1.<api-domain>

.PHONY: aws-healthcheck
aws-healthcheck:
	curl -i -X 'GET' \
		'https://${AWS_API_v1_DOMAIN}/shop/v1/healthcheck' \
		-H 'accept: application/json'

.PHONY: aws-create-user aws-get-me
aws-create-user:
	curl -i -X 'POST' \
		'https://${AWS_API_v1_DOMAIN}/shop/v1/users' \
		-H 'accept: application/json' \
		-d ''

aws-get-me:
	curl -i -X 'GET' \
		'https://${AWS_API_v1_DOMAIN}/shop/v1/users/me' \
		-H 'accept: application/json'
