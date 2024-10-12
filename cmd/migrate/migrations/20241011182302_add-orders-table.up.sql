CREATE TABLE IF NOT EXISTS orders (
  id INT PRIMARY  KEY ,
  userId INT NOT NULL,
  total DECIMAL(10, 2),
  status VARCHAR(20) DEFAULT 'pending',
  CHECK (status IN ('pending', 'completed', 'cancelled')),
  address TEXT NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (userId) REFERENCES user(id)
);
