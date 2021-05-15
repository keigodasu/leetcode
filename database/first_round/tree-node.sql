select
t1.id,
case 
    when sum(t1.id = t2.p_id) = 0  then 'Leaf'
    when sum(t1.p_id is null) > 0 then 'Root'
    else 'Inner'
end as type
from
tree t1
join
tree t2
group by t1.id
order by t1.id;


SELECT DISTINCT t1.id, (
    CASE
    WHEN t1.p_id IS NULL  THEN 'Root'
    WHEN t1.p_id IS NOT NULL AND t2.id IS NOT NULL THEN 'Inner'
    WHEN t1.p_id IS NOT NULL AND t2.id IS NULL THEN 'Leaf'
    END
) AS Type 
FROM tree t1
LEFT JOIN tree t2
ON t1.id = t2.p_id;
