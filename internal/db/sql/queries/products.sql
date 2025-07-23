
--Companies--------------------------------------------------------------

-- name: CreateCompanies :one
INSERT INTO companies (id,name,status) VALUES ($1,$2,$3) RETURNING *;

-- name: GetCompanyByName :one
SELECT * FROM companies WHERE name=$1;

-- name: GetAllCompanies :many
SELECT * FROM companies;


--Models --------------------------------------------------------------



-- name: CreateModel :one
INSERT INTO models (id,company_id,name) VALUES ($1,$2,$3) RETURNING *;

-- name: GetModelByName :one
SELECT * FROM models WHERE name=$1;



-- name: CreateModelVariant :one
INSERT INTO model_variants(id,model_id,model_type,model_image) VALUES($1,$2,$3,$4) RETURNING *;

-- name: GetModelVariantsById :one
SELECT * FROM model_variants WHERE model_id=$1;

--Brands---------------------------------------------------------------

-- name: CreateBrand :one 

INSERT INTO brands(id,name) VALUES ($1,$2) RETURNING *;

-- name: GetBrandByName :one

SELECT * FROM brands WHERE name=$1;

--Categories -------------------------------------------------------------

-- name: CreateCategories :one

INSERT INTO categories(id,name,image) VALUES ($1,$2,$3) RETURNING *;

-- name: GetCategoriesByName :one
SELECT * FROM categories WHERE name=$1;


--Products ----------------------------------------------------------------

-- name: CreateProductParts :one 

INSERT INTO product_parts(id,company_id,model_id,brand_id,category_id,part_no,is_active,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING *;

-- name: GetProductPartsByProductId :one
SELECT * FROM product_parts WHERE company_id=$1;

-- name: GetProductPartsByModelId :one
SELECT * FROM product_parts WHERE model_id=$1;

-- name: GetProductPartsByCategoryId :one
SELECT * FROM product_parts WHERE category_id=$1;

-- name: GetProductByBrandName :one
SELECT * FROM product_parts WHERE brand_id=$1;

-- name: GetProductPartsByPartNo :one
SELECT * FROM product_parts WHERE part_no=$1;

--Customer-----------------------------------------------------------------------


-- name: CreateCustomer :one
INSERT INTO customers (id,customer_company_name,contact_person,mobile,type,customer_designation,address,flat,street,city,state,pincode,created_at,updated_at,deleted_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING *;

-- name: GetAllCustomers :many
SELECT * FROM customers;