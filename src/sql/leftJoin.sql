SELECT usuarios.username,notas.nota_prova FROM usuarios LEFT JOIN notas on notas.usuario_id=usuarios.id;
