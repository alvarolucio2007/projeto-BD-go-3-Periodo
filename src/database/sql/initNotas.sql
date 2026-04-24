CREATE TABLE IF NOT EXISTS notas(
  id SERIAL PRIMARY KEY,
  usuario_id INTEGER NOT NULL,
  prova_id INTEGER NOT NULL,
  nota_prova NUMERIC(3,1) CHECK(nota_prova>=0 AND nota_prova<=10),
  CONSTRAINT fk_aluno FOREIGN KEY(usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
  CONSTRAINT fk_prova FOREIGN KEY(prova_id) REFERENCES provas(id) ON DELETE CASCADE,
  CONSTRAINT nota_unica_por_aluno UNIQUE (usuario_id,prova_id)
);
