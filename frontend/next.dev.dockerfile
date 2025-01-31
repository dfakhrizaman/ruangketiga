# frontend/next.dev.dockerfile

# Stage 1: Base Image
FROM node:18-alpine AS base

WORKDIR /app

# Stage 2: Install Dependencies
FROM base AS deps

# Install libc6-compat if needed
RUN apk add --no-cache libc6-compat

# Copy dependency definitions
COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* ./

# Install dependencies based on the available lock file
RUN \
  if [ -f yarn.lock ]; then yarn install --frozen-lockfile; \
  elif [ -f package-lock.json ]; then npm ci; \
  elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && pnpm install --frozen-lockfile; \
  else echo "Lockfile not found." && exit 1; \
  fi

# Stage 3: Development Builder
FROM base AS builder

WORKDIR /app

# Copy installed dependencies
COPY --from=deps /app/node_modules ./node_modules

# Copy the rest of the application code
COPY . .

# Expose Next.js default port
EXPOSE 3000

# Start Next.js in development mode
CMD ["yarn", "dev"]
