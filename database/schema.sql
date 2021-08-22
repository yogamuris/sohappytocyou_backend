CREATE TABLE user(
                     id INTEGER PRIMARY KEY,
                     username VARCHAR(50) UNIQUE NOT NULL,
                     email VARCHAR(200) UNIQUE NOT NULL,
                     password VARCHAR(255) NOT NULL,
                     created_at DATETIME,
                     verified_at DATETIME
);

CREATE TABLE page(
                     id INTEGER PRIMARY KEY,
                     id_user INTEGER NOT NULL,
                     username VARCHAR(50) UNIQUE NOT NULL,
                     backgroud VARCHAR(10),
                     photo VARCHAR(255),
                     description VARCHAR(255),
                     FOREIGN KEY (id_user)
                         REFERENCES user (id)
);

CREATE TABLE link(
                     id INTEGER PRIMARY KEY,
                     id_page INTEGER NOT NULL,
                     url VARCHAR(255) NOT NULL,
                     Visited INTEGER DEFAULT 0,
                     FOREIGN KEY (id_page)
                         REFERENCES page (id)
);