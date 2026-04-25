INSERT INTO usuarios (username,password,role) VALUES ($1,$2,$3::role_usuario) RETURNING id;
