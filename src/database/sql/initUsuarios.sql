DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_usuario') THEN
    CREATE TYPE role_usuario AS ENUM ('admin','professor','aluno');
  END IF;
END $$;

CREATE TABLE IF NOT EXISTS usuarios(
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role role_usuario NOT NULL
);

