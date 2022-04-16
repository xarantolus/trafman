FROM postgres
ENV POSTGRES_DB trafman_database
COPY create_database_scheme.sql /docker-entrypoint-initdb.d/
