UPDATE provas 
	SET 
    nome_prova = COALESCE(NULLIF($1, ''), nome_prova),
    turma_prova = COALESCE(NULLIF($2, ''), turma_prova),
    materia_prova = COALESCE(NULLIF($3, ''), materia_prova),
    data_prova = COALESCE($4, data_prova)
	WHERE id = $5;
