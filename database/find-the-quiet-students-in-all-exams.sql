with cte as (select s.student_id,
	student_name,rank() over (partition by exam_id order by score) min_exam,
	rank() over (partition by exam_id order by score desc ) max_exam, score
	from Student s join exam e
	on e.student_id=s.student_id)
select student_id STUDENT_ID,student_name STUDENT_NAME from cte
group by student_name
having min(min_exam)>1
and min(max_exam)>1
order by student_id
