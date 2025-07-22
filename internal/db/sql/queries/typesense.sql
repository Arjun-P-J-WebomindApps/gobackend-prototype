
-- name: GetProductTableMetaData :many
SELECT 
    pp.id,
    co.name AS company,
    m.name AS model,
    mv.model_type,
    b.name AS brand,
    c.name AS category,
    pp.part_no
FROM product_parts pp
    JOIN companies co on co.id=pp.company_id
    JOIN models m on m.id = pp.model_id
    LEFT JOIN model_variants mv ON mv.model_id=m.id
    JOIN brands b ON b.id = pp.brand_id
    JOIN categories c ON c.id =  pp.category_id;