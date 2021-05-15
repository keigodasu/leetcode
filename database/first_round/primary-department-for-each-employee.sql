SELECT 
employee_id, 
CASE WHEN Count(*) = 1 THEN Max(department_id) ELSE Max(CASE WHEN primary_flag = 'Y' THEN department_id end) end AS department_id FROM employee 
GROUP BY employee_id;
