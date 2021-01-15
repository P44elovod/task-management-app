CREATE TABLE "project" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar(500) NOT NULL,
  "description" varchar(1000),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "column" (
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
  "name" varchar(500) NOT NULL,
  "description" varchar(5000),
  "task_id" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "column" ADD FOREIGN KEY ("project_id") REFERENCES "project" ("id");

ALTER TABLE "task" ADD FOREIGN KEY ("column_id") REFERENCES "column" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id");