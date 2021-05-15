select
product_id,
p.product_name
from
sales s
join
product p
using(product_id)
group by product_id
having min(sale_date) >= '2019-01-01' and max(sale_date) <= '2019-03-31'


