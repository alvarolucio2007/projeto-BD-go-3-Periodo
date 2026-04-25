
SELECT usuarios.username,provas.nome_prova,notas.nota_prova,provas.data_prova
FROM usuarios 
LEFT JOIN notas on notas.usuario_id=usuarios.id 
LEFT JOIN provas on provas.id=notas.prova_id;
