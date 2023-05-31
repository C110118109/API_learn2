set DB_REPLICAS=0
set SOURCES_HOST=localhost
set SOURCES_PORT=5432
set SOURCES_USER=postgres
set SOURCES_PASSWORD=Kk918523
set SOURCES_DATABASE=homework
set SOURCES_SSLMODE=disable

set PUBLIC = ""
set PRIVATE = ""

@REM echo %SOURCES_DATABASE%
air.exe -c .air.toml

