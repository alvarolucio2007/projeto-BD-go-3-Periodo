package database

func LerQuantidadeProvaAluno() (map[string]int, error) {
	query := "SELECT usuarios.username,COUNT(notas.id) as total_provas FROM usuarios LEFT JOIN notas ON usuarios.id=notas.usuario_id WHERE usuarios.role='aluno' GROUP BY usuarios.username ORDER BY total_provas DESC;"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stats := make(map[string]int)
	for rows.Next() {
		var (
			user string
			qtd  int
		)
		if err := rows.Scan(&user, &qtd); err != nil {
			return nil, err
		}
		stats[user] = qtd
	}
	return stats, nil
}
