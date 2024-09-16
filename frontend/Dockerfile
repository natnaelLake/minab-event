# Use Node.js 18 as the base image
FROM node:18

# Set the working directory
WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN yarn install

# Copy the rest of the application code
COPY . .

# Build the application
RUN yarn run build

# Expose the port for the frontend
EXPOSE 3000

# Start the application
CMD ["yarn", "run", "dev"]
