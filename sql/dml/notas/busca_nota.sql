SELECT u.username, p.nome_prova, n.nota_prova, p.data_prova
        FROM notas n
        INNER JOIN usuarios u ON u.id = n.usuario_id
        INNER JOIN provas p ON p.id = n.prova_id;
