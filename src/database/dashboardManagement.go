package database

import "github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"

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

func LerQuantidadeNotaProvaAluno() (map[string]models.EstatisticaAluno, error) {
	query := `SELECT 
    u.username,
    COUNT(n.id) AS total_provas,
    COALESCE(AVG(n.nota), 0.0) AS media 
		FROM usuarios u
		LEFT JOIN notas n ON u.id = n.usuario_id
		LEFT JOIN provas p ON p.id = n.prova_id
		WHERE u.role = 'aluno' 
		GROUP BY u.username 
		ORDER BY total_provas ASC;`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]models.EstatisticaAluno)
	for rows.Next() {
		var (
			username   string
			quantidade int
			media      float64
		)
		if err := rows.Scan(&username, &quantidade, &media); err != nil {
			return nil, err
		}
		result[username] = models.EstatisticaAluno{QuantidadeProva: quantidade, MediaProvas: media}
	}
	return result, nil
}
