# Use Node.js for building
FROM node:22 as builder

# Set the working directory
WORKDIR /app

# Copy source code and install dependencies
COPY . .
RUN npm install

# Build the app for production
RUN npm run build

# Use a lightweight server for serving the static files
FROM nginx:1.25

# Copy the built files
COPY --from=builder /app/dist /usr/share/nginx/html

COPY --from=builder /app/nginx.conf /etc/nginx/nginx.conf

# Expose the port
EXPOSE 80

# Command to start Nginx
CMD ["nginx", "-g", "daemon off;"]
