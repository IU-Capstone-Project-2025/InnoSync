package com.innosync;

import io.github.cdimascio.dotenv.Dotenv;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Info;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class BackendApplication {

	private static final Logger logger = LoggerFactory.getLogger(BackendApplication.class);

	public static void main(String[] args) {
		// Load .env file only if it exists (for local development)
		try {
			Dotenv dotenv = Dotenv.configure().ignoreIfMissing().load();
			if (dotenv.get("DB_URL") != null) {
				System.setProperty("DB_URL", dotenv.get("DB_URL"));
			}
			if (dotenv.get("DB_USERNAME") != null) {
				System.setProperty("DB_USERNAME", dotenv.get("DB_USERNAME"));
			}
			if (dotenv.get("DB_PASSWORD") != null) {
				System.setProperty("DB_PASSWORD", dotenv.get("DB_PASSWORD"));
			}
			logger.info("Environment variables loaded and application starting");
		} catch (Exception e) {
			logger.info("No .env file found, using system environment variables");
		}
		SpringApplication.run(BackendApplication.class, args);
	}

	@Bean
	public OpenAPI customOpenAPI() {
		return new OpenAPI()
				.info(new Info()
						.title("My API")
						.version("1.0")
						.description("Spring Boot REST API"));
	}
}
