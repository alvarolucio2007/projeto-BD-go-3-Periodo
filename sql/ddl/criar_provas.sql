
CREATE TABLE IF NOT EXISTS provas(
  id SERIAL PRIMARY KEY,
  nome_prova VARCHAR(50) NOT NULL,
  turma_prova VARCHAR(50) NOT NULL,
  materia_prova VARCHAR(50) NOT NULL,
  data_prova TIMESTAMPTZ NOT NULL
);

