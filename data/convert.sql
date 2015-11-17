use depositauth_development;
delete from deposit_auth;
insert into deposit_auth select 0, empl_id, computing_id, first_name, middle_name, last_name, academic_career, academic_program, academic_plan, plan_degree, title_of_work, milestone, IF(ISNULL(libra_id),'', libra_id), date_approved, libra_approved, libra_wrote_to_sis, created_at, updated_at from sis_records;
