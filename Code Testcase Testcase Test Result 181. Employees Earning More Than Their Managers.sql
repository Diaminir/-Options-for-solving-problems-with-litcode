SELECT a.name AS employee
FROM employee a
JOIN employee b ON (b.id = a.managerid)
WHERE a.salary>b.salary
