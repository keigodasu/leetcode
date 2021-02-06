with cte as (
select
    Id,
        month,
            salary
            from
            (
            select
            Id,
            month,
            salary,
            row_number() over(partition by Id order by month desc) rnk
            from employee
            ) t
            where rnk > 1
            )

select
id,
month,
sum(salary) over(partition by id order by month asc rows 2 preceding) as salary
from
cte
order by id asc, month desc
