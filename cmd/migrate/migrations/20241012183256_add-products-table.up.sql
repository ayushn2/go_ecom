CREATE TABLE IF NOT EXISTS products (
  id INT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  image VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  quantity INT  NOT NULL CHECK (quantity >= 0),
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
