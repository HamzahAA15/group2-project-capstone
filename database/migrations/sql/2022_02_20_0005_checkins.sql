DROP TABLE IF EXISTS checkins;

CREATE TABLE checkins (
	id VARCHAR(50) PRIMARY KEY,
	attendance_id VARCHAR(50) NOT NULL,
    temprature FLOAT NOT NULL,
	checkin_time TIMESTAMP DEFAULT now() NOT NULL,
	checkout_time TIMESTAMP DEFAULT NULL,
	FOREIGN KEY (attendance_id) REFERENCES attendances(id)
);