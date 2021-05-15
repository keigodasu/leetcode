select
b.book_id,
b.name
from
books b
left join
orders o
on b.book_id = o.book_id and o.dispatch_date between '2018-06-23' and '2019-06-23'
where datediff('2019-06-23', b.available_from) > 30
group by b.book_id
having ifnull(sum(o.quantity), 0) < 10;

