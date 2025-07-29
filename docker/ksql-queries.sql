CREATE OR REPLACE STREAM analytic_products_filtered_table_stream (
    product_id VARCHAR KEY,
    name VARCHAR,
    description VARCHAR,
    price STRUCT<amount INT, currency VARCHAR>,
    category VARCHAR,
    brand VARCHAR,
    stock STRUCT<available INT, reserved INT>,
    sku VARCHAR,
    tags ARRAY<STRING>,
    images ARRAY<STRUCT<url VARCHAR, alt VARCHAR>>,
    specifications STRUCT<weight VARCHAR, dimensions VARCHAR, battery_life VARCHAR, water_resistance VARCHAR>,
    created_at VARCHAR,
    updated_at VARCHAR,
    index VARCHAR,
    store_id VARCHAR
)
WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='analytic_products_filtered-table', PARTITIONS=2);

CREATE OR REPLACE STREAM analytic_products_find_stream (
    id VARCHAR KEY,
    user_id VARCHAR,
    find VARCHAR
)
WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='analytic_products_find', PARTITIONS=2);

CREATE OR REPLACE STREAM personal_recom
WITH (PARTITIONS=2, REPLICAS=1) AS
SELECT *
FROM analytic_products_filtered_table_stream ap
LEFT JOIN analytic_products_find_stream af WITHIN 1 HOUR ON (ap.name = af.find);