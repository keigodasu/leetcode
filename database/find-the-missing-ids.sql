recursive seq_ids as (
select 1 as seq_id
union
select seq_id + 1
from seq_ids
where seq_id < (select max(customer_id) from customers)
)

select seq_id ids from seq_ids
where seq_id not in (select customer_id from customers)
