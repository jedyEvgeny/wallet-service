CREATE TABLE wallets (
    id SERIAL PRIMARY KEY,
    wallet_id UUID NOT NULL UNIQUE,
    amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_wallet_id ON wallets(wallet_id);
