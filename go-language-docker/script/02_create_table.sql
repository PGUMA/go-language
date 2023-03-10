-- DB切り替え
\c dev

-- テーブル作成
CREATE TABLE  devschema.user (
    id bigserial PRIMARY KEY,
    name VARCHAR(10) NOT NULL
);

-- 権限追加
GRANT ALL PRIVILEGES ON devschema.user TO dev;

-- サンプルレコード作成
INSERT INTO devschema.user (name) VALUES ('test_01');
INSERT INTO devschema.user (name) VALUES ('test_02');
INSERT INTO devschema.user (name) VALUES ('test_03');
