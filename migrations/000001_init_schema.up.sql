CREATE TABLE users (                                                                      
    id SERIAL PRIMARY KEY,                                                                
    email VARCHAR(255) UNIQUE NOT NULL,                                                   
    password_hash VARCHAR(255) NOT NULL,                                                  
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP                         
);                                                                                        
                                                                                              
CREATE TABLE product(                                                                   
    article_id VARCHAR(50) PRIMARY KEY,                                                   
    product_code VARCHAR(50),
    prod_name VARCHAR(255),
    product_type_no INTEGER,
    product_type_name VARCHAR(255),
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);