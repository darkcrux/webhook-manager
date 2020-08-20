CREATE TABLE transaction_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    description TEXT,
    SAMPLE_PAYLOAD TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    customer_external_id TEXT,
    unique_key VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP WITHOUT TIME ZONE default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE webhooks (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    transaction_type_id INTEGER NOT NULL,
    webhook_url TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
    FOREIGN KEY (customer_id) REFERENCES customers (id),
    FOREIGN KEY (transaction_type_id) REFERENCES transaction_types (id),
    CONSTRAINT unq_customer_transaction_type UNIQUE (customer_id, transaction_type_id)
);

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
        idem_key TEXT NOT NULL,
    webhook_id INTEGER NOT NULL,
    payload JSON NOT NULL,
    status VARCHAR(10) NOT NULL,
    fail_reason TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
    FOREIGN KEY (webhook_id) REFERENCES webhooks (id)
);