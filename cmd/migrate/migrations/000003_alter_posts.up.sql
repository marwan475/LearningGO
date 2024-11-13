ALTER TABLE
    posts
ADD 
    CONSTRAINT fk_user FOREIGN KEY (userid) REFERENCES users (id);