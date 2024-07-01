-- Tabla de GiftCards
CREATE TABLE GiftCards (
    giftcard_id SERIAL PRIMARY KEY,
    type VARCHAR(50),
    balance DECIMAL(10, 2),
    expiration_date DATE,
    status VARCHAR(50),
    is_promotional BOOLEAN DEFAULT FALSE,
    campaign_id INT REFERENCES Campaigns(campaign_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Campaigns
CREATE TABLE Campaigns (
    campaign_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    start_date DATE,
    end_date DATE,
    discount_percentage DECIMAL(5, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Inventory
CREATE TABLE Inventory (
    inventory_id SERIAL PRIMARY KEY,
    giftcard_id INT REFERENCES GiftCards(giftcard_id),
    location_type VARCHAR(50),
    bin_location VARCHAR(50),
    quantity INT,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Transactions
CREATE TABLE Transactions (
    transaction_id SERIAL PRIMARY KEY,
    giftcard_id INT REFERENCES GiftCards(giftcard_id),
    amount DECIMAL(10, 2),
    transaction_type VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Orders
CREATE TABLE Orders (
    order_id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES Customers(customer_id),
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount DECIMAL(10, 2),
    status VARCHAR(50)
);

-- Tabla de Customers
CREATE TABLE Customers (
    customer_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(20),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Companies
CREATE TABLE Companies (
    company_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    address TEXT,
    contact_email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de CompanyOrders
CREATE TABLE CompanyOrders (
    company_order_id SERIAL PRIMARY KEY,
    company_id INT REFERENCES Companies(company_id),
    order_id INT REFERENCES Orders(order_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Reports
CREATE TABLE Reports (
    report_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    filters TEXT,
    generated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de APIs
CREATE TABLE APIs (
    api_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    endpoint VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de AuditLogs
CREATE TABLE AuditLogs (
    log_id SERIAL PRIMARY KEY,
    action TEXT,
    user_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de Settings
CREATE TABLE Settings (
    setting_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    value TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
