# Build stage
FROM ubuntu:latest AS build

# Install dependencies
RUN apt-get update && apt-get install -y \
      curl git unzip xz-utils zip libglu1-mesa \
  && rm -rf /var/lib/apt/lists/*

# Install Flutter SDK
RUN git clone https://github.com/flutter/flutter.git /flutter
ENV PATH="/flutter/bin:${PATH}"

# Switch to stable channel and ensure web support
RUN flutter channel stable \
  && flutter upgrade \
  && flutter config --enable-web \
  && flutter precache --web

# Copy source and build
WORKDIR /app
COPY . .
RUN flutter pub get \
  && flutter clean \
  && flutter build web --release

# Serve with nginx
FROM nginx:alpine
COPY --from=build /app/build/web /usr/share/nginx/html

# Optional: set cache headers, SPA fallback
RUN printf 'server { \
    listen 80; root /usr/share/nginx/html; index index.html; \
    location / { try_files $uri $uri/ /index.html; } \
    location ~* \\.(js|css|png|jpg|jpeg|gif|ico|svg)$ { expires 1y; add_header Cache-Control "public, max-age=31536000"; } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
