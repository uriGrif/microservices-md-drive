# Use Maven to build the JAR file inside the container
FROM maven:3.9-eclipse-temurin-17 as builder

# Set the working directory
WORKDIR /app

# Copy source code and build
COPY . .
RUN mvn clean package -DskipTests

# Use only the JAR in the final image
FROM openjdk:17

WORKDIR /app

COPY --from=builder /app/target/profile-service.jar .

CMD ["java", "-jar", "profile-service.jar"]
