SELECT usuarios.username,provas.nome_prova,notas.nota_prova FROM usuarios INNER JOIN notas ON usuarios.id=notas.usuario_id INNER JOIN provas on provas.id=notas.prova_id;
