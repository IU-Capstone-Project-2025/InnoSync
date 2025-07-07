FROM eclipse-temurin:21-jdk
WORKDIR /app
COPY . .
RUN ./mvnw clean package -DskipTests

FROM eclipse-temurin:21-jre
WORKDIR /app
COPY --from=0 /app/target/*.jar app.jar
#COPY src/main/resources/static /app/static
EXPOSE 8080
CMD ["java", "-jar", "app.jar"]