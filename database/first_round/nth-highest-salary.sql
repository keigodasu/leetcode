CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      select 
        distinct salary
      from (
          select
           salary,
           dense_rank() over(order by salary desc) rnk
          from
          employee
      ) t
      where rnk = N
  );
END
