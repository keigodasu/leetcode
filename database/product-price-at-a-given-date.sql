select
distinct p.product_id, ifnull(t.new_price, 10) price
from
products p
left join
(
select
*
from
products
where (product_id, change_date) in 
(
select
    product_id,
    max(change_date)
from
    products
where 
    change_date <= '2019-08-16'
group by product_id 
) 
)t
using(product_id);

