CREATE TABLE IF NOT EXISTS veiculo (
    "id"    uuid            PRIMARY KEY     NOT NULL    DEFAULT gen_random_uuid(),
    "placa" VARCHAR(7)                    NOT NULL
);