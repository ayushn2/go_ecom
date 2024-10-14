CREATE TABLE IF NOT EXISTS order_items (
  id INT PRIMARY KEY  ,
  orderId INT  NOT NULL CHECK(orderId >0 ),
  productId INT NOT NULL CHECK (productId >0 ),
  quantity INT NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  

  FOREIGN KEY (orderId) REFERENCES orders(id),
  FOREIGN KEY (productId) REFERENCES products(id)
);