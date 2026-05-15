package database

import (
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func LerQuantidadeProvaAluno(nomeBusca string) (map[string]int64, error) { // Essa é para um gráfico de barras
	query := `SELECT 
		usuarios.username,COUNT(notas.id) as total_provas 
		FROM usuarios u
		LEFT JOIN notas ON u.id=notas.usuario_id 
		WHERE u.role='aluno' AND u.username ILIKE $1
		GROUP BY u.username 
		ORDER BY total_provas DESC;`
	rows, err := DB.Query(query, "%"+nomeBusca+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stats := make(map[string]int64)
	for rows.Next() {
		var (
			user string
			qtd  int
		)
		if err := rows.Scan(&user, &qtd); err != nil {
			return nil, err
		}
		stats[user] = int64(qtd)
	}
	return stats, nil
}

func LerQuantidadeNotaProvaAluno(nomeBusca string) (map[string]models.EstatisticaAluno, error) { // Essa função é prum gráfico de dispersão
	query := `SELECT 
    u.username,
    COUNT(n.id) AS total_provas,
    COALESCE(AVG(n.nota), 0.0) AS media 
		FROM usuarios u
		LEFT JOIN notas n ON u.id = n.usuario_id
		LEFT JOIN provas p ON p.id = n.prova_id
		WHERE u.role = 'aluno' AND u.username ILIKE $1
		GROUP BY u.username 
		ORDER BY total_provas ASC;`
	rows, err := DB.Query(query, "%"+nomeBusca+"%")
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

func LerMediaNotaMateria(nomeCategoria string) (map[string]models.EstatisticaAluno, error) { // sim estou reutilizando o models.EstatisticaAluno, se cabe direito consigo...
	// ah e também é gráfico de barras horizontal
	query := `SELECT 
    p.materia_prova,
    COALESCE(AVG(n.nota), 0.0) AS media_materia,
		COUNT(n.id) AS total_provas
		FROM provas p
		LEFT JOIN notas n ON p.id = n.prova_id
		WHERE p.materia_prova ILIKE $1
		GROUP BY p.materia_prova
		ORDER BY media_materia ASC;`
	rows, err := DB.Query(query, "%"+nomeCategoria+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]models.EstatisticaAluno)
	for rows.Next() {
		var (
			nomeMateria  string
			mediaMateria float64
			totalProvas  int
		)
		if err := rows.Scan(&nomeMateria, &mediaMateria, &totalProvas); err != nil {
			return nil, err
		}
		result[nomeMateria] = models.EstatisticaAluno{QuantidadeProva: totalProvas, MediaProvas: mediaMateria}
	}
	return result, nil
}

func LerDistribuicaoStatusAluno() (map[string]int64, error) { // Gráfico de pizza, bem básico...
	query := `
	SELECT
		CASE
			WHEN nota>=7 THEN 'Aprovado'
			WHEN nota>=5 THEN 'Recuperação'
			ELSE 'Reprovado'
		END AS status,
		COUNT(*)
	FROM notas
	GROUP BY status;`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]int64)
	for rows.Next() {
		var (
			status string
			valor  int64
		)
		if err := rows.Scan(&status, &valor); err != nil {
			return nil, err
		}
		result[status] = valor
	}
	return result, nil
}
