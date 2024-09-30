FROM golang:1.23-alpine AS builder
WORKDIR /root/
RUN git clone https://github.com/willroberts/minecraft-client.git
WORKDIR /root/minecraft-client/
RUN CGO_ENABLED=0 go build -o /root/ ./cmd/shell/main.go

FROM eclipse-temurin:21-jre-noble
WORKDIR /src/server_files
COPY --from=builder /root/main /root/
CMD java -Xms4G -Xmx4G -XX:+UseG1GC -XX:+ParallelRefProcEnabled -XX:MaxGCPauseMillis=200 -XX:+UnlockExperimentalVMOptions -XX:+DisableExplicitGC -XX:+AlwaysPreTouch -XX:G1NewSizePercent=30 -XX:G1MaxNewSizePercent=40 -XX:G1HeapRegionSize=8M -XX:G1ReservePercent=20 -XX:G1HeapWastePercent=5 -XX:G1MixedGCCountTarget=4 -XX:InitiatingHeapOccupancyPercent=15 -XX:G1MixedGCLiveThresholdPercent=90 -XX:G1RSetUpdatingPauseTimePercent=5 -XX:SurvivorRatio=32 -XX:+PerfDisableSharedMem -XX:MaxTenuringThreshold=1 -Dusing.aikars.flags=https://mcflags.emc.gs -Daikars.new.flags=true -jar /src/server_files/minecraft_server.jar nogui