-- name: GetVeiculo :one
SELECT
    "id", "placa"
FROM veiculo
WHERE id = $1;

-- name: InsertVeiculo :one
INSERT INTO veiculo
    ( "placa" ) VALUES
    ( $1 )
RETURNING "id";

-- name: UpdateVeiculo :one
UPDATE veiculo
SET
    placa = $1
WHERE
    id = $2
RETURNING "id";

-- name: DeleteVeiculo :exec
DELETE FROM veiculo
WHERE id = $1;

-- name: GetMotorista :one
SELECT
    "id", "nome", "veiculo_id"
FROM motorista
WHERE
    id = $1;

-- name: GetMotoristaByVeiculo :many
SELECT
    "id", "nome", "veiculo_id"
FROM motorista
WHERE
    veiculo_id = $1;    

-- name: InsertMotorista :one
INSERT INTO motorista
    ( "nome" ) VALUES
    ( $1 )    
RETURNING "id";

-- name: UpdateMotorista :one
UPDATE motorista
SET
    nome = $1
WHERE
    id = $2
RETURNING "id";

-- name: SetMotoristaVeiculo :one
UPDATE motorista
SET
    veiculo_id = $1
WHERE
    id = $2
RETURNING "id";

-- name: DeleteMotorista :exec
DELETE FROM motorista
WHERE id = $1;