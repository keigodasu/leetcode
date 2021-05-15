with s as (
select
    product_name,
    buyer_id
from
Sales s
left join
Product p
on
s.product_id = p.product_id
)

select
distinct buyer_id
from
s
where
buyer_id in (
    select buyer_id from s where product_name = 'S8'
)
and
buyer_id not in (
    select buyer_id from s where product_name = 'iPhone' 
);

select
    distinct buyer_id
from
    product p
join sales s
on p.product_id = s.product_id
group by buyer_id
having sum(p.product_name = 'S8') > 0 and sum(p.product_name = 'iPhone') = 0;
