SELECT id,nome_prova,turma_prova,materia_prova,data_prova FROM provas WHERE nome_prova ILIKE $1;
