select
p.name name,
sum(i.rest) rest,
sum(i.paid) paid,
sum(i.canceled) canceled,
sum(i.refunded) refunded
from
invoice i
join
product p
using(product_id)
group by
product_id
order by p.name
