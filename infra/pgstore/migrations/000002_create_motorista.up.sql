CREATE TABLE IF NOT EXISTS motorista (
    "id"                uuid            PRIMARY KEY     NOT NULL    DEFAULT gen_random_uuid(),
    "veiculo_id"        uuid,                 
    "nome"              VARCHAR(255)                    NOT NULL,    

    FOREIGN KEY (veiculo_id) REFERENCES veiculo(id)
);