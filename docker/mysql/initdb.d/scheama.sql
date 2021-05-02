-- golang_dbという名前のデータベースを作成
CREATE DATABASE IF NOT EXISTS golang_db;
-- golang_dbをアクティブ
use golang_db;
-- usersテーブルを作成。名前とパスワード
CREATE TABLE users (
    id INT(11) AUTO_INCREMENT NOT NULL,
    uuid VARCHAR(64),
    name VARCHAR(64) NOT NULL,
    email VARCHAR(64),
    password CHAR(30) NOT NULL,
    created_at DATETIME
);
-- usersテーブルに２つレコードを追加
INSERT INTO users (name, password) VALUES ("gophar", "5555","kouropa","ko@com","password");
INSERT INTO users (name, password) VALUES ("octcat", "0000","kou","kk@com");