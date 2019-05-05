FROM lessthanjake328/quickstart
CMD echo "Hello world! This is my first Docker image."
RUN ls
COPY . .
RUN pwd
RUN ls
