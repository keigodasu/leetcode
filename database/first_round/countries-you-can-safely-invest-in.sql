with callHistory as (
    select
    caller_id as call_id,
    duration
    from
    calls
    union all
    select
    callee_id as call_id,
    duration
    from
    calls
)

select
country_name as country
from
callHistory as c
join
(
    select
        id,
        country_code,
        c.name as country_name
    from
    person p
    join
    country c
    on left(p.phone_number, 3) = c.country_code
) as p
on c.call_id = p.id
group by p.country_code
having avg(duration) > (select avg(duration) from callHistory)

