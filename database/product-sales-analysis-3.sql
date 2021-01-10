select
product_id,
year as first_year,
quantity,
price
from
sales
where
(product_id, year) in
(
select
product_id,
min(year) first_year
from
sales
join
product
using(product_id)
group by product_id
);

