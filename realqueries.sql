RESOURCES="select rv.num_int_ressource
from ressource_rv rv
inner join societe_pos pos
	on pos.num_int_societe_pos = rv.num_int_societe_pos
inner join ressource_rv_work_type wt
	on wt.num_int_ressource = rv.num_int_ressource
where rv.is_actif = 1
	and rv.is_web_agenda = 1
	and pos.num_int_societe = $1
	and wt.num_int_work_type = 1;"

TYRE="select pdd.largeur, pdd.serie, pdd.diametre, pdd.is_sur_jante, pdd.is_jante_alu
from pneus_depot pd
inner join pneus_depot_det pdd
	on pdd.num_int_depot = pd.num_int_depot
where pd.num_int_depot = $1
	and pdd.hotel_status >= 1
limit 1;"

WEBPARAMS="select web.heure_debut,
		web.heure_fin,
		web.heure_debut_we,
		web.heure_fin_we,
		web.nb_jour_min, 
		web.nb_jour_max,
		web.intervalle_min
from parametres_webagenda web
where web.num_int_societe = $1
	and web.num_int_work_type = 1;"

TASKPARAMS="select task.resource,
		task.diameter_min,
		task.diameter_max,
		task.is_sur_jante,
		task.is_jante_alu,
		task.rof,
		task.duration
from timemanager_task task
where task.num_int_societe = $1
	and task.num_int_work_type = 1
	and (task.resource = $2 or task.resource is null)
order by task.resource,
		task.diameter_min;"

SCHEDULER="select rvh.jour,
		rvh.heure_debut,
		rvh.heure_fin
from rendez_vous_horaire rvh
where rvh.num_int_societe = $1
	and rvh.is_exclude = 0
order by rvh.jour,
		rvh.heure_debut;"
