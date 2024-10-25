-- name: GetParentOrg :many
WITH RECURSIVE orgs_q as
(
	select id, name, parent_id
	from organizations o
	where o.name = $1 
	UNION ALL
select p.id, p.name, p.parent_id
from organizations p
INNER JOIN orgs_q q ON o.id = q.parent_id 
)
select * from orgs_q;
