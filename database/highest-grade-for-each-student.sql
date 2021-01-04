select
    student_id,
    min(course_id) course_id,
    grade
from
(
select
    student_id,
    course_id,
    grade,
    rank() over(partition by student_id order by grade desc) rnk
from
    enrollments
) t
where t.rnk = 1
group by student_id;
