CREATE TABLE "public"."users"(
    "uuid"	uuid PRIMARY KEY DEFAULT uuid_generate_v1mc(),
    "name"	VARCHAR(60) NOT NULL UNIQUE,
    "email"	VARCHAR(120) NOT NULL UNIQUE,
    "password" VARCHAR(200) NOT NULL,
    "create_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "deleted_at" timestamp
);

CREATE TABLE "public"."articles"(
    "uuid"	uuid PRIMARY KEY DEFAULT uuid_generate_v1mc(),
    "user_uuid" uuid NOT NULL,
    "title"	VARCHAR(60) NOT NULL,
    "text"	VARCHAR(2000) NOT NULL,
    "create_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "deleted_at" timestamp
);

CREATE TABLE "public"."comments"(
    "uuid"	uuid PRIMARY KEY DEFAULT uuid_generate_v1mc(),
    "article_uuid" uuid NOT NULL,
    "user_uuid" uuid NOT NULL,
    "text"	VARCHAR(2000) NOT NULL,
    "create_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "deleted_at" timestamp
);

CREATE TABLE "public"."article_likes"(
"uuid"	uuid PRIMARY KEY DEFAULT uuid_generate_v1mc(),
"article_uuid" uuid NOT NULL,
"user_uuid" uuid NOT NULL
);

CREATE TABLE "public"."comment_likes"(
"uuid"	uuid PRIMARY KEY DEFAULT uuid_generate_v1mc(),
"comment_uuid" uuid NOT NULL,
"user_uuid" uuid NOT NULL
);

Alter TABLE "public"."article_likes" add constraint fk_articles_like_uuid foreign key ("article_uuid")
references "public"."articles"(uuid) on delete restrict on update cascade;

Alter TABLE "public"."article_likes" add constraint fk_users_like_uuid foreign key ("user_uuid")
references "public"."users"(uuid) on delete restrict on update cascade;

Alter TABLE "public"."comment_likes" add constraint fk_articles_like_uuid foreign key ("comment_uuid")
references "public"."comments"(uuid) on delete restrict on update cascade;

Alter TABLE "public"."comment_likes" add constraint fk_users_like_uuid foreign key ("user_uuid")
references "public"."users"(uuid) on delete restrict on update cascade;

Alter TABLE "public"."articles" add constraint fk_articles_user_uuid foreign key ("user_uuid")
references "public"."users"(uuid) on delete restrict on update cascade;

Alter TABLE "public"."comments" add constraint fk_comments_user_uuid foreign key ("user_uuid")
references "public"."users"(uuid) on delete restrict on update cascade;

Alter TABLE "public"."comments" add constraint fk_comments_article_uuid foreign key ("article_uuid")
references "public"."articles"(uuid) on delete restrict on update cascade;