DROP SCHEMA IF EXISTS qr_order_system;
CREATE SCHEMA qr_order_system;
USE qr_order_system;

/** ショップ */
DROP TABLE IF EXISTS shops;
CREATE TABLE shops(
    id INT(10) NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

/** メニュー大分類 */
DROP TABLE IF EXISTS broad_categories;
CREATE TABLE broad_categories(
    id INT(10) NOT NULL AUTO_INCREMENT,
    shop_id INT(10) NOT NULL,
    name VARCHAR(100) NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(shop_id) REFERENCES shops(id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

/** メニュー小分類 */
DROP TABLE IF EXISTS sub_categories;
CREATE TABLE sub_categories(
    id INT(10) NOT NULL AUTO_INCREMENT,
    broad_category_id INT(10) NOT NULL,
    name VARCHAR(100) NOT NULL,
    price INT(10),
    image_url VARCHAR(200), 
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(broad_category_id) REFERENCES broad_categories(id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;


/** 初期データ挿入 */
insert into shops (name) values ('テストショップ');