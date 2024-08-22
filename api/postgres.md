## Postgres Notes
/*
docker run \
--name myPostgresDb \
-p 5455:5432 \
-e POSTGRES_USER=postgresUser \
-e POSTGRES_PASSWORD=postgresPW \
-e POSTGRES_DB=postgresDB \
-d \
postgres:16-alpine
*/

// docker exec -it fa58d2a22133 sh

// psql -d postgresDB -U postgresUser

/*
CREATE TABLE points (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	timestamp timestamp,
	user_id uuid,
	point_type_id uuid,
	value real,
	created_by uuid
);
*/

/*
INSERT INTO points (timestamp, user_id, point_type_id, value, created_by) VALUES
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 20.01, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 12.01, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 13.01, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 14.01, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 15.01, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 20.02, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 12.02, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 13.02, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 14.02, '550e8400-e29b-41d4-a716-446655440000'),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 15.02, '550e8400-e29b-41d4-a716-446655440000');
*/

/*
CREATE TABLE users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name text,
	email text
);
*/

/*
INSERT INTO users (id, name, email) VALUES
('d6f0bf56-0abb-4278-9e2c-c3e0bfc18c1d', 'Nadia Ward', 'bszuderaw@gmail.com'),
('550e8400-e29b-41d4-a716-446655440000', 'Bruce Szudera Wienand', 'tenbounce.official@gmail.com'),
('123e4567-e89b-12d3-a456-426614174000', 'Derek Therrien', 'dtherrien2503@gmail.com'),
('987fbc97-4bed-5078-889f-8c6e44d66b00', 'Lourens Willekes', 'lourw95@gmail.com');
*/

/*
CREATE TABLE point_types (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name text
);
*/

/*
INSERT INTO point_types (id, name) VALUES
('4e4b2b1c-5063-425a-a409-71b431068f78', 'Compulsory Routine'),
('0d1b30ef-00d4-41d6-8581-b8d554752816', 'Optional Routine'),
('dade4383-d869-4562-a680-88cb38f9972a', 'Tenboounce'),
('8640f8e9-0cf6-4be4-b182-d40c21a44067', 'Ten Doubles');
*/
