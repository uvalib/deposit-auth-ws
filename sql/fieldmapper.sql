-- drop the table if it exists
DROP TABLE IF EXISTS fieldmapper;

-- and create the new one
CREATE TABLE fieldmapper(
   id          INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
   field_class VARCHAR( 32 ) NOT NULL DEFAULT '',
   field_name  VARCHAR( 255 ) NOT NULL DEFAULT '',
   field_value VARCHAR( 255 ) NOT NULL DEFAULT '',
   create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) CHARACTER SET utf8 COLLATE utf8_bin;
-- ) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- set of degree mapping values (manually created)
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "DNP", "DNP (Doctor of Nursing Practice)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "EDD", "EDD (Doctor of Education)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "MA", "MA (Master of Arts)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "MAPE", "MAPE (Master of Arts in Physics Education)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "MAR", "MAR (Master of Architecture)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "MARH", "MARH (Master of Architectural History)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "ME", "ME (Master of Engineering)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "MFA", "MFA (Master of Fine Arts)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "MS", "MS (Master of Science)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "MUEP", "MUEP (Master of Urban and Environmental Planning)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "PHD", "PHD (Doctor of Philosophy)" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "SJD", "SJD (Doctor of Juridical Science)" );

-- set of department mapping values (use script/helpers/mappingsql.ksh to generate from the existing libra department_facet_mappings table
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ADMIN-EDD", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ADMIN-MED", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ANTHRO-PHD", "Department of Anthropology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ARCH-MAR", "Department of Architectural History" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ARH-MARH", "Department of Architectural History" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ARTARC-MA", "Department of Art" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ARTARC-PHD", "Department of Art" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ASTRON-PHD", "Department of Astronomy" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOL-MA", "Department of Biology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOL-MS", "Department of Biology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOL-PHD", "Department of Biology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOMEN-ME", "Department of Biomedical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOMEN-MS", "Department of Biomedical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOMEN-PHD", "Department of Biomedical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOMOL-PHD", "Department of Biochemistry and Molecular Genetics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOP-PHD", "Department of Biophysics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CELL-PHD", "Department of Molecular" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CHEM-MS", "Department of Chemistry" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CHEM-PHD", "Department of Chemistry" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CHEMEN-CGE", "Department of Chemical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CHEMEN-ME", "Department of Chemical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CHEMEN-MS", "Department of Chemical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CHEMEN-PHD", "Department of Chemical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CIVIL-CGE", "Department of Civil Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CIVIL-ME", "Department of Civil Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CIVIL-MS", "Department of Civil Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CIVIL-PHD", "Department of Civil Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CLAS-PHD", "Department of Classics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CLNPSY-PHD", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "COMPEN-ME", "Department of Computer Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "COMPEN-MS", "Department of Computer Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "COMPEN-PHD", "Department of Computer Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "COMPSC-MCS", "Department of Computer Science" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "COMPSC-MS", "Department of Computer Science" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "COMPSC-PHD", "Department of Computer Science" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "COUNS-EDD", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "COUNS-MED", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CURRIN-EDD", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CURRIN-MED", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "Civil & Env Engr", "Department of Civil Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "DRAMA-MFA", "Department of Drama" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EASIAN-MA", "Department of East Asian Studies" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ECON-PHD", "Department of Economics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EDPSYC-EDD", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EDPSYC-MED", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EDUC-PHD", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ELECT-CGE", "Department of Electrical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ELECT-ME", "Department of Electrical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ELECT-MS", "Department of Electrical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ELECT-PHD", "Department of Electrical Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ENGL-MA", "Department of English" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ENGL-PHD", "Department of English" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ENGPHY-CGE", "Department of Engineering Physics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ENGPHY-MEP", "Department of Engineering Physics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ENGPHY-MS", "Department of Engineering Physics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ENGPHY-PHD", "Department of Engineering Physics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EVSC-MA", "Department of Environmental Sciences" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EVSC-MS", "Department of Environmental Sciences" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EVSC-PHD", "Department of Environmental Sciences" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EXPATH-PHD", "Department of Pathology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "FORAFF-MA", "Department of Foreign Affairs" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "FORAFF-PHD", "Department of Politics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "FRENCH-PHD", "Department of French" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "GERMAN-MS", "Department of Germanic Languages and Literatures" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "GERMAN-PHD", "Department of Germanic Languages and Literatures" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "GOVT-MA", "Department of Politics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "GOVT-PHD", "Department of Politics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "HIGHED-EDD", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "HIGHED-MED", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "HIST-MA", "Department of History" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "HIST-PHD", "Department of History" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ITAL-MA", "Department of Spanish" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "KINES-MED", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MAE-CGE", "Department of Mechanical and Aerospace Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MAE-ME", "Department of Mechanical and Aerospace Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MAE-MS", "Department of Mechanical and Aerospace Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MAE-PHD", "Department of Mechanical and Aerospace Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MATH-PHD", "Department of Mathematics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MATSC-CGE", "Department of Materials Science and Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MATSC-MMSE", "Department of Materials Science and Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MATSCI-MS", "Department of Materials Science and Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MATSCI-PHD", "Department of Materials Science and Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MICRO-PHD", "Department of Microbiology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MUSIC-PHD", "Department of Music" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "NEURO-PHD", "Department of Neuroscience" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "NURS-DNP", "School of Nursing" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "NURS-PHD", "School of Nursing" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PHARM-PHD", "Department of Pharmacology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PHIL-PHD", "Department of Philosophy" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PHY-PHD", "Department of Molecular Physiology and Biological Physics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PHYS-MS", "Department of Physics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PHYS-PHD", "Department of Physics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PLAN-MUEP", "Department of Urban and Environmental Planning" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PSYCH-MA", "Department of Psychology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PSYCH-PHD", "Department of Psychology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "RELIG-MA", "Department of Religious Studies" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "RELIG-PHD", "Department of Religious Studies" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SLAVIC-MA", "Department of Slavic Languages and Literatures" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SLAVIC-PHD", "Department of Slavic Languages and Literatures" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SOCIOL-MA", "Department of Sociology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SOCIOL-PHD", "Department of Sociology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SPAN-PHD", "Department of Spanish" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SPATH-MED", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SPCED-MED", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "STATS-PHD", "Department of Statistics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SYSTEM-AM", "Department of Systems Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SYSTEM-CGE", "Department of Systems Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SYSTEM-ME", "Department of Systems Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SYSTEM-MS", "Department of Systems Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SYSTEM-PHD", "Department of Systems Engineering" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "University of Virginia Libraries", "University of Virginia Library" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "WRITE-MFA", "Department of English" );

-- manually added later after processing SIS files
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "ANTHRO-MA", "Department of Anthropology" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "BIOLPS-MS", "Department of Biological and Physical Sciences" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CHEM-MA", "Department of Chemistry" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "CLAS-MA", "Department of Classics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "EDPOL-EDD", "Department of Education Policy Studies" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "FRENCH-MA", "Department of French" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "GERMAN-MA", "Department of Germanic Languages and Literatures" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "KINES-EDD", "School of Education and Human Development" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "LAW-SJD", "School of Law" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MATH-MA", "Department of Mathematics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MATH-MS", "Department of Mathematics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MESA-MA", "Department of Middle Eastern and South Asian Languages and Cultures" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "MUSIC-MA", "Department of Music" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "PHYS-MAPE", "Department of Physics" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "SPAN-MA", "Department of Spanish" );
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "department", "STATS-MS", "Department of Statistics" );

-- added for new programs
INSERT INTO fieldmapper( field_class, field_name, field_value ) VALUES( "degree", "MED", "MED (Master of Education)" );
