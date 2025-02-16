# =================================================================
# API Request
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
	get-charge-histories \
	get-customer \
	create-customer \
	delete-customer \
	get-product-comments \
	create-product-comment \
	get-product-my-comment \
	update-product-my-comment \
	delete-product-my-comment \
	create-like-product-comment \
	delete-like-product-comment

# 生存確認API
.PHONY: healthcheck
healthcheck:
	curl -i -X 'GET' \
		'http://localhost:8080/shop/healthcheck' \
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
get-products:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products?cursor=MjAwMDEwMDE%3D&limit=10' \
		-H 'accept: application/json' | jq .

get-product-by-id:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products/20001001' \
		-H 'accept: application/json' | jq .

# コメント系API
get-product-comments:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products/10001001/comments?cursor=NTQwMDkwMTI%3D&limit=10' \
		-H 'accept: application/json' | jq .

create-product-comment:
	curl -i -sX 'POST' \
		'http://localhost:8080/shop/v1/products/10001001/comments' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '{"title": "とても良い商品です", "content": "この商品は非常に良いです。特にデザインが素晴らしい。", "rate": 4}'

get-product-my-comment:
	curl -sX 'GET' \ \
	curl -sX 'GET' \ \
		'http://localhost:8080/shop/v1/products/10001001/users/comments/54009221' \
		-H 'accept: application/json' | jq .

update-product-my-comment:
	curl -i -X 'PUT' \
		'http://localhost:8080/shop/v1/products/10001001/users/comments/54009221' \
		-H 'accept: */*' \
		-H 'Content-Type: application/json' \
		-d '{"title": "まずまず商品でした", "content": "価格が少し高いです。", "rate": 3}'

delete-product-my-comment:
	curl -i -X 'DELETE' \
		'http://localhost:8080/shop/v1/products/10001001/users/comments/54009221' \
		-H 'accept: */*'

# いいね系API
create-like-product-comment:
	curl -i -X 'POST' \
		'http://localhost:8080/shop/v1/products/10001001/comments/54009221/like' \
		-H 'accept: */*' \
		-d ''

delete-like-product-comment:
	curl -i -X 'DELETE' \
		'http://localhost:8080/shop/v1/products/10001001/comments/54009221/like' \
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
