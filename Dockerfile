FROM golang:1.7

# Create the directory where the application will reside
RUN mkdir /ms

# Copy the application files (needed for production)
ADD testMicroservice /ms/testMicroservice

# Set the working directory to the app directory
WORKDIR /ms

# Expose the application on port 8080.
# This should be the same as in the app.conf file
EXPOSE 8080

# Set the entry point of the container to the application executable
ENTRYPOINT /ms/testMicroservice