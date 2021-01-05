select
    product_name,
    product_id,
    order_id,
    order_date
from
(
select
    product_name,
    product_id,
    order_id,
    order_date,
    rank() over(partition by product_id order by order_date desc) rnk
from
    orders
    join
    products
    using(product_id)
) t
where
    rnk = 1
order by product_name asc, product_id asc, order_id asc;


with orders_with_rank as (
select
    product_name,
    product_id,
    order_id,
    order_date,
    rank() over(partition by product_id order by order_date desc) rnk
from
    orders
    join
    products
    using(product_id)
)
select
    product_name,
    product_id,
    order_id,
    order_date
from orders_with_rank
where
    rnk = 1
order by product_name, product_id, order_id;
