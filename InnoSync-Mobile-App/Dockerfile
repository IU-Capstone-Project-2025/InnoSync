FROM ubuntu:20.04

# Avoid interactive prompts during package installation
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && apt-get install -y \
    curl \
    git \
    unzip \
    xz-utils \
    zip \
    libglu1-mesa \
    openjdk-11-jdk \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Set up environment variables
ENV ANDROID_HOME=/opt/android-sdk
ENV PATH=$PATH:$ANDROID_HOME/cmdline-tools/latest/bin:$ANDROID_HOME/platform-tools

# Download and install Android SDK
RUN mkdir -p $ANDROID_HOME/cmdline-tools && \
    wget -q https://dl.google.com/android/repository/commandlinetools-linux-8512546_latest.zip -O cmdline-tools.zip && \
    unzip cmdline-tools.zip -d $ANDROID_HOME/cmdline-tools && \
    mv $ANDROID_HOME/cmdline-tools/cmdline-tools $ANDROID_HOME/cmdline-tools/latest && \
    rm cmdline-tools.zip

# Accept licenses and install required SDK components
RUN yes | sdkmanager --licenses && \
    sdkmanager "platform-tools" "platforms;android-33" "build-tools;33.0.0"

# Install Flutter
RUN git clone https://github.com/flutter/flutter.git /opt/flutter && \
    cd /opt/flutter && \
    git checkout 3.32.6 && \
    export PATH="$PATH:/opt/flutter/bin" && \
    flutter doctor

ENV PATH=$PATH:/opt/flutter/bin

# Set working directory to the project root
WORKDIR /app

# Copy the entire project
COPY . .

# Change to frontend directory and get dependencies
WORKDIR /app/frontend
RUN flutter pub get

# Build APK
CMD ["flutter", "build", "apk", "--release"] 