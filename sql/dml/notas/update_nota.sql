UPDATE notas SET 
    usuario_id = COALESCE(NULLIF($1, 0), usuario_id),
    prova_id = COALESCE(NULLIF($2, 0), prova_id),
    nota_prova = COALESCE(NULLIF($3, -1), nota_prova) -- Use -1 se 0 for nota válida
WHERE id = $4;
