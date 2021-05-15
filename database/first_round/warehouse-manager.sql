select
w.name warehouse_name,
sum(p.width * p.length * p.height * w.units) volume
from
warehouse w
join
products p
using(product_id)
group by w.name
