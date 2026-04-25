UPDATE usuarios 
  SET 
      username = COALESCE(NULLIF($1, ''), username),
      password = COALESCE(NULLIF($2, ''), password),
      role     = COALESCE(NULLIF($3, '')::role_usuario, role)
  WHERE id = $4;
