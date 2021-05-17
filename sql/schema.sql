-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS "article";
CREATE TABLE "article" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "title" text,
  "remark" text,
  "pic" text,
  "content" text
);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "username" TEXT NOT NULL,
  "password" TEXT NOT NULL,
  "avatar" TEXT
);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS "comment";
CREATE TABLE "comment" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "pid" integer,
  "article_id" integer NOT NULL,
  "time" text,
  "content" text,
  "username" text,
  "avatar" text
);
