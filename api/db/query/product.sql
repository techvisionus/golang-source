-- name: ListProduct :many
SELECT * FROM products AS P
WHERE P.code = $1 OR LOWER(P.name) LIKE '%' || $2 || '%'
ORDER BY P.id DESC
OFFSET $3
LIMIT $4;

-- name: CreateProduct :one
INSERT INTO products (
  code, name, price
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET
code = $1,
name = $2,
price = $3
WHERE id = $4
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1;

-- name: GetProductByCode :one
SELECT * FROM products
WHERE code = $1;

-- name: DeleteAuthor :exec
DELETE FROM products WHERE id = $1;
