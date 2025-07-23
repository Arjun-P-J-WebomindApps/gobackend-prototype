-- +goose Up


--Companies-----------------------------------------------------------------------

CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    status BOOLEAN NOT NULL
);

--Models--------------------------------------------------------------------------

CREATE TABLE models(
    id UUID PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id),
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE model_variants(
    id UUID PRIMARY KEY,
    model_id UUID NOT NULL REFERENCES models(id),
    model_type TEXT,
    model_image TEXT
);

--Brands-------------------------------------------------------------------------

CREATE TABLE brands (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

--Categories -----------------------------------------------------------------------

CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    image TEXT
);

-- Products -------------------------------------------------------------------------

CREATE TABLE product_parts(
    id UUID PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id),
    model_id UUID NOT NULL REFERENCES models(id),
    brand_id UUID NOT  NULL REFERENCES brands(id),
    category_id UUID NOT NULL REFERENCES categories(id),
    part_no TEXT NOT NULL UNIQUE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Customers ----------------------------------------------------------------------------

CREATE TABLE customers(
    id UUID PRIMARY KEY,
    customer_company_name VARCHAR(255) NOT NULL,
    contact_person VARCHAR(255) NOT NULL ,
    mobile VARCHAR(20) UNIQUE NOT NULL,
    type VARCHAR(255) NOT NULL,
    customer_designation VARCHAR(100),
    address TEXT,
    flat VARCHAR(100),
    street VARCHAR(100),
    city VARCHAR(100),
    state VARCHAR(100),
    pincode VARCHAR(20),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- +goose Down

DROP TABLE customers;
DROP TABLE product_parts;
DROP TABLE categories;
DROP TABLE brands;
DROP TABLE model_variants;
DROP TABLE models;
DROP TABLE companies;
