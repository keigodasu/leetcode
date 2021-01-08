select
month,
country,
sum(case when state = 'approved' then 1 else 0 end) approved_count,
sum(case when state='approved' then amount else 0 end) approved_amount,
sum(case when state = 'back' then 1 else 0 end) chargeback_count,
sum(case when state='back' then amount else 0 end) chargeback_amount
from
(
select left(c.trans_date, 7) month, country, amount, "back" state
    from chargebacks c
    join
    transactions t
    on
    t.id = c.trans_id
    union all
    select left(trans_date, 7) month, country, amount,  state
    from transactions
    where state = 'approved'
) temp
group by month, country
order by month;

