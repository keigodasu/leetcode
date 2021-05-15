select
s.student_id,
s.student_name,
ss.subject_name,
count(e.subject_name) attended_exams
from
Students s
join
Subjects ss
left outer join
Examinations e
on 
e.student_id = s.student_id and e.subject_name = ss.subject_name
group by s.student_id, ss.subject_name
order by s.student_id, ss.subject_name
