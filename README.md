# Tekkit Classic Docker ![Docker Build](https://github.com/lolPants/tekkit-classic-docker/workflows/Docker%20Build/badge.svg)

## 🚀 Running the Image
```sh
$ docker run -d \
    -p 25565:25565 \
    -e SERVER_OP=your_mc_username \
    --name tekkit-classic \
    docker.pkg.github.com/lolpants/tekkit-classic-docker/tekkit-classic-server:<tag>
```

Be sure to replace `<tag>` with the latest tag. See the [packages](https://github.com/lolPants/tekkit-classic-docker/packages) page for the latest.
