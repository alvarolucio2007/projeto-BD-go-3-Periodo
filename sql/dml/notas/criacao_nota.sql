INSERT INTO notas (usuario_id,prova_id,nota_prova) SELECT $1,$2,$3 FROM usuarios WHERE id=$1 AND role='aluno' RETURNING id;
