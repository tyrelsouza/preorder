# Dockerfile for building the preorder service image
FROM debian:latest

# Update the package list and install MySQL client
RUN apt-get update && apt-get install -y \
    mariadb-client ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy the preorder binary from the local directory into the container
COPY ./preorder /usr/local/bin/preorder

# Set executable permissions on the binary
RUN chmod +x /usr/local/bin/preorder


# Copy the entrypoint script
#COPY entrypoint.sh /usr/local/bin/entrypoint.sh
#RUN chmod +x /usr/local/bin/entrypoint.sh
#
## Copy the SQL script to the container
#COPY create_database.sql /docker-entrypoint-initdb.d/create_database.sql
#
## Set the entrypoint script
#ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
#
# Specify the command to run when the container starts
CMD ["/usr/local/bin/preorder"]

