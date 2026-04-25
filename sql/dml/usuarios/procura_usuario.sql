SELECT id, username,password, role FROM usuarios WHERE username ILIKE $1;
