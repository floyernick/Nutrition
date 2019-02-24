CREATE TABLE sessions
(
  id      UUID PRIMARY KEY,
  user_id UUID NOT NULL
);

CREATE TABLE users
(
  id            UUID PRIMARY KEY,
  name          VARCHAR(50) UNIQUE NOT NULL,
  password      VARCHAR(100)       NOT NULL,
  calories      INT                NOT NULL,
  proteins      FLOAT              NOT NULL,
  carbohydrates FLOAT              NOT NULL,
  fats          FLOAT              NOT NULL,
  salt          FLOAT              NOT NULL,
  sugar         FLOAT              NOT NULL
);

CREATE TABLE products
(
  id            UUID PRIMARY KEY,
  name          VARCHAR(150) NOT NULL,
  image_url     VARCHAR(250) NOT NULL,
  description   TEXT         NOT NULL,
  calories      INT          NOT NULL,
  proteins      FLOAT        NOT NULL,
  carbohydrates FLOAT        NOT NULL,
  fats          FLOAT        NOT NULL,
  salt          FLOAT        NOT NULL,
  sugar         FLOAT        NOT NULL
);

CREATE TABLE records
(
  id            UUID PRIMARY KEY,
  user_id       UUID         NOT NULL,
  name          VARCHAR(150) NOT NULL,
  calories      INT          NOT NULL,
  proteins      FLOAT        NOT NULL,
  carbohydrates FLOAT        NOT NULL,
  fats          FLOAT        NOT NULL,
  salt          FLOAT        NOT NULL,
  sugar         FLOAT        NOT NULL,
  created       TIMESTAMP    NOT NULL
);