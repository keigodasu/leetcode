with cte as (
select
id,
visit_date,
people,
lead(people, 1) over(order by id) next,
lead(people, 2) over(order by id) next2,
lag(people, 1) over(order by id) prev,
lag(people, 2) over(order by id) prev2
from stadium
)

select
id,
visit_date,
people
from
cte
where
(people >= 100 and next >= 100 and next2 >= 100)
or
(people >= 100 and prev >= 100 and next >= 100)
or
(people >= 100 and prev >= 100 and prev2 >= 100)
