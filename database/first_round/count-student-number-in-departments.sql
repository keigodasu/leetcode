select
d.dept_name,
count(s.student_id) as student_number
from
department d
left join
student s
using(dept_id)
group by d.dept_id
order by student_number desc
