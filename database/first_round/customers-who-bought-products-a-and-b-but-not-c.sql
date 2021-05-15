select
c.customer_id,
c.customer_name
from
orders o
join
customers c
using(customer_id)
group by c.customer_id
having sum(o.product_name = 'A') >= 1
and
sum(o.product_name = 'B') >= 1
and
sum(o.product_name = 'C') = 0
order by customer_id
