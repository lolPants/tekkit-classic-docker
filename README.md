# Tekkit Classic Docker ![Docker Build](https://github.com/lolPants/tekkit-classic-docker/workflows/Docker%20Build/badge.svg)

## 🚀 Running the Image
```sh
$ docker run -it -d \
    -p 25565:25565 \
    -e SERVER_OP=your_mc_username \
    --name tekkit-classic \
    docker.pkg.github.com/lolpants/tekkit-classic-docker/tekkit-classic-server:<tag>
```

Be sure to replace `<tag>` with the latest tag. See the [packages](https://github.com/lolPants/tekkit-classic-docker/packages) page for the latest.

### 📁 Volumes
Server data is stored in the `/minecraft` directory inside the container. You may wish to mount `/minecraft/world` to enable world persistence.

To configure plugins, you can mount `/minecraft/plugins`. To configure Tekkit mods, you can mount `/minecraft/config`.

### 🛠️ Configuration
All keys in `server.properties` can be configured at container runtime using environment variables. Available variables and their types are listed below.

```env
ALLOW_NETHER=<bool>
LEVEL_NAME=<string>
ENABLE_QUERY=<bool>
ALLOW_FLIGHT=<bool>
RCON_PASSWORD=<string>
SERVER_PORT=<int>
LEVEL_TYPE=<string>
ENABLE_RCON=<bool>
LEVEL_SEED=<string>
SERVER_IP=<string>
MAX_BUILD_HEIGHT=<int>
SPAWN_NPCS=<bool>
DEBUG=<bool>
WHITE_LIST=<bool>
SPAWN_ANIMALS=<bool>
ONLINE_MODE=<bool>
PVP=<bool>
DIFFICULTY=<int>
GAMEMODE=<int>
MAX_PLAYERS=<int>
RCON_PORT=<int>
SPAWN_MONSTERS=<bool>
GENERATE_STRUCTURES=<bool>
VIEW_DISTANCE=<int>
MOTD=<string>
```
