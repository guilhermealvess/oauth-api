CREATE TABLE IF NOT EXISTS apis(
    id UUID NOT NULL,
    name VARCHAR NOT NULL,
    slug VARCHAR NOT NULL,
    is_active BOOLEAN DEFAULT false,
    description VARCHAR NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updasted_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT api_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS permissions(
    id UUID NOT NULL,
    api_id UUID NOT NULL,
    api_resource VARCHAR NOT NULL,
    is_active BOOLEAN DEFAULT false,
    action VARCHAR NULL NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updasted_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(api_id) REFERENCES apis(id),
    CONSTRAINT application_api_pkey PRIMARY KEY (id)
);

CREATE INDEX permission_idx ON permissions USING btree (api_id, api_resource, action);

CREATE TABLE IF NOT EXISTS clients(
    id UUID NOT NULL,
    name VARCHAR NOT NULL,
    slug VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    client_key VARCHAR NOT NULL,
    hash_client_secret VARCHAR NOT NULL,
    round_hash_client_secret INTEGER NOT NULL,
    salt_hash_client_secret VARCHAR NOT NULL,
    is_active BOOLEAN DEFAULT false,
    last_login TIMESTAMP,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updasted_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT client_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS client_permissions(
    id UUID NOT NULL,
    client_id UUID NOT NULL,
    permission_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(client_id) REFERENCES clients(id),
    FOREIGN KEY(permission_id) REFERENCES permissions(id),
    CONSTRAINT client_permission_pkey PRIMARY KEY (id)
);