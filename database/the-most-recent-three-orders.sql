select
    customer_name,
    customer_id,
    order_id,
    order_date
from (
select
    c.name customer_name,
    c.customer_id customer_id,
    o.order_id order_id,
    o.order_date order_date,
    rank() over(partition by customer_id order by order_date desc) rnk
from
orders o
join customers c
using(customer_id)
) t
where rnk <= 3
order by customer_name, customer_id, order_date desc
