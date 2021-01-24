select
e.*,
case 
when operator = '>' then (case when v1.value > v2.value then 'true' else 'false' end) 
when operator = '<' then (case when v1.value < v2.value then 'true' else 'false' end) 
when operator = '=' then (case when v1.value = v2.value then 'true' else 'false' end) 
end as value
from
expressions e
join
variables v1
on 
v1.name = e.left_operand 
join
variables v2
on
v2.name = e.right_operand
