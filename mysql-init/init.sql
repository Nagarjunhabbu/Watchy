CREATE DATABASE IF NOT EXISTS eventdb;

USE eventdb;

CREATE TABLE IF NOT EXISTS users (
                                        
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(20),
    gender varchar(1) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS events (
                                        
    id INT AUTO_INCREMENT PRIMARY KEY,
    event_title VARCHAR(255) NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    video_id VARCHAR(255) NOT NULL,
    action VARCHAR(20),
    duration int,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

INSERT INTO users(user_id,name,location,gender) VALUES("user123","Nags","Mysore","M");
INSERT INTO users(user_id,name,location,gender) VALUES("user456","Virat","Banaglore","M");