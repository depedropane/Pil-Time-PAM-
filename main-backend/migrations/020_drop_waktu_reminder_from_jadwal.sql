-- Drop waktu_reminder_pagi and waktu_reminder_malam from jadwal table
ALTER TABLE jadwal 
DROP COLUMN IF EXISTS waktu_reminder_pagi,
DROP COLUMN IF EXISTS waktu_reminder_malam;
