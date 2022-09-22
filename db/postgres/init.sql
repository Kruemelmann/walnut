-- walnut db schema

CREATE TABLE IF NOT EXISTS "User" (
    "ID" int   NOT NULL,
    "Username" VARCHAR(45)   NOT NULL,
    "Password" VARCHAR(45)   NOT NULL,
    "Role" int   NOT NULL,
    "LastLoggedin" date   NULL,
    CONSTRAINT "pk_User" PRIMARY KEY (
        "ID"
    )
);

CREATE TABLE IF NOT EXISTS "UserTeam" (
    "ID" int   NOT NULL,
    "Team" int   NOT NULL,
    "User" int   NOT NULL,
    CONSTRAINT "pk_UserTeam" PRIMARY KEY (
        "ID"
    )
);

CREATE TABLE IF NOT EXISTS "Team" (
    "ID" int   NOT NULL,
    "Title" VARCHAR(45)   NOT NULL,
    CONSTRAINT "pk_Team" PRIMARY KEY (
        "ID"
    )
);

CREATE TABLE IF NOT EXISTS "Role" (
    "ID" int   NOT NULL,
    "Title" VARCHAR(45)   NOT NULL,
    CONSTRAINT "pk_Role" PRIMARY KEY (
        "ID"
    )
);

CREATE TABLE IF NOT EXISTS "Module" (
    "ID" int   NOT NULL,
    "Kind" int   NOT NULL,
    "Creator" int   NOT NULL,
    CONSTRAINT "pk_Module" PRIMARY KEY (
        "ID"
    )
);

CREATE TABLE IF NOT EXISTS "Kind" (
    "ID" int   NOT NULL,
    "Title" VARCHAR(45)   NOT NULL,
    CONSTRAINT "pk_Kind" PRIMARY KEY (
        "ID"
    )
);

ALTER TABLE "User" ADD CONSTRAINT "fk_User_Role" FOREIGN KEY("Role") REFERENCES "Role" ("ID");

ALTER TABLE "UserTeam" ADD CONSTRAINT "fk_UserTeam_Team" FOREIGN KEY("Team") REFERENCES "Team" ("ID");

ALTER TABLE "UserTeam" ADD CONSTRAINT "fk_UserTeam_User" FOREIGN KEY("User") REFERENCES "User" ("ID");

ALTER TABLE "Module" ADD CONSTRAINT "fk_Module_Kind" FOREIGN KEY("Kind") REFERENCES "Kind" ("ID");

ALTER TABLE "Module" ADD CONSTRAINT "fk_Module_Creator" FOREIGN KEY("Creator") REFERENCES "User" ("ID");

CREATE INDEX "idx_User_Username" ON "User" ("Username");




