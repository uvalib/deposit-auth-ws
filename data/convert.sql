use depositauth_development;
delete from deposit_auth;
insert into deposit_auth select 0, computing_id, milestone, IF(ISNULL(libra_id),'', libra_id), title_of_work, academic_program, libra_approved, libra_wrote_to_sis, created_at, updated_at from sis_records;
