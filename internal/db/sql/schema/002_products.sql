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
    name TEXT NOT NULL
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
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down

DROP TABLE product_parts;
DROP TABLE categories;
DROP TABLE brands;
DROP TABLE model_variants;
DROP TABLE models;
DROP TABLE companies;
