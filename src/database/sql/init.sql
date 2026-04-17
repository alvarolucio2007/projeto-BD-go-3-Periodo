DROP TYPE IF EXISTS role_usuario CASCADE;
CREATE TYPE role_usuario AS ENUM ('admin','professor','aluno');
CREATE TABLE IF NOT EXISTS usuarios(
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role role_usuario NOT NULL
);
CREATE TABLE IF NOT EXISTS provas(
  id SERIAL PRIMARY KEY,
  nome_prova VARCHAR(50) NOT NULL,
  turma_prova VARCHAR(50) NOT NULL,
  materia_prova VARCHAR(50) NOT NULL,
  data_prova TIMESTAMPTZ NOT NULL
);
CREATE TABLE IF NOT EXISTS notas(
  id SERIAL PRIMARY KEY,
  usuario_id INTEGER NOT NULL,
  prova_id INTEGER NOT NULL,
  nota_prova NUMERIC(3,1) CHECK(nota_prova>=0 AND nota_prova<=10),
  CONSTRAINT fk_aluno FOREIGN KEY(usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
  CONSTRAINT fk_prova FOREIGN KEY(prova_id) REFERENCES provas(id) ON DELETE CASCADE,
  CONSTRAINT nota_unica_por_aluno UNIQUE (usuario_id,prova_id)
);
