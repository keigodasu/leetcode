select
username,
activity,
startDate,
endDate
from
(
select
*,
row_number() over (partition by username order by startDate desc) as rnk,
count(username) over (partition by username) as counts
from
useractivity 
) t
where
rnk = 2 or counts = 1

