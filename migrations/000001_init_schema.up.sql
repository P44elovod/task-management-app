CREATE TABLE "project" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(500) NOT NULL,
  "description" varchar(1000),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "project_column" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(255) UNIQUE NOT NULL,
  "project_id" bigint,
  "position" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "task" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(500) NOT NULL,
  "description" varchar(5000),
  "column_id" bigint,
  "position" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "comment" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "text" varchar(5000) NOT NULL,
  "task_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "project_column" ADD FOREIGN KEY ("project_id") REFERENCES "project" ("id") ON DELETE CASCADE;

ALTER TABLE "task" ADD FOREIGN KEY ("column_id") REFERENCES "project_column" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE;