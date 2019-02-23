CREATE TABLE sessions
(
  id      UUID PRIMARY KEY,
  user_id UUID NOT NULL
);

CREATE TABLE users
(
  id            UUID PRIMARY KEY,
  name          VARCHAR(100) NOT NULL,
  password      VARCHAR(100) NOT NULL,
  calories      INT          NOT NULL,
  proteins      FLOAT        NOT NULL,
  carbohydrates FLOAT        NOT NULL,
  fats          FLOAT        NOT NULL,
  salt          FLOAT        NOT NULL,
  sugar         FLOAT        NOT NULL
);
