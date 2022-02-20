DROP TABLE IF EXISTS attendances;

CREATE TABLE attendances (
	id VARCHAR(50) PRIMARY KEY,
	day_id VARCHAR(50) NOT NULL,
	user_id VARCHAR(50) NOT NULL,
	status ENUM('pending', 'approved', 'rejected') NOT NULL,
    notes VARCHAR(255) NOT NULL,
	admin_id VARCHAR(50) NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	FOREIGN KEY (day_id) REFERENCES days(id),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (admin_id) REFERENCES users(id)
);