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
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    code VARCHAR(50) UNIQUE NOT NULL,
    initial_balance DECIMAL(10, 2),
    activation_date TIMESTAMP,
    last_used_date TIMESTAMP,
    pin_code VARCHAR(6),
    max_uses INT,
    current_uses INT DEFAULT 0
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

-- Tabla para diseños de GiftCards
CREATE TABLE GiftCardDesigns (
    design_id SERIAL PRIMARY KEY,
    giftcard_id INT REFERENCES GiftCards(giftcard_id),
    template_url VARCHAR(255),
    custom_message TEXT,
    brand_logo VARCHAR(255),
    color_scheme VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla para reglas de redención
CREATE TABLE RedemptionRules (
    rule_id SERIAL PRIMARY KEY,
    giftcard_id INT REFERENCES GiftCards(giftcard_id),
    restriction_type VARCHAR(50), -- product, category, location, time
    restriction_details JSONB,    -- Permite almacenar reglas complejas
    valid_from TIMESTAMP,
    valid_until TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla para beneficios de campaña
CREATE TABLE CampaignBenefits (
    benefit_id SERIAL PRIMARY KEY,
    campaign_id INT REFERENCES Campaigns(campaign_id),
    benefit_type VARCHAR(50),     -- cashback, bonus, double_points
    benefit_details JSONB,        -- Detalles específicos del beneficio
    is_stackable BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla para usuarios de empresas
CREATE TABLE CompanyUsers (
    user_id SERIAL PRIMARY KEY,
    company_id INT REFERENCES Companies(company_id),
    role_id INT REFERENCES Roles(role_id),
    email VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de roles y permisos
CREATE TABLE Roles (
    role_id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    permissions JSONB,            -- Almacena permisos en formato JSON
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla para historial de GiftCards
CREATE TABLE GiftCardHistory (
    history_id SERIAL PRIMARY KEY,
    giftcard_id INT REFERENCES GiftCards(giftcard_id),
    action VARCHAR(50),
    old_values JSONB,
    new_values JSONB,
    created_by INT,              -- Usuario que realizó el cambio
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla para registros de API
CREATE TABLE APILogs (
    log_id SERIAL PRIMARY KEY,
    api_id INT REFERENCES APIs(api_id),
    request_data JSONB,
    response_data JSONB,
    response_code INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Índices para mejorar el rendimiento
CREATE INDEX idx_giftcards_code ON GiftCards(code);
CREATE INDEX idx_giftcards_campaign ON GiftCards(campaign_id);
CREATE INDEX idx_transactions_giftcard ON Transactions(giftcard_id);
CREATE INDEX idx_company_orders_company ON CompanyOrders(company_id);
