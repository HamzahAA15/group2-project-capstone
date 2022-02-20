DROP TABLE IF EXISTS certificates;

CREATE TABLE certificates (
	id VARCHAR(50) PRIMARY KEY,
	user_id VARCHAR(50) NOT NULL,
	image VARCHAR(255) NOT NULL,
	dosage INT NOT NULL,
	status ENUM('approved', 'rejected') NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	FOREIGN KEY (user_id) REFERENCES users(id)
);