-- Imported records carried a third-party product name in labels shown in the UI. Relabel them
-- with the neutral wording the importer now writes. Only exact generated values are changed;
-- anything a user edited is left alone.

UPDATE public.companies SET division = 'Импортировано из отчёта' WHERE division = 'Imported from WellPlan';
UPDATE public.companies SET name = 'Импортированные кейсы' WHERE name = 'Imported (WellPlan)';
UPDATE public.trajectories
   SET description = 'Импортировано из отчёта: ' || substr(description, length('Imported from WellPlan: ') + 1)
 WHERE description LIKE 'Imported from WellPlan: %';
UPDATE public.cases SET case_description = 'Импортировано из отчёта' WHERE case_description = 'Imported from WellPlan';
UPDATE public.strings SET name = 'Рабочая колонна (импорт)' WHERE name = 'WellPlan work string';
