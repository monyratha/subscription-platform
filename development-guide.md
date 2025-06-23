# migrate create -ext sql -dir migrations -seq auth_service_schema
# migrate create -ext sql -dir migrations -seq seed_auth_data



# Docker

# List all containers (including stopped ones)
docker ps -a

# 2. Start a specific container by name or ID:
docker start <container_name_or_id>

# 3. Start and attach to the container (to see output):
docker start -a <container_name_or_id>

# 4. Start multiple containers:
docker start <container1> <container2>