with a as ( 
select
v.user_id,
v.visit_date,
if(transaction_date is null, 0, count(*)) as transaction_count
from
visits v
left join
transactions t
on v.user_id = t.user_id and v.visit_date = t.transaction_date
group by 1,2
),
b as ( 
select
row_number() over() as rn from transactions union select 0
)

select
rn as transactions_count, if(transaction_count is null, 0, count(*)) as visits_count
from
a right join b on transaction_count = rn
where rn <= (select max(transaction_count) from a)
group by rn order by 1
