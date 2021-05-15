with history as (
select
fail_date as exec_date,
'failed' as result
from failed 
union all
select
success_date as exec_date,
'succeeded' as result
from succeeded
),
history_with_rownum as (
select
exec_date,
result,
row_number() over(order by result, exec_date asc) as seq
from
history
where exec_date between '2019-01-01' and '2019-12-31'
),
history_with_seqstart as (
select
exec_date,
result,
seq,
date_add(exec_date, interval -seq day) as seqStart
from
history_with_rownum
)

select
result as period_state,
min(exec_date) as start_date,
max(exec_date) as end_date
from
history_with_seqstart
group by seqStart, result
order by start_date asc
