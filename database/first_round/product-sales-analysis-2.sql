select
p.product_id,
sum(s.quantity) total_quantity
from
sales s
left join
product p
using(product_id)
group by p.product_id
