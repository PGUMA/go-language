-- DB作成
CREATE DATABASE dev; 

-- 作成したDBへ切り替え
\c dev

-- スキーマ作成
CREATE SCHEMA devschema;

-- ロールの作成
CREATE ROLE dev WITH LOGIN PASSWORD 'postgres';

-- 権限追加
GRANT ALL PRIVILEGES ON SCHEMA devschema TO dev;